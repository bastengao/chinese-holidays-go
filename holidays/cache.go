package holidays

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const dataEndpoint = "http://chinese-holidays-data.basten.me/data"

var _ Queryer = (*cache)(nil)

type cache struct {
	l             sync.RWMutex
	b             *book
	indexChecksum [sha256.Size]byte
}

type entry struct {
	Year         int    `json:"year"`
	LastModified string `json:"last_modified"`
}

// NewCacheQueryer returns a new Queryer that fetches online data and check updates every day.
func NewCacheQueryer() (Queryer, error) {
	url := fmt.Sprintf("%s/index.json", dataEndpoint)
	b, err := downloadData(url)
	if err != nil {
		return nil, err
	}

	book, err := newBookFromEntries(b)
	if err != nil {
		return nil, err
	}

	cache := &cache{
		b:             book,
		indexChecksum: sha256.Sum256(b),
	}
	go cache.updateInterval()

	return cache, nil
}

func (c *cache) IsHoliday(date time.Time) (bool, error) {
	c.l.RLock()
	b, err := c.b.IsHoliday(date)
	c.l.RUnlock()
	return b, err
}

func (c *cache) IsWorkingday(date time.Time) (bool, error) {
	c.l.RLock()
	b, err := c.b.IsWorkingday(date)
	c.l.RUnlock()
	return b, err
}

func (c *cache) updateInterval() {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			err := c.update()
			if err != nil {
				log.Println(err)
			}
		}
	}()
}

func (c *cache) update() error {
	url := fmt.Sprintf("%s/index.json", dataEndpoint)
	b, err := downloadData(url)
	if err != nil {
		return err
	}

	checkSum := sha256.Sum256(b)
	if c.indexChecksum == checkSum { // same, skip update
		return nil
	}

	newBook, err := newBookFromEntries(b)
	if err != nil {
		return err
	}

	c.l.Lock()
	c.b = newBook
	c.indexChecksum = checkSum
	c.l.Unlock()
	return nil
}

func newBookFromEntries(data []byte) (*book, error) {
	var entries []entry
	err := json.Unmarshal(data, &entries)
	if err != nil {
		return nil, err
	}

	var events []event
	for _, entry := range entries {
		url := fmt.Sprintf("%s/%d.json", dataEndpoint, entry.Year)
		b, err := downloadData(url)
		if err != nil {
			return nil, err
		}

		e, err := parseEvents(b)
		if err != nil {
			return nil, err
		}

		events = append(events, e...)
	}

	book, err := newBookfromEvents(events)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func downloadData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
