package datetime

import (
	"time"
)

var Layout string = "2006-01-02"

func GetLowDateLimit() time.Time {
	LowLimit, _ := time.Parse(Layout, "1900-01-01")
	return LowLimit
}

func GetHighDateLimit() time.Time {
	HighLimit, _ := time.Parse(Layout, "2050-01-01")
	return HighLimit
}

func DateStringToTime(date string) (time.Time, error) {
	Date, err := time.Parse(Layout, date[:10])
	if err != nil {
		return time.Time{}, err
	}
	return Date, nil
}

func IsInRange(StartDate time.Time, EndDate time.Time, Date time.Time) bool {
	return Date.After(StartDate) && Date.Before(EndDate)
}
