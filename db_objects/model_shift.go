package db_objects

import (
	"github.com/cass-dlcm/SplatStatsGo/enums"
)

type Shift struct {
	UserId                  int64                           `json:"user_id"`
	PlayerSplatnetId        string                          `json:"player_splatnet_id"`
	JobId                   int64                           `json:"job_id"`
	SplatnetJson            *int64                          `json:"splatnet_json,omitempty"` // Links to ShiftSplatnet
	StatInkJson             *int64                          `json:"stat_ink_json,omitempty"` // Links to ShiftStatInk
	StartTime               int64                           `json:"start_time"`
	PlayTime                int64                           `json:"play_time"`
	EndTime                 int64                           `json:"end_time"`
	DangerRate              float64                         `json:"danger_rate"`
	IsClear                 bool                            `json:"is_clear"`
	JobFailureReason        *enums.FailureReasonEnum        `json:"job_failure_reason,omitempty"`
	FailureWave             *int                            `json:"failure_wave,omitempty"`
	GradePoint              int                             `json:"grade_point"`
	GradePointDelta         int                             `json:"grade_point_delta"`
	JobScore                int                             `json:"job_score"`
	DrizzlerCount           int                             `json:"drizzler_count"`
	FlyfishCount            int                             `json:"flyfish_count"`
	GoldieCount             int                             `json:"goldie_count"`
	GrillerCount            int                             `json:"griller_count"`
	MawsCount               int                             `json:"maws_count"`
	ScrapperCount           int                             `json:"scrapper_count"`
	SteelEelCount           int                             `json:"steel_eel_count"`
	SteelheadCount          int                             `json:"steelhead_count"`
	StingerCount            int                             `json:"stinger_count"`
	Stage                   enums.SalmonStageEnum           `json:"stage"`
	PlayerName              string                          `json:"player_name"`
	PlayerDeathCount        int                             `json:"player_death_count"`
	PlayerReviveCount       int                             `json:"player_revive_count"`
	PlayerGoldenEggs        int                             `json:"player_golden_eggs"`
	PlayerPowerEggs         int                             `json:"player_power_eggs"`
	PlayerSpecial           enums.SalmonSpecial             `json:"player_special"`
	PlayerTitle             enums.SalmonTitle               `json:"player_title"`
	PlayerSpecies           enums.SpeciesEnum               `json:"player_species"`
	PlayerGender            enums.GenderEnum                `json:"player_gender"`
	PlayerW1Specials        int                             `json:"player_w1_specials"`
	PlayerW2Specials        *int                            `json:"player_w2_specials,omitempty"`
	PlayerW3Specials        *int                            `json:"player_w3_specials,omitempty"`
	PlayerW1Weapon          enums.SalmonWeaponEnum          `json:"player_w1_weapon"`
	PlayerW2Weapon          *enums.SalmonWeaponEnum         `json:"player_w2_weapon,omitempty"`
	PlayerW3Weapon          *enums.SalmonWeaponEnum         `json:"player_w3_weapon,omitempty"`
	PlayerDrizzlerKills     int                             `json:"player_drizzler_kills"`
	PlayerFlyfishKills      int                             `json:"player_flyfish_kills"`
	PlayerGoldieKills       int                             `json:"player_goldie_kills"`
	PlayerGrillerKills      int                             `json:"player_griller_kills"`
	PlayerMawsKills         int                             `json:"player_maws_kills"`
	PlayerScrapperKills     int                             `json:"player_scrapper_kills"`
	PlayerSteelEelKills     int                             `json:"player_steel_eel_kills"`
	PlayerSteelheadKills    int                             `json:"player_steelhead_kills"`
	PlayerStingerKills      int                             `json:"player_stinger_kills"`
	Teammate0SplatnetId     *string                         `json:"teammate_0_splatnet_id,omitempty"`
	Teammate0Name           *string                         `json:"teammate_0_name,omitempty"`
	Teammate0DeathCount     *int                            `json:"teammate_0_death_count,omitempty"`
	Teammate0ReviveCount    *int                            `json:"teammate_0_revive_count,omitempty"`
	Teammate0GoldenEggs     *int                            `json:"teammate_0_golden_eggs,omitempty"`
	Teammate0PowerEggs      *int                            `json:"teammate_0_power_eggs,omitempty"`
	Teammate0Special        *enums.SalmonSpecial            `json:"teammate_0_special,omitempty"`
	Teammate0Species        *enums.SpeciesEnum              `json:"teammate_0_species,omitempty"`
	Teammate0Gender         *enums.GenderEnum               `json:"teammate_0_gender,omitempty"`
	Teammate0W1Specials     *int                            `json:"teammate_0_w1_specials,omitempty"`
	Teammate0W2Specials     *int                            `json:"teammate_0_w2_specials,omitempty"`
	Teammate0W3Specials     *int                            `json:"teammate_0_w3_specials,omitempty"`
	Teammate0W1Weapon       *enums.SalmonWeaponEnum         `json:"teammate_0_w1_weapon,omitempty"`
	Teammate0W2Weapon       *enums.SalmonWeaponEnum         `json:"teammate_0_w2_weapon,omitempty"`
	Teammate0W3Weapon       *enums.SalmonWeaponEnum         `json:"teammate_0_w3_weapon,omitempty"`
	Teammate0DrizzlerKills  *int                            `json:"teammate_0_drizzler_kills,omitempty"`
	Teammate0FlyfishKills   *int                            `json:"teammate_0_flyfish_kills,omitempty"`
	Teammate0GoldieKills    *int                            `json:"teammate_0_goldie_kills,omitempty"`
	Teammate0GrillerKills   *int                            `json:"teammate_0_griller_kills,omitempty"`
	Teammate0MawsKills      *int                            `json:"teammate_0_maws_kills,omitempty"`
	Teammate0ScrapperKills  *int                            `json:"teammate_0_scrapper_kills,omitempty"`
	Teammate0SteelEelKills  *int                            `json:"teammate_0_steel_eel_kills,omitempty"`
	Teammate0SteelheadKills *int                            `json:"teammate_0_steelhead_kills,omitempty"`
	Teammate0StingerKills   *int                            `json:"teammate_0_stinger_kills,omitempty"`
	Teammate1SplatnetId     *string                         `json:"teammate_1_splatnet_id,omitempty"`
	Teammate1Name           *string                         `json:"teammate_1_name,omitempty"`
	Teammate1DeathCount     *int                            `json:"teammate_1_death_count,omitempty"`
	Teammate1ReviveCount    *int                            `json:"teammate_1_revive_count,omitempty"`
	Teammate1GoldenEggs     *int                            `json:"teammate_1_golden_eggs,omitempty"`
	Teammate1PowerEggs      *int                            `json:"teammate_1_power_eggs,omitempty"`
	Teammate1Special        *enums.SalmonSpecial            `json:"teammate_1_special,omitempty"`
	Teammate1Species        *enums.SpeciesEnum              `json:"teammate_1_species,omitempty"`
	Teammate1Gender         *enums.GenderEnum               `json:"teammate_1_gender,omitempty"`
	Teammate1W1Specials     *int                            `json:"teammate_1_w1_specials,omitempty"`
	Teammate1W2Specials     *int                            `json:"teammate_1_w2_specials,omitempty"`
	Teammate1W3Specials     *int                            `json:"teammate_1_w3_specials,omitempty"`
	Teammate1W1Weapon       *enums.SalmonWeaponEnum         `json:"teammate_1_w1_weapon,omitempty"`
	Teammate1W2Weapon       *enums.SalmonWeaponEnum         `json:"teammate_1_w2_weapon,omitempty"`
	Teammate1W3Weapon       *enums.SalmonWeaponEnum         `json:"teammate_1_w3_weapon,omitempty"`
	Teammate1DrizzlerKills  *int                            `json:"teammate_1_drizzler_kills,omitempty"`
	Teammate1FlyfishKills   *int                            `json:"teammate_1_flyfish_kills,omitempty"`
	Teammate1GoldieKills    *int                            `json:"teammate_1_goldie_kills,omitempty"`
	Teammate1GrillerKills   *int                            `json:"teammate_1_griller_kills,omitempty"`
	Teammate1MawsKills      *int                            `json:"teammate_1_maws_kills,omitempty"`
	Teammate1ScrapperKills  *int                            `json:"teammate_1_scrapper_kills,omitempty"`
	Teammate1SteelEelKills  *int                            `json:"teammate_1_steel_eel_kills,omitempty"`
	Teammate1SteelheadKills *int                            `json:"teammate_1_steelhead_kills,omitempty"`
	Teammate1StingerKills   *int                            `json:"teammate_1_stinger_kills,omitempty"`
	Teammate2SplatnetId     *string                         `json:"teammate_2_splatnet_id,omitempty"`
	Teammate2Name           *string                         `json:"teammate_2_name,omitempty"`
	Teammate2DeathCount     *int                            `json:"teammate_2_death_count,omitempty"`
	Teammate2ReviveCount    *int                            `json:"teammate_2_revive_count,omitempty"`
	Teammate2GoldenEggs     *int                            `json:"teammate_2_golden_eggs,omitempty"`
	Teammate2PowerEggs      *int                            `json:"teammate_2_power_eggs,omitempty"`
	Teammate2Special        *enums.SalmonSpecial            `json:"teammate_2_special,omitempty"`
	Teammate2Species        *enums.SpeciesEnum              `json:"teammate_2_species,omitempty"`
	Teammate2Gender         *enums.GenderEnum               `json:"teammate_2_gender,omitempty"`
	Teammate2W1Specials     *int                            `json:"teammate_2_w1_specials,omitempty"`
	Teammate2W2Specials     *int                            `json:"teammate_2_w2_specials,omitempty"`
	Teammate2W3Specials     *int                            `json:"teammate_2_w3_specials,omitempty"`
	Teammate2W1Weapon       *enums.SalmonWeaponEnum         `json:"teammate_2_w1_weapon,omitempty"`
	Teammate2W2Weapon       *enums.SalmonWeaponEnum         `json:"teammate_2_w2_weapon,omitempty"`
	Teammate2W3Weapon       *enums.SalmonWeaponEnum         `json:"teammate_2_w3_weapon,omitempty"`
	Teammate2DrizzlerKills  *int                            `json:"teammate_2_drizzler_kills,omitempty"`
	Teammate2FlyfishKills   *int                            `json:"teammate_2_flyfish_kills,omitempty"`
	Teammate2GoldieKills    *int                            `json:"teammate_2_goldie_kills,omitempty"`
	Teammate2GrillerKills   *int                            `json:"teammate_2_griller_kills,omitempty"`
	Teammate2MawsKills      *int                            `json:"teammate_2_maws_kills,omitempty"`
	Teammate2ScrapperKills  *int                            `json:"teammate_2_scrapper_kills,omitempty"`
	Teammate2SteelEelKills  *int                            `json:"teammate_2_steel_eel_kills,omitempty"`
	Teammate2SteelheadKills *int                            `json:"teammate_2_steelhead_kills,omitempty"`
	Teammate2StingerKills   *int                            `json:"teammate_2_stinger_kills,omitempty"`
	ScheduleEndTime         *int64                          `json:"schedule_end_time,omitempty"`
	ScheduleStartTime       *int64                          `json:"schedule_start_time,omitempty"`
	ScheduleWeapon0         *enums.SalmonWeaponScheduleEnum `json:"schedule_weapon_0,omitempty"`
	ScheduleWeapon1         *enums.SalmonWeaponScheduleEnum `json:"schedule_weapon_1,omitempty"`
	ScheduleWeapon2         *enums.SalmonWeaponScheduleEnum `json:"schedule_weapon_2,omitempty"`
	ScheduleWeapon3         *enums.SalmonWeaponScheduleEnum `json:"schedule_weapon_3,omitempty"`
	Wave1WaterLevel         enums.SalmonWaterLevel          `json:"wave_1_water_level"`
	Wave1EventType          enums.SalmonEvent               `json:"wave_1_event_type"`
	Wave1GoldenIkuraNum     int                             `json:"wave_1_golden_ikura_num"`
	Wave1GoldenIkuraPopNum  int                             `json:"wave_1_golden_ikura_pop_num"`
	Wave1IkuraNum           int                             `json:"wave_1_ikura_num"`
	Wave1QuotaNum           int                             `json:"wave_1_quota_num"`
	Wave2WaterLevel         *enums.SalmonWaterLevel         `json:"wave_2_water_level,omitempty"`
	Wave2EventType          *enums.SalmonEvent              `json:"wave_2_event_type,omitempty"`
	Wave2GoldenIkuraNum     *int                            `json:"wave_2_golden_ikura_num,omitempty"`
	Wave2GoldenIkuraPopNum  *int                            `json:"wave_2_golden_ikura_pop_num,omitempty"`
	Wave2IkuraNum           *int                            `json:"wave_2_ikura_num,omitempty"`
	Wave2QuotaNum           *int                            `json:"wave_2_quota_num,omitempty"`
	Wave3WaterLevel         *enums.SalmonWaterLevel         `json:"wave_3_water_level,omitempty"`
	Wave3EventType          *enums.SalmonEvent              `json:"wave_3_event_type,omitempty"`
	Wave3GoldenIkuraNum     *int                            `json:"wave_3_golden_ikura_num,omitempty"`
	Wave3GoldenIkuraPopNum  *int                            `json:"wave_3_golden_ikura_pop_num,omitempty"`
	Wave3IkuraNum           *int                            `json:"wave_3_ikura_num,omitempty"`
	Wave3QuotaNum           *int                            `json:"wave_3_quota_num,omitempty"`
}

