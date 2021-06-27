/* This file is part of SplatStatsGo.
 *
 * SplatStatsGo is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * SplatStatsGo is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with SplatStatsGo.  If not, see <https://www.gnu.org/licenses/>.
 */

package enums

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/text/message"
)

type SalmonStageEnum string

const (
	AnySalmonStage SalmonStageEnum = "Any Stage"
	smokeyard      SalmonStageEnum = "Salmonid Smokeyard"
	polaris        SalmonStageEnum = "Ruins of Ark Polaris"
	grounds        SalmonStageEnum = "Spawning Grounds"
	bay            SalmonStageEnum = "Marooner's Bay"
	outpost        SalmonStageEnum = "Lost Outpost"
)

type SalmonSplatnetScheduleStageImageEnum string

const (
	smokeyardSplatnetImg SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/e9f7c7b35e6d46778cd3cbc0d89bd7e1bc3be493.png"
	polarisSplatnetImg   SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/50064ec6e97aac91e70df5fc2cfecf61ad8615fd.png"
	groundsSplatnetImg   SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/65c68c6f0641cc5654434b78a6f10b0ad32ccdee.png"
	baySplatnetImg       SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/e07d73b7d9f0c64e552b34a2e6c29b8564c63388.png"
	outpostSplatnetImg   SalmonSplatnetScheduleStageImageEnum = "/images/coop_stage/6d68f5baa75f3a94e5e9bfb89b82e7377e3ecd2c.png"
)

func (sse *SalmonStageEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SSE SalmonStageEnum
	r := (*SSE)(sse)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *sse {
	case smokeyard, polaris, grounds, bay, outpost:
		return nil
	}
	return errors.New("Invalid SalmonStageEnum. Got: " + fmt.Sprint(*sse))
}

func (sse SalmonStageEnum) GetDisplay(printer *message.Printer) string {
	return printer.Sprintf("%s", sse)
}

func GetSalmonStage() []SalmonStageEnum {
	return []SalmonStageEnum{
		AnySalmonStage, smokeyard, polaris, grounds, bay, outpost,
	}
}

func (ssssie *SalmonSplatnetScheduleStageImageEnum) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SSSSIE SalmonSplatnetScheduleStageImageEnum
	r := (*SSSSIE)(ssssie)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *ssssie {
	case smokeyardSplatnetImg, polarisSplatnetImg, groundsSplatnetImg, baySplatnetImg, outpostSplatnetImg:
		return nil
	}
	return errors.New("Invalid SalmonSplatnetScheduleStageImageEnum. Got: " + fmt.Sprint(*ssssie))
}
