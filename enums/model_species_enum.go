package enums

import (
	"encoding/json"
	"errors"
	"fmt"
)

type SpeciesEnum string

const (
	inklings    SpeciesEnum = "inklings"
	octolings   SpeciesEnum = "octolings"
	speciesNone SpeciesEnum = ""
)

func (se *SpeciesEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SE SpeciesEnum
	r := (*SE)(se)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *se {
	case inklings, octolings, speciesNone:
		return nil
	}
	return errors.New("Invalid SpeciesEnum. Got: " + fmt.Sprint(*se))
}
