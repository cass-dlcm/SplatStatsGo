package enums

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/message"
)

type Rule string

const (
	AllRules Rule = "any"
	sz       Rule = "splat_zones"
	rm       Rule = "rainmaker"
	cb       Rule = "clam_blitz"
	tc       Rule = "tower_control"
	TurfWar  Rule = "turf_war"
)

func (rule Rule) GetDisplay(printer *message.Printer) string {
	switch rule {
	case AllRules:
		return printer.Sprintf("Any Rules")
	case sz:
		return printer.Sprintf("Splat Zones")
	case rm:
		return printer.Sprintf("Rainmaker")
	case cb:
		return printer.Sprintf("Clam Blitz")
	case tc:
		return printer.Sprintf("Tower Control")
	case TurfWar:
		return printer.Sprintf("Turf War")
	}
	return ""
}

func GetRule() []Rule {
	return []Rule{
		AllRules, sz, rm, cb, tc, TurfWar,
	}
}

func (rule *Rule) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type R Rule
	r := (*R)(rule)
	if err := json.Unmarshal(b, &r); err != nil {
		panic(err)
	}
	switch *rule {
	case AllRules, sz, cb, rm, tc, TurfWar:
		return nil
	}
	return errors.New("Invalid FailureReasonEnum. Got: " + fmt.Sprint(*rule))
}
