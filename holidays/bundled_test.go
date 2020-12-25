package holidays

import (
	"testing"
)

func TestLoadData(t *testing.T) {
	events, err := loadData()
	if err != nil {
		t.Error(err)
	}

	if len(events) != 75 {
		t.Error(len(events))
	}
}
