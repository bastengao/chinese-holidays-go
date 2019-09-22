package holidays

import "time"

type book struct {
	events []event
	index  map[string]event
}

func newBookfromEvents(events []event) (book, error) {
	index := make(map[string]event)
	for _, e := range events {
		days, err := e.days()
		if err != nil {
			return book{}, err
		}

		for _, day := range days {
			index[_key(day)] = e
		}
	}
	return book{events, index}, nil
}

func (b book) isHoliday(d time.Time) bool {
	e := b.findEvent(d)

	if e == nil {
		return isWeekend(d)
	}

	return e.isHoliday()
}

func (b book) isWorkingday(d time.Time) bool {
	e := b.findEvent(d)

	if e == nil {
		return !isWeekend(d)
	}

	return e.isWorkingday()
}

func (b book) findEvent(d time.Time) *event {
	e, ok := b.index[_key(d)]
	if !ok {
		return nil
	}

	return &e
}

func isWeekend(d time.Time) bool {
	day := d.Weekday()
	return day == 6 || day == 0
}

func _key(d time.Time) string {
	return d.Format("2006-01-02")
}
