package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	start time.Time
	end   time.Time
}

func NewDateRange(start, end time.Time) (DateRange, error) {
	if start.IsZero() || end.IsZero() {
		return DateRange{}, errors.New("dates can't be zero")
	}

	if end.Before(start) {
		return DateRange{}, errors.New("end before start")
	}

	return DateRange{
		start: start,
		end:   end,
	}, nil
}

func (d DateRange) Hours() float64 {
	return d.end.Sub(d.start).Hours()
}

func main() {
	lifetime, _ := NewDateRange(
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	)

	fmt.Println(lifetime.Hours())

	travelInTime, _ := NewDateRange(
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	)

	fmt.Println(travelInTime.Hours())
}
