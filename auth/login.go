package auth

import (
	"github.com/cass-dlcm/SplatStatsGo/obj_sql"
	"github.com/cass-dlcm/SplatStatsGo/site_objects"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"text/template"
)

var loginPage *template.Template

func initTemplates() error {
	var err error
	loginPage, err = template.ParseFiles("tmpl/base.gohtml", "tmpl/auth/login.gohtml", "tmpl/main_site.gohtml")
	return err
}

func Signin(w http.ResponseWriter, r *http.Request) {
	if loginPage == nil {
		if err := initTemplates(); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	if err := loginPage.Execute(w, site_objects.LoginAuth{
		Auth: site_objects.AuthInfo{
			Authenticated: false,
			UserId:        -1,
		},
	}); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func SigninSubmitted(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userName := r.Form.Get("username")
	password := r.Form.Get("password")

	user, err := obj_sql.ReadUser(userName)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !user.EmailVerified {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
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
	if err = obj_sql.AddSessionToken(sessionToken, userName); err != nil {
		// If there is an error in setting the cache, return an internal server error
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
	http.Redirect(w, r, "/two_battles", http.StatusFound)
}
