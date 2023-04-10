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
	"github.com/cass-dlcm/SplatStatsGo/db_objects"
	"github.com/cass-dlcm/SplatStatsGo/obj_sql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func GetUserByNamePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

type VerificationOutput struct {
	SID              string                   `json:"sid"`
	ServiceSID       string                   `json:"service_sid"`
	AccountSID       string                   `json:"account_sid"`
	To               string                   `json:"to"`
	Channel          string                   `json:"channel"`
	Status           string                   `json:"status"`
	Valid            bool                     `json:"valid"`
	Lookup           interface{}              `json:"lookup"`
	Amount           string                   `json:"amount"`
	Payee            string                   `json:"payee"`
	SendCodeAttempts []map[string]interface{} `json:"send_code_attempts"`
	DateCreated      time.Time                `json:"date_created"`
	DateUpdated      time.Time                `json:"date_updated"`
	URL              string                   `json:"url"`
}

type TwilioError struct {
	Code     int    `json:"code"`
	Detail   string `json:"detail"`
	Message  string `json:"message"`
	MoreInfo string `json:"more_info"`
	Status   int    `json:"status"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &db_objects.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Next, insert the username, along with the hashed password into the database
	if err = obj_sql.WriteNewUser(&db_objects.User{
		Username: user.Username,
		Password: string(hashedPassword),
		Email:    user.Email,
	}); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds db_objects.User
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := obj_sql.ReadUser(creds.Username)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	uuidVal, err := uuid.NewRandom()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Create a new random session token
	sessionToken := uuidVal.String()
	// Set the token in the cache, along with the user whom it represents
	if err = obj_sql.AddSessionToken(sessionToken, creds.Username); err != nil {
		// If there is an error in setting the cache, return an internal server error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: sessionToken,
		Path:  "/api",
	})
}
