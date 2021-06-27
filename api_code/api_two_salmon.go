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

func AddShift(w http.ResponseWriter, r *http.Request) {
	sessionToken := r.Header.Get("session_token")
	if sessionToken == "" {
		log.Println("Session cookie not included")
		w.Header().Set("WWW-Authenticate", "/api/v0-4-0/auth/signin")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user, err := obj_sql.CheckSessionToken(sessionToken)
	if err != nil {
		log.Println("Called by AddShift(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.Header().Set("WWW-Authenticate", "/api/v0-4-0/auth/signin")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var shift api_objects.Shift
	if err := json.NewDecoder(r.Body).Decode(&shift); err != nil {
		log.Println("Called by AddShift(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shift.UserId = *user
	err = obj_sql.WriteNewShift(&shift)
	if fmt.Sprint(err) == "shift already exists" {
		w.Header().Set("Location", "/api/v0-4-0/two_salmon/"+fmt.Sprint(shift.UserId)+"/"+fmt.Sprint(shift.JobId))
		w.WriteHeader(http.StatusFound)
		return
	}
	if err != nil {
		log.Println("Called by AddShift(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", "/api/v0-4-0/two_salmon/"+fmt.Sprint(shift.UserId)+"/"+fmt.Sprint(shift.JobId))
	w.WriteHeader(http.StatusCreated)
}

func GetShiftById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	dir, file := path.Split(r.URL.Path)
	userNum, err := strconv.ParseInt(path.Base(dir), 0, 64)
	if err != nil {
		log.Println("Called by GetShiftById(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	splatnetNum, err := strconv.ParseInt(file, 0, 64)
	if err != nil {
		log.Println("Called by GetShiftById(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	shift, err := obj_sql.GetShiftLean(userNum, splatnetNum)
	if err != nil {
		log.Println("Called by GetShiftById(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	shiftMarshalled, err := json.Marshal(*shift)
	if err != nil {
		log.Println("Called by GetShiftById(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(shiftMarshalled); err != nil {
		log.Println("Called by GetShiftById(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetShifts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	_, salmonQuery, columns := common.GetSalmonQuery(r)
	_, timeQuery := common.GetTimeQuery(r)
	startAt, endAt := common.GetBounds(r)
	salmonQuery = append(salmonQuery, timeQuery[0], timeQuery[1])
	columns = append(columns, "play_timefrom", "play_timeto")
	shifts, err := obj_sql.GetShiftsLean(salmonQuery, columns, startAt, endAt, common.GetSort(r))
	if err != nil {
		log.Println("Called by GetShifts(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	shiftsMarshalled, err := json.Marshal(shifts)
	if err != nil {
		log.Println("Called by GetShifts(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(shiftsMarshalled); err != nil {
		log.Println("Called by GetShifts(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetShiftsForUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	userNum, err := strconv.ParseInt(path.Base(r.URL.Path), 0, 64)
	if err != nil {
		log.Println("Called by GetShiftsForUser(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	startAt, endAt := common.GetBounds(r)
	shifts, err := obj_sql.GetShiftsUser(userNum, startAt, endAt, common.GetSort(r))
	if err != nil {
		log.Println("Called by GetShiftsForUser(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	shiftsMarshalled, err := json.Marshal(shifts)
	if err != nil {
		log.Println("Called by GetShiftsForUser(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(shiftsMarshalled); err != nil {
		log.Println("Called by GetShiftsForUser(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
