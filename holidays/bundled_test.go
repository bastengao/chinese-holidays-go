package holidays

import (
	"testing"
)

func TestLoadData(t *testing.T) {
	_, err := loadData()
	if err != nil {
		t.Error(err)
	}
}
