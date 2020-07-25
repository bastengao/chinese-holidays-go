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

func TestDays(t *testing.T) {
	e := event{
		Name:  "元旦",
		Type:  "holiday",
		Range: []string{"2016-12-31", "2017-01-02"},
	}

	day0, _ := parseTime("2016-12-31")
	day1, _ := parseTime("2017-01-01")
	day2, _ := parseTime("2017-01-02")
	days, err := e.days()
	if err != nil {
		t.Fail()
	}

	if len(days) != 3 {
		t.Fail()
	}

	if !days[0].Equal(day0) || !days[1].Equal(day1) || !days[2].Equal(day2) {
		t.Fail()
	}
}
