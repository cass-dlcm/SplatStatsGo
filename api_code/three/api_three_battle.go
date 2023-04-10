package three

import (
	"encoding/json"
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/api_objects"
	"github.com/cass-dlcm/SplatStatsGo/obj_sql"
	"log"
	"net/http"
)

func AddBattle(w http.ResponseWriter, r *http.Request) {
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
	var battle api_objects.Battle3
	if err := json.NewDecoder(r.Body).Decode(&battle); err != nil {
		log.Println("Called by AddShift(" + fmt.Sprint(w) + ", " + fmt.Sprint(r) + ")")
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userId := *user
	originId := battle.Data.VsHistoryDetail.ID
	err = obj_sql.WriteNewBattle3(&battle, userId)
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
