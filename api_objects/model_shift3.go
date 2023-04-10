package api_objects

import (
	"encoding/base64"
	"fmt"
	"time"
)

type Shift3 struct {
	Data struct {
		CoopHistoryDetail struct {
			Typename              string          `json:"__typename"`
			ID                    string          `json:"id"`
			AfterGrade            AfterGrade      `json:"afterGrade"`
			MyResult              MemberResults   `json:"myResult"`
			MemberResults         []MemberResults `json:"memberResults"`
			BossResult            *BossResult     `json:"bossResult"`
			EnemyResults          []EnemyResult   `json:"enemyResults"`
			WaveResults           []Wave3         `json:"waveResults"`
			ResultWave            int             `json:"resultWave"`
			PlayedTime            time.Time       `json:"playedTime"`
			Rule                  string          `json:"rule"`
			CoopStage             CoopStage       `json:"coopStage"`
			DangerRate            float64         `json:"dangerRate"`
			ScenarioCode          *string         `json:"scenarioCode"`
			SmellMeter            int             `json:"smellMeter"`
			Weapons               []Weapon        `json:"weapons"`
			AfterGradePoint       int             `json:"afterGradePoint"`
			Scale                 *Scale          `json:"scale"`
			JobPoint              int             `json:"jobPoint"`
			JobScore              int             `json:"jobScore"`
			JobRate               float64         `json:"jobRate"`
			JobBonus              int             `json:"jobBonus"`
			NextHistoryDetail     HistoryDetail   `json:"nextHistoryDetail"`
			PreviousHistoryDetail HistoryDetail   `json:"previousHistoryDetail"`
		} `json:"coopHistoryDetail"`
	} `json:"data"`
}

func (s *Shift3) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(s.Data.CoopHistoryDetail.ID)
	if err != nil {
		return err
	}
	s.Data.CoopHistoryDetail.ID = string(id[57:])
	return nil
}

func (s *Shift3) EncodeId() {
	s.Data.CoopHistoryDetail.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("CoopHistoryDetail-u-%s:%s_%s", s.Data.CoopHistoryDetail.MyResult.Player.ID, s.Data.CoopHistoryDetail.PlayedTime.Format("20060102T150405"), s.Data.CoopHistoryDetail.ID)))
}

type CoopStage struct {
	Name  string `json:"name"`
	Image Image  `json:"image"`
	ID    string `json:"id"`
}

func (cs *CoopStage) DecodeId() error {
	stage, err := base64.StdEncoding.DecodeString(cs.ID)
	if err != nil {
		return err
	}
	cs.ID = string(stage[10:])
	return nil
}

func (cs *CoopStage) FillFromId() {
	switch cs.ID {
	case "1":
		cs.Name = "Spawning Grounds"
	case "2":
		cs.Name = "Sockeye Station"
	case "6":
		cs.Name = "Marooner's Bay"
	case "7":
		cs.Name = "Gone Fission Hydroplant"
	case "100":
		cs.Name = "Wahoo World"
	case "102":
		cs.Name = "Inkblot Art Academy"
	}
	cs.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("CoopStage-%s", cs.ID)))
}

type AfterGrade struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (ag *AfterGrade) DecodeId() error {
	afterGrade, err := base64.StdEncoding.DecodeString(ag.ID)
	if err != nil {
		return err
	}
	ag.ID = string(afterGrade[10:])
	return nil
}

func (ag *AfterGrade) FillFromId() {
	switch ag.ID {
	case "8":
		ag.Name = "Eggsecutive VP"
	case "7":
		ag.Name = "Profreshional +3"
	case "6":
		ag.Name = "Profreshional +2"
	case "5":
		ag.Name = "Profreshional +1"
	case "4":
		ag.Name = "Profreshional Part-Timer"
	case "3":
		ag.Name = "Overachiever"
	case "2":
		ag.Name = "Go-Getter"
	case "1":
		ag.Name = "Part-Timer"
	case "0":
		ag.Name = "Apprentice"
	}
	ag.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("CoopGrade-%s", ag.ID)))
}

type Scale struct {
	Gold   int `json:"gold"`
	Silver int `json:"silver"`
	Bronze int `json:"bronze"`
}

type BossResult struct {
	HasDefeatBoss bool  `json:"hasDefeatBoss"`
	Boss          Enemy `json:"boss"`
}

type Enemy struct {
	Name  string `json:"name"`
	ID    string `json:"id"`
	Image Image  `json:"image"`
}

type EnemyResult struct {
	DefeatCount     int   `json:"defeatCount"`
	TeamDefeatCount int   `json:"teamDefeatCount"`
	PopCount        int   `json:"popCount"`
	Enemy           Enemy `json:"enemy"`
}

func (e *Enemy) DecodeId() error {
	enemy, err := base64.StdEncoding.DecodeString(e.ID)
	if err != nil {
		return err
	}
	e.ID = string(enemy[10:])
	return nil
}

func (e *Enemy) FillFromId() {
	switch e.ID {
	case "4":
		e.Name = "Steelhead"
	case "5":
		e.Name = "Flyfish"
	case "6":
		e.Name = "Scrapper"
	case "7":
		e.Name = "Steel Eel"
	case "8":
		e.Name = "Stinger"
	case "9":
		e.Name = "Maws"
	case "10":
		e.Name = "Drizzler"
	case "11":
		e.Name = "Fish Stick"
	case "12":
		e.Name = "Flipper-Flopper"
	case "13":
		e.Name = "Big Shot"
	case "14":
		e.Name = "Slammin' Lid"
	case "15":
		e.Name = "Goldie"
	case "17":
		e.Name = "Griller"
	case "20":
		e.Name = "Mudmouth"
	case "23":
		e.Name = "Cohozuna"
	case "24":
		e.Name = "Horrorboros"
	}
	e.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("CoopEnemy-%s", e.ID)))
}

