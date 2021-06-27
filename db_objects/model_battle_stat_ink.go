package db_objects

import "github.com/cass-dlcm/SplatStatsGo/enums"

type BattleStatInk struct {
	Id             int    `json:"id"`
	SplatnetNumber int    `json:"splatnet_number"`
	Url            string `json:"url"`
	User           int64  `json:"user"`                // Links to battleStatInkUser
	Lobby          int64  `json:"lobby"`               // Links to statInkKeyName
	Mode           int64  `json:"mode"`                // Links to statInkKeyName
	Rule           int64  `json:"rule"`                // Links to statInkKeyName
	Map            int64  `json:"map"`                 // Links to battleStatInkMap
	Weapon         int64  `json:"weapon"`              // Links to BattleStatInkWeapon
	Freshness      *int64 `json:"freshness,omitempty"` // Links to BattleStatInkFreshness
	Rank           *int64 `json:"rank,omitempty"`      // Links to BattleStatInkRank
	RankExp        *int   `json:"rank_exp,omitempty"`
	RankAfter      *int64 `json:"rank_after,omitempty"` // Links to battleStatInkRank
	// XPower                         *interface{} `json:"x_power,omitempty"`
	// XPowerAfter                    *interface{} `json:"x_power_after,omitempty"`
	// EstimateXPower                 *interface{} `json:"estimate_x_power,omitempty"`
	Level        int     `json:"level"`
	LevelAfter   int     `json:"level_after"`
	StarRank     int     `json:"star_rank"`
	Result       string  `json:"result"`
	KnockOut     bool    `json:"knock_out"`
	RankInTeam   int     `json:"rank_in_team"`
	Kill         int     `json:"kill"`
	Death        int     `json:"death"`
	KillOrAssist int     `json:"kill_or_assist"`
	Special      int     `json:"special"`
	KillRatio    float64 `json:"kill_ratio"`
	KillRate     float64 `json:"kill_rate"`
	//MaxKillCombo                   *interface{} `json:"max_kill_combo,omitempty"`
	//MaxKillStreak                  *interface{} `json:"max_kill_streak,omitempty"`
	//DeathReasons                   *interface{} `json:"death_reasons,omitempty"`
	MyPoint                    int     `json:"my_point"`
	EstimateGachiPower         *int    `json:"estimate_gachi_power,omitempty"`
	LeaguePoint                *string `json:"league_point,omitempty"`
	MyTeamEstimateLeaguePoint  *int    `json:"my_team_estimate_league_point,omitempty"`
	HisTeamEstimateLeaguePoint *int    `json:"his_team_estimate_league_point,omitempty"`
	//MyTeamPoint                    *interface{} `json:"my_team_point,omitempty"`
	//HisTeamPoint                   *interface{} `json:"his_team_point,omitempty"`
	MyTeamPercent                  *string  `json:"my_team_percent,omitempty"`
	HisTeamPercent                 *string  `json:"his_team_percent,omitempty"`
	MyTeamId                       *string  `json:"my_team_id,omitempty"`
	HisTeamId                      *string  `json:"his_team_id,omitempty"`
	Species                        int64    `json:"species"`              // Links to statInkKeyName
	Gender                         int64    `json:"gender"`               // Links to statInkGender
	FestTitle                      *int64   `json:"fest_title,omitempty"` // Links to statInkKeyName
	FestExp                        *int     `json:"fest_exp,omitempty"`
	FestTitleAfter                 *int64   `json:"fest_title_after,omitempty"` // Links to statInkKeyName
	FestExpAfter                   *int     `json:"fest_exp_after,omitempty"`
	FestPower                      *string  `json:"fest_power,omitempty"`
	MyTeamEstimateFestPower        *int     `json:"my_team_estimate_fest_power,omitempty"`
	HisTeamMyTeamEstimateFestPower *int     `json:"his_team_my_team_estimate_fest_power,omitempty"`
	MyTeamFestTheme                *string  `json:"my_team_fest_theme,omitempty"`
	MyTeamNickname                 *string  `json:"my_team_nickname,omitempty"`
	HisTeamNickname                *string  `json:"his_team_nickname,omitempty"`
	Clout                          *int     `json:"clout,omitempty"`
	TotalClout                     *int     `json:"total_clout,omitempty"`
	TotalCloutAfter                *int     `json:"total_clout_after,omitempty"`
	MyTeamWinStreak                *int     `json:"my_team_win_streak,omitempty"`
	HisTeamWinStreak               *int     `json:"his_team_win_streak,omitempty"`
	SynergyBonus                   *float64 `json:"synergy_bonus,omitempty"`
	SpecialBattle                  *int64   `json:"special_battle,omitempty"` // Links to statInkKeyName
	ImageResult                    *string  `json:"image_result"`
	ImageGear                      *string  `json:"image_gear"`
	Gears                          int64    `json:"gears"` // Links to StatInkGears
	Period                         int      `json:"period"`
	PeriodRange                    string   `json:"period_range"`
	//Events                         *interface{} `json:"events,omitempty"`
	//SplatnetJson                   *interface{} `json:"splatnet_json,omitempty"`
	Agent     int64 `json:"agent"` // Links to BattleStatInkAgent
	Automated bool  `json:"automated"`
	//Environment                    *interface{} `json:"environment,omitempty"`
	LinkUrl string `json:"link_url"`
	//Note                           *interface{} `json:"note,omitempty"`
	GameVersion   string `json:"game_version"`
	NawabariBonus *int   `json:"nawabari_bonus,omitempty"`
	StartAt       int64  `json:"start_at,omitempty"`    // Links to statInkTime
	EndAt         int64  `json:"end_at,omitempty"`      // Links to statInkTime
	RegisterAt    int64  `json:"register_at,omitempty"` // Links to statInkTime
}

