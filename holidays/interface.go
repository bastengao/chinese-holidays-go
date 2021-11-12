package holidays

import "time"

// Queryer is the interface that wraps the Query method.
type Queryer interface {
	// IsHoliday checks given date is holiday or not.
	IsHoliday(date time.Time) (bool, error)

	// IsWorkingday checks given date is working day or not.
	IsWorkingday(date time.Time) (bool, error)
}
