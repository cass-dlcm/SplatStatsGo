package db_objects

import "github.com/cass-dlcm/SplatStatsGo/enums"

type BattleSplatnet struct {
	Udemae             *int64 `json:"udemae,omitempty"` // Links to battleSplatnetUdemae
	Stage              int64  `json:"stage"`            // Links to splatnetTriple
	OtherTeamCount     *int   `json:"other_team_count,omitempty"`
	MyTeamCount        *int   `json:"my_team_count,omitempty"`
	StarRank           int    `json:"star_rank"`
	Rule               int64  `json:"rule"`          // Links to battleSplatnetRule
	PlayerResult       int64  `json:"player_result"` // Links to battleSplatnetPlayerResult
	EstimateGachiPower *int   `json:"estimate_gachi_power,omitempty"`
	ElapsedTime        int    `json:"elapsed_time"`
	StartTime          int    `json:"start_time"`
	GameMode           int64  `json:"game_mode"` // Links to splatnetDouble
	//	XPower              interface{}  `json:"x_power,omitempty"`
	BattleNumber string `json:"battle_number"`
	Type         string `json:"type"`
	PlayerRank   int    `json:"player_rank"`
	//CrownPlayers        interface{}  `json:"crown_players,omitempty"`
	WeaponPaintPoint int `json:"weapon_paint_point"`
	//Rank                interface{}  `json:"rank,omitempty"`
	MyTeamResult int64 `json:"my_team_result"` // Links to splatnetDouble
	//EstimateXPower      *interface{} `json:"estimate_x_power,omitempty"`
	OtherTeamResult     int64    `json:"other_team_result"` // Links to splatnetDouble
	LeaguePoint         *float64 `json:"league_point,omitempty"`
	WinMeter            *float64 `json:"win_meter,omitempty"`
	MyTeamPercentage    *float64 `json:"my_team_percentage,omitempty"`
	OtherTeamPercentage *float64 `json:"other_team_percentage,omitempty"`
	TagId               *string  `json:"tag_id,omitempty"`
}

type BattleSplatnetTeamMember struct {
	MyTeam       bool
	Parent       int64 // Links to BattleSplatnet
	PlayerResult int64 // Links to BattleSplatnetPlayerResult
}

type BattleSplatnetRule struct {
	Key           string `json:"key"`
	Name          string `json:"name"`
	MultilineName string `json:"multiline_name"`
}

type BattleSplatnetPlayerResult struct {
	DeathCount     int   `json:"death_count"`
	GamePaintPoint int   `json:"game_paint_point"`
	KillCount      int   `json:"kill_count"`
	SpecialCount   int   `json:"special_count"`
	AssistCount    int   `json:"assist_count"`
	SortScore      int   `json:"sort_score"`
	Player         int64 `json:"player"` // Links to battleSplatnetPlayerResultPlayer
}

type BattleSplatnetPlayerResultPlayer struct {
	HeadSkills    int64  `json:"head_skills"`    // Links to battleSplatnetPlayerResultPlayerSkills
	ShoesSkills   int64  `json:"shoes_skills"`   // Links to battleSplatnetPlayerResultPlayerSkills
	ClothesSkills int64  `json:"clothes_skills"` // Links to battleSplatnetPlayerResultPlayerSkills
	PlayerRank    int    `json:"player_rank"`
	StarRank      int    `json:"star_rank"`
	Nickname      string `json:"nickname"`
	PlayerType    int64  `json:"player_type"` // Links to splatnetPlayerType
	PrincipalId   string `json:"principal_id"`
	Head          int64  `json:"head"`             // Links to battleSplatnetPlayerResultPlayerClothing
	Clothes       int64  `json:"clothes"`          // Links to battleSplatnetPlayerResultPlayerClothing
	Shoes         int64  `json:"shoes"`            // Links to battleSplatnetPlayerResultPlayerClothing
	Udemae        *int64 `json:"udemae,omitempty"` // Links to BattleSplatnetUdemae
	Weapon        int64  `json:"weapon"`           // links to BattleSplatnetPlayerResultPlayerWeapon
}

type BattleSplatnetPlayerResultPlayerSkills struct {
	Main int64 `json:"main"` // links to SplatnetTriple
}

type BattleSplatnetPlayerResultPlayerSkillsSubContainer struct {
	Parent int64 `json:"parent"`
	Sub    int64 `json:"sub"` // links to SplatnetTriple
}

type BattleSplatnetPlayerResultPlayerClothing struct {
	Id        string `json:"id"`
	Image     string `json:"image"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Kind      string `json:"kind"`
	Rarity    int    `json:"rarity"`
	Brand     int64  `json:"brand"` // Links to BattleSplatnetPlayerResultPlayerClothingBrand
}

type BattleSplatnetPlayerResultPlayerClothingBrand struct {
	Id            string `json:"id"`
	Image         string `json:"image"`
	Name          string `json:"name"`
	FrequentSkill int64  `json:"frequent_skill"` // Links to splatnetTriple
}

type BattleSplatnetPlayerResultPlayerWeapon struct {
	Id        enums.BattleWeaponEnum `json:"id"`
	Image     string                 `json:"image"`
	Name      string                 `json:"name"`
	Thumbnail string                 `json:"thumbnail"`
	Sub       int64                  `json:"sub"`     // Links to splatnetQuad
	Special   int64                  `json:"special"` // Links to splatnetQuad
}