type BattleStatInkUser struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Url        string `json:"url"`
	JoinAt     int64  `json:"join_at"` // Links to statInkTime
	Profile    int64  `json:"profile"` // Links to statInkProfile
	//Stat       *interface{} `json:"stat,omitempty"`
	Stats int64 `json:"stats"` // links to battleStatInkUserStats
}

type BattleStatInkUserStats struct {
	// V1 *interface{} `json:"v1,omitempty"`
	V2 *int64 `json:"v2,omitempty"` // Links to battleStatInkUserStatsV2
}

type BattleStatInkUserStatsV2 struct {
	UpdatedAt int64 `json:"updated_at"` // Links to statInkTime
	Entire    int64 `json:"entire"`     // Links to BattleStatInkUserStatsV2Entire
	Nawabari  int64 `json:"nawabari"`   // Links to BattleStatInkUserStatsV2Nawabari
	Gachi     int64 `json:"gachi"`      // Links to BattleStatInkUserStatsV2Gachi
}

type BattleStatInkUserStatsV2Nawabari struct {
	Battles     int     `json:"battles"`
	WinPct      float64 `json:"win_pct"`
	KillRatio   float64 `json:"kill_ratio"`
	KillTotal   int     `json:"kill_total"`
	KillAvg     float64 `json:"kill_avg"`
	KillPerMin  float64 `json:"kill_per_min"`
	DeathTotal  int     `json:"death_total"`
	DeathAvg    float64 `json:"death_avg"`
	DeathPerMin float64 `json:"death_per_min"`
	TotalInked  int     `json:"total_inked"`
	MaxInked    int     `json:"max_inked"`
	AvgInked    float64 `json:"avg_inked"`
}

type BattleStatInkUserStatsV2Gachi struct {
	Battles     int     `json:"battles"`
	WinPct      float64 `json:"win_pct"`
	KillRatio   float64 `json:"kill_ratio"`
	KillTotal   int     `json:"kill_total"`
	KillAvg     float64 `json:"kill_avg"`
	KillPerMin  float64 `json:"kill_per_min"`
	DeathTotal  int     `json:"death_total"`
	DeathAvg    float64 `json:"death_avg"`
	DeathPerMin float64 `json:"death_per_min"`
	Rules       int64   `json:"rules,omitempty"` // Links to battleStatInkUserStatsV2GachiRules
}

type BattleStatInkUserStatsV2GachiRules struct {
	Area   int64 `json:"area"`   // Links to BattleStatInkUserStatsV2GachiRulesSub
	Yagura int64 `json:"yagura"` // Links to BattleStatInkUserStatsV2GachiRulesSub
	Hoko   int64 `json:"hoko"`   // Links to BattleStatInkUserStatsV2GachiRulesSub
	Asari  int64 `json:"asari"`  // Links to BattleStatInkUserStatsV2GachiRulesSub
}