type ShiftSplatnet struct {
	JobId           int64
	DangerRate      float64 `json:"danger_rate"`
	JobResult       int64   `json:"job_result"` // Links to ShiftSplatnetJobResult
	JobScore        int     `json:"job_score"`
	JobRate         int     `json:"job_rate"`
	GradePoint      int     `json:"grade_point"`
	GradePointDelta int     `json:"grade_point_delta"`
	KumaPoint       int     `json:"kuma_point"`
	StartTime       int64   `json:"start_time"`
	PlayerType      int64   `json:"player_type"` // Links to SplatnetPlayerType
	PlayTime        int64   `json:"play_time"`
	BossCounts      int64   `json:"boss_counts"` // Links to ShiftSplatnetBossCounts
	EndTime         int64   `json:"end_time"`
	MyResult        int64   `json:"my_result"` // Links to ShiftSplatnetPlayer
	Grade           int64   `json:"grade"`     // Links to ShiftSplatnetGrade
	Schedule        int64   `json:"schedule"`  // Links to ShiftSplatnetSchedule
}

type ShiftSplatnetPlayerContainer struct {
	Parent int64
	Player int64
}

type ShiftSplatnetPlayer struct {
	Special        int64  `json:"special"` // Links to SplatnetQuad
	Pid            string `json:"pid"`
	PlayerType     int64  `json:"player_type"` // Links to SplatnetPlayerType
	Name           string `json:"name"`
	DeadCount      int    `json:"dead_count"`
	GoldenIkuraNum int    `json:"golden_ikura_num"`
	BossKillCounts int64  `json:"boss_kill_counts"` // Links to ShiftSplatnetBossCounts
	IkuraNum       int    `json:"ikura_num"`
	HelpCount      int    `json:"help_count"`
}

