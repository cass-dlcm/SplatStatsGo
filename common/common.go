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

package common

import (
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/enums"
	"golang.org/x/text/message"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

func GetBounds(r *http.Request) (int64, int64) {
	startAtStr, ok := r.URL.Query()["startAt"]
	if !ok || len(startAtStr) < 1 {
		startAtStr = []string{"0"}
	}
	startAt, err := strconv.ParseInt(startAtStr[0], 0, 64)
	if err != nil {
		log.Println(err)
		startAt = 0
	}
	endAtStr, ok := r.URL.Query()["endAt"]
	if !ok || len(endAtStr) < 1 {
		endAtStr = []string{"50"}
	}
	endAt, err := strconv.ParseInt(endAtStr[0], 0, 64)
	if err != nil {
		log.Println(err)
		endAt = 50
	}
	return startAt, endAt
}

func GetSort(r *http.Request) string {
	sortStringArr, ok := r.URL.Query()["sort"]
	if !ok || len(sortStringArr) < 1 {
		return "desc"
	}
	if strings.ToLower(sortStringArr[0]) == "desc" || strings.ToLower(sortStringArr[0]) == "asc" {
		return sortStringArr[0]
	}
	return "desc"
}

func GetLanguage(r *http.Request) string {
	lang := r.Header.Get("Accept-Language")
	if lang == "" {
		return "en-US"
	}
	return strings.Split(lang, ",")[0]
}

func CheckAllFlag(r *http.Request) bool {
	allArr, ok := r.URL.Query()["all"]
	if !ok || len(allArr) < 1 {
		return false
	}
	if strings.ToLower(allArr[0]) == "true" {
		return true
	}
	return false
}

func CalcMedian(nums []float64) float64 {
	if len(nums) == 0 {
		return -1
	}
	sort.Float64s(nums) // sort the numbers

	mNumber := len(nums) / 2

	if len(nums)%2 == 1 {
		return nums[mNumber]
	}

	return (nums[mNumber-1] + nums[mNumber]) / 2
}

func Translate(s string, printer *message.Printer) string {
	return printer.Sprintf(s)
}

func GetTimeQuery(r *http.Request) (string, []time.Time) {
	fromStrArr := r.URL.Query()["from"]
	if len(fromStrArr) < 1 {
		fromStrArr = []string{""}
	}
	from, err := time.Parse("2006-01-02", fromStrArr[0])
	if err != nil {
		from, _ = time.Parse("2006-01-02", "2017-07-21")
	}
	toStrArr := r.URL.Query()["to"]
	if len(toStrArr) < 1 {
		toStrArr = []string{""}
	}
	to, err := time.Parse("2006-01-02", toStrArr[0])
	if err != nil {
		to = time.Now()
	}
	return fmt.Sprintf("from=%s&to=%s",
			from.Format("2006-01-02"), to.Format("2006-01-02"),
		), []time.Time{
			from, to,
		}
}

func GetBattlesQuery(r *http.Request) (string, []interface{}, []string) {
	ruleStringArr, ok := r.URL.Query()["rule"]
	var rule enums.Rule
	if !ok || len(ruleStringArr) < 1 {
		rule = enums.AllRules
	} else {
		for _, item := range enums.GetRule() {
			if item == enums.Rule(ruleStringArr[0]) {
				rule = enums.Rule(ruleStringArr[0])
			}
		}
		if rule == "" {
			rule = enums.AllRules
		}
	}
	matchTypeStringArr, ok := r.URL.Query()["match_type"]
	var matchType enums.Lobby
	if !ok || len(matchTypeStringArr) < 1 {
		matchType = enums.AnyLobby
	} else {
		for _, item := range enums.GetLobby() {
			if item == enums.Lobby(matchTypeStringArr[0]) {
				matchType = enums.Lobby(matchTypeStringArr[0])
			}
		}
		if matchType == "" {
			matchType = enums.AnyLobby
		}
	}
	rankStringArr, ok := r.URL.Query()["rank"]
	var rank enums.Rank
	if !ok || len(rankStringArr) < 1 {
		rank = enums.AnyRank
	} else {
		for _, item := range enums.GetRanks() {
			if item == enums.Rank(strings.ToUpper(rankStringArr[0])) {
				rank = enums.Rank(strings.ToUpper(rankStringArr[0]))
			}
		}
		if rank == "" {
			rank = enums.AnyRank
		}
	}
	weaponStringArr, ok := r.URL.Query()["weapon"]
	var weapon enums.BattleWeaponEnum
	if !ok || len(weaponStringArr) < 1 {
		weapon = enums.AnyWeapon
	} else {
		for _, item := range enums.GetBattleWeaponEnum() {
			if item == enums.BattleWeaponEnum(weaponStringArr[0]) {
				weapon = enums.BattleWeaponEnum(weaponStringArr[0])
			}
		}
		if weapon == "" {
			weapon = enums.AnyWeapon
		}
	}
	stageStringArr, ok := r.URL.Query()["stage"]
	var stage enums.BattleStage
	if !ok || len(stageStringArr) < 1 {
		stage = enums.AnyStage
	} else {
		for _, item := range enums.GetStages() {
			if item == enums.BattleStage(stageStringArr[0]) {
				stage = enums.BattleStage(stageStringArr[0])
			}
		}
		if stage == "" {
			stage = enums.AnyStage
		}
	}
	winStringArr, ok := r.URL.Query()["win"]
	var win enums.TrinaryBool
	if !ok || len(winStringArr) < 1 {
		win = enums.AnyBool
	} else {
		switch enums.TrinaryBool(strings.ToLower(winStringArr[0])) {
		case enums.BoolT, enums.BoolTrue:
			win = enums.BoolT
		case enums.BoolF, enums.BoolFalse:
			win = enums.BoolF
		default:
			win = enums.AnyBool
		}
	}
	hasDCStringArr, ok := r.URL.Query()["has_dc"]
	var hasDC enums.TrinaryBool
	if !ok || len(hasDCStringArr) < 1 {
		hasDC = enums.AnyBool
	} else {
		switch enums.TrinaryBool(strings.ToLower(hasDCStringArr[0])) {
		case enums.BoolT, enums.BoolTrue:
			hasDC = enums.BoolT
		case enums.BoolF, enums.BoolFalse:
			hasDC = enums.BoolF
		default:
			hasDC = enums.AnyBool
		}
	}
	return fmt.Sprintf("rule=%s&match_type=%s&rank=%s&weapon=%s&stage=%s&win=%s&has_dc=%s",
			rule, matchType, rank, weapon, stage, win, hasDC,
		), []interface{}{
			rule,
			matchType,
			rank,
			weapon,
			stage,
			win,
			hasDC,
		}, []string{
			"rule",
			"match_type",
			"player_rank",
			"player_weapon",
			"stage",
			"win",
			"has_disconnected_player",
		}
}

func GetSalmonQuery(r *http.Request) (string, []interface{}, []string) {
	stageStringArr, ok := r.URL.Query()["stage"]
	var stage enums.SalmonStageEnum
	if !ok || len(stageStringArr) < 1 {
		stage = enums.AnySalmonStage
	} else {
		for _, item := range enums.GetSalmonStage() {
			if item == enums.SalmonStageEnum(stageStringArr[0]) {
				stage = enums.SalmonStageEnum(stageStringArr[0])
			}
		}
		if stage == "" {
			stage = enums.AnySalmonStage
		}
	}
	specialStringArr, ok := r.URL.Query()["special"]
	var special enums.SalmonSpecial
	if !ok || len(specialStringArr) < 1 {
		special = enums.AnySalmonSpecial
	} else {
		for _, item := range enums.GetSalmonSpecials() {
			if item == enums.SalmonSpecial(specialStringArr[0]) {
				special = enums.SalmonSpecial(specialStringArr[0])
			}
		}
		if special == "" {
			special = enums.AnySalmonSpecial
		}
	}
	clearStringArr, ok := r.URL.Query()["cleared"]
	var clear enums.TrinaryBool
	if !ok || len(clearStringArr) < 1 {
		clear = enums.AnyBool
	} else {
		switch enums.TrinaryBool(strings.ToLower(clearStringArr[0])) {
		case enums.BoolT, enums.BoolTrue:
			clear = enums.BoolT
		case enums.BoolF, enums.BoolFalse:
			clear = enums.BoolF
		default:
			clear = enums.AnyBool
		}
	}
	failReasonStringArr, ok := r.URL.Query()["fail_reason"]
	var failReason enums.FailureReasonEnum
	if !ok || len(failReasonStringArr) < 1 {
		failReason = enums.AnyFailureReason
	} else {
		for _, item := range enums.GetFailureReasons() {
			if item == enums.FailureReasonEnum(failReasonStringArr[0]) {
				failReason = enums.FailureReasonEnum(failReasonStringArr[0])
			}
		}
		if failReason == "" {
			failReason = enums.AnyFailureReason
		}
	}
	return fmt.Sprintf("stage=%s&special=%s&cleared=%s&fail_reason=%s",
			stage, special, clear, failReason,
		), []interface{}{
			stage, special, clear, failReason,
		}, []string{
			"stage", "player_special", "is_clear", "job_failure_reason",
		}
}

func MakeSlice(args ...interface{}) []interface{} {
	return args
}

func IntOrZero(num *int) int {
	if num != nil {
		return *num
	}
	return 0
}

func SalmonEventOrEmpty(event *enums.SalmonEvent) enums.SalmonEvent {
	if event != nil {
		return *event
	}
	return ""
}
