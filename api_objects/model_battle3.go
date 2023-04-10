package api_objects

import (
	"encoding/base64"
	"fmt"
	"time"
)

type Battle3 struct {
	Data struct {
		VsHistoryDetail struct {
			Typename              string        `json:"__typename"`
			ID                    string        `json:"id"`
			VsRule                VsRule        `json:"vsRule"`
			VsMode                VsMode        `json:"vsMode"`
			Player                Player        `json:"player"`
			Judgement             string        `json:"judgement"`
			MyTeam                MyTeam        `json:"myTeam"`
			VsStage               VsStage       `json:"vsStage"`
			FestMatch             any           `json:"festMatch"`
			Knockout              string        `json:"knockout"`
			OtherTeams            []OtherTeams  `json:"otherTeams"`
			BankaraMatch          any           `json:"bankaraMatch"`
			XMatch                any           `json:"xMatch"`
			Duration              int           `json:"duration"`
			PlayedTime            time.Time     `json:"playedTime"`
			Awards                []Awards      `json:"awards"`
			LeagueMatch           any           `json:"leagueMatch"`
			NextHistoryDetail     HistoryDetail `json:"nextHistoryDetail"`
			PreviousHistoryDetail HistoryDetail `json:"previousHistoryDetail"`
		} `json:"vsHistoryDetail"`
	} `json:"data"`
}

func (b *Battle3) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(b.Data.VsHistoryDetail.ID)
	if err != nil {
		return err
	}
	b.Data.VsHistoryDetail.ID = string(id[63:])
	return nil
}

func (b *Battle3) EncodeId() {
	if b.Data.VsHistoryDetail.VsMode.Mode == "FEST" {
		b.Data.VsHistoryDetail.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("VsHistoryDetail-u-%s:REGULAR:%s_%s", b.Data.VsHistoryDetail.Player.ID, b.Data.VsHistoryDetail.PlayedTime.Format("20060102T150405"), b.Data.VsHistoryDetail.ID)))
	} else {
		b.Data.VsHistoryDetail.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("VsHistoryDetail-u-%s:%s:%s_%s", b.Data.VsHistoryDetail.Player.ID, b.Data.VsHistoryDetail.VsMode.Mode, b.Data.VsHistoryDetail.PlayedTime.Format("20060102T150405"), b.Data.VsHistoryDetail.ID)))
	}
}

type VsRule struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Rule string `json:"rule"`
}

func (vr *VsRule) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(vr.ID)
	if err != nil {
		return err
	}
	vr.ID = string(id[:7])
	return nil
}

func (vr *VsRule) EncodeId() {
	switch vr.ID {
	case "0":
		vr.Name = "Turf War"
		vr.Rule = "TURF_WAR"
	case "1":
		vr.Name = "Splat Zones"
		vr.Rule = "AREA"
	case "2":
		vr.Name = "Tower Control"
		vr.Rule = "LOFT"
	case "3":
		vr.Name = "Rainmaker"
		vr.Rule = "GOAL"
	case "4":
		vr.Name = "Clam Blitz"
		vr.Rule = "CLAM"
	}
	vr.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("VsRule-%s", vr.ID)))
}

type VsMode struct {
	Mode string `json:"mode"`
	ID   string `json:"id"`
}

func (vm *VsMode) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(vm.ID)
	if err != nil {
		return err
	}
	vm.ID = string(id[7:])
	return nil
}

func (vm *VsMode) EncodeId() {
	switch vm.ID {
	case "1":
		vm.Mode = "REGULAR"
	case "2", "51":
		vm.Mode = "BANKARA"
	case "5":
		vm.Mode = "PRIVATE"
	case "7":
		vm.Mode = "FEST"
	}
	vm.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("VsMode-%s", vm.ID)))
}

