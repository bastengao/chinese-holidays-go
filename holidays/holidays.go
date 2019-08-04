package holidays

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/bastengao/chinese-holidays-go/holidays/statik" // load data
	"github.com/rakyll/statik/fs"
)

type event struct {
	Name  string
	Type  string
	Range []string
}

func loadData() ([]event, error) {
	statikFS, err := fs.New()
	if err != nil {
		return nil, err
	}
	var events []event
	err = fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		if path == "/" {
			return err
		}
		b, err := fs.ReadFile(statikFS, path)
		if err != nil {
			return nil
		}
		e, err := parseEvents(b)
		events = append(events, e...)

		fmt.Println(path)
		return err
	})
	if err != nil {
		return nil, err
	}
	return events, nil
}

func parseEvents(b []byte) ([]event, error) {
	var events []event
	err := json.Unmarshal(b, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