type ShiftSplatnetPlayerWeaponList struct {
	Parent int64
	Id     string `json:"id"`
	Weapon int64  `json:"weapon"` // Links to ShiftSplatnetPlayerWeaponListWeapon
}

type ShiftSplatnetBossCounts struct {
	Var3  int64 `json:"3"`  // Links to ShiftSplatnetBossCountsBoss
	Var6  int64 `json:"6"`  // Links to ShiftSplatnetBossCountsBoss
	Var9  int64 `json:"9"`  // Links to ShiftSplatnetBossCountsBoss
	Var12 int64 `json:"12"` // Links to ShiftSplatnetBossCountsBoss
	Var13 int64 `json:"13"` // Links to ShiftSplatnetBossCountsBoss
	Var14 int64 `json:"14"` // Links to ShiftSplatnetBossCountsBoss
	Var15 int64 `json:"15"` // Links to ShiftSplatnetBossCountsBoss
	Var16 int64 `json:"16"` // Links to ShiftSplatnetBossCountsBoss
	Var21 int64 `json:"21"` // Links to ShiftSplatnetBossCountsBoss
}

type ShiftSplatnetBossCountsBoss struct {
	Boss  int64 `json:"boss"` // Links to SplatnetDouble
	Count int   `json:"count"`
}

