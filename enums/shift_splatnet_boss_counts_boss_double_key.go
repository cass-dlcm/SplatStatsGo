package enums

import (
	"encoding/json"
	"errors"
	"fmt"
)

type ShiftSplatnetBossCountsBossDoubleKeyEnum string

const (
	goldieKey    ShiftSplatnetBossCountsBossDoubleKeyEnum = "sakelien-golden"
	steelheadKey ShiftSplatnetBossCountsBossDoubleKeyEnum = "sakelien-bomber"
	flyfishKey   ShiftSplatnetBossCountsBossDoubleKeyEnum = "sakelien-cup-twins"
	scrapperKey  ShiftSplatnetBossCountsBossDoubleKeyEnum = "sakelien-shield"
	steelEelKey  ShiftSplatnetBossCountsBossDoubleKeyEnum = "sakelien-snake"
	stingerKey   ShiftSplatnetBossCountsBossDoubleKeyEnum = "sakelien-tower"
	mawsKey      ShiftSplatnetBossCountsBossDoubleKeyEnum = "sakediver"
	grillerKey   ShiftSplatnetBossCountsBossDoubleKeyEnum = "sakedozer"
	drizzlerKey  ShiftSplatnetBossCountsBossDoubleKeyEnum = "sakerocket"
)

func (ssbcbdke *ShiftSplatnetBossCountsBossDoubleKeyEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SSBCBDKE ShiftSplatnetBossCountsBossDoubleKeyEnum
	r := (*SSBCBDKE)(ssbcbdke)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *ssbcbdke {
	case goldieKey, steelheadKey, flyfishKey, scrapperKey, steelEelKey, stingerKey, mawsKey, grillerKey, drizzlerKey:
		return nil
	}
	return errors.New("Invalid ShiftSplatnetBossCountsBossDoubleKeyEnum. Got: " + fmt.Sprint(*ssbcbdke))
}
