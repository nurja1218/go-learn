package rms

import "time"

type DateRangeType struct {
	from time.Time
	to   time.Time
}

type Search struct {
	text      string
	dateRange DateRangeType
	types     []string
}

func NewSearch() Search {
	format := "1900-01-01 00:00:00"
	s := Search{}
	s.text = ""
	s.dateRange.from, _ = time.Parse(format, format)
	s.dateRange.to = time.Now()
	s.types = []string{""}

	return s
}
