package holidays

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

const dataEndpoint = "http://chinese-holidays-data.basten.me/data"

var userAgent = fmt.Sprintf("chinese-holidays-go/%s", version)

var errNotModified = errors.New("index not modified")

var _ Queryer = (*cache)(nil)

type cache struct {
	l             sync.RWMutex
	b             *book
	indexChecksum [sha256.Size]byte
	indexMtime    string
}

type entry struct {
	Year         int    `json:"year"`
	LastModified string `json:"last_modified"`
}

// NewCacheQueryer returns a new Queryer that fetches online data and check updates every day.
func NewCacheQueryer() (Queryer, error) {
	url := fmt.Sprintf("%s/index.json", dataEndpoint)
	b, mtime, err := downloadData(url, "")
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
		indexMtime:    mtime,
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
	b, mtime, err := downloadData(url, c.indexMtime)
	// 如果内容没有变化直接退出
	if errors.Is(err, errNotModified) {
		return nil
	}
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
	c.indexMtime = mtime
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
		b, _, err := downloadData(url, "")
		if err != nil {
			return nil, err
		}

		e, err := parseEvents(b)
		if err != nil {
			return nil, fmt.Errorf("parse year %d: %w", entry.Year, err)
		}

		events = append(events, e...)
	}

	book, err := newBookfromEvents(events)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// downloadData 返回内容和 Last-Modified 头
func downloadData(url string, mtime string) ([]byte, string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, "", err
	}
	req.Header.Set("UserAgent", userAgent)
	if mtime != "" {
		req.Header.Set("If-Modified-Since", mtime)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotModified {
		return nil, "", errNotModified
	} else if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("download error: %d", resp.StatusCode)
	}

	newMtime := resp.Header.Get("Last-Modified")
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	return b, newMtime, nil
}
