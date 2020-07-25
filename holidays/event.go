package holidays

import (
	"fmt"
	"time"
)

const local = "Asia/Shanghai"
const typeHoliday = "holiday"
const typeWorkingday = "workingday"

var location, lerr = time.LoadLocation(local)

func init() {
	if lerr != nil {
		fmt.Println(lerr)
	}
}

type event struct {
	Name  string
	Type  string
	Range []string
}

func (e event) isHoliday() bool {
	return e.Type == typeHoliday
}

func (e event) isWorkingday() bool {
	return e.Type == typeWorkingday
}

func (e event) days() ([]time.Time, error) {
	if len(e.Range) == 1 {
		t, err := parseTime(e.Range[0])
		if err != nil {
			return nil, err
		}
		return []time.Time{t}, nil
	} else if len(e.Range) == 2 {
		start, err := parseTime(e.Range[0])
		if err != nil {
			return nil, err
		}

		end, err := parseTime(e.Range[1])
		if err != nil {
			return nil, err
		}
		return rangeDates(start, end), nil
	}

	return nil, fmt.Errorf("Wrong Range %v", e.Range)
}

func parseTime(s string) (time.Time, error) {
	var t time.Time
	t, err := time.ParseInLocation("2006-01-02", s, location)
	if err != nil {
		return t, err
	}

	return t, err
}

func rangeDates(start, end time.Time) []time.Time {
	y, m, d := start.Date()
	date := time.Date(y, m, d, 0, 0, 0, 0, location)
	dates := []time.Time{}
	for ; !date.After(end); date = date.AddDate(0, 0, 1) {
		dates = append(dates, date)
	}
	return dates
}
