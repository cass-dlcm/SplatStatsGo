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

package api_objects

import (
	"github.com/cass-dlcm/SplatStatsGo/enums"
	"time"
)

type StatInkGender struct {
	Key     enums.GenderEnum `json:"key"`
	Name    StatInkName      `json:"name"`
	Iso5218 int              `json:"iso5218"`
}

type StatInkKeyName struct {
	Key  string      `json:"key"`
	Name StatInkName `json:"name"`
}

type StatInkName struct {
	DeDE string `json:"de_DE,omitempty"`
	EnGB string `json:"en_GB,omitempty"`
	EnUS string `json:"en_US,omitempty"`
	EsES string `json:"es_ES,omitempty"`
	EsMX string `json:"es_MX,omitempty"`
	FrCA string `json:"fr_CA,omitempty"`
	FrFR string `json:"fr_FR,omitempty"`
	ItIT string `json:"it_IT,omitempty"`
	JaJP string `json:"ja_JP,omitempty"`
	NlNL string `json:"nl_NL,omitempty"`
	RuRU string `json:"ru_RU,omitempty"`
	ZhCN string `json:"zh_CN,omitempty"`
	ZhTW string `json:"zh_TW,omitempty"`
}

type StatInkTime struct {
	Time    int       `json:"time"`
	Iso8601 time.Time `json:"iso8601"`
}

type StatInkProfile struct {
	//Nnid        *interface{} `json:"nnid,omitempty"`
	FriendCode *string `json:"friend_code,omitempty"`
	Twitter    *string `json:"twitter,omitempty"`
	//Ikanakama   *interface{} `json:"ikanakama,omitempty"`
	//Ikanakama2  *interface{} `json:"ikanakama2,omitempty"`
	//Environment *interface{} `json:"environment,omitempty"`
}