type Wave3 struct {
	WaveNumber       int        `json:"waveNumber"`
	WaterLevel       int        `json:"waterLevel"`
	EventWave        *EventWave `json:"eventWave"`
	DeliverNorm      int        `json:"deliverNorm"`
	GoldenPopCount   int        `json:"goldenPopCount"`
	TeamDeliverCount int        `json:"teamDeliverCount"`
	SpecialWeapons   []Special  `json:"specialWeapons"`
}

type EventWave struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (ew *EventWave) DecodeId() error {
	event, err := base64.StdEncoding.DecodeString(ew.ID)
	if err != nil {
		return err
	}
	ew.ID = string(event[14:])
	return nil
}

func (ew *EventWave) FillFromId() {
	switch ew.ID {
	case "1":
		ew.Name = "Rush"
	case "2":
		ew.Name = "Goldie Seeking"
	case "3":
		ew.Name = "The Griller"
	case "4":
		ew.Name = "The Mothership"
	case "5":
		ew.Name = "Fog"
	case "6":
		ew.Name = "Cohock Charge"
	case "7":
		ew.Name = "Giant Tornado"
	case "8":
		ew.Name = "Mudmouth Eruptions"
	}
	ew.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("CoopEventWave-%s", ew.ID)))
}

type Special struct {
	Name  string `json:"name"`
	Image Image  `json:"image"`
	ID    string `json:"ID"`
}

func (s *Special) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(s.ID)
	if err != nil {
		return err
	}
	s.ID = string(id[14:])
	return nil
}

func (s *Special) FillFromId() {
	switch s.ID {
	case "20006":
		s.Name = "Booyah Bomb"
	case "20007":
		s.Name = "Wave Breaker"
	case "20009":
		s.Name = "Killer Wail 5.1"
	case "20010":
		s.Name = "Inkjet"
	case "20012":
		s.Name = "Crab Tank"
	case "20013":
		s.Name = "Reefslider"
	case "20014":
		s.Name = "Triple Inkstrike"
	}
	s.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("SpecialWeapon-%s", s.ID)))
}

type Weapon struct {
	Name  string `json:"name"`
	Image Image  `json:"image"`
}

type MemberResults struct {
	Player struct {
		IsPlayer  string `json:"__isPlayer"`
		Byname    string `json:"byname"`
		Name      string `json:"name"`
		NameID    string `json:"nameId"`
		Nameplate struct {
			Badges     []Badge    `json:"badges"`
			Background Background `json:"background"`
		} `json:"nameplate"`
		Uniform Uniform `json:"uniform"`
		ID      string  `json:"id"`
		Species string  `json:"species"`
	} `json:"player"`
	Weapons            []Weapon      `json:"weapons"`
	SpecialWeapon      SpecialWeapon `json:"specialWeapon"`
	DefeatEnemyCount   int           `json:"defeatEnemyCount"`
	DeliverCount       int           `json:"deliverCount"`
	GoldenAssistCount  int           `json:"goldenAssistCount"`
	GoldenDeliverCount int           `json:"goldenDeliverCount"`
	RescueCount        int           `json:"rescueCount"`
	RescuedCount       int           `json:"rescuedCount"`
}

func (mr *MemberResults) DecodeId() error {
	playerId, err := base64.StdEncoding.DecodeString(mr.Player.ID)
	if err != nil {
		return err
	}
	mr.Player.ID = string(playerId[89:])
	return nil
}

func (mr *MemberResults) EncodeId(userGameId, shiftId string, playedTime time.Time) {
	mr.Player.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("CoopPlayer-u-%s:%s_%s:u-%s", userGameId, playedTime.Format("20060102T150405"), shiftId, mr.Player.ID)))
}

type SpecialWeapon struct {
	Name     string `json:"name"`
	Image    Image  `json:"image"`
	WeaponID int    `json:"weaponId"`
}

func (sw *SpecialWeapon) FillFromId() {
	switch sw.WeaponID {
	case 20006:
		sw.Name = "Booyah Bomb"
	case 20007:
		sw.Name = "Wave Breaker"
	case 20009:
		sw.Name = "Killer Wail 5.1"
	case 20010:
		sw.Name = "Inkjet"
	case 20012:
		sw.Name = "Crab Tank"
	case 20013:
		sw.Name = "Reefslider"
	case 20014:
		sw.Name = "Triple Inkstrike"
	}
}

type Uniform struct {
	Name  string `json:"name"`
	Image Image  `json:"image"`
	ID    string `json:"id"`
}

func (u *Uniform) DecodeId() error {
	uniform, err := base64.StdEncoding.DecodeString(u.ID)
	if err != nil {
		return err
	}
	u.ID = string(uniform[12:])
	return nil
}

func (u *Uniform) FillFromId() {
	switch u.ID {
	case "1":
		u.Name = "Orange Slopsuit"
	case "2":
		u.Name = "Green Slopsuit"
	case "3":
		u.Name = "Yellow Slopsuit"
	case "4":
		u.Name = "Pink Slopsuit"
	case "5":
		u.Name = "Blue Slopsuit"
	case "6":
		u.Name = "Black Slopsuit"
	case "7":
		u.Name = "White Slopsuit"
	case "8":
		u.Name = "Orange Gloopsuit"
	case "9":
		u.Name = "Black Gloopsuit"
	case "10":
		u.Name = "Yellow Gloopsuit"
	case "11":
		u.Name = "Brown Gloopsuit"
	}
	u.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("CoopUniform-%s", u.ID)))
}
