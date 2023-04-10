package routing

import (
	"github.com/cass-dlcm/SplatStatsGo/api_code"
	"github.com/cass-dlcm/SplatStatsGo/api_code/three"
	"github.com/cass-dlcm/SplatStatsGo/api_code/two"
	"github.com/cass-dlcm/SplatStatsGo/auth"
	"github.com/cass-dlcm/SplatStatsGo/two/battles"
	"github.com/cass-dlcm/SplatStatsGo/two/salmon"
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
		handler = api_code.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{
	Route{
		"AddBattle",
		strings.ToUpper("Post"),
		"/api/two_battles/",
		two.AddBattle,
	},

	Route{
		"GetBattleById",
		strings.ToUpper("Get"),
		"/api/two_battles/{userId}/{splatnetNum}",
		two.GetBattleById,
	},

	Route{
		"GetBattles",
		strings.ToUpper("Get"),
		"/api/two_battles/",
		two.GetBattles,
	},

	Route{
		"GetBattlesForUser",
		strings.ToUpper("Get"),
		"/api/two_battles/{userId}",
		two.GetBattlesForUser,
	},

	Route{
		"GetShiftsForUser",
		strings.ToUpper("Get"),
		"/api/two_salmon/{userId}",
		two.GetShiftsForUser,
	},

	Route{
		"AddShift3",
		strings.ToUpper("Post"),
		"/api/three_salmon/",
		three.AddShift,
	},

	Route{
		"GetShift3",
		strings.ToUpper("Get"),
		"/api/three_salmon/{userId}/{shiftId}",
		three.GetShift,
	},

	Route{
		"GetShifts3",
		strings.ToUpper("Get"),
		"/api/three_salmon/",
		three.GetShifts,
	},

	Route{
		"GetUserByNamePage",
		strings.ToUpper("Get"),
		"/api/auth/{username}",
		api_code.GetUserByNamePage,
	},

	Route{
		"ApiSignUp",
		strings.ToUpper("Post"),
		"/api/auth/signup",
		api_code.SignUp,
	},

	Route{
		"ApiSignIn",
		strings.ToUpper("Post"),
		"/api/auth/signin",
		api_code.Signin,
	},

	Route{
		"SignInPost",
		strings.ToUpper("Post"),
		"/auth/signin",
		auth.SigninSubmitted,
	},

	Route{
		"SignInGet",
		strings.ToUpper("Get"),
		"/auth/signin",
		auth.Signin,
	},

	Route{
		"TwoBattlesIndex",
		strings.ToUpper("Get"),
		"/two_battles",
		battles.Index,
	},

	Route{
		"TwoBattlesIndexUser",
		strings.ToUpper("Get"),
		"/two_battles/{userId}",
		battles.IndexUser,
	},

	Route{
		"TwoBattlesBattleDetail",
		strings.ToUpper("Get"),
		"/two_battles/{userId}/{splatnetNum}",
		battles.Detail,
	},

	Route{
		"TwoSalmonIndex",
		strings.ToUpper("Get"),
		"/two_salmon",
		salmon.Index,
	},

	Route{
		"TwoSalmonIndexUser",
		strings.ToUpper("Get"),
		"/two_salmon/{userId}",
		salmon.IndexUser,
	},

	Route{
		"TwoSalmonShiftDetail",
		strings.ToUpper("Get"),
		"/two_salmon/{userId}/{splatnetNum}",
		salmon.Detail,
	},
}
