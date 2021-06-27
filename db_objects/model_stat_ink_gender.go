package db_objects

import "github.com/cass-dlcm/SplatStatsGo/enums"

type StatInkGender struct {
	Key     enums.GenderEnum `json:"key"`
	Name    int64            `json:"name"` // Links to statInkName
	Iso5218 int              `json:"iso5218"`
}
