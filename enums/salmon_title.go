package enums

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/message"
)

type SalmonTitle string

const (
	AnyTitle      SalmonTitle = "any"
	intern        SalmonTitle = "0"
	apprentice    SalmonTitle = "1"
	partTimer     SalmonTitle = "2"
	goGetter      SalmonTitle = "3"
	overachiever  SalmonTitle = "4"
	profreshional SalmonTitle = "5"
)

func (st *SalmonTitle) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type ST SalmonTitle
	r := (*ST)(st)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *st {
	case intern, apprentice, partTimer, goGetter, overachiever, profreshional:
		return nil
	}
	return errors.New("Invalid SalmonTitle. Got: " + fmt.Sprint(*st))
}

func (st SalmonTitle) GetDisplay(printer *message.Printer) string {
	switch st {
	case AnyTitle:
		return printer.Sprintf("Any Title")
	case intern:
		return printer.Sprintf("Intern")
	case apprentice:
		return printer.Sprintf("Apprentice")
	case partTimer:
		return printer.Sprintf("Part-Timer")
	case goGetter:
		return printer.Sprintf("Go-Getter")
	case overachiever:
		return printer.Sprintf("Overachiever")
	case profreshional:
		return printer.Sprintf("Profreshional")
	}
	return ""
}

func GetTitles() []SalmonTitle {
	return []SalmonTitle{
		AnyTitle, intern, apprentice, partTimer, goGetter, overachiever, profreshional,
	}
}
