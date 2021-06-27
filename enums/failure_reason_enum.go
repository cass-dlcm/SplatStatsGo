package enums

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/message"
)

type FailureReasonEnum string

const (
	AnyFailureReason FailureReasonEnum = "any"
	wipeOut          FailureReasonEnum = "wipe_out"
	timeLimit        FailureReasonEnum = "time_limit"
)

func (fre *FailureReasonEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type FRE FailureReasonEnum
	r := (*FRE)(fre)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *fre {
	case wipeOut, timeLimit:
		return nil
	}
	return errors.New("Invalid FailureReasonEnum. Got: " + fmt.Sprint(*fre))
}

func (fre FailureReasonEnum) GetDisplay(printer *message.Printer) string {
	switch fre {
	case AnyFailureReason:
		return printer.Sprintf("Any Failure Reason")
	case wipeOut:
		return printer.Sprintf("Wipe out")
	case timeLimit:
		return printer.Sprint("Time is up")
	}
	return ""
}

func GetFailureReasons() []FailureReasonEnum {
	return []FailureReasonEnum{AnyFailureReason, wipeOut, timeLimit}
}