type ShiftSplatnetWave struct {
	Parent            int64
	WaterLevel        int64 `json:"water_level"` // Links to SplatnetDouble
	EventType         int64 `json:"event_type"`  // Links to SplatnetDouble
	GoldenIkuraNum    int   `json:"golden_ikura_num"`
	GoldenIkuraPopNum int   `json:"golden_ikura_pop_num"`
	IkuraNum          int   `json:"ikura_num"`
	QuotaNum          int   `json:"quota_num"`
}

type ShiftSplatnetSchedule struct {
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
	Stage     int64 `json:"stage"` // Links to ShiftSplatnetScheduleStage
}

type ShiftSplatnetScheduleWeapon struct {
	Parent            int64
	Id                string `json:"id,omitempty"`
	Weapon            *int64 `json:"weapon,omitempty"` // Links to ShiftSplatnetScheduleWeaponWeapon
	CoopSpecialWeapon *int64 `json:"weapon,omitempty"` // Links to ShiftSplatnetScheduleWeaponSpecialWeapon
}

type ShiftStatInk struct {
	Id             int    `json:"id"`
	Uuid           string `json:"uuid"`
	SplatnetNumber int64  `json:"splatnet_number"`
	Url            string `json:"url"`
	ApiEndpoint    string `json:"api_endpoint"`
	User           int64  `json:"user"`  // Links to ShiftStatInkUser
	Stage          int64  `json:"stage"` // Links to ShiftStatInkTriple
	IsCleared      bool   `json:"is_cleared"`
	FailReason     *int64 `json:"fail_reason,omitempty"` // Links to StatInkKeyName
	ClearWaves     int    `json:"clear_waves"`
	DangerRate     string `json:"danger_rate"`
	Title          int64  `json:"title"` // Links to ShiftStatInkTitle
	TitleExp       int    `json:"title_exp"`
	TitleAfter     int64  `json:"title_after"` // Links to ShiftStatInkTitle
	TitleExpAfter  int    `json:"title_exp_after"`
	MyData         int64  `json:"my_data"` // Links to ShiftStatInkPlayer
	Agent          int64  `json:"agent"`   // Links to ShiftStatInkAgent
	Automated      bool   `json:"automated"`
	//Note           *interface{} `json:"note,omitempty"`
	//LinkUrl        *interface{} `json:"link_url,omitempty"`
	ShiftStartAt int64 `json:"shift_start_at"` // Links to StatInkTime
	StartAt      int64 `json:"start_at"`       // Links to StatInkTime
	//EndAt          *interface{} `json:"end_at,omitempty"`
	RegisterAt int64 `json:"register_at"` // Links to StatInkTime
}

