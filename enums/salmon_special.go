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

import "golang.org/x/text/message"

type SalmonSpecial string

const (
	AnySalmonSpecial        SalmonSpecial = "any"
	salmonSplatBombLauncher SalmonSpecial = "2"
	salmonStingRay          SalmonSpecial = "7"
	salmonInkjet            SalmonSpecial = "8"
	salmonSplashdown        SalmonSpecial = "9"
)

func (ss SalmonSpecial) GetDisplay(printer *message.Printer) string {
	switch ss {
	case AnySalmonSpecial:
		return printer.Sprintf("Any Special")
	case salmonSplatBombLauncher:
		return printer.Sprintf("Splat-Bomb Launcher")
	case salmonStingRay:
		return printer.Sprintf("Sting Ray")
	case salmonInkjet:
		return printer.Sprintf("Inkjet")
	case salmonSplashdown:
		return printer.Sprintf("Splashdown")
	}
	return ""
}

func GetSalmonSpecials() []SalmonSpecial {
	return []SalmonSpecial{
		AnySalmonSpecial, salmonSplatBombLauncher, salmonStingRay, salmonInkjet, salmonSplashdown,
	}
}
