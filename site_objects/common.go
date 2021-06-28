package site_objects

import "golang.org/x/text/message"

type StatSummary struct {
	Min    float64
	Median float64
	Mean   float64
	Max    float64
	Vals   []float64
	Sum    float64
}

type Page struct {
	StartAt int64
	EndAt   int64
}

type Navigation struct {
	HasPrev     bool
	HasNext     bool
	PrevPage    Page
	NextPage    Page
	LastPage    Page
	CurrentPage Page
	Sort        string
	Query       string
	Time        string
}

type FuncUtils struct {
	Printer   *message.Printer
	Translate func(string, *message.Printer) string
	Auth      AuthInfo
	MakeSlice func(...interface{}) []interface{}
}
