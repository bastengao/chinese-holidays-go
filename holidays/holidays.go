// Package holidays supply holidays query.
package holidays

import (
	"fmt"
	"time"
)

var b *book

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

// IsHoliday checks given date is holiday or not.
func IsHoliday(date time.Time) (bool, error) {
	err := checkInitBook()
	if err != nil {
		return false, err
	}

	return b.isHoliday(date), nil
}

// IsWorkingday checks given date is working day or not.
func IsWorkingday(date time.Time) (bool, error) {
	err := checkInitBook()
	if err != nil {
		return false, err
	}

	return b.isWorkingday(date), nil
}

func checkInitBook() error {
	if b == nil {
		return fmt.Errorf("book initialize failed")
	}
	return nil
}
