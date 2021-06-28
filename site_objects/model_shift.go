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

package site_objects

import (
	"github.com/cass-dlcm/SplatStatsGo/enums"
	"time"
)

type ShiftResults struct {
	Utils            FuncUtils
	HasShifts        bool
	Clears           int
	Wave2Clears      int
	Wave1Clears      int
	ClearPercent     float64
	WaveTwoPercent   float64
	WaveOnePercent   float64
	PlayerGoldenEggs StatSummary
	TeamGoldenEggs   StatSummary
	PlayerPowerEggs  StatSummary
	TeamPowerEggs    StatSummary
	PlayerRevives    StatSummary
	PlayerDeaths     StatSummary
	DangerRate       StatSummary
	ShiftSummaries   []ShiftInfo
	Nav              Navigation
	CurrentDate      string
	Stages           []enums.SalmonStageEnum
	Specials         []enums.SalmonSpecial
	FailureReasons   []enums.FailureReasonEnum
}

type ShiftInfo struct {
	PlayerName   string
	PlayerId     int64
	JobId        int64
	Stage        enums.SalmonStageEnum
	IsClear      bool
	PlayerGolden int
	TeamGolden   int
	TeamPower    int
	DangerRate   float64
	PlayerTitle  enums.SalmonTitle
	GradePoint   int
	Time         time.Time
}

type ShiftDetails struct {
	Utils             FuncUtils
	JobId             int64
	Stage             enums.SalmonStageEnum
	Result            string
	Title             enums.SalmonTitle
	GradePoint        int
	DangerRate        float64
	ScheduleStartTime time.Time
	HasEndTime        bool
	ScheduleEndTime   time.Time
	PlayTime          time.Time
	Waves             []ShiftWave
	Players           []ShiftPlayer
	Names             []string
	Bosses            []ShiftBoss
}

type ShiftWave struct {
	Num         int
	Event       enums.SalmonEvent
	WaterLevel  enums.SalmonWaterLevel
	Quota       int
	Delivers    int
	Appearances int
	PowerEggs   int
}

type ShiftPlayer struct {
	Name     string
	Weapons  []enums.SalmonWeaponEnum
	Special  enums.SalmonSpecial
	Specials []int
	Rescues  int
	Deaths   int
	Golden   int
	Power    int
}

type ShiftBoss struct {
	Name        enums.SalmonBoss
	Appearances int
	Kills       []int
}
