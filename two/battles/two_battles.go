package battles

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
var battleTmpl *template.Template

func initTemplates() error {
	var err error
	indexTmpl, err = template.ParseFiles("tmpl/base.gohtml", "tmpl/main_site.gohtml", "tmpl/two_battles/index.gohtml", "tmpl/two_battles/filter.gohtml")
	if err != nil {
		return err
	}
	battleTmpl, err = template.ParseFiles("tmpl/base.gohtml", "tmpl/main_site.gohtml", "tmpl/two_battles/battle.gohtml")
	return err
}

func Index(w http.ResponseWriter, r *http.Request) {
	langHeader := common.GetLanguage(r)
	sort := common.GetSort(r)
	battlesQueryStr, battlesQuery, columns := common.GetBattlesQuery(r)
	timeStr, timeQuery := common.GetTimeQuery(r)
	startAt, endAt := common.GetBounds(r)
	battlesQuery = append(battlesQuery, timeQuery[0], timeQuery[1])
	columns = append(columns, "timefrom", "timeto")
	battles, err := obj_sql.GetBattlesLean(battlesQuery, columns, startAt, endAt, sort)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pks, err := obj_sql.ReadKeyArrayWithCondition(battlesQuery, columns, "pk", "two_battles_battle", sort, "time")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	results := site_objects.BattleResults{
		Rules:     enums.GetRule(),
		Lobbies:   enums.GetLobby(),
		Ranks:     enums.GetRanks(),
		Weapons:   enums.GetBattleWeaponEnum(),
		Stages:    enums.GetStages(),
		Utils: 	site_objects.FuncUtils{
			MakeSlice: common.MakeSlice,
			Translate: common.Translate,
		},
		Nav: site_objects.Navigation{
			Sort: sort,
			CurrentPage: site_objects.Page{
				StartAt: startAt,
				EndAt:   endAt - 1,
			},
			Query: battlesQueryStr,
			Time:  timeStr,
		},
		HasBattles: false,
		Wins:       0,
		Kills: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		Deaths: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		Assists: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		Specials: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		Inked: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		BattleSummaries: make([]site_objects.BattleInfo, len(battles)),
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
		battle, err := obj_sql.GetBattleLeanPk(pk)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		results.Kills.Min = math.Min(results.Kills.Min, float64(battle.PlayerKills))
		results.Kills.Max = math.Max(results.Kills.Max, float64(battle.PlayerKills))
		results.Kills.Vals[i] = float64(battle.PlayerKills)
		results.Kills.Sum += float64(battle.PlayerKills)
		results.Deaths.Min = math.Min(results.Deaths.Min, float64(battle.PlayerDeaths))
		results.Deaths.Max = math.Max(results.Deaths.Max, float64(battle.PlayerDeaths))
		results.Deaths.Vals[i] = float64(battle.PlayerDeaths)
		results.Deaths.Sum += float64(battle.PlayerDeaths)
		results.Assists.Min = math.Min(results.Assists.Min, float64(battle.PlayerAssists))
		results.Assists.Max = math.Max(results.Assists.Max, float64(battle.PlayerAssists))
		results.Assists.Vals[i] = float64(battle.PlayerAssists)
		results.Assists.Sum += float64(battle.PlayerAssists)
		results.Specials.Min = math.Min(results.Specials.Min, float64(battle.PlayerSpecials))
		results.Specials.Max = math.Max(results.Specials.Max, float64(battle.PlayerSpecials))
		results.Specials.Vals[i] = float64(battle.PlayerSpecials)
		results.Specials.Sum += float64(battle.PlayerSpecials)
		results.Inked.Min = math.Min(results.Inked.Min, float64(battle.PlayerGamePaintPoint))
		results.Inked.Max = math.Max(results.Inked.Max, float64(battle.PlayerGamePaintPoint))
		results.Inked.Vals[i] = float64(battle.PlayerGamePaintPoint)
		results.Inked.Sum += float64(battle.PlayerGamePaintPoint)
		if battle.Win {
			results.Wins += 1
		}
	}
	for i := range battles {
		results.HasBattles = true
		results.BattleSummaries[i] = site_objects.BattleInfo{
			PlayerName:        battles[i].PlayerName,
			UserId:            battles[i].UserId,
			BattleNumber:      battles[i].BattleNumber,
			Rule:              battles[i].Rule.GetDisplay(printer),
			Stage:             battles[i].Stage.GetDisplay(printer),
			PlayerWeaponImage: "/static/two_battles/weapons/" + fmt.Sprint(battles[i].PlayerWeapon) + ".png",
			PlayerWeapon:      battles[i].PlayerWeapon.GetDisplay(printer),
			PlayerKills:       battles[i].PlayerKills,
			PlayerDeaths:      battles[i].PlayerDeaths,
			Time:              time.Unix(int64(battles[i].Time), 0),
		}
		if battles[i].Win {
			results.BattleSummaries[i].Result = "Win"
		} else {
			results.BattleSummaries[i].Result = "Lose"
		}
	}
	if results.HasBattles {
		results.WinRate = float64(results.Wins) / float64(len(pks)) * 100
		results.Kills.Mean = results.Kills.Sum / float64(len(pks))
		results.Deaths.Mean = results.Deaths.Sum / float64(len(pks))
		results.Assists.Mean = results.Assists.Sum / float64(len(pks))
		results.Specials.Mean = results.Specials.Sum / float64(len(pks))
		results.Inked.Mean = results.Inked.Sum / float64(len(pks))
		results.Kills.Median = common.CalcMedian(results.Kills.Vals)
		results.Deaths.Median = common.CalcMedian(results.Deaths.Vals)
		results.Assists.Median = common.CalcMedian(results.Assists.Vals)
		results.Specials.Median = common.CalcMedian(results.Specials.Vals)
		results.Inked.Median = common.CalcMedian(results.Inked.Vals)
	}
	if indexTmpl == nil {
		if err := initTemplates(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	w.Header().Add("Content-Language", fmt.Sprint(language.Make(langHeader)))
	if err := indexTmpl.Execute(w, results); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func IndexUser(w http.ResponseWriter, r *http.Request) {
	langHeader := common.GetLanguage(r)
	sort := common.GetSort(r)
	userNum, err := strconv.ParseInt(path.Base(r.URL.Path), 0, 64)
	battlesQueryStr, battlesQuery, columns := common.GetBattlesQuery(r)
	startAt, endAt := common.GetBounds(r)
	timeStr, timeQuery := common.GetTimeQuery(r)
	battlesQuery = append(battlesQuery, timeQuery[0], timeQuery[1])
	columns = append(columns, "timefrom", "timeto")
	battles, err := obj_sql.GetBattlesLeanUser(userNum, battlesQuery, columns, startAt, endAt, sort)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	battlesQuery = append(battlesQuery, userNum)
	columns = append(columns, "user_id")
	pks, err := obj_sql.ReadKeyArrayWithCondition(battlesQuery, columns, "pk", "two_battles_battle", sort, "time")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	results := site_objects.BattleResults{
		Utils: site_objects.FuncUtils{
			Translate:         common.Translate,
			MakeSlice: common.MakeSlice,
		},
		Rules:     enums.GetRule(),
		Lobbies:   enums.GetLobby(),
		Ranks:     enums.GetRanks(),
		Weapons:   enums.GetBattleWeaponEnum(),
		Stages:    enums.GetStages(),
		Nav: site_objects.Navigation{
			Sort: sort,
			CurrentPage: site_objects.Page{
				StartAt: startAt,
				EndAt:   endAt - 1,
			},
			Query: battlesQueryStr,
			Time:  timeStr,
		},
		HasBattles: false,
		Wins:       0,
		Kills: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		Deaths: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		Assists: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		Specials: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		Inked: site_objects.StatSummary{
			Min:  math.MaxFloat64,
			Max:  0.0,
			Vals: make([]float64, len(pks)),
			Sum:  0.0,
		},
		BattleSummaries: make([]site_objects.BattleInfo, len(battles)),
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
		battle, err := obj_sql.GetBattleLeanPk(pk)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		results.Kills.Min = math.Min(results.Kills.Min, float64(battle.PlayerKills))
		results.Kills.Max = math.Max(results.Kills.Max, float64(battle.PlayerKills))
		results.Kills.Vals[i] = float64(battle.PlayerKills)
		results.Kills.Sum += float64(battle.PlayerKills)
		results.Deaths.Min = math.Min(results.Deaths.Min, float64(battle.PlayerDeaths))
		results.Deaths.Max = math.Max(results.Deaths.Max, float64(battle.PlayerDeaths))
		results.Deaths.Vals[i] = float64(battle.PlayerDeaths)
		results.Deaths.Sum += float64(battle.PlayerDeaths)
		results.Assists.Min = math.Min(results.Assists.Min, float64(battle.PlayerAssists))
		results.Assists.Max = math.Max(results.Assists.Max, float64(battle.PlayerAssists))
		results.Assists.Vals[i] = float64(battle.PlayerAssists)
		results.Assists.Sum += float64(battle.PlayerAssists)
		results.Specials.Min = math.Min(results.Specials.Min, float64(battle.PlayerSpecials))
		results.Specials.Max = math.Max(results.Specials.Max, float64(battle.PlayerSpecials))
		results.Specials.Vals[i] = float64(battle.PlayerSpecials)
		results.Specials.Sum += float64(battle.PlayerSpecials)
		results.Inked.Min = math.Min(results.Inked.Min, float64(battle.PlayerGamePaintPoint))
		results.Inked.Max = math.Max(results.Inked.Max, float64(battle.PlayerGamePaintPoint))
		results.Inked.Vals[i] = float64(battle.PlayerGamePaintPoint)
		results.Inked.Sum += float64(battle.PlayerGamePaintPoint)
		if battle.Win {
			results.Wins += 1
		}
	}
	for i := range battles {
		results.HasBattles = true
		results.BattleSummaries[i] = site_objects.BattleInfo{
			PlayerName:        battles[i].PlayerName,
			UserId:            battles[i].UserId,
			BattleNumber:      battles[i].BattleNumber,
			Rule:              battles[i].Rule.GetDisplay(printer),
			Stage:             battles[i].Stage.GetDisplay(printer),
			PlayerWeaponImage: "/static/two_battles/weapons/" + fmt.Sprint(battles[i].PlayerWeapon) + ".png",
			PlayerWeapon:      battles[i].PlayerWeapon.GetDisplay(printer),
			PlayerKills:       battles[i].PlayerKills,
			PlayerDeaths:      battles[i].PlayerDeaths,
			Time:              time.Unix(int64(battles[i].Time), 0),
		}
		if battles[i].Win {
			results.BattleSummaries[i].Result = "Win"
		} else {
			results.BattleSummaries[i].Result = "Lose"
		}
	}
	if results.HasBattles {
		results.WinRate = float64(results.Wins) / float64(len(pks)) * 100
		results.Kills.Mean = results.Kills.Sum / float64(len(pks))
		results.Deaths.Mean = results.Deaths.Sum / float64(len(pks))
		results.Assists.Mean = results.Assists.Sum / float64(len(pks))
		results.Specials.Mean = results.Specials.Sum / float64(len(pks))
		results.Inked.Mean = results.Inked.Sum / float64(len(pks))
		results.Kills.Median = common.CalcMedian(results.Kills.Vals)
		results.Deaths.Median = common.CalcMedian(results.Deaths.Vals)
		results.Assists.Median = common.CalcMedian(results.Assists.Vals)
		results.Specials.Median = common.CalcMedian(results.Specials.Vals)
		results.Inked.Median = common.CalcMedian(results.Inked.Vals)
	}
	if indexTmpl == nil {
		if err := initTemplates(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	w.Header().Add("Content-Language", fmt.Sprint(language.Make(langHeader)))
	if err := indexTmpl.Execute(w, results); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func mainOrNil(in *string) string {
	if in != nil {
		return fmt.Sprintf("/static/two_battles/abilities/mains/%s.png", *in)
	}
	return ""
}

func Detail(w http.ResponseWriter, r *http.Request) {
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
	battleObj, err := obj_sql.GetBattleLean(userNum, splatnetNum)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var result string
	if (*battleObj).Win {
		result = "Win"
	} else {
		result = "Lose"
	}
	var endResult string
	if (*battleObj).Rule != enums.TurfWar && (*battleObj).ElapsedTime < 300 {
		endResult = "Knockout"
	} else {
		endResult = "Time"
	}
	elapsedTimeMinSec := fmt.Sprintf("%d:%02d", (*battleObj).ElapsedTime/60, (*battleObj).ElapsedTime%60)
	playerCount := 1
	if (*battleObj).Teammate0SplatnetId != nil {
		playerCount += 1
		if (*battleObj).Teammate1SplatnetId != nil {
			playerCount += 1
			if (*battleObj).Teammate2SplatnetId != nil {
				playerCount += 1
			}
		}
	}
	if (*battleObj).Opponent0SplatnetId != nil {
		playerCount += 1
		if (*battleObj).Opponent1SplatnetId != nil {
			playerCount += 1
			if (*battleObj).Opponent2SplatnetId != nil {
				playerCount += 1
				if (*battleObj).Opponent3SplatnetId != nil {
					playerCount += 1
				}
			}
		}
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
	battleDetailsObj := site_objects.BattleDetails{
		BattleNumber:      (*battleObj).BattleNumber,
		Rule:              (*battleObj).Rule.GetDisplay(printer),
		MatchType:         (*battleObj).MatchType.GetDisplay(printer),
		Stage:             (*battleObj).Stage.GetDisplay(printer),
		Result:            result,
		EndResult:         endResult,
		MyTeamCount:       (*battleObj).MyTeamCount,
		OtherTeamCount:    (*battleObj).OtherTeamCount,
		StartTime:         time.Unix(int64((*battleObj).Time), 0),
		EndTime:           time.Unix(int64((*battleObj).Time+(*battleObj).ElapsedTime), 0),
		ElapsedTimeMinSec: elapsedTimeMinSec,
		ElapsedTime:       (*battleObj).ElapsedTime,
		Players:           make([]site_objects.PlayerDetails, playerCount),
		Utils:			   site_objects.FuncUtils{
			Translate:         common.Translate,
			Printer:           printer,
		},
	}
	if user != nil {
		battleDetailsObj.Utils.Auth.Authenticated = true
		battleDetailsObj.Utils.Auth.UserId = *user
		battleDetailsObj.Utils.Auth.Username = (*userObj).Username
	}
	playerRecorded := []bool{false, false, false, false, false, false, false, false}
	i := 0
	for i < playerCount {
		if !playerRecorded[0] {
			battleDetailsObj.Players[i].Name = (*battleObj).PlayerName
			battleDetailsObj.Players[i].WeaponIcon = fmt.Sprintf("/static/two_battles/weapons/%s.png", (*battleObj).PlayerWeapon)
			battleDetailsObj.Players[i].Weapon = (*battleObj).PlayerWeapon.GetDisplay(printer)
			battleDetailsObj.Players[i].LevelStar = (*battleObj).PlayerLevelStar
			battleDetailsObj.Players[i].Level = (*battleObj).PlayerLevel
			if (*battleObj).PlayerRank != nil {
				battleDetailsObj.Players[i].Rank = *(*battleObj).PlayerRank
			} else {
				battleDetailsObj.Players[i].Rank = ""
			}
			battleDetailsObj.Players[i].GamePaintPoint = (*battleObj).PlayerGamePaintPoint
			battleDetailsObj.Players[i].KA = (*battleObj).PlayerKills + (*battleObj).PlayerAssists
			battleDetailsObj.Players[i].Assists = (*battleObj).PlayerAssists
			battleDetailsObj.Players[i].Specials = (*battleObj).PlayerSpecials
			battleDetailsObj.Players[i].Kills = (*battleObj).PlayerKills
			battleDetailsObj.Players[i].Deaths = (*battleObj).PlayerDeaths
			battleDetailsObj.Players[i].HeadgearMain = fmt.Sprintf("/static/two_battles/abilities/mains/%s.png", (*battleObj).PlayerHeadgearMain)
			battleDetailsObj.Players[i].ClothesMain = fmt.Sprintf("/static/two_battles/abilities/mains/%s.png", (*battleObj).PlayerClothesMain)
			battleDetailsObj.Players[i].ShoesMain = fmt.Sprintf("/static/two_battles/abilities/mains/%s.png", (*battleObj).PlayerShoesMain)
			if (*battleObj).PlayerHeadgearSub0 != nil {
				battleDetailsObj.Players[i].HeadgearSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).PlayerHeadgearSub0)
				if (*battleObj).PlayerHeadgearSub1 != nil {
					battleDetailsObj.Players[i].HeadgearSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).PlayerHeadgearSub1)
					if (*battleObj).PlayerHeadgearSub2 != nil {
						battleDetailsObj.Players[i].HeadgearSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).PlayerHeadgearSub2)
					}
				}
			}
			if (*battleObj).PlayerClothesSub0 != nil {
				battleDetailsObj.Players[i].ClothesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).PlayerClothesSub0)
				if (*battleObj).PlayerClothesSub1 != nil {
					battleDetailsObj.Players[i].ClothesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).PlayerClothesSub1)
					if (*battleObj).PlayerClothesSub2 != nil {
						battleDetailsObj.Players[i].ClothesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).PlayerClothesSub2)
					}
				}
			}
			if (*battleObj).PlayerShoesSub0 != nil {
				battleDetailsObj.Players[i].ShoesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).PlayerShoesSub0)
				if (*battleObj).PlayerShoesSub1 != nil {
					battleDetailsObj.Players[i].ShoesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).PlayerShoesSub1)
					if (*battleObj).PlayerShoesSub2 != nil {
						battleDetailsObj.Players[i].ShoesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).PlayerShoesSub2)
					}
				}
			}
			playerRecorded[0] = true
			i++
		} else if !playerRecorded[1] && (*battleObj).Teammate0SplatnetId != nil {
			battleDetailsObj.Players[i].Name = *(*battleObj).Teammate0Name
			battleDetailsObj.Players[i].WeaponIcon = fmt.Sprintf("/static/two_battles/weapons/%s.png", *(*battleObj).Teammate0Weapon)
			battleDetailsObj.Players[i].Weapon = (*(*battleObj).Teammate0Weapon).GetDisplay(printer)
			battleDetailsObj.Players[i].LevelStar = *(*battleObj).Teammate0LevelStar
			battleDetailsObj.Players[i].Level = *(*battleObj).Teammate0Level
			if (*battleObj).Teammate0Rank != nil {
				battleDetailsObj.Players[i].Rank = *(*battleObj).Teammate0Rank
			} else {
				battleDetailsObj.Players[i].Rank = ""
			}
			battleDetailsObj.Players[i].GamePaintPoint = *(*battleObj).Teammate0GamePaintPoint
			battleDetailsObj.Players[i].KA = *(*battleObj).Teammate0Kills + *(*battleObj).Teammate0Assists
			battleDetailsObj.Players[i].Assists = *(*battleObj).Teammate0Assists
			battleDetailsObj.Players[i].Specials = *(*battleObj).Teammate0Specials
			battleDetailsObj.Players[i].Kills = *(*battleObj).Teammate0Kills
			battleDetailsObj.Players[i].Deaths = *(*battleObj).Teammate0Deaths
			battleDetailsObj.Players[i].HeadgearMain = mainOrNil((*battleObj).Teammate0HeadgearMain)
			battleDetailsObj.Players[i].ClothesMain = mainOrNil((*battleObj).Teammate0ClothesMain)
			battleDetailsObj.Players[i].ShoesMain = mainOrNil((*battleObj).Teammate0ShoesMain)
			if (*battleObj).Teammate0HeadgearSub0 != nil {
				battleDetailsObj.Players[i].HeadgearSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate0HeadgearSub0)
				if (*battleObj).Teammate0HeadgearSub1 != nil {
					battleDetailsObj.Players[i].HeadgearSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate0HeadgearSub1)
					if (*battleObj).Teammate0HeadgearSub2 != nil {
						battleDetailsObj.Players[i].HeadgearSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate0HeadgearSub2)
					}
				}
			}
			if (*battleObj).Teammate0ClothesSub0 != nil {
				battleDetailsObj.Players[i].ClothesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate0ClothesSub0)
				if (*battleObj).Teammate0ClothesSub1 != nil {
					battleDetailsObj.Players[i].ClothesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate0ClothesSub1)
					if (*battleObj).Teammate0ClothesSub2 != nil {
						battleDetailsObj.Players[i].ClothesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate0ClothesSub2)
					}
				}
			}
			if (*battleObj).Teammate0ShoesSub0 != nil {
				battleDetailsObj.Players[i].ShoesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate0ShoesSub0)
				if (*battleObj).Teammate0ShoesSub1 != nil {
					battleDetailsObj.Players[i].ShoesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate0ShoesSub1)
					if (*battleObj).Teammate0ShoesSub2 != nil {
						battleDetailsObj.Players[i].ShoesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate0ShoesSub2)
					}
				}
			}
			playerRecorded[1] = true
			i++
		} else if !playerRecorded[2] && (*battleObj).Teammate1SplatnetId != nil {
			battleDetailsObj.Players[i].Name = *(*battleObj).Teammate1Name
			battleDetailsObj.Players[i].WeaponIcon = fmt.Sprintf("/static/two_battles/weapons/%s.png", *(*battleObj).Teammate1Weapon)
			battleDetailsObj.Players[i].Weapon = (*(*battleObj).Teammate1Weapon).GetDisplay(printer)
			battleDetailsObj.Players[i].LevelStar = *(*battleObj).Teammate1LevelStar
			battleDetailsObj.Players[i].Level = *(*battleObj).Teammate1Level
			if (*battleObj).Teammate1Rank != nil {
				battleDetailsObj.Players[i].Rank = *(*battleObj).Teammate1Rank
			} else {
				battleDetailsObj.Players[i].Rank = ""
			}
			battleDetailsObj.Players[i].GamePaintPoint = *(*battleObj).Teammate1GamePaintPoint
			battleDetailsObj.Players[i].KA = *(*battleObj).Teammate1Kills + *(*battleObj).Teammate1Assists
			battleDetailsObj.Players[i].Assists = *(*battleObj).Teammate1Assists
			battleDetailsObj.Players[i].Specials = *(*battleObj).Teammate1Specials
			battleDetailsObj.Players[i].Kills = *(*battleObj).Teammate1Kills
			battleDetailsObj.Players[i].Deaths = *(*battleObj).Teammate1Deaths
			battleDetailsObj.Players[i].HeadgearMain = mainOrNil((*battleObj).Teammate1HeadgearMain)
			battleDetailsObj.Players[i].ClothesMain = mainOrNil((*battleObj).Teammate1ClothesMain)
			battleDetailsObj.Players[i].ShoesMain = mainOrNil((*battleObj).Teammate1ShoesMain)
			if (*battleObj).Teammate1HeadgearSub0 != nil {
				battleDetailsObj.Players[i].HeadgearSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate1HeadgearSub0)
				if (*battleObj).Teammate1HeadgearSub1 != nil {
					battleDetailsObj.Players[i].HeadgearSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate1HeadgearSub1)
					if (*battleObj).Teammate1HeadgearSub2 != nil {
						battleDetailsObj.Players[i].HeadgearSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate1HeadgearSub2)
					}
				}
			}
			if (*battleObj).Teammate1ClothesSub0 != nil {
				battleDetailsObj.Players[i].ClothesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate1ClothesSub0)
				if (*battleObj).Teammate1ClothesSub1 != nil {
					battleDetailsObj.Players[i].ClothesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate1ClothesSub1)
					if (*battleObj).Teammate1ClothesSub2 != nil {
						battleDetailsObj.Players[i].ClothesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate1ClothesSub2)
					}
				}
			}
			if (*battleObj).Teammate1ShoesSub0 != nil {
				battleDetailsObj.Players[i].ShoesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate1ShoesSub0)
				if (*battleObj).Teammate1ShoesSub1 != nil {
					battleDetailsObj.Players[i].ShoesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate1ShoesSub1)
					if (*battleObj).Teammate1ShoesSub2 != nil {
						battleDetailsObj.Players[i].ShoesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate1ShoesSub2)
					}
				}
			}
			playerRecorded[2] = true
			i++
		} else if !playerRecorded[3] && (*battleObj).Teammate2SplatnetId != nil {
			battleDetailsObj.Players[i].Name = *(*battleObj).Teammate2Name
			battleDetailsObj.Players[i].WeaponIcon = fmt.Sprintf("/static/two_battles/weapons/%s.png", *(*battleObj).Teammate2Weapon)
			battleDetailsObj.Players[i].Weapon = (*(*battleObj).Teammate2Weapon).GetDisplay(printer)
			battleDetailsObj.Players[i].LevelStar = *(*battleObj).Teammate2LevelStar
			battleDetailsObj.Players[i].Level = *(*battleObj).Teammate2Level
			if (*battleObj).Teammate2Rank != nil {
				battleDetailsObj.Players[i].Rank = *(*battleObj).Teammate2Rank
			} else {
				battleDetailsObj.Players[i].Rank = ""
			}
			battleDetailsObj.Players[i].GamePaintPoint = *(*battleObj).Teammate2GamePaintPoint
			battleDetailsObj.Players[i].KA = *(*battleObj).Teammate2Kills + *(*battleObj).Teammate2Assists
			battleDetailsObj.Players[i].Assists = *(*battleObj).Teammate2Assists
			battleDetailsObj.Players[i].Specials = *(*battleObj).Teammate2Specials
			battleDetailsObj.Players[i].Kills = *(*battleObj).Teammate2Kills
			battleDetailsObj.Players[i].Deaths = *(*battleObj).Teammate2Deaths
			battleDetailsObj.Players[i].HeadgearMain = mainOrNil((*battleObj).Teammate2HeadgearMain)
			battleDetailsObj.Players[i].ClothesMain = mainOrNil((*battleObj).Teammate2ClothesMain)
			battleDetailsObj.Players[i].ShoesMain = mainOrNil((*battleObj).Teammate2ShoesMain)
			if (*battleObj).Teammate2HeadgearSub0 != nil {
				battleDetailsObj.Players[i].HeadgearSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate2HeadgearSub0)
				if (*battleObj).Teammate2HeadgearSub1 != nil {
					battleDetailsObj.Players[i].HeadgearSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate2HeadgearSub1)
					if (*battleObj).Teammate2HeadgearSub2 != nil {
						battleDetailsObj.Players[i].HeadgearSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate2HeadgearSub2)
					}
				}
			}
			if (*battleObj).Teammate2ClothesSub0 != nil {
				battleDetailsObj.Players[i].ClothesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate2ClothesSub0)
				if (*battleObj).Teammate2ClothesSub1 != nil {
					battleDetailsObj.Players[i].ClothesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate2ClothesSub1)
					if (*battleObj).Teammate2ClothesSub2 != nil {
						battleDetailsObj.Players[i].ClothesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate2ClothesSub2)
					}
				}
			}
			if (*battleObj).Teammate2ShoesSub0 != nil {
				battleDetailsObj.Players[i].ShoesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate2ShoesSub0)
				if (*battleObj).Teammate2ShoesSub1 != nil {
					battleDetailsObj.Players[i].ShoesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate2ShoesSub1)
					if (*battleObj).Teammate2ShoesSub2 != nil {
						battleDetailsObj.Players[i].ShoesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Teammate2ShoesSub2)
					}
				}
			}
			playerRecorded[3] = true
			i++
		} else if !playerRecorded[4] && (*battleObj).Opponent0SplatnetId != nil {
			battleDetailsObj.Players[i].Name = *(*battleObj).Opponent0Name
			battleDetailsObj.Players[i].WeaponIcon = fmt.Sprintf("/static/two_battles/weapons/%s.png", *(*battleObj).Opponent0Weapon)
			battleDetailsObj.Players[i].Weapon = (*(*battleObj).Opponent0Weapon).GetDisplay(printer)
			battleDetailsObj.Players[i].LevelStar = *(*battleObj).Opponent0LevelStar
			battleDetailsObj.Players[i].Level = *(*battleObj).Opponent0Level
			if (*battleObj).Opponent0Rank != nil {
				battleDetailsObj.Players[i].Rank = *(*battleObj).Opponent0Rank
			} else {
				battleDetailsObj.Players[i].Rank = ""
			}
			battleDetailsObj.Players[i].GamePaintPoint = *(*battleObj).Opponent0GamePaintPoint
			battleDetailsObj.Players[i].KA = *(*battleObj).Opponent0Kills + *(*battleObj).Opponent0Assists
			battleDetailsObj.Players[i].Assists = *(*battleObj).Opponent0Assists
			battleDetailsObj.Players[i].Specials = *(*battleObj).Opponent0Specials
			battleDetailsObj.Players[i].Kills = *(*battleObj).Opponent0Kills
			battleDetailsObj.Players[i].Deaths = *(*battleObj).Opponent0Deaths
			battleDetailsObj.Players[i].HeadgearMain = mainOrNil((*battleObj).Opponent0HeadgearMain)
			battleDetailsObj.Players[i].ClothesMain = mainOrNil((*battleObj).Opponent0ClothesMain)
			battleDetailsObj.Players[i].ShoesMain = mainOrNil((*battleObj).Opponent0ShoesMain)
			if (*battleObj).Opponent0HeadgearSub0 != nil {
				battleDetailsObj.Players[i].HeadgearSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent0HeadgearSub0)
				if (*battleObj).Opponent0HeadgearSub1 != nil {
					battleDetailsObj.Players[i].HeadgearSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent0HeadgearSub1)
					if (*battleObj).Opponent0HeadgearSub2 != nil {
						battleDetailsObj.Players[i].HeadgearSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent0HeadgearSub2)
					}
				}
			}
			if (*battleObj).Opponent0ClothesSub0 != nil {
				battleDetailsObj.Players[i].ClothesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent0ClothesSub0)
				if (*battleObj).Opponent0ClothesSub1 != nil {
					battleDetailsObj.Players[i].ClothesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent0ClothesSub1)
					if (*battleObj).Opponent0ClothesSub2 != nil {
						battleDetailsObj.Players[i].ClothesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent0ClothesSub2)
					}
				}
			}
			if (*battleObj).Opponent0ShoesSub0 != nil {
				battleDetailsObj.Players[i].ShoesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent0ShoesSub0)
				if (*battleObj).Opponent0ShoesSub1 != nil {
					battleDetailsObj.Players[i].ShoesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent0ShoesSub1)
					if (*battleObj).Opponent0ShoesSub2 != nil {
						battleDetailsObj.Players[i].ShoesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent0ShoesSub2)
					}
				}
			}
			playerRecorded[4] = true
			i++
		} else if !playerRecorded[5] && (*battleObj).Opponent1SplatnetId != nil {
			battleDetailsObj.Players[i].Name = *(*battleObj).Opponent1Name
			battleDetailsObj.Players[i].WeaponIcon = fmt.Sprintf("/static/two_battles/weapons/%s.png", *(*battleObj).Opponent1Weapon)
			battleDetailsObj.Players[i].Weapon = (*(*battleObj).Opponent1Weapon).GetDisplay(printer)
			battleDetailsObj.Players[i].LevelStar = *(*battleObj).Opponent1LevelStar
			battleDetailsObj.Players[i].Level = *(*battleObj).Opponent1Level
			if (*battleObj).Opponent1Rank != nil {
				battleDetailsObj.Players[i].Rank = *(*battleObj).Opponent1Rank
			} else {
				battleDetailsObj.Players[i].Rank = ""
			}
			battleDetailsObj.Players[i].GamePaintPoint = *(*battleObj).Opponent1GamePaintPoint
			battleDetailsObj.Players[i].KA = *(*battleObj).Opponent1Kills + *(*battleObj).Opponent1Assists
			battleDetailsObj.Players[i].Assists = *(*battleObj).Opponent1Assists
			battleDetailsObj.Players[i].Specials = *(*battleObj).Opponent1Specials
			battleDetailsObj.Players[i].Kills = *(*battleObj).Opponent1Kills
			battleDetailsObj.Players[i].Deaths = *(*battleObj).Opponent1Deaths
			battleDetailsObj.Players[i].HeadgearMain = mainOrNil((*battleObj).Opponent1HeadgearMain)
			battleDetailsObj.Players[i].ClothesMain = mainOrNil((*battleObj).Opponent1ClothesMain)
			battleDetailsObj.Players[i].ShoesMain = mainOrNil((*battleObj).Opponent1ShoesMain)
			if (*battleObj).Opponent1HeadgearSub0 != nil {
				battleDetailsObj.Players[i].HeadgearSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent1HeadgearSub0)
				if (*battleObj).Opponent1HeadgearSub1 != nil {
					battleDetailsObj.Players[i].HeadgearSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent1HeadgearSub1)
					if (*battleObj).Opponent1HeadgearSub2 != nil {
						battleDetailsObj.Players[i].HeadgearSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent1HeadgearSub2)
					}
				}
			}
			if (*battleObj).Opponent1ClothesSub0 != nil {
				battleDetailsObj.Players[i].ClothesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent1ClothesSub0)
				if (*battleObj).Opponent1ClothesSub1 != nil {
					battleDetailsObj.Players[i].ClothesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent1ClothesSub1)
					if (*battleObj).Opponent1ClothesSub2 != nil {
						battleDetailsObj.Players[i].ClothesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent1ClothesSub2)
					}
				}
			}
			if (*battleObj).Opponent1ShoesSub0 != nil {
				battleDetailsObj.Players[i].ShoesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent1ShoesSub0)
				if (*battleObj).Opponent1ShoesSub1 != nil {
					battleDetailsObj.Players[i].ShoesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent1ShoesSub1)
					if (*battleObj).Opponent1ShoesSub2 != nil {
						battleDetailsObj.Players[i].ShoesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent1ShoesSub2)
					}
				}
			}
			playerRecorded[5] = true
			i++
		} else if !playerRecorded[6] && (*battleObj).Opponent2SplatnetId != nil {
			battleDetailsObj.Players[i].Name = *(*battleObj).Opponent2Name
			battleDetailsObj.Players[i].WeaponIcon = fmt.Sprintf("/static/two_battles/weapons/%s.png", *(*battleObj).Opponent2Weapon)
			battleDetailsObj.Players[i].Weapon = (*(*battleObj).Opponent2Weapon).GetDisplay(printer)
			battleDetailsObj.Players[i].LevelStar = *(*battleObj).Opponent2LevelStar
			battleDetailsObj.Players[i].Level = *(*battleObj).Opponent2Level
			if (*battleObj).Opponent2Rank != nil {
				battleDetailsObj.Players[i].Rank = *(*battleObj).Opponent2Rank
			} else {
				battleDetailsObj.Players[i].Rank = ""
			}
			battleDetailsObj.Players[i].GamePaintPoint = *(*battleObj).Opponent2GamePaintPoint
			battleDetailsObj.Players[i].KA = *(*battleObj).Opponent2Kills + *(*battleObj).Opponent2Assists
			battleDetailsObj.Players[i].Assists = *(*battleObj).Opponent2Assists
			battleDetailsObj.Players[i].Specials = *(*battleObj).Opponent2Specials
			battleDetailsObj.Players[i].Kills = *(*battleObj).Opponent2Kills
			battleDetailsObj.Players[i].Deaths = *(*battleObj).Opponent2Deaths
			battleDetailsObj.Players[i].HeadgearMain = mainOrNil((*battleObj).Opponent2HeadgearMain)
			battleDetailsObj.Players[i].ClothesMain = mainOrNil((*battleObj).Opponent2ClothesMain)
			battleDetailsObj.Players[i].ShoesMain = mainOrNil((*battleObj).Opponent2ShoesMain)
			if (*battleObj).Opponent2HeadgearSub0 != nil {
				battleDetailsObj.Players[i].HeadgearSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent2HeadgearSub0)
				if (*battleObj).Opponent2HeadgearSub1 != nil {
					battleDetailsObj.Players[i].HeadgearSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent2HeadgearSub1)
					if (*battleObj).Opponent2HeadgearSub2 != nil {
						battleDetailsObj.Players[i].HeadgearSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent2HeadgearSub2)
					}
				}
			}
			if (*battleObj).Opponent2ClothesSub0 != nil {
				battleDetailsObj.Players[i].ClothesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent2ClothesSub0)
				if (*battleObj).Opponent2ClothesSub1 != nil {
					battleDetailsObj.Players[i].ClothesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent2ClothesSub1)
					if (*battleObj).Opponent2ClothesSub2 != nil {
						battleDetailsObj.Players[i].ClothesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent2ClothesSub2)
					}
				}
			}
			if (*battleObj).Opponent2ShoesSub0 != nil {
				battleDetailsObj.Players[i].ShoesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent2ShoesSub0)
				if (*battleObj).Opponent2ShoesSub1 != nil {
					battleDetailsObj.Players[i].ShoesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent2ShoesSub1)
					if (*battleObj).Opponent2ShoesSub2 != nil {
						battleDetailsObj.Players[i].ShoesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent2ShoesSub2)
					}
				}
			}
			playerRecorded[6] = true
			i++
		} else if !playerRecorded[7] && (*battleObj).Opponent3SplatnetId != nil {
			battleDetailsObj.Players[i].Name = *(*battleObj).Opponent3Name
			battleDetailsObj.Players[i].WeaponIcon = fmt.Sprintf("/static/two_battles/weapons/%s.png", *(*battleObj).Opponent3Weapon)
			battleDetailsObj.Players[i].Weapon = (*(*battleObj).Opponent3Weapon).GetDisplay(printer)
			battleDetailsObj.Players[i].LevelStar = *(*battleObj).Opponent3LevelStar
			battleDetailsObj.Players[i].Level = *(*battleObj).Opponent3Level
			if (*battleObj).Opponent3Rank != nil {
				battleDetailsObj.Players[i].Rank = *(*battleObj).Opponent3Rank
			} else {
				battleDetailsObj.Players[i].Rank = ""
			}
			battleDetailsObj.Players[i].GamePaintPoint = *(*battleObj).Opponent3GamePaintPoint
			battleDetailsObj.Players[i].KA = *(*battleObj).Opponent3Kills + *(*battleObj).Opponent3Assists
			battleDetailsObj.Players[i].Assists = *(*battleObj).Opponent3Assists
			battleDetailsObj.Players[i].Specials = *(*battleObj).Opponent3Specials
			battleDetailsObj.Players[i].Kills = *(*battleObj).Opponent3Kills
			battleDetailsObj.Players[i].Deaths = *(*battleObj).Opponent3Deaths
			battleDetailsObj.Players[i].HeadgearMain = mainOrNil((*battleObj).Opponent3HeadgearMain)
			battleDetailsObj.Players[i].ClothesMain = mainOrNil((*battleObj).Opponent3ClothesMain)
			battleDetailsObj.Players[i].ShoesMain = mainOrNil((*battleObj).Opponent3ShoesMain)
			if (*battleObj).Opponent3HeadgearSub0 != nil {
				battleDetailsObj.Players[i].HeadgearSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent3HeadgearSub0)
				if (*battleObj).Opponent3HeadgearSub1 != nil {
					battleDetailsObj.Players[i].HeadgearSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent3HeadgearSub1)
					if (*battleObj).Opponent3HeadgearSub2 != nil {
						battleDetailsObj.Players[i].HeadgearSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent3HeadgearSub2)
					}
				}
			}
			if (*battleObj).Opponent3ClothesSub0 != nil {
				battleDetailsObj.Players[i].ClothesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent3ClothesSub0)
				if (*battleObj).Opponent3ClothesSub1 != nil {
					battleDetailsObj.Players[i].ClothesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent3ClothesSub1)
					if (*battleObj).Opponent3ClothesSub2 != nil {
						battleDetailsObj.Players[i].ClothesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent3ClothesSub2)
					}
				}
			}
			if (*battleObj).Opponent3ShoesSub0 != nil {
				battleDetailsObj.Players[i].ShoesSub0 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent3ShoesSub0)
				if (*battleObj).Opponent3ShoesSub1 != nil {
					battleDetailsObj.Players[i].ShoesSub1 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent3ShoesSub1)
					if (*battleObj).Opponent3ShoesSub2 != nil {
						battleDetailsObj.Players[i].ShoesSub2 = fmt.Sprintf("/static/two_battles/abilities/subs/%s.png", *(*battleObj).Opponent3ShoesSub2)
					}
				}
			}
			playerRecorded[7] = true
			i++
		}
	}
	if battleTmpl == nil {
		if err := initTemplates(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
	w.Header().Add("Content-Language", fmt.Sprint(language.Make(langHeader)))
	if err := battleTmpl.Execute(w, battleDetailsObj); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
