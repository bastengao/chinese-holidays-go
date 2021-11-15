package holidays

import (
	"testing"
	"time"
)

func TestNewMultipleQueryer(t *testing.T) {
	bundleQueryer, err := BundleQueryer()
	if err != nil {
		t.Error(err)
	}

	cacheQueryer, err := NewCacheQueryer()
	if err != nil {
		t.Error(err)
	}

	queryer := NewMultipleQueryer(cacheQueryer, bundleQueryer)
	d := time.Date(2019, 10, 1, 0, 0, 0, 0, location)
	r, err := queryer.IsHoliday(d)
	if err != nil {
		t.Error(err)
	}
	if !r {
		t.Fail()
	}
}