type ShiftStatInkBossData struct {
	Parent      int64
	ParentTable string
	Boss        int64 `json:"boss"` // Links to ShiftStatInkBossDataBoss
	Count       int   `json:"count"`
}

type ShiftStatInkWave struct {
	Parent               int64
	KnownOccurrence      *int64 `json:"known_occurrence,omitempty"` // Links to ShiftStatInkTriple
	WaterLevel           *int64 `json:"water_level,omitempty"`      // Links to ShiftStatInkTriple
	GoldenEggQuota       int    `json:"golden_egg_quota"`
	GoldenEggAppearances int    `json:"golden_egg_appearances"`
	GoldenEggDelivered   int    `json:"golden_egg_delivered"`
	PowerEggCollected    int    `json:"power_egg_collected"`
}

type ShiftStatInkTriple struct {
	Key      string `json:"key"`
	Name     int64  `json:"name"` // Links to StatInkName
	Splatnet int    `json:"splatnet"`
}

type ShiftStatInkTripleString struct {
	Key      string `json:"key"`
	Name     int64  `json:"name"` // Links to StatInkName
	Splatnet string `json:"splatnet"`
}

type ShiftStatInkTripleContainer struct {
	Parent int64
	Triple int64
}

type ShiftStatInkPlayer struct {
	SplatnetId         string `json:"splatnet_id"`
	Name               string `json:"name"`
	Special            int64  `json:"special"` // Links to ShiftStatInkTriple
	Rescue             int    `json:"rescue"`
	Death              int    `json:"death"`
	GoldenEggDelivered int    `json:"golden_egg_delivered"`
	PowerEggCollected  int    `json:"power_egg_collected"`
	Species            int64  `json:"species"` // Links to StatInkKeyName
	Gender             int64  `json:"gender"`  // Links to StatInkGender
}

type ShiftStatInkPlayerContainer struct {
	Parent int64
	Player int64
}

type ShiftStatInkUser struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
	Url        string `json:"url"`
	SalmonUrl  string `json:"salmon_url"`
	BattleUrl  string `json:"battle_url"`
	JoinAt     int64  `json:"join_at"` // Links to StatInkTime
	Profile    int64  `json:"profile"` // Links to StatInkProfile
	Stats      int64  `json:"stats"`   // Links to ShiftStatInkUserStats
}

type ShiftStatInkUserStats struct {
	WorkCount       int   `json:"work_count"`
	TotalGoldenEggs int   `json:"total_golden_eggs"`
	TotalEggs       int   `json:"total_eggs"`
	TotalRescued    int   `json:"total_rescued"`
	TotalPoint      int   `json:"total_point"`
	AsOf            int64 `json:"as_of"`         // Links to StatInkTime
	RegisteredAt    int64 `json:"registered_at"` // Links to StatInkTime
}

type ShiftStatInkTitle struct {
	Splatnet    int   `json:"splatnet"`
	GenericName int64 `json:"generic_name"` // Links to StatInkName
}

type ShiftStatInkFailReason struct {
	Key  enums.FailureReasonEnum `json:"key"`
	Name int64                   `json:"name"` // Links to statInkName
}
