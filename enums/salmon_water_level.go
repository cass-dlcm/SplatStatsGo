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

type SalmonWaterLevel string

const (
	low    SalmonWaterLevel = "low"
	normal SalmonWaterLevel = "normal"
	high   SalmonWaterLevel = "high"
)

func (swl SalmonWaterLevel) GetDisplay(printer *message.Printer) string {
	switch swl {
	case low:
		return printer.Sprintf("Low Tide")
	case normal:
		return printer.Sprintf("Mid Tide")
	case high:
		return printer.Sprintf("High Tide")
	}
	return ""
}

func (swl *SalmonWaterLevel) UnmarshalJSON(b []byte) error {
	// Define a secondary type to avoid ending up with a recursive call to json.Unmarshal
	type SWL SalmonWaterLevel
	r := (*SWL)(swl)
	err := json.Unmarshal(b, &r)
	if err != nil {
		panic(err)
	}
	switch *swl {
	case low, normal, high:
		return nil
	}
	return errors.New("Invalid SalmonWaterLevel. Got: " + fmt.Sprint(*swl))
}
