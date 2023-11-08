package holidays

import (
	"embed"
	"encoding/json"
	"io/fs"
)

//go:embed data/*
var dfs embed.FS

func loadData() ([]event, error) {
	var events []event
	err := fs.WalkDir(dfs, ".", func(path string, d fs.DirEntry, err error) error {
		if path == "/" {
			return err
		}
		b, err := dfs.ReadFile(path)
		if err != nil {
			return nil
		}
		e, err := parseEvents(b)
		events = append(events, e...)

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