type Nameplate struct {
	Badges     []Badge    `json:"badges"`
	Background Background `json:"background"`
}
type PrimaryGearPower struct {
	Name  string `json:"name"`
	Image Image  `json:"image"`
}
type AdditionalGearPowers struct {
	Name  string `json:"name"`
	Image Image  `json:"image"`
}
type OriginalImage struct {
	URL string `json:"url"`
}
type UsualGearPower struct {
	Name        string `json:"name"`
	Desc        string `json:"desc"`
	Image       Image  `json:"image"`
	IsEmptySlot bool   `json:"isEmptySlot"`
}
type Brand struct {
	Name           string         `json:"name"`
	Image          Image          `json:"image"`
	UsualGearPower UsualGearPower `json:"usualGearPower"`
	ID             string         `json:"id"`
}

func (b *Brand) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(b.ID)
	if err != nil {
		return err
	}
	b.ID = string(id[6:])
	return nil
}

func (b *Brand) EncodeId() {
	b.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Brand-%s", b.ID)))
}

type HeadGear struct {
	Name                 string                 `json:"name"`
	Image                Image                  `json:"image"`
	IsGear               string                 `json:"__isGear"`
	PrimaryGearPower     PrimaryGearPower       `json:"primaryGearPower"`
	AdditionalGearPowers []AdditionalGearPowers `json:"additionalGearPowers"`
	OriginalImage        OriginalImage          `json:"originalImage"`
	Brand                Brand                  `json:"brand"`
}
type ClothingGear struct {
	Name                 string                 `json:"name"`
	Image                Image                  `json:"image"`
	IsGear               string                 `json:"__isGear"`
	PrimaryGearPower     PrimaryGearPower       `json:"primaryGearPower"`
	AdditionalGearPowers []AdditionalGearPowers `json:"additionalGearPowers"`
	OriginalImage        OriginalImage          `json:"originalImage"`
	Brand                Brand                  `json:"brand"`
}
type ShoesGear struct {
	Name                 string                 `json:"name"`
	Image                Image                  `json:"image"`
	IsGear               string                 `json:"__isGear"`
	PrimaryGearPower     PrimaryGearPower       `json:"primaryGearPower"`
	AdditionalGearPowers []AdditionalGearPowers `json:"additionalGearPowers"`
	OriginalImage        OriginalImage          `json:"originalImage"`
	Brand                Brand                  `json:"brand"`
}
type Player struct {
	IsPlayer     string       `json:"__isPlayer"`
	Byname       string       `json:"byname"`
	Name         string       `json:"name"`
	NameID       string       `json:"nameId"`
	Nameplate    Nameplate    `json:"nameplate"`
	ID           string       `json:"id"`
	HeadGear     HeadGear     `json:"headGear"`
	ClothingGear ClothingGear `json:"clothingGear"`
	ShoesGear    ShoesGear    `json:"shoesGear"`
	Paint        int          `json:"paint"`
}

func (p *Player) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(p.ID)
	if err != nil {
		return err
	}
	p.ID = string(id[95:])
	return nil
}

func (p *Player) EncodeId(mode string, timeStamp time.Time, battleId string) {
	if mode == "FEST" {
		p.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("VsPlayer-u-%s:REGULAR:%s_%s:u-%s", p.ID, timeStamp.Format("20060102T150405"), battleId, p.ID)))
	} else {
		p.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("VsPlayer-u-%s:%s:%s_%s:u-%s", p.ID, mode, timeStamp.Format("20060102T150405"), battleId, p.ID)))
	}
}

