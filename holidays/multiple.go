package holidays

import "time"

var _ Queryer = multipleQueryer{}

type multipleQueryer struct {
	queryers []Queryer
}

// NewMultipleQueryer returns a new Queryer that delegates query to underlying multiple Queryers.
// Try each queryers in order until one returns a result.
func NewMultipleQueryer(queryers ...Queryer) Queryer {
	return multipleQueryer{queryers}
}

func (m multipleQueryer) IsHoliday(date time.Time) (b bool, err error) {
	for _, q := range m.queryers {
		b, err = q.IsHoliday(date)
		if err == nil {
			return b, nil
		}
	}

	return false, err
}

func (m multipleQueryer) IsWorkingday(date time.Time) (b bool, err error) {
	for _, q := range m.queryers {
		b, err = q.IsWorkingday(date)
		if err == nil {
			return b, nil
		}
	}

	return false, err
}
