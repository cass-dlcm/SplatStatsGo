package api_code

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
	"encoding/json"
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/api_objects"
	"github.com/cass-dlcm/SplatStatsGo/common"
	"github.com/cass-dlcm/SplatStatsGo/obj_sql"
	"log"
	"net/http"
	"path"
	"strconv"
)

func AddBattle(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.Header.Get("session_token")
	if sessionToken == "" {
		log.Println("Session cookie not included")
		w.Header().Set("WWW-Authenticate", "/api/v0-4-0/auth/signin")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user, err := obj_sql.CheckSessionToken(sessionToken)
	if err != nil {
		log.Println(err)
		w.Header().Set("WWW-Authenticate", "/api/v0-4-0/auth/signin")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var battle api_objects.Battle
	if err := json.NewDecoder(r.Body).Decode(&battle); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	battle.UserId = *user
	err = obj_sql.WriteNewBattle(&battle)
	if fmt.Sprint(err) == "battle already exists" {
		w.Header().Set("Location", "/api/v0-4-0/two_battles/"+fmt.Sprint(battle.UserId)+"/"+fmt.Sprint(battle.BattleNumber))
		w.WriteHeader(http.StatusFound)
		return
	}
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", "/api/v0-4-0/two_battles/"+fmt.Sprint(battle.UserId)+"/"+fmt.Sprint(battle.BattleNumber))
	w.WriteHeader(http.StatusCreated)
}

func GetBattleById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	dir, file := path.Split(r.URL.Path)
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
	battle, err := obj_sql.GetBattleLean(userNum, splatnetNum)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	battleMarshalled, err := json.Marshal(*battle)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(battleMarshalled); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetBattles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, battlesQuery, columns := common.GetBattlesQuery(r)
	startAt, endAt := common.GetBounds(r)
	_, timeQuery := common.GetTimeQuery(r)
	battlesQuery = append(battlesQuery, timeQuery[0], timeQuery[1])
	columns = append(columns, "timefrom", "timeto")
	battles, err := obj_sql.GetBattlesLean(battlesQuery, columns, startAt, endAt, common.GetSort(r))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	battlesMarshalled, err := json.Marshal(battles)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(battlesMarshalled); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetBattlesForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	startAt, endAt := common.GetBounds(r)
	_, battlesQuery, columns := common.GetBattlesQuery(r)
	userNum, err := strconv.ParseInt(path.Base(r.URL.Path), 0, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var battles []api_objects.Battle
	if common.CheckAllFlag(r) {
		battles, err = obj_sql.GetBattlesUser(userNum, battlesQuery, columns, startAt, endAt, common.GetSort(r))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		battles, err = obj_sql.GetBattlesLeanUser(userNum, battlesQuery, columns, startAt, endAt, common.GetSort(r))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	battlesMarshalled, err := json.Marshal(battles)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(battlesMarshalled); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
