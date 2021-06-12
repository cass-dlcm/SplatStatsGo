/*
 * SplatStats
 *
 * This is the API documentation for a future revision of SplatStats.
 *
 * API version: 0.4.0
 * Contact: splatstats@cass-dlcm.dev
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package api_code

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/v0-4-0/",
		Index,
	},

	Route{
		"AddBattle",
		strings.ToUpper("Post"),
		"/api/v0-4-0/two_battles/",
		AddBattle,
	},

	Route{
		"GetBattleById",
		strings.ToUpper("Get"),
		"/api/v0-4-0/api/v0-3-0/two_battles/{userId}/{splatnetNum}",
		GetBattleById,
	},

	Route{
		"GetBattles",
		strings.ToUpper("Get"),
		"/api/v0-4-0/two_battles/",
		GetBattles,
	},

	Route{
		"GetBattlesForUser",
		strings.ToUpper("Get"),
		"/api/v0-4-0/two_battles/{userId}",
		GetBattlesForUser,
	},

	Route{
		"AddShift",
		strings.ToUpper("Post"),
		"/api/v0-4-0/two_salmon/",
		AddShift,
	},

	Route{
		"GetShiftById",
		strings.ToUpper("Get"),
		"/api/v0-4-0/two_salmon/{userId}/{splatnetNum}",
		GetShiftById,
	},

	Route{
		"GetShifts",
		strings.ToUpper("Get"),
		"/api/v0-4-0/two_salmon/",
		GetShifts,
	},

	Route{
		"GetShiftsForUser",
		strings.ToUpper("Get"),
		"/api/v0-4-0/two_salmon/{userId}",
		GetShiftsForUser,
	},

	Route{
		"GetToken",
		strings.ToUpper("Get"),
		"/api/v0-4-0/api/v0-4-0/user/get_token",
		GetToken,
	},

	Route{
		"GetUserByName",
		strings.ToUpper("Get"),
		"/api/v0-4-0/api/v0-4-0/user/{username}",
		GetUserByName,
	},
}
