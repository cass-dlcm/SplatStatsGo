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

type StatInkPlayer struct {
	Team string `json:"team,omitempty"`

	IsMe bool `json:"is_me,omitempty"`

	Weapon *StatInkWeapon `json:"weapon,omitempty"`

	Level int32 `json:"level,omitempty"`

	Rank *AllOfstatInkPlayerRank `json:"rank,omitempty"`

	StarRank int32 `json:"star_rank,omitempty"`

	RankInTeam int32 `json:"rank_in_team,omitempty"`

	Kill int32 `json:"kill,omitempty"`

	Death int32 `json:"death,omitempty"`

	KillOrAssist int32 `json:"kill_or_assist,omitempty"`

	Special int32 `json:"special,omitempty"`

	MyKill *interface{} `json:"my_kill,omitempty"`

	Point int32 `json:"point,omitempty"`

	Name string `json:"name,omitempty"`

	Species *StatInkKeyName `json:"species,omitempty"`

	Gender *StatInkGender `json:"gender,omitempty"`

	FestTitle *StatInkKeyName `json:"fest_title,omitempty"`

	SplatnetId string `json:"splatnet_id,omitempty"`

	Top500 bool `json:"top_500,omitempty"`

	Icon string `json:"icon,omitempty"`
}
