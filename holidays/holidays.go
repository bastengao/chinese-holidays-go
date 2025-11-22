// Package holidays supply holidays query.
package holidays

import (
	"fmt"
	"time"
)

const version = "1.7.0"

var b Queryer

func init() {
	events, err := loadData()
	if err != nil {
		fmt.Println(err)
	}

	bk, err := newBookfromEvents(events)
	if err != nil {
		fmt.Println(err)
	} else {
		b = &bk
	}
}

// BundleQueryer returns a bundle queryer.
func BundleQueryer() (Queryer, error) {
	err := checkInitBook()
	if err != nil {
		return nil, err
	}

	return b, nil
}

// IsHoliday checks given date is holiday or not.
//
// Deprecated: Use Queryer.IsHoliday instead.
func IsHoliday(date time.Time) (bool, error) {
	err := checkInitBook()
	if err != nil {
		return false, err
	}

	return b.IsHoliday(date)
}

// IsWorkingday checks given date is working day or not.
//
// Deprecated: Use Queryer.IsWorkingday instead.
func IsWorkingday(date time.Time) (bool, error) {
	err := checkInitBook()
	if err != nil {
		return false, err
	}

	return b.IsWorkingday(date)
}

func checkInitBook() error {
	if b == nil {
		return fmt.Errorf("book initialize failed")
	}
	return nil
}
