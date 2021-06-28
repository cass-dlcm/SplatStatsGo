package site_objects

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

import (
	"github.com/cass-dlcm/SplatStatsGo/enums"
	"time"
)

type BattleResults struct {
	HasBattles      bool
	Wins            int
	WinRate         float64
	Kills           StatSummary
	Deaths          StatSummary
	Assists         StatSummary
	Specials        StatSummary
	Inked           StatSummary
	BattleSummaries []BattleInfo
	Nav             Navigation
	Query           string
	Rules           []enums.Rule
	Lobbies         []enums.Lobby
	Weapons         []enums.BattleWeaponEnum
	Ranks           []enums.Rank
	Stages          []enums.BattleStage
	CurrentDate     string
	Utils           FuncUtils
}

type BattleQuery struct {
	Rule      enums.Rule
	MatchType enums.Lobby
	Rank      enums.Rank
	Weapon    enums.BattleWeaponEnum
}

type BattleInfo struct {
	PlayerName        string
	UserId            int64
	BattleNumber      int64
	Rule              string
	Stage             string
	PlayerWeaponImage string
	PlayerWeapon      string
	Result            string
	PlayerKills       int
	PlayerDeaths      int
	Time              time.Time
}

type BattleDetails struct {
	Utils             FuncUtils
	BattleNumber      int64
	Rule              string
	MatchType         string
	Stage             string
	Result            string
	EndResult         string
	MyTeamCount       float64
	OtherTeamCount    float64
	StartTime         time.Time
	EndTime           time.Time
	ElapsedTimeMinSec string
	ElapsedTime       int
	Players           []PlayerDetails
}

type PlayerDetails struct {
	Name           string
	WeaponIcon     string
	Weapon         string
	LevelStar      int
	Level          int
	Rank           string
	GamePaintPoint int
	KA             int
	Assists        int
	Specials       int
	Kills          int
	Deaths         int
	HeadgearMain   string
	HeadgearSub0   string
	HeadgearSub1   string
	HeadgearSub2   string
	ClothesMain    string
	ClothesSub0    string
	ClothesSub1    string
	ClothesSub2    string
	ShoesMain      string
	ShoesSub0      string
	ShoesSub1      string
	ShoesSub2      string
}
