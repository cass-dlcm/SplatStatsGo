/*
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package enums

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/message"
)

type SalmonEvent string

const (
	cohock       SalmonEvent = "cohock-charge"
	fog          SalmonEvent = "fog"
	goldieSeek   SalmonEvent = "goldie-seeking"
	grillerEvent SalmonEvent = "griller"
	rush         SalmonEvent = "rush"
	mothership   SalmonEvent = "the-mothership"
	waterLevels  SalmonEvent = "water-levels"
)

func (se SalmonEvent) GetDisplay(printer *message.Printer) string {
	switch se {
	case cohock:
		return printer.Sprintf("Cohock Charge")
	case fog:
		return printer.Sprintf("Fog")
	case goldieSeek:
		return printer.Sprintf("Goldie Seeking")
	case grillerEvent:
		return printer.Sprintf("The Griller")
	case rush:
		return printer.Sprintf("Rush")
	case mothership:
		return printer.Sprintf("The Mothership")
	case waterLevels:
		return ""
	}
	return ""
}

func (se *SalmonEvent) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SE SalmonEvent
	r := (*SE)(se)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *se {
	case cohock, fog, goldieSeek, grillerEvent, rush, mothership, waterLevels:
		return nil
	}
	return errors.New("Invalid SalmonEvent. Got: " + fmt.Sprint(*se))
}
