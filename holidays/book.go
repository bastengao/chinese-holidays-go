package holidays

import "time"

var _ Queryer = (*book)(nil)

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

func (b *book) IsHoliday(d time.Time) (bool, error) {
	e := b.findEvent(d)

	if e == nil {
		return isWeekend(d), nil
	}

	return e.isHoliday(), nil
}

func (b *book) IsWorkingday(d time.Time) (bool, error) {
	e := b.findEvent(d)

	if e == nil {
		return !isWeekend(d), nil
	}

	return e.isWorkingday(), nil
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
