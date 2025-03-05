package holidays

import (
	"testing"
	"time"
)

func TestNewCacheQueryer(t *testing.T) {
	queryer, err := NewCacheQueryer()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if queryer == nil {
		t.FailNow()
	}

	d := time.Date(2019, 10, 1, 0, 0, 0, 0, location)
	r, err := queryer.IsHoliday(d)
	if err != nil {
		t.Error(err)
	}
	if !r {
		t.Fail()
	}
}

func TestNewCacheQueryer_update(t *testing.T) {
	queryer, err := NewCacheQueryer()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if queryer == nil {
		t.FailNow()
	}

	cache, ok := queryer.(*cache)
	if !ok {
		t.FailNow()
	}
	err = cache.update()
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	d := time.Date(2019, 10, 1, 0, 0, 0, 0, location)
	r, err := queryer.IsHoliday(d)
	if err != nil {
		t.Error(err)
	}
	if !r {
		t.Fail()
	}
}
