package holidays

import (
	"testing"
)

func TestIsWeekend(t *testing.T) {
	sat, err := parseTime("2019-08-03")
	if err != nil {
		t.Error(err)
	}

	r := isWeekend(sat)
	if !r {
		t.Fail()
	}

	sun, err := parseTime("2019-08-04")
	if err != nil {
		t.Error(err)
	}

	r = isWeekend(sun)
	if !r {
		t.Fail()
	}
}

func TestIsWeekendFalse(t *testing.T) {
	friday, err := parseTime("2019-08-02")
	if err != nil {
		t.Error(err)
	}

	r := isWeekend(friday)
	if r {
		t.Fail()
	}
}