type Color struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
	G float64 `json:"g"`
	R float64 `json:"r"`
}
type TeamResult struct {
	PaintRatio *float64 `json:"paintRatio"`
	Score      *int     `json:"score"`
	Noroshi    any      `json:"noroshi"`
}
type MaskingImage struct {
	Width           int    `json:"width"`
	Height          int    `json:"height"`
	MaskImageURL    string `json:"maskImageUrl"`
	OverlayImageURL string `json:"overlayImageUrl"`
}
type BattleSpecialWeapon struct {
	MaskingImage MaskingImage `json:"maskingImage"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Image        Image        `json:"image"`
}
type Image3D struct {
	URL string `json:"url"`
}
type Image2D struct {
	URL string `json:"url"`
}
type Image3DThumbnail struct {
	URL string `json:"url"`
}
type Image2DThumbnail struct {
	URL string `json:"url"`
}
type SubWeapon struct {
	Name  string `json:"name"`
	Image Image  `json:"image"`
	ID    string `json:"id"`
}
type BattleWeapon struct {
	Name             string              `json:"name"`
	Image            Image               `json:"image"`
	SpecialWeapon    BattleSpecialWeapon `json:"specialWeapon"`
	ID               string              `json:"id"`
	Image3D          Image3D             `json:"image3d"`
	Image2D          Image2D             `json:"image2d"`
	Image3DThumbnail Image3DThumbnail    `json:"image3dThumbnail"`
	Image2DThumbnail Image2DThumbnail    `json:"image2dThumbnail"`
	SubWeapon        SubWeapon           `json:"subWeapon"`
}

func (bw *BattleWeapon) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(bw.ID)
	if err != nil {
		return err
	}
	bw.ID = string(id[7:])
	return nil
}

func (bw *BattleWeapon) EncodeId() {
	switch bw.ID {
	case "0":
		bw.Name = "Sploosh-o-matic"
	case "1":
		bw.Name = "Neo Sploosh-o-matic"
	case "10":
		bw.Name = "Splattershot Jr."
	case "11":
		bw.Name = "Custom Splattershot Jr."
	case "20":
		bw.Name = "Splash-o-matic"
	case "21":
		bw.Name = "Neo Splash-o-matic"
	case "30":
		bw.Name = "Aerospray MG"
	case "31":
		bw.Name = "Aerospray RG"
	case "40":
		bw.Name = "Splattershot"
	case "41":
		bw.Name = "Tentatek Splattershot"
	case "45":
		bw.Name = "Hero Shot Replica"
	case "50":
		bw.Name = ".52 Gal"
	case "60":
		bw.Name = "N-ZAP '85"
	case "61":
		bw.Name = "N-ZAP '89"
	case "70":
		bw.Name = "Splattershot Pro"
	case "71":
		bw.Name = "Forge Splattershot Pro"
	case "80":
		bw.Name = ".96 Gal"
	case "81":
		bw.Name = ".96 Gal Deco"
	case "90":
		bw.Name = "Jet Squelcher"
	case "91":
		bw.Name = "Custom Jet Squelcher"
	case "100":
		bw.Name = "Splattershot Nova"
	case "200":
		bw.Name = "Luna Blaster"
	case "201":
		bw.Name = "Luna Blaster Neo"
	case "210":
		bw.Name = "Blaster"
	case "220":
		bw.Name = "Range Blaster"
	case "230":
		bw.Name = "Clash Blaster"
	case "231":
		bw.Name = "Clash Blaster Neo"
	case "240":
		bw.Name = "Rapid Blaster"
	case "241":
		bw.Name = "Rapid Blaster Deco"
	case "250":
		bw.Name = "Rapid Blaster Pro"
	case "300":
		bw.Name = "L-3 Nozzlenose"
	case "301":
		bw.Name = "L-3 Nozzlenose D"
	case "400":
		bw.Name = "Squeezer"
	case "1000":
		bw.Name = "Carbon Roller"
	case "1001":
		bw.Name = "Carbon Roller Deco"
	case "1010":
		bw.Name = "Splat Roller"
	case "1011":
		bw.Name = "Krak-On Splat Roller"
	case "1020":
		bw.Name = "Dynamo Roller"
	case "1030":
		bw.Name = "Flingza Roller"
	case "1040":
		bw.Name = "Big Swig Roller"
	case "1100":
		bw.Name = "Inkbrush"
	case "1101":
		bw.Name = "Inkbrush Nouveau"
	case "1110":
		bw.Name = "Octobrush"
	case "2000":
		bw.Name = "Classic Squiffer"
	case "2010":
		bw.Name = "Splat Charger"
	case "2011":
		bw.Name = "Z+F Splat Charger"
	case "2020":
		bw.Name = "Splatterscope"
	case "2021":
		bw.Name = "Z+F Splatterscope"
	case "2030":
		bw.Name = "E-liter 4K"
	case "2040":
		bw.Name = "E-liter 4K Scope"
	case "2050":
		bw.Name = "Bamboozler 14 Mk I"
	case "2060":
		bw.Name = "Goo Tuber"
	case "2070":
		bw.Name = "Snipewriter 5H"
	case "3000":
		bw.Name = "Slosher"
	case "3001":
		bw.Name = "Slosher Deco"
	case "3010":
		bw.Name = "Tri-Slosher"
	case "3011":
		bw.Name = "Tri-Slosher Nouveau"
	case "3020":
		bw.Name = "Sloshing Machine"
	case "3030":
		bw.Name = "Bloblobber"
	case "3040":
		bw.Name = "Explosher"
	case "4000":
		bw.Name = "Mini Splatling"
	case "4001":
		bw.Name = "Zink Mini Splatling"
	case "4010":
		bw.Name = "Heavy Splatling"
	case "4020":
		bw.Name = "Hydra Splatling"
	case "4030":
		bw.Name = "Ballpoint Splatling"
	case "4040":
		bw.Name = "Nautilus 47"
	case "5000":
		bw.Name = "Dapple Dualies"
	case "5001":
		bw.Name = "Dapple Dualies Nouveau"
	case "5010":
		bw.Name = "Splat Dualies"
	case "5020":
		bw.Name = "Glooga Dualies"
	case "5030":
		bw.Name = "Dualie Squelchers"
	case "5040":
		bw.Name = "Dark Tetra Dualies"
	case "6000":
		bw.Name = "Splat Brella"
	case "6010":
		bw.Name = "Tenta Brella"
	case "6020":
		bw.Name = "Undercover Brella"
	case "7010":
		bw.Name = "Tri-Stringer"
	case "7020":
		bw.Name = "REEF-LUX 450"
	case "8000":
		bw.Name = "Splatana Stamper"
	case "8010":
		bw.Name = "Splatana Wiper"
	}
	switch bw.ID {
	case "0", "231", "300", "1010", "7020":
		bw.SubWeapon = SubWeapon{
			Name: "Curling Bomb",
			ID:   "U3ViV2VhcG9uLTY=",
		}
	case "1", "1011", "5000", "6010":
		bw.SubWeapon = SubWeapon{
			Name: "Squid Beakon",
			ID:   "U3ViV2VhcG9uLTg=",
		}
	case "10", "41", "200", "230", "1100", "2010", "2020", "3000", "5030":
		bw.SubWeapon = SubWeapon{
			Name: "Splat Bomb",
			ID:   "U3ViV2VhcG9uLTA=",
		}
	case "11", "2060", "5001", "8010":
		bw.SubWeapon = SubWeapon{
			Name: "Torpedo",
			ID:   "U3ViV2VhcG9uLTEz",
		}
	case "20", "301", "1001", "4000", "8000":
		bw.SubWeapon = SubWeapon{
			Name: "Burst Bomb",
			ID:   "U3ViV2VhcG9uLTI=",
		}
	case "21", "40", "45", "60", "71", "220", "1110", "5010":
		bw.SubWeapon = SubWeapon{
			Name: "Suction Bomb",
			ID:   "U3ViV2VhcG9uLTE=",
		}
	case "30", "201", "3011", "3020", "4030":
		bw.SubWeapon = SubWeapon{
			Name: "Fizzy Bomb",
			ID:   "U3ViV2VhcG9uLTU=",
		}
	case "31", "80", "1020", "2070", "3030", "4010", "6000":
		bw.SubWeapon = SubWeapon{
			Name: "Sprinkler",
			ID:   "U3ViV2VhcG9uLTM=",
		}
	case "50", "400", "1040", "2011", "2021", "5020":
		bw.SubWeapon = SubWeapon{
			Name: "Splash Wall",
			ID:   "U3ViV2VhcG9uLTQ=",
		}
	case "61", "210", "1000", "2050", "4020", "5040":
		bw.SubWeapon = SubWeapon{
			Name: "Autobomb",
			ID:   "U3ViV2VhcG9uLTc=",
		}
	case "70", "90", "3001":
		bw.SubWeapon = SubWeapon{
			Name: "Angle Shooter",
			ID:   "U3ViV2VhcG9uLTEy",
		}
	case "91", "250", "3010", "4001", "7010":
		bw.SubWeapon = SubWeapon{
			Name: "Toxic Mist",
			ID:   "U3ViV2VhcG9uLTEx",
		}
	case "100", "2000", "3040", "4040":
		bw.SubWeapon = SubWeapon{
			Name: "Point Sensor",
			ID:   "U3ViV2VhcG9uLTk=",
		}
	case "240", "1030", "1101", "2030", "2040", "6020":
		bw.SubWeapon = SubWeapon{
			Name: "Ink Mine",
			ID:   "U3ViV2VhcG9uLTEw",
		}
	}
	switch bw.ID {
	case "0", "201", "301", "1101", "4000", "8010":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0xMQ==",
			Name: "Ultra Stamp",
		}
	case "1", "50", "100", "1100", "2050", "7010":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi05",
			Name: "Killer Wail 5.1",
		}
	case "10", "210", "1010", "2000", "4001":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0y",
			Name: "Big Bubbler",
		}
	case "11", "220", "2030", "2040", "4010", "5030":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi03",
			Name: "Wave Breaker",
		}
	case "20", "70", "300", "5010":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0xMg==",
			Name: "Crab Tank",
		}
	case "21", "41", "240", "2011", "2021", "3000", "6000":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0xNA==",
			Name: "Triple Inkstrike",
		}
	case "30", "5001", "5040", "6020":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0xMw==",
			Name: "Reefslider",
		}
	case "31", "71", "3020", "4020", "5020":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi02",
			Name: "Booyah Bomb",
		}
	case "40", "45", "230", "400", "1001":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0x",
			Name: "Trizooka",
		}
	case "60", "1020", "2070", "3011", "5000":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0xNQ==",
			Name: "Tacticooler",
		}
	case "61", "231":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0xNg==",
			Name: "Super Chump",
		}
	case "80", "90", "250", "1040", "2010", "2020", "6010":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi04",
			Name: "Ink Vac",
		}
	case "81", "1011":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0xNw==",
			Name: "Kraken Royale",
		}
	case "91", "3030", "3040", "4040":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi01",
			Name: "Ink Storm",
		}
	case "200", "1000", "1110", "3001", "8000":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0z",
			Name: "Zipcaster",
		}
	case "241", "3010", "4030":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi0xMA==",
			Name: "Inkjet",
		}
	case "1030", "2060", "7020":
		bw.SpecialWeapon = BattleSpecialWeapon{
			ID:   "U3BlY2lhbFdlYXBvbi00",
			Name: "Tenta Missiles",
		}
	}
	bw.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("Weapon-%s", bw.ID)))
}

type ThumbnailImage struct {
	URL string `json:"url"`
}
type PlayerResult struct {
	Kill       int `json:"kill"`
	Death      int `json:"death"`
	Assist     int `json:"assist"`
	Special    int `json:"special"`
	NoroshiTry any `json:"noroshiTry"`
}
type Players struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	IsMyself       bool         `json:"isMyself"`
	Byname         string       `json:"byname"`
	Weapon         BattleWeapon `json:"weapon"`
	Species        string       `json:"species"`
	IsPlayer       string       `json:"__isPlayer"`
	NameID         string       `json:"nameId"`
	Nameplate      Nameplate    `json:"nameplate"`
	HeadGear       HeadGear     `json:"headGear"`
	ClothingGear   ClothingGear `json:"clothingGear"`
	ShoesGear      ShoesGear    `json:"shoesGear"`
	Paint          int          `json:"paint"`
	Result         PlayerResult `json:"result"`
	Crown          bool         `json:"crown"`
	FestDragonCert string       `json:"festDragonCert"`
}

func (p *Players) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(p.ID)
	if err != nil {
		return err
	}
	p.ID = string(id[95:])
	return nil
}

func (p *Players) EncodeId(userPlayerId, mode string, timeStamp time.Time, battleId string) {
	if mode == "FEST" {
		p.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("VsPlayer-u-%s:REGULAR:%s_%s:u-%s", userPlayerId, timeStamp.Format("20060102T150405"), battleId, p.ID)))
	} else {
		p.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("VsPlayer-u-%s:%s:%s_%s:u-%s", userPlayerId, mode, timeStamp.Format("20060102T150405"), battleId, p.ID)))
	}
}

type MyTeam struct {
	Color                Color      `json:"color"`
	Result               TeamResult `json:"result"`
	TricolorRole         any        `json:"tricolorRole"`
	FestTeamName         any        `json:"festTeamName"`
	FestUniformBonusRate any        `json:"festUniformBonusRate"`
	Judgement            string     `json:"judgement"`
	Players              []Players  `json:"players"`
	Order                int        `json:"order"`
	FestStreakWinCount   any        `json:"festStreakWinCount"`
	FestUniformName      any        `json:"festUniformName"`
}
type VsStage struct {
	Name  string `json:"name"`
	Image Image  `json:"image"`
	ID    string `json:"id"`
}

func (vs *VsStage) DecodeId() error {
	id, err := base64.StdEncoding.DecodeString(vs.ID)
	if err != nil {
		return err
	}
	vs.ID = string(id[8:])
	return nil
}

func (vs *VsStage) EncodeId() {
	switch vs.ID {
	case "1":
		vs.Name = "Scorch Gorge"
	case "2":
		vs.Name = "Eeltail Alley"
	case "3":
		vs.Name = "Hagglefish Market"
	case "4":
		vs.Name = "Undertow Spillway"
	case "5":
		vs.Name = "Um'ami Ruins"
	case "6":
		vs.Name = "Mincemeat Metalworks"
	case "7":
		vs.Name = "Brinewater Springs"
	case "9":
		vs.Name = "Flounder Heights"
	case "10":
		vs.Name = "Hammerhead Bridge"
	case "11":
		vs.Name = "Museum d'Alfonsino"
	case "12":
		vs.Name = "Mahi-Mahi Resort"
	case "13":
		vs.Name = "Inkblot Art Academy"
	case "14":
		vs.Name = "Sturgeon Shipyard"
	case "15":
		vs.Name = "MakoMart"
	case "16":
		vs.Name = "Wahoo World"
	case "18":
		vs.Name = "Manta Maria"
	}
	vs.ID = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("VsStage-%s", vs.ID)))
}

type OtherTeams struct {
	Color              Color      `json:"color"`
	Result             TeamResult `json:"result"`
	TricolorRole       any        `json:"tricolorRole"`
	Judgement          string     `json:"judgement"`
	Players            []Players  `json:"players"`
	Order              int        `json:"order"`
	FestTeamName       any        `json:"festTeamName"`
	FestStreakWinCount any        `json:"festStreakWinCount"`
	FestUniformName    any        `json:"festUniformName"`
}
type Awards struct {
	Name string `json:"name"`
	Rank string `json:"rank"`
}
