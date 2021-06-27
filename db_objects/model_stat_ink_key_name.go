package db_objects

type StatInkKeyName struct {
	Key  string `json:"key"`
	Name int64  `json:"name"` // Links to statInkName
}