type BattleStatInkMap struct {
	Key       string `json:"key"`
	Name      int64  `json:"name"` // Links to StatInkName
	Splatnet  int    `json:"splatnet"`
	Area      int    `json:"area"`
	ReleaseAt int64  `json:"release_at"` // Links to statInkTime
	ShortName int64  `json:"short_name"` // Links to StatInkName
}

type BattleStatInkWeapon struct {
	Key         enums.BattleStatinkWeaponEnum `json:"key"`
	Name        int64                         `json:"name"` // Links to StatInkName
	Splatnet    int                           `json:"splatnet"`
	Type        int64                         `json:"type"` // Links to BattleStatInkWeaponType
	ReskinOf    *string                       `json:"reskin_of"`
	MainRef     string                        `json:"main_ref"`
	Sub         int64                         `json:"sub"`           // Links to statInkKeyName
	Special     int64                         `json:"special"`       // Links to statInkKeyName
	MainPowerUp int64                         `json:"main_power_up"` // Links to statInkKeyName
}

type BattleStatInkWeaponType struct {
	Key      string `json:"key"`
	Name     int64  `json:"name"`     // Links to statInkName
	Category int64  `json:"category"` // Links to statInkKeyName
}

type BattleStatInkFreshness struct {
	Freshness float64 `json:"freshness"`
	Title     int64   `json:"title"` // Links to statInkName
}

type BattleStatInkRank struct {
	Key  string `json:"key"`
	Name int64  `json:"name"` // Links to statInkName
	Zone int64  `json:"zone"` // Links to statInkKeyName
}

type BattleStatInkGears struct {
	Headgear int64 `json:"headgear"` // Links to battleStatInkGearsClothes
	Clothing int64 `json:"clothing"` // Links to battleStatInkGearsClothes
	Shoes    int64 `json:"shoes"`    // Links to battleStatInkGearsClothes
}

type BattleStatInkGearsClothes struct {
	Gear           int64 `json:"gear"`            // Links to battleStatInkGearsClothesGear
	PrimaryAbility int64 `json:"primary_ability"` // Links to statInkKeyName
}

type BattleStatInkGearsClothesSecondaryAbilityContainer struct {
	Parent           int64 // Links to BattleStatInkGearsClothes
	SecondaryAbility int64 // Links to StatInkKeyName
}

type BattleStatInkGearsClothesGear struct {
	Key            string `json:"key"`
	Name           int64  `json:"name"` // Links to statInkName
	Splatnet       int    `json:"splatnet"`
	Type           int64  `json:"type"`            // Links to statInkKeyName
	Brand          int64  `json:"brand"`           // Links to statInkKeyName
	PrimaryAbility int64  `json:"primary_ability"` // Links to statInkKeyName
}

type BattleStatInkPlayer struct {
	Parent       int64
	Team         string `json:"team"`
	IsMe         bool   `json:"is_me"`
	Weapon       int64  `json:"weapon"` // Links to BattleStatInkWeapon
	Level        int    `json:"level"`
	Rank         *int64 `json:"rank,omitempty"` // Links to battleStatInkRank
	StarRank     int    `json:"star_rank"`
	RankInTeam   int    `json:"rank_in_team"`
	Kill         int    `json:"kill"`
	Death        int    `json:"death"`
	KillOrAssist int    `json:"kill_or_assist"`
	Special      int    `json:"special"`
	// MyKill       *interface{} `json:"my_kill,omitempty"`
	Point      int    `json:"point"`
	Name       string `json:"name"`
	Species    *int64 `json:"species"`              // Links to statInkKeyName
	Gender     int64  `json:"gender"`               // Links to StatInkGender
	FestTitle  *int64 `json:"fest_title,omitempty"` // Links to statInkKeyName
	SplatnetId string `json:"splatnet_id"`
	Top500     bool   `json:"top_500"`
	Icon       string `json:"icon"`
}

type BattleStatInkAgent struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	// GameVersion     *interface{} `json:"game_version,omitempty"`
	// GameVersionDate *interface{} `json:"game_version_date,omitempty"`
	// Custom          *interface{} `json:"custom,omitempty"`
	Variables *int64 `json:"variables,omitempty"` // Links to StatInkVariables
}
