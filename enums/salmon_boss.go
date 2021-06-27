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

import "golang.org/x/text/message"

type SalmonBoss string

const (
	Goldie    SalmonBoss = "sakelien-golden"
	Steelhead SalmonBoss = "sakelien-bomber"
	Flyfish   SalmonBoss = "sakelien-cup-twins"
	Scrapper  SalmonBoss = "sakelien-shield"
	SteelEel  SalmonBoss = "sakelien-snake"
	Stinger   SalmonBoss = "sakelien-tower"
	Maws      SalmonBoss = "sakediver"
	Griller   SalmonBoss = "sakedozer"
	Drizzler  SalmonBoss = "sakerocket"
)

func (sb SalmonBoss) GetDisplay(printer *message.Printer) string {
	switch sb {
	case Goldie:
		return printer.Sprintf("Goldie")
	case Steelhead:
		return printer.Sprintf("Steelhead")
	case Flyfish:
		return printer.Sprintf("Flyfish")
	case Scrapper:
		return printer.Sprintf("Scrapper")
	case SteelEel:
		return printer.Sprintf("Steel Eel")
	case Stinger:
		return printer.Sprintf("Stinger")
	case Maws:
		return printer.Sprintf("Maws")
	case Griller:
		return printer.Sprintf("Griller")
	case Drizzler:
		return printer.Sprintf("Drizzler")
	}
	return ""
}
