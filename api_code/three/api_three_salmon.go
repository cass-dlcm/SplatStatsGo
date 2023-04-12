// Package three (of [api_code] contains HTTP handlers for Splatoon 3 SplatStats API endpoints.
package three

import (
	"encoding/json"
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/api_objects"
	"github.com/cass-dlcm/SplatStatsGo/obj_sql"
	"log"
	"net/http"
	"path"
	"strconv"
	"time"
)

func AddShift(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		log.Println("Session cookie not included")
		w.Header().Set("WWW-Authenticate", "/api/auth/signin")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	user, err := obj_sql.CheckSessionToken(sessionToken.Value)
	if err != nil {
		log.Println(err)
		w.Header().Set("WWW-Authenticate", "/api/auth/signin")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var shift api_objects.Shift3
	if err := json.NewDecoder(r.Body).Decode(&shift); err != nil {
		log.Println("Called by AddShift(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := *user
	originId := shift.Data.CoopHistoryDetail.ID
	err = obj_sql.WriteNewShift3(&shift, userId)
	if fmt.Sprint(err) == "shift already exists" {
		w.Header().Set("Location", "/api/three_salmon/"+fmt.Sprint(userId)+"/"+fmt.Sprint(originId))
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
	w.Header().Set("Location", "/api/three_salmon/"+fmt.Sprint(userId)+"/"+fmt.Sprint(originId))
	w.WriteHeader(http.StatusCreated)
}

func GetShift(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pathFront, shiftId := path.Split(r.RequestURI)
	userId, err := strconv.ParseInt(path.Base(pathFront), 0, 64)
	shift, err := obj_sql.GetShift3(userId, shiftId)
	if err != nil {
		log.Println("Called by GetShift(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	shiftMarshalled, err := json.Marshal(shift)
	if err != nil {
		log.Println("Called by GetShift(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(shiftMarshalled); err != nil {
		log.Println("Called by GetShift(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetShifts(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	userId := q.Get("userId")
	stage := q.Get("stage")
	namesOnly := q.Get("namesOnly") == "true"
	var resultWave int64
	var err error
	var timeFrom, timeTo time.Time
	if q.Get("time_from") != "" {
		timeFrom, err = time.Parse(time.RFC3339, q.Get("time_from"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		timeFrom = time.Unix(0, 0)
	}
	if q.Get("time_to") != "" {
		timeTo, err = time.Parse(time.RFC3339, q.Get("time_from"))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		timeTo = time.Now()
	}
	if q.Get("resultWave") != "" {
		resultWave, err = strconv.ParseInt(q.Get("resultWave"), 10, 32)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		resultWave = -1
	}
	if namesOnly {
		shifts, err := obj_sql.GetShiftNames3(userId, stage, int(resultWave), timeFrom, timeTo)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		shiftsMarshalled, err := json.Marshal(shifts)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(shiftsMarshalled); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	shifts, err := obj_sql.GetShiftStubs3(userId, stage, int(resultWave), timeFrom, timeTo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	shiftsMarshalled, err := json.Marshal(shifts)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(shiftsMarshalled); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
