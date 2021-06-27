package enums

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Rank string

const (
	AnyRank Rank = "any"
	cMinus  Rank = "C-"
	cRank   Rank = "C"
	cPlus   Rank = "C+"
	bMinus  Rank = "B-"
	bRank   Rank = "B"
	bPlus   Rank = "B+"
	aMinus  Rank = "A-"
	aRank   Rank = "A"
	aPlus   Rank = "A+"
	sRank   Rank = "S"
	sPlus   Rank = "S+"
	xRank   Rank = "X"
)

func GetRanks() []Rank {
	return []Rank{
		AnyRank, cMinus, cRank, cPlus, bMinus, bRank, bPlus, aMinus, aRank, aPlus, sRank, sPlus, xRank,
	}
}

func (rank *Rank) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type R Rank
	r := (*R)(rank)
	if err := json.Unmarshal(b, &r); err != nil {
		panic(err)
	}
	switch *rank {
	case AnyRank, cMinus, cRank, cPlus, bMinus, bRank, bPlus, aMinus, aRank, aPlus, sRank, sPlus, xRank:
		return nil
	}
	return errors.New("Invalid FailureReasonEnum. Got: " + fmt.Sprint(*rank))
}
