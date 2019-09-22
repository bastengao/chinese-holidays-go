package holidays

import (
	"testing"
)

func TestParseTime(t *testing.T) {
	d, err := parseTime("2019-08-04")
	if err != nil {
		t.Error(err)
	}

	if d.Year() != 2019 || d.Month() != 8 || d.Day() != 4 {
		t.Fail()
	}
}

func TestRangeDates(t *testing.T) {
	s, _ := parseTime("2019-08-01")
	e, _ := parseTime("2019-08-03")
	days := rangeDates(s, e)
	if len(days) != 3 {
		t.Fail()
	}

	if days[0] != s || days[2] != e {
		t.Fail()
	}
}
