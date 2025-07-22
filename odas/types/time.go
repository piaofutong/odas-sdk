package types

import "time"

// YMDTime represents a date in YYYYMMDD format
type YMDTime int64

// Time converts YMDTime to time.Time
func (t YMDTime) Time() time.Time {
	return time.Date(int(t/10000), time.Month(t%10000/100), int(t%100), 0, 0, 0, 0, time.Local)
}

// Format formats YMDTime according to the given TimeFilterType
func (t YMDTime) Format(dType TimeFilterType) string {
	switch dType {
	case TimeFilterYear:
		return t.Time().Format("2006")
	case TimeFilterMonth:
		return t.Time().Format("2006-01")
	}
	return t.Time().Format("2006-01-02")
}

// TimeFilterType represents the type of time filtering
type TimeFilterType int

const (
	TimeFilterDay TimeFilterType = iota
	TimeFilterMonth
	TimeFilterYear
)

// Page represents pagination parameters
type Page struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// TimeBetween represents a time range
type TimeBetween struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// TimeSpan represents a time span with filtering options
type TimeSpan struct {
	Start        time.Time      `json:"start"`
	End          time.Time      `json:"end"`
	Type         TimeFilterType `json:"type"`
	IntervalDays int            `json:"intervalDays"`
}

// Time is an alias for time.Time for compatibility
type Time = time.Time

// OrderType represents the type of order
type OrderType int

// BasedOnOrderEnum represents the based on order enumeration
type BasedOnOrderEnum int

// RegionType_Enums represents the region type enumeration
type RegionType_Enums int

// PassedTimeBetween represents a time range for passed time
type PassedTimeBetween struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}
