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

package salmon

import (
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/common"
	"github.com/cass-dlcm/SplatStatsGo/db_objects"
	"github.com/cass-dlcm/SplatStatsGo/enums"
	"github.com/cass-dlcm/SplatStatsGo/obj_sql"
	"github.com/cass-dlcm/SplatStatsGo/site_objects"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"math"
	"net/http"
	"path"
	"strconv"
	"text/template"
	"time"
)

var indexTmpl *template.Template
var detailTmpl *template.Template

func initTemplates() error {
	var err error
	indexTmpl, err = template.ParseFiles("tmpl/base.gohtml", "tmpl/main_site.gohtml", "tmpl/two_salmon/index.gohtml", "tmpl/two_salmon/filter.gohtml")
	if err != nil {
		return err
	}
	detailTmpl, err = template.ParseFiles("tmpl/base.gohtml", "tmpl/main_site.gohtml", "tmpl/two_salmon/shift.gohtml")
	return err
}

func Index(w http.ResponseWriter, r *http.Request) {
	if indexTmpl == nil {
		if err := initTemplates(); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	langHeader := common.GetLanguage(r)
	sort := common.GetSort(r)
	salmonQueryStr, salmonQuery, columns := common.GetSalmonQuery(r)
	timeStr, timeQuery := common.GetTimeQuery(r)
	startAt, endAt := common.GetBounds(r)
	salmonQuery = append(salmonQuery, timeQuery[0], timeQuery[1])
	columns = append(columns, "play_timefrom", "play_timeto")
	shifts, err := obj_sql.GetShiftsLean(salmonQuery, columns, startAt, endAt, sort)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pks, err := obj_sql.ReadKeyArrayWithCondition(salmonQuery, columns, "pk", "two_salmon_shift", sort, "play_time")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	results := site_objects.ShiftResults{
		PlayerGoldenEggs: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		TeamGoldenEggs: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		PlayerPowerEggs: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		TeamPowerEggs: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		PlayerRevives: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		PlayerDeaths: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		DangerRate: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		ShiftSummaries: make([]site_objects.ShiftInfo, len(shifts)),
		Nav: site_objects.Navigation{
			CurrentPage: site_objects.Page{
				StartAt: startAt,
				EndAt:   endAt,
			},
			Sort:  sort,
			Query: salmonQueryStr,
			Time:  timeStr,
		},
		Utils: site_objects.FuncUtils{
			Translate: common.Translate,
		},
		CurrentDate:    time.Now().Format("2006-01-02"),
		Stages:         enums.GetSalmonStage(),
		Specials:       enums.GetSalmonSpecials(),
		FailureReasons: enums.GetFailureReasons(),
	}
	userCookie, err := r.Cookie("session_token")
	var user *int64
	var printer *message.Printer
	if err != nil {
		log.Println(err)
		printer = message.NewPrinter(language.Make(langHeader))
	} else {
		user, err = obj_sql.CheckSessionToken(userCookie.Value)
		if err != nil {
			printer = message.NewPrinter(language.Make(langHeader))
			log.Println(err)
		} else {
			if user != nil {
				results.Utils.Auth.Authenticated = true
				results.Utils.Auth.UserId = *user
				userObj, err := obj_sql.ReadUserById(*user)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				results.Utils.Auth.Username = (*userObj).Username
				printer = message.NewPrinter(language.Make((*userObj).Locale))
			} else {
				printer = message.NewPrinter(language.Make(langHeader))
			}
		}
	}
	results.Utils.Printer = printer
	if (2*startAt - endAt) >= 0 {
		results.Nav.HasPrev = true
		results.Nav.PrevPage = site_objects.Page{
			StartAt: 2*startAt - endAt,
			EndAt:   startAt,
		}
	} else if startAt > 0 {
		results.Nav.HasPrev = true
		results.Nav.PrevPage = site_objects.Page{
			StartAt: 0,
			EndAt:   endAt - startAt,
		}
	}
	if int64(len(pks)) > endAt {
		results.Nav.HasNext = true
		results.Nav.NextPage = site_objects.Page{
			StartAt: endAt,
			EndAt:   (endAt - startAt) + endAt,
		}
		results.Nav.LastPage = site_objects.Page{
			StartAt: ((int64(len(pks)) - startAt) / (endAt - startAt)) * (endAt - startAt),
			EndAt:   int64(len(pks)),
		}
	}
	for i, pk := range pks {
		shift, err := obj_sql.GetShiftLeanPk(pk)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if shift.IsClear {
			results.Clears += 1
			results.Wave2Clears += 1
			results.Wave1Clears += 1
		} else if *shift.FailureWave == 3 {
			results.Wave2Clears += 1
			results.Wave1Clears += 1
		} else if *shift.FailureWave == 2 {
			results.Wave1Clears += 1
		}
		results.PlayerGoldenEggs.Min = math.Min(results.PlayerGoldenEggs.Min, float64(shift.PlayerGoldenEggs))
		results.PlayerGoldenEggs.Max = math.Max(results.PlayerGoldenEggs.Max, float64(shift.PlayerGoldenEggs))
		results.PlayerGoldenEggs.Vals[i] = float64(shift.PlayerGoldenEggs)
		results.PlayerGoldenEggs.Sum += float64(shift.PlayerGoldenEggs)
		teamGoldenDelivered := shift.Wave1GoldenDelivered
		if shift.Wave2GoldenDelivered != nil {
			teamGoldenDelivered += *shift.Wave2GoldenDelivered
			if shift.Wave3GoldenDelivered != nil {
				teamGoldenDelivered += *shift.Wave3GoldenDelivered
			}
		}
		results.TeamGoldenEggs.Min = math.Min(results.TeamGoldenEggs.Min, float64(teamGoldenDelivered))
		results.TeamGoldenEggs.Max = math.Max(results.TeamGoldenEggs.Max, float64(teamGoldenDelivered))
		results.TeamGoldenEggs.Vals[i] = float64(teamGoldenDelivered)
		results.TeamGoldenEggs.Sum += float64(teamGoldenDelivered)
		results.PlayerPowerEggs.Min = math.Min(results.PlayerPowerEggs.Min, float64(shift.PlayerPowerEggs))
		results.PlayerPowerEggs.Max = math.Max(results.PlayerPowerEggs.Max, float64(shift.PlayerPowerEggs))
		results.PlayerPowerEggs.Vals[i] = float64(shift.PlayerPowerEggs)
		results.PlayerPowerEggs.Sum += float64(shift.PlayerPowerEggs)
		teamPower := shift.Wave1PowerEggs
		if shift.Wave2PowerEggs != nil {
			teamPower += *shift.Wave2PowerEggs
			if shift.Wave3PowerEggs != nil {
				teamPower += *shift.Wave3PowerEggs
			}
		}
		results.TeamPowerEggs.Min = math.Min(results.TeamPowerEggs.Min, float64(teamPower))
		results.TeamPowerEggs.Max = math.Max(results.TeamPowerEggs.Max, float64(teamPower))
		results.TeamPowerEggs.Vals[i] = float64(teamPower)
		results.TeamPowerEggs.Sum += float64(teamPower)
		results.PlayerRevives.Min = math.Min(results.PlayerRevives.Min, float64(shift.PlayerReviveCount))
		results.PlayerRevives.Max = math.Max(results.PlayerRevives.Max, float64(shift.PlayerReviveCount))
		results.PlayerRevives.Vals[i] = float64(shift.PlayerReviveCount)
		results.PlayerRevives.Sum += float64(shift.PlayerReviveCount)
		results.PlayerDeaths.Min = math.Min(results.PlayerDeaths.Min, float64(shift.PlayerDeathCount))
		results.PlayerDeaths.Max = math.Max(results.PlayerDeaths.Max, float64(shift.PlayerDeathCount))
		results.PlayerDeaths.Vals[i] = float64(shift.PlayerDeathCount)
		results.PlayerDeaths.Sum += float64(shift.PlayerDeathCount)
		results.DangerRate.Min = math.Min(results.DangerRate.Min, shift.DangerRate)
		results.DangerRate.Max = math.Max(results.DangerRate.Max, shift.DangerRate)
		results.DangerRate.Vals[i] = shift.DangerRate
		results.DangerRate.Sum += shift.DangerRate
	}
	for i := range shifts {
		results.HasShifts = true
		teamGoldenDelivered := shifts[i].Wave1GoldenDelivered
		if shifts[i].Wave2GoldenDelivered != nil {
			teamGoldenDelivered += *shifts[i].Wave2GoldenDelivered
			if shifts[i].Wave3GoldenDelivered != nil {
				teamGoldenDelivered += *shifts[i].Wave3GoldenDelivered
			}
		}
		teamPower := shifts[i].Wave1PowerEggs
		if shifts[i].Wave2PowerEggs != nil {
			teamPower += *shifts[i].Wave2PowerEggs
			if shifts[i].Wave3PowerEggs != nil {
				teamPower += *shifts[i].Wave3PowerEggs
			}
		}
		results.ShiftSummaries[i] = site_objects.ShiftInfo{
			PlayerName:   shifts[i].PlayerName,
			PlayerId:     shifts[i].UserId,
			JobId:        shifts[i].JobId,
			Stage:        shifts[i].Stage,
			IsClear:      shifts[i].IsClear,
			PlayerGolden: shifts[i].PlayerGoldenEggs,
			TeamGolden:   teamGoldenDelivered,
			TeamPower:    teamPower,
			DangerRate:   shifts[i].DangerRate,
			PlayerTitle:  shifts[i].PlayerTitle,
			GradePoint:   shifts[i].GradePoint,
			Time:         time.Unix(shifts[i].PlayTime, 0),
		}
	}
	if results.HasShifts {
		results.ClearPercent = float64(results.Clears) / float64(len(pks)) * 100
		results.WaveTwoPercent = float64(results.Wave2Clears) / float64(len(pks)) * 100
		results.WaveOnePercent = float64(results.Wave1Clears) / float64(len(pks)) * 100
		results.PlayerGoldenEggs.Mean = results.PlayerGoldenEggs.Sum / float64(len(pks))
		results.TeamGoldenEggs.Mean = results.TeamGoldenEggs.Sum / float64(len(pks))
		results.PlayerPowerEggs.Mean = results.PlayerPowerEggs.Sum / float64(len(pks))
		results.TeamPowerEggs.Mean = results.TeamPowerEggs.Sum / float64(len(pks))
		results.PlayerRevives.Mean = results.PlayerRevives.Sum / float64(len(pks))
		results.PlayerDeaths.Mean = results.PlayerDeaths.Sum / float64(len(pks))
		results.DangerRate.Mean = results.DangerRate.Sum / float64(len(pks))
		results.PlayerGoldenEggs.Median = common.CalcMedian(results.PlayerGoldenEggs.Vals)
		results.TeamGoldenEggs.Median = common.CalcMedian(results.TeamGoldenEggs.Vals)
		results.PlayerPowerEggs.Median = common.CalcMedian(results.PlayerPowerEggs.Vals)
		results.TeamPowerEggs.Median = common.CalcMedian(results.TeamPowerEggs.Vals)
		results.PlayerRevives.Median = common.CalcMedian(results.PlayerRevives.Vals)
		results.PlayerDeaths.Median = common.CalcMedian(results.PlayerDeaths.Vals)
		results.DangerRate.Median = common.CalcMedian(results.DangerRate.Vals)
	}
	if err := indexTmpl.Execute(w, results); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func IndexUser(w http.ResponseWriter, r *http.Request) {
	if indexTmpl == nil {
		if err := initTemplates(); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	langHeader := common.GetLanguage(r)
	sort := common.GetSort(r)
	userNum, err := strconv.ParseInt(path.Base(r.URL.Path), 0, 64)
	salmonQueryStr, salmonQuery, columns := common.GetSalmonQuery(r)
	timeStr, timeQuery := common.GetTimeQuery(r)
	startAt, endAt := common.GetBounds(r)
	salmonQuery = append(salmonQuery, timeQuery[0], timeQuery[1])
	columns = append(columns, "play_timefrom", "play_timeto")
	shifts, err := obj_sql.GetShiftsLeanUser(userNum, salmonQuery, columns, startAt, endAt, sort)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	salmonQuery = append(salmonQuery, userNum)
	columns = append(columns, "user_id")
	pks, err := obj_sql.ReadKeyArrayWithCondition(salmonQuery, columns, "pk", "two_salmon_shift", sort, "play_time")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	results := site_objects.ShiftResults{
		PlayerGoldenEggs: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		TeamGoldenEggs: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		PlayerPowerEggs: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		TeamPowerEggs: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		PlayerRevives: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		PlayerDeaths: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		DangerRate: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Vals: make([]float64, len(pks)),
		},
		ShiftSummaries: make([]site_objects.ShiftInfo, len(shifts)),
		Nav: site_objects.Navigation{
			CurrentPage: site_objects.Page{
				StartAt: startAt,
				EndAt:   endAt,
			},
			Sort:  sort,
			Query: salmonQueryStr,
			Time:  timeStr,
		},
		Utils: site_objects.FuncUtils{
			Translate: common.Translate,
		},
		CurrentDate:    time.Now().Format("2006-01-02"),
		Stages:         enums.GetSalmonStage(),
		Specials:       enums.GetSalmonSpecials(),
		FailureReasons: enums.GetFailureReasons(),
	}
	userCookie, err := r.Cookie("session_token")
	var user *int64
	var printer *message.Printer
	if err != nil {
		log.Println(err)
		printer = message.NewPrinter(language.Make(langHeader))
	} else {
		user, err = obj_sql.CheckSessionToken(userCookie.Value)
		if err != nil {
			printer = message.NewPrinter(language.Make(langHeader))
			log.Println(err)
		} else {
			if user != nil {
				results.Utils.Auth.Authenticated = true
				results.Utils.Auth.UserId = *user
				userObj, err := obj_sql.ReadUserById(*user)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				results.Utils.Auth.Username = (*userObj).Username
				printer = message.NewPrinter(language.Make((*userObj).Locale))
			} else {
				printer = message.NewPrinter(language.Make(langHeader))
			}
		}
	}
	results.Utils.Printer = printer
	if (2*startAt - endAt) >= 0 {
		results.Nav.HasPrev = true
		results.Nav.PrevPage = site_objects.Page{
			StartAt: 2*startAt - endAt,
			EndAt:   startAt,
		}
	} else if startAt > 0 {
		results.Nav.HasPrev = true
		results.Nav.PrevPage = site_objects.Page{
			StartAt: 0,
			EndAt:   endAt - startAt,
		}
	}
	if int64(len(pks)) > endAt {
		results.Nav.HasNext = true
		results.Nav.NextPage = site_objects.Page{
			StartAt: endAt,
			EndAt:   (endAt - startAt) + endAt,
		}
		results.Nav.LastPage = site_objects.Page{
			StartAt: ((int64(len(pks)) - startAt) / (endAt - startAt)) * (endAt - startAt),
			EndAt:   int64(len(pks)),
		}
	}
	for i, pk := range pks {
		shift, err := obj_sql.GetShiftLeanPk(pk)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		if shift.IsClear {
			results.Clears += 1
			results.Wave2Clears += 1
			results.Wave1Clears += 1
		} else if *shift.FailureWave == 3 {
			results.Wave2Clears += 1
			results.Wave1Clears += 1
		} else if *shift.FailureWave == 2 {
			results.Wave1Clears += 1
		}
		results.PlayerGoldenEggs.Min = math.Min(results.PlayerGoldenEggs.Min, float64(shift.PlayerGoldenEggs))
		results.PlayerGoldenEggs.Max = math.Max(results.PlayerGoldenEggs.Max, float64(shift.PlayerGoldenEggs))
		results.PlayerGoldenEggs.Vals[i] = float64(shift.PlayerGoldenEggs)
		results.PlayerGoldenEggs.Sum += float64(shift.PlayerGoldenEggs)
		teamGoldenDelivered := shift.Wave1GoldenDelivered
		if shift.Wave2GoldenDelivered != nil {
			teamGoldenDelivered += *shift.Wave2GoldenDelivered
			if shift.Wave3GoldenDelivered != nil {
				teamGoldenDelivered += *shift.Wave3GoldenDelivered
			}
		}
		results.TeamGoldenEggs.Min = math.Min(results.TeamGoldenEggs.Min, float64(teamGoldenDelivered))
		results.TeamGoldenEggs.Max = math.Max(results.TeamGoldenEggs.Max, float64(teamGoldenDelivered))
		results.TeamGoldenEggs.Vals[i] = float64(teamGoldenDelivered)
		results.TeamGoldenEggs.Sum += float64(teamGoldenDelivered)
		results.PlayerPowerEggs.Min = math.Min(results.PlayerPowerEggs.Min, float64(shift.PlayerPowerEggs))
		results.PlayerPowerEggs.Max = math.Max(results.PlayerPowerEggs.Max, float64(shift.PlayerPowerEggs))
		results.PlayerPowerEggs.Vals[i] = float64(shift.PlayerPowerEggs)
		results.PlayerPowerEggs.Sum += float64(shift.PlayerPowerEggs)
		teamPower := shift.Wave1PowerEggs
		if shift.Wave2PowerEggs != nil {
			teamPower += *shift.Wave2PowerEggs
			if shift.Wave3PowerEggs != nil {
				teamPower += *shift.Wave3PowerEggs
			}
		}
		results.TeamPowerEggs.Min = math.Min(results.TeamPowerEggs.Min, float64(teamPower))
		results.TeamPowerEggs.Max = math.Max(results.TeamPowerEggs.Max, float64(teamPower))
		results.TeamPowerEggs.Vals[i] = float64(teamPower)
		results.TeamPowerEggs.Sum += float64(teamPower)
		results.PlayerRevives.Min = math.Min(results.PlayerRevives.Min, float64(shift.PlayerReviveCount))
		results.PlayerRevives.Max = math.Max(results.PlayerRevives.Max, float64(shift.PlayerReviveCount))
		results.PlayerRevives.Vals[i] = float64(shift.PlayerReviveCount)
		results.PlayerRevives.Sum += float64(shift.PlayerReviveCount)
		results.PlayerDeaths.Min = math.Min(results.PlayerDeaths.Min, float64(shift.PlayerDeathCount))
		results.PlayerDeaths.Max = math.Max(results.PlayerDeaths.Max, float64(shift.PlayerDeathCount))
		results.PlayerDeaths.Vals[i] = float64(shift.PlayerDeathCount)
		results.PlayerDeaths.Sum += float64(shift.PlayerDeathCount)
		results.DangerRate.Min = math.Min(results.DangerRate.Min, shift.DangerRate)
		results.DangerRate.Max = math.Max(results.DangerRate.Max, shift.DangerRate)
		results.DangerRate.Vals[i] = shift.DangerRate
		results.DangerRate.Sum += shift.DangerRate
	}
	for i := range shifts {
		results.HasShifts = true
		teamGoldenDelivered := shifts[i].Wave1GoldenDelivered
		if shifts[i].Wave2GoldenDelivered != nil {
			teamGoldenDelivered += *shifts[i].Wave2GoldenDelivered
			if shifts[i].Wave3GoldenDelivered != nil {
				teamGoldenDelivered += *shifts[i].Wave3GoldenDelivered
			}
		}
		teamPower := shifts[i].Wave1PowerEggs
		if shifts[i].Wave2PowerEggs != nil {
			teamPower += *shifts[i].Wave2PowerEggs
			if shifts[i].Wave3PowerEggs != nil {
				teamPower += *shifts[i].Wave3PowerEggs
			}
		}
		results.ShiftSummaries[i] = site_objects.ShiftInfo{
			PlayerName:   shifts[i].PlayerName,
			PlayerId:     shifts[i].UserId,
			JobId:        shifts[i].JobId,
			Stage:        shifts[i].Stage,
			IsClear:      shifts[i].IsClear,
			PlayerGolden: shifts[i].PlayerGoldenEggs,
			TeamGolden:   teamGoldenDelivered,
			TeamPower:    teamPower,
			DangerRate:   shifts[i].DangerRate,
			PlayerTitle:  shifts[i].PlayerTitle,
			GradePoint:   shifts[i].GradePoint,
			Time:         time.Unix(shifts[i].PlayTime, 0),
		}
	}
	if results.HasShifts {
		results.ClearPercent = float64(results.Clears) / float64(len(pks)) * 100
		results.WaveTwoPercent = float64(results.Wave2Clears) / float64(len(pks)) * 100
		results.WaveOnePercent = float64(results.Wave1Clears) / float64(len(pks)) * 100
		results.PlayerGoldenEggs.Mean = results.PlayerGoldenEggs.Sum / float64(len(pks))
		results.TeamGoldenEggs.Mean = results.TeamGoldenEggs.Sum / float64(len(pks))
		results.PlayerPowerEggs.Mean = results.PlayerPowerEggs.Sum / float64(len(pks))
		results.TeamPowerEggs.Mean = results.TeamPowerEggs.Sum / float64(len(pks))
		results.PlayerRevives.Mean = results.PlayerRevives.Sum / float64(len(pks))
		results.PlayerDeaths.Mean = results.PlayerDeaths.Sum / float64(len(pks))
		results.DangerRate.Mean = results.DangerRate.Sum / float64(len(pks))
		results.PlayerGoldenEggs.Median = common.CalcMedian(results.PlayerGoldenEggs.Vals)
		results.TeamGoldenEggs.Median = common.CalcMedian(results.TeamGoldenEggs.Vals)
		results.PlayerPowerEggs.Median = common.CalcMedian(results.PlayerPowerEggs.Vals)
		results.TeamPowerEggs.Median = common.CalcMedian(results.TeamPowerEggs.Vals)
		results.PlayerRevives.Median = common.CalcMedian(results.PlayerRevives.Vals)
		results.PlayerDeaths.Median = common.CalcMedian(results.PlayerDeaths.Vals)
		results.DangerRate.Median = common.CalcMedian(results.DangerRate.Vals)
	}
	if err := indexTmpl.Execute(w, results); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func Detail(w http.ResponseWriter, r *http.Request) {
	if detailTmpl == nil {
		if err := initTemplates(); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	langHeader := common.GetLanguage(r)
	dir, file := path.Split(r.RequestURI)
	userNum, err := strconv.ParseInt(path.Base(dir), 0, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	splatnetNum, err := strconv.ParseInt(file, 0, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shiftObj, err := obj_sql.GetShiftLean(userNum, splatnetNum)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var user *int64
	var userObj *db_objects.User
	var printer *message.Printer
	userCookie, err := r.Cookie("session_token")
	if err != nil {
		log.Println(err)
		printer = message.NewPrinter(language.Make(langHeader))
	} else {
		user, err = obj_sql.CheckSessionToken(userCookie.Value)
		if err != nil {
			log.Println(err)
			printer = message.NewPrinter(language.Make(langHeader))
		} else {
			userObj, err = obj_sql.ReadUserById(*user)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			printer = message.NewPrinter(language.Make((*userObj).Locale))
		}
	}
	playerCount := 1
	if shiftObj.Teammate0SplatnetId != nil {
		playerCount += 1
		if shiftObj.Teammate1SplatnetId != nil {
			playerCount += 1
			if shiftObj.Teammate2SplatnetId != nil {
				playerCount += 1
			}
		}
	}
	waveCount := 1
	if shiftObj.Wave2Quota != nil {
		waveCount += 1
		if shiftObj.Wave3Quota != nil {
			waveCount += 1
		}
	}
	shiftDetailsObj := site_objects.ShiftDetails{
		Utils: site_objects.FuncUtils{
			Printer:   printer,
			Translate: common.Translate,
			MakeSlice: common.MakeSlice,
		},
		JobId:             shiftObj.JobId,
		Stage:             shiftObj.Stage,
		Title:             shiftObj.PlayerTitle,
		GradePoint:        shiftObj.GradePoint,
		DangerRate:        shiftObj.DangerRate,
		ScheduleStartTime: time.Unix(shiftObj.ScheduleStartTime, 0),
		ScheduleEndTime:   time.Time{},
		PlayTime:          time.Unix(shiftObj.PlayTime, 0),
		Names:             make([]string, playerCount),
		Bosses: []site_objects.ShiftBoss{
			{Name: enums.Goldie, Appearances: shiftObj.GoldieCount},
			{Name: enums.Steelhead, Appearances: shiftObj.SteelheadCount},
			{Name: enums.Flyfish, Appearances: shiftObj.FlyfishCount},
			{Name: enums.Scrapper, Appearances: shiftObj.ScrapperCount},
			{Name: enums.SteelEel, Appearances: shiftObj.SteelEelCount},
			{Name: enums.Stinger, Appearances: shiftObj.StingerCount},
			{Name: enums.Maws, Appearances: shiftObj.MawsCount},
			{Name: enums.Griller, Appearances: shiftObj.GrillerCount},
			{Name: enums.Drizzler, Appearances: shiftObj.DrizzlerCount},
		},
		Players: make([]site_objects.ShiftPlayer, playerCount),
		Waves:   make([]site_objects.ShiftWave, waveCount),
	}
	if user != nil {
		shiftDetailsObj.Utils.Auth.Authenticated = true
		shiftDetailsObj.Utils.Auth.UserId = *user
		shiftDetailsObj.Utils.Auth.Username = (*userObj).Username
	}
	for bossId := range shiftDetailsObj.Bosses {
		shiftDetailsObj.Bosses[bossId].Kills = make([]int, playerCount)
	}
	playerRecorded := []bool{false, false, false, false}
	for i := range shiftDetailsObj.Players {
		if !playerRecorded[0] {
			shiftDetailsObj.Names[i] = shiftObj.PlayerName
			shiftDetailsObj.Players[i].Name = shiftObj.PlayerName
			shiftDetailsObj.Players[i].Weapons = make([]enums.SalmonWeaponEnum, waveCount)
			shiftDetailsObj.Players[i].Weapons[0] = shiftObj.PlayerW1Weapon
			if shiftObj.PlayerW2Weapon != nil {
				shiftDetailsObj.Players[i].Weapons[1] = *shiftObj.PlayerW2Weapon
				if shiftObj.PlayerW3Weapon != nil {
					shiftDetailsObj.Players[i].Weapons[2] = *shiftObj.PlayerW3Weapon
				}
			}
			shiftDetailsObj.Players[i].Special = shiftObj.PlayerSpecial
			shiftDetailsObj.Players[i].Specials = make([]int, waveCount)
			shiftDetailsObj.Players[i].Specials[0] = shiftObj.PlayerW1Specials
			if shiftObj.PlayerW2Specials != nil {
				shiftDetailsObj.Players[i].Specials[1] = *shiftObj.PlayerW2Specials
				if shiftObj.PlayerW3Specials != nil {
					shiftDetailsObj.Players[i].Specials[2] = *shiftObj.PlayerW3Specials
				}
			}
			shiftDetailsObj.Players[i].Rescues = shiftObj.PlayerReviveCount
			shiftDetailsObj.Players[i].Deaths = shiftObj.PlayerDeathCount
			shiftDetailsObj.Players[i].Golden = shiftObj.PlayerGoldenEggs
			shiftDetailsObj.Players[i].Power = shiftObj.PlayerPowerEggs
			shiftDetailsObj.Bosses[0].Kills[i] = shiftObj.PlayerGoldieKills
			shiftDetailsObj.Bosses[1].Kills[i] = shiftObj.PlayerSteelheadKills
			shiftDetailsObj.Bosses[2].Kills[i] = shiftObj.PlayerFlyfishKills
			shiftDetailsObj.Bosses[3].Kills[i] = shiftObj.PlayerScrapperKills
			shiftDetailsObj.Bosses[4].Kills[i] = shiftObj.PlayerSteelEelKills
			shiftDetailsObj.Bosses[5].Kills[i] = shiftObj.PlayerStingerKills
			shiftDetailsObj.Bosses[6].Kills[i] = shiftObj.PlayerMawsKills
			shiftDetailsObj.Bosses[7].Kills[i] = shiftObj.PlayerGrillerKills
			shiftDetailsObj.Bosses[8].Kills[i] = shiftObj.PlayerDrizzlerKills
			playerRecorded[0] = true
		} else if !playerRecorded[1] && shiftObj.Teammate0SplatnetId != nil {
			shiftDetailsObj.Names[i] = *shiftObj.Teammate0Name
			shiftDetailsObj.Players[i].Name = *shiftObj.Teammate0Name
			shiftDetailsObj.Players[i].Weapons = make([]enums.SalmonWeaponEnum, waveCount)
			if shiftObj.Teammate0W1Weapon != nil {
				shiftDetailsObj.Players[i].Weapons[0] = *shiftObj.Teammate0W1Weapon
				if shiftObj.Teammate0W2Weapon != nil {
					shiftDetailsObj.Players[i].Weapons[1] = *shiftObj.Teammate0W2Weapon
					if shiftObj.Teammate0W3Weapon != nil {
						shiftDetailsObj.Players[i].Weapons[2] = *shiftObj.Teammate0W3Weapon
					}
				}
			}
			shiftDetailsObj.Players[i].Special = *shiftObj.Teammate0Special
			shiftDetailsObj.Players[i].Specials = make([]int, waveCount)
			shiftDetailsObj.Players[i].Specials[0] = *shiftObj.Teammate0W1Specials
			if shiftObj.Teammate0W2Specials != nil {
				shiftDetailsObj.Players[i].Specials[1] = *shiftObj.Teammate0W2Specials
				if shiftObj.Teammate0W3Specials != nil {
					shiftDetailsObj.Players[i].Specials[2] = *shiftObj.Teammate0W3Specials
				}
			}
			shiftDetailsObj.Players[i].Rescues = *shiftObj.Teammate0ReviveCount
			shiftDetailsObj.Players[i].Deaths = *shiftObj.Teammate0DeathCount
			shiftDetailsObj.Players[i].Golden = *shiftObj.Teammate0GoldenEggs
			shiftDetailsObj.Players[i].Power = *shiftObj.Teammate0PowerEggs
			shiftDetailsObj.Bosses[0].Kills[i] = common.IntOrZero(shiftObj.Teammate0GoldieKills)
			shiftDetailsObj.Bosses[1].Kills[i] = common.IntOrZero(shiftObj.Teammate0SteelheadKills)
			shiftDetailsObj.Bosses[2].Kills[i] = common.IntOrZero(shiftObj.Teammate0FlyfishKills)
			shiftDetailsObj.Bosses[3].Kills[i] = common.IntOrZero(shiftObj.Teammate0ScrapperKills)
			shiftDetailsObj.Bosses[4].Kills[i] = common.IntOrZero(shiftObj.Teammate0SteelEelKills)
			shiftDetailsObj.Bosses[5].Kills[i] = common.IntOrZero(shiftObj.Teammate0StingerKills)
			shiftDetailsObj.Bosses[6].Kills[i] = common.IntOrZero(shiftObj.Teammate0MawsKills)
			shiftDetailsObj.Bosses[7].Kills[i] = common.IntOrZero(shiftObj.Teammate0GrillerKills)
			shiftDetailsObj.Bosses[8].Kills[i] = common.IntOrZero(shiftObj.Teammate0DrizzlerKills)
			playerRecorded[1] = true
		} else if !playerRecorded[2] && shiftObj.Teammate1SplatnetId != nil {
			shiftDetailsObj.Names[i] = *shiftObj.Teammate1Name
			shiftDetailsObj.Players[i].Name = *shiftObj.Teammate1Name
			shiftDetailsObj.Players[i].Weapons = make([]enums.SalmonWeaponEnum, waveCount)
			if shiftObj.Teammate1W1Weapon != nil {
				shiftDetailsObj.Players[i].Weapons[0] = *shiftObj.Teammate1W1Weapon
				if shiftObj.Teammate1W2Weapon != nil {
					shiftDetailsObj.Players[i].Weapons[1] = *shiftObj.Teammate1W2Weapon
					if shiftObj.Teammate1W3Weapon != nil {
						shiftDetailsObj.Players[i].Weapons[2] = *shiftObj.Teammate1W3Weapon
					}
				}
			}
			shiftDetailsObj.Players[i].Special = *shiftObj.Teammate1Special
			shiftDetailsObj.Players[i].Specials = make([]int, waveCount)
			shiftDetailsObj.Players[i].Specials[0] = *shiftObj.Teammate1W1Specials
			if shiftObj.Teammate1W2Specials != nil {
				shiftDetailsObj.Players[i].Specials[1] = *shiftObj.Teammate1W2Specials
				if shiftObj.Teammate1W3Specials != nil {
					shiftDetailsObj.Players[i].Specials[2] = *shiftObj.Teammate1W3Specials
				}
			}
			shiftDetailsObj.Players[i].Rescues = *shiftObj.Teammate1ReviveCount
			shiftDetailsObj.Players[i].Deaths = *shiftObj.Teammate1DeathCount
			shiftDetailsObj.Players[i].Golden = *shiftObj.Teammate1GoldenEggs
			shiftDetailsObj.Players[i].Power = *shiftObj.Teammate1PowerEggs
			shiftDetailsObj.Bosses[0].Kills[i] = common.IntOrZero(shiftObj.Teammate1GoldieKills)
			shiftDetailsObj.Bosses[1].Kills[i] = common.IntOrZero(shiftObj.Teammate1SteelheadKills)
			shiftDetailsObj.Bosses[2].Kills[i] = common.IntOrZero(shiftObj.Teammate1FlyfishKills)
			shiftDetailsObj.Bosses[3].Kills[i] = common.IntOrZero(shiftObj.Teammate1ScrapperKills)
			shiftDetailsObj.Bosses[4].Kills[i] = common.IntOrZero(shiftObj.Teammate1SteelEelKills)
			shiftDetailsObj.Bosses[5].Kills[i] = common.IntOrZero(shiftObj.Teammate1StingerKills)
			shiftDetailsObj.Bosses[6].Kills[i] = common.IntOrZero(shiftObj.Teammate1MawsKills)
			shiftDetailsObj.Bosses[7].Kills[i] = common.IntOrZero(shiftObj.Teammate1GrillerKills)
			shiftDetailsObj.Bosses[8].Kills[i] = common.IntOrZero(shiftObj.Teammate1DrizzlerKills)
			playerRecorded[2] = true
		} else if !playerRecorded[3] && shiftObj.Teammate2SplatnetId != nil {
			shiftDetailsObj.Names[i] = *shiftObj.Teammate2Name
			shiftDetailsObj.Players[i].Name = *shiftObj.Teammate2Name
			shiftDetailsObj.Players[i].Weapons = make([]enums.SalmonWeaponEnum, waveCount)
			if shiftObj.Teammate2W1Weapon != nil {
				shiftDetailsObj.Players[i].Weapons[0] = *shiftObj.Teammate2W1Weapon
				if shiftObj.Teammate2W2Weapon != nil {
					shiftDetailsObj.Players[i].Weapons[1] = *shiftObj.Teammate2W2Weapon
					if shiftObj.Teammate2W3Weapon != nil {
						shiftDetailsObj.Players[i].Weapons[2] = *shiftObj.Teammate2W3Weapon
					}
				}
			}
			shiftDetailsObj.Players[i].Special = *shiftObj.Teammate2Special
			shiftDetailsObj.Players[i].Specials = make([]int, waveCount)
			shiftDetailsObj.Players[i].Specials[0] = *shiftObj.Teammate2W1Specials
			if shiftObj.Teammate2W2Specials != nil {
				shiftDetailsObj.Players[i].Specials[1] = *shiftObj.Teammate2W2Specials
				if shiftObj.Teammate2W3Specials != nil {
					shiftDetailsObj.Players[i].Specials[2] = *shiftObj.Teammate2W3Specials
				}
			}
			shiftDetailsObj.Players[i].Rescues = *shiftObj.Teammate2ReviveCount
			shiftDetailsObj.Players[i].Deaths = *shiftObj.Teammate2DeathCount
			shiftDetailsObj.Players[i].Golden = *shiftObj.Teammate2GoldenEggs
			shiftDetailsObj.Players[i].Power = *shiftObj.Teammate2PowerEggs
			shiftDetailsObj.Bosses[0].Kills[i] = common.IntOrZero(shiftObj.Teammate2GoldieKills)
			shiftDetailsObj.Bosses[1].Kills[i] = common.IntOrZero(shiftObj.Teammate2SteelheadKills)
			shiftDetailsObj.Bosses[2].Kills[i] = common.IntOrZero(shiftObj.Teammate2FlyfishKills)
			shiftDetailsObj.Bosses[3].Kills[i] = common.IntOrZero(shiftObj.Teammate2ScrapperKills)
			shiftDetailsObj.Bosses[4].Kills[i] = common.IntOrZero(shiftObj.Teammate2SteelEelKills)
			shiftDetailsObj.Bosses[5].Kills[i] = common.IntOrZero(shiftObj.Teammate2StingerKills)
			shiftDetailsObj.Bosses[6].Kills[i] = common.IntOrZero(shiftObj.Teammate2MawsKills)
			shiftDetailsObj.Bosses[7].Kills[i] = common.IntOrZero(shiftObj.Teammate2GrillerKills)
			shiftDetailsObj.Bosses[8].Kills[i] = common.IntOrZero(shiftObj.Teammate2DrizzlerKills)
			playerRecorded[3] = true
		}
	}
	if shiftObj.IsClear {
		shiftDetailsObj.Result = printer.Sprintf("Cleared")
	} else {
		shiftDetailsObj.Result = fmt.Sprintf("%s %d - %s", printer.Sprintf("Failed on wave"), *shiftObj.FailureWave, shiftObj.JobFailureReason.GetDisplay(printer))
	}
	for i := range shiftDetailsObj.Waves {
		if i == 0 {
			shiftDetailsObj.Waves[i].Num = 1
			shiftDetailsObj.Waves[i].Quota = shiftObj.Wave1Quota
			shiftDetailsObj.Waves[i].PowerEggs = shiftObj.Wave1PowerEggs
			shiftDetailsObj.Waves[i].Delivers = shiftObj.Wave1GoldenDelivered
			shiftDetailsObj.Waves[i].Appearances = shiftObj.Wave1GoldenAppear
			shiftDetailsObj.Waves[i].WaterLevel = shiftObj.Wave1WaterLevel
			shiftDetailsObj.Waves[i].Event = shiftObj.Wave1EventType
		} else if i == 1 {
			shiftDetailsObj.Waves[i].Num = 2
			shiftDetailsObj.Waves[i].Quota = *shiftObj.Wave2Quota
			shiftDetailsObj.Waves[i].PowerEggs = *shiftObj.Wave2PowerEggs
			shiftDetailsObj.Waves[i].Delivers = *shiftObj.Wave2GoldenDelivered
			shiftDetailsObj.Waves[i].Appearances = *shiftObj.Wave2GoldenAppear
			shiftDetailsObj.Waves[i].WaterLevel = *shiftObj.Wave2WaterLevel
			shiftDetailsObj.Waves[i].Event = common.SalmonEventOrEmpty(shiftObj.Wave2EventType)
		} else if i == 2 {
			shiftDetailsObj.Waves[i].Num = 3
			shiftDetailsObj.Waves[i].Quota = *shiftObj.Wave3Quota
			shiftDetailsObj.Waves[i].PowerEggs = *shiftObj.Wave3PowerEggs
			shiftDetailsObj.Waves[i].Delivers = *shiftObj.Wave3GoldenDelivered
			shiftDetailsObj.Waves[i].Appearances = *shiftObj.Wave3GoldenAppear
			shiftDetailsObj.Waves[i].WaterLevel = *shiftObj.Wave3WaterLevel
			shiftDetailsObj.Waves[i].Event = common.SalmonEventOrEmpty(shiftObj.Wave3EventType)
		}
	}
	w.Header().Add("Content-Language", fmt.Sprint(language.Make(langHeader)))
	if err := detailTmpl.Execute(w, shiftDetailsObj); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
