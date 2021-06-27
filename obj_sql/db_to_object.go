package obj_sql

import (
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/api_objects"
	"github.com/cass-dlcm/SplatStatsGo/db_objects"
	"github.com/cass-dlcm/SplatStatsGo/enums"
	"log"
	"runtime/debug"
)

func GetBattlesLean(values []interface{}, conditionColumns []string, startAt, endAt int64, sort string) ([]api_objects.Battle, error) {
	if endAt-startAt > 50 {
		endAt = startAt + 50
	}
	pks, err := ReadKeyArrayKeyCondition(values, conditionColumns, "pk", "two_battles_battle", sort, "time")
	if err != nil {
		return nil, err
	}
	var battles []api_objects.Battle
	if endAt-startAt < int64(len(pks)) {
		battles = make([]api_objects.Battle, endAt-startAt)
	} else {
		battles = make([]api_objects.Battle, len(pks))
	}
	for i := startAt; i < endAt && i < int64(len(pks)); i++ {
		battle, err := GetBattleLeanPk(pks[i])
		if err != nil {
			return nil, err
		}
		battles[i-startAt] = *battle
	}
	return battles, nil
}

func GetBattlesLeanUser(userId int64, values []interface{}, conditionColumns []string, startAt, endAt int64, sort string) ([]api_objects.Battle, error) {
	if endAt-startAt > 50 {
		endAt = startAt + 50
	}
	values = append(values, userId)
	conditionColumns = append(conditionColumns, "user_id")
	pks, err := ReadKeyArrayKeyCondition(values, conditionColumns, "pk", "two_battles_battle", sort, "time")
	if err != nil {
		return nil, err
	}
	var battles []api_objects.Battle
	if endAt-startAt < int64(len(pks)) {
		battles = make([]api_objects.Battle, endAt-startAt)
	} else {
		battles = make([]api_objects.Battle, len(pks))
	}
	for i := startAt; i < endAt && i < int64(len(pks)); i++ {
		battle, err := GetBattleLeanPk(pks[i])
		if err != nil {
			return nil, err
		}
		battles[i-startAt] = *battle
	}
	return battles, nil
}

func GetShiftsLean(values []interface{}, conditionColumns []string, startAt, endAt int64, sort string) ([]api_objects.Shift, error) {
	if endAt-startAt > 50 {
		endAt = startAt + 50
	}
	pks, err := ReadKeyArrayKeyCondition(values, conditionColumns, "pk", "two_salmon_shift", sort, "play_time")
	if err != nil {
		return nil, err
	}
	var shifts []api_objects.Shift
	if endAt-startAt < int64(len(pks)) {
		shifts = make([]api_objects.Shift, endAt-startAt)
	} else {
		shifts = make([]api_objects.Shift, len(pks))
	}
	for i := startAt; i < endAt && i < int64(len(pks)); i++ {
		shift, err := GetShiftLeanPk(pks[i])
		if err != nil {
			return nil, err
		}
		shifts[i-startAt] = *shift
	}
	return shifts, nil
}

func GetBattlesUser(userId int64, values []interface{}, conditionColumns []string, startAt, endAt int64, sort string) ([]api_objects.Battle, error) {
	if endAt-startAt > 50 {
		endAt = startAt + 50
	}
	values = append(values, userId)
	conditionColumns = append(conditionColumns, "user_id")
	pks, err := ReadKeyArrayKeyCondition(values, conditionColumns, "pk", "two_battles_battle", sort, "time")
	if err != nil {
		return nil, err
	}
	var battles []api_objects.Battle
	if endAt-startAt < int64(len(pks)) {
		battles = make([]api_objects.Battle, endAt-startAt)
	} else {
		battles = make([]api_objects.Battle, len(pks))
	}
	for i := startAt; i < endAt; i++ {
		battleTemp, err := GetBattlePk(pks[i])
		if err != nil {
			return nil, err
		}
		battles[i-startAt] = *battleTemp
	}
	return battles, nil
}

func GetShiftsUser(userId, startAt, endAt int64, sort string) ([]api_objects.Shift, error) {
	if endAt-startAt > 50 {
		endAt = startAt + 50
	}
	splatnetNums, err := ReadKeyArrayWithKey(fmt.Sprint(userId), "user_id", "splatnet_number", "two_salmon_shift", sort, "play_time")
	if err != nil {
		return nil, err
	}
	var shifts []api_objects.Shift
	if endAt-startAt < int64(len(splatnetNums)) {
		shifts = make([]api_objects.Shift, endAt-startAt)
	} else {
		shifts = make([]api_objects.Shift, len(splatnetNums))
	}
	for i := startAt; i < endAt && i < int64(len(splatnetNums)); i++ {
		shiftTemp, err := GetShift(userId, splatnetNums[i])
		if err != nil {
			return nil, err
		}
		shifts[i-startAt] = *shiftTemp
	}
	return shifts, nil
}

func GetShiftsLeanUser(userId int64, values []interface{}, conditionColumns []string, startAt, endAt int64, sort string) ([]api_objects.Shift, error) {
	if endAt-startAt > 50 {
		endAt = startAt + 50
	}
	values = append(values, userId)
	conditionColumns = append(conditionColumns, "user_id")
	pks, err := ReadKeyArrayKeyCondition(values, conditionColumns, "pk", "two_salmon_shift", sort, "play_time")
	if err != nil {
		return nil, err
	}
	var shifts []api_objects.Shift
	if endAt-startAt < int64(len(pks)) {
		shifts = make([]api_objects.Shift, endAt-startAt)
	} else {
		shifts = make([]api_objects.Shift, len(pks))
	}
	for i := startAt; i < endAt && i < int64(len(pks)); i++ {
		shiftTemp, err := GetShiftLeanPk(pks[i])
		if err != nil {
			return nil, err
		}
		shifts[i-startAt] = *shiftTemp
	}
	return shifts, nil
}

func GetBattle(userId int64, splatnetNumber int64) (*api_objects.Battle, error) {
	var battle db_objects.Battle
	var pk int
	if err := readObjWithUserSplatnet(userId, splatnetNumber, "two_battles_battle").Scan(
		&pk, &battle.UserId, &battle.SplatnetJson, &battle.SplatnetUpload, &battle.StatInkJson, &battle.StatInkUpload,
		&battle.SplatnetNumber, &battle.PlayerSplatnetId, &battle.ElapsedTime, &battle.HasDisconnectedPlayer,
		&battle.LeaguePoint, &battle.MatchType, &battle.Rule, &battle.MyTeamCount, &battle.OtherTeamCount,
		&battle.SplatfestPoint, &battle.SplatfestTitle, &battle.Stage, &battle.TagId, &battle.Time, &battle.Win,
		&battle.WinMeter, &battle.Opponent0SplatnetId, &battle.Opponent0Name, &battle.Opponent0Rank,
		&battle.Opponent0LevelStar, &battle.Opponent0Level, &battle.Opponent0Weapon, &battle.Opponent0Gender,
		&battle.Opponent0Species, &battle.Opponent0Assists, &battle.Opponent0Deaths, &battle.Opponent0GamePaintPoint,
		&battle.Opponent0Kills, &battle.Opponent0Specials, &battle.Opponent0Headgear, &battle.Opponent0HeadgearMain,
		&battle.Opponent0HeadgearSub0, &battle.Opponent0HeadgearSub1, &battle.Opponent0HeadgearSub2,
		&battle.Opponent0Clothes, &battle.Opponent0ClothesMain, &battle.Opponent0ClothesSub0,
		&battle.Opponent0ClothesSub1, &battle.Opponent0ClothesSub2, &battle.Opponent0Shoes, &battle.Opponent0ShoesMain,
		&battle.Opponent0ShoesSub0, &battle.Opponent0ShoesSub1, &battle.Opponent0ShoesSub2, &battle.Opponent1SplatnetId,
		&battle.Opponent1Name, &battle.Opponent1Rank, &battle.Opponent1LevelStar, &battle.Opponent1Level,
		&battle.Opponent1Weapon, &battle.Opponent1Gender, &battle.Opponent1Species, &battle.Opponent1Assists,
		&battle.Opponent1Deaths, &battle.Opponent1GamePaintPoint, &battle.Opponent1Kills, &battle.Opponent1Specials,
		&battle.Opponent1Headgear, &battle.Opponent1HeadgearMain, &battle.Opponent1HeadgearSub0,
		&battle.Opponent1HeadgearSub1, &battle.Opponent1HeadgearSub2, &battle.Opponent1Clothes,
		&battle.Opponent1ClothesMain, &battle.Opponent1ClothesSub0, &battle.Opponent1ClothesSub1,
		&battle.Opponent1ClothesSub2, &battle.Opponent1Shoes, &battle.Opponent1ShoesMain, &battle.Opponent1ShoesSub0,
		&battle.Opponent1ShoesSub1, &battle.Opponent1ShoesSub2, &battle.Opponent2SplatnetId, &battle.Opponent2Name,
		&battle.Opponent2Rank, &battle.Opponent2LevelStar, &battle.Opponent2Level, &battle.Opponent2Weapon,
		&battle.Opponent2Gender, &battle.Opponent2Species, &battle.Opponent2Assists, &battle.Opponent2Deaths,
		&battle.Opponent2GamePaintPoint, &battle.Opponent2Kills, &battle.Opponent2Specials, &battle.Opponent2Headgear,
		&battle.Opponent2HeadgearMain, &battle.Opponent2HeadgearSub0, &battle.Opponent2HeadgearSub1,
		&battle.Opponent2HeadgearSub2, &battle.Opponent2Clothes, &battle.Opponent2ClothesMain,
		&battle.Opponent2ClothesSub0, &battle.Opponent2ClothesSub1, &battle.Opponent2ClothesSub2,
		&battle.Opponent2Shoes, &battle.Opponent2ShoesMain, &battle.Opponent2ShoesSub0, &battle.Opponent2ShoesSub1,
		&battle.Opponent2ShoesSub2, &battle.Opponent3SplatnetId, &battle.Opponent3Name, &battle.Opponent3Rank,
		&battle.Opponent3LevelStar, &battle.Opponent3Level, &battle.Opponent3Weapon, &battle.Opponent3Gender,
		&battle.Opponent3Species, &battle.Opponent3Assists, &battle.Opponent3Deaths, &battle.Opponent3GamePaintPoint,
		&battle.Opponent3Kills, &battle.Opponent3Specials, &battle.Opponent3Headgear, &battle.Opponent3HeadgearMain,
		&battle.Opponent3HeadgearSub0, &battle.Opponent3HeadgearSub1, &battle.Opponent3HeadgearSub2,
		&battle.Opponent3Clothes, &battle.Opponent3ClothesMain, &battle.Opponent3ClothesSub0,
		&battle.Opponent3ClothesSub1, &battle.Opponent3ClothesSub2, &battle.Opponent3Shoes, &battle.Opponent3ShoesMain,
		&battle.Opponent3ShoesSub0, &battle.Opponent3ShoesSub1, &battle.Opponent3ShoesSub2, &battle.Teammate0SplatnetId,
		&battle.Teammate0Name, &battle.Teammate0Rank, &battle.Teammate0LevelStar, &battle.Teammate0Level,
		&battle.Teammate0Weapon, &battle.Teammate0Gender, &battle.Teammate0Species, &battle.Teammate0Assists,
		&battle.Teammate0Deaths, &battle.Teammate0GamePaintPoint, &battle.Teammate0Kills, &battle.Teammate0Specials,
		&battle.Teammate0Headgear, &battle.Teammate0HeadgearMain, &battle.Teammate0HeadgearSub0,
		&battle.Teammate0HeadgearSub1, &battle.Teammate0HeadgearSub2, &battle.Teammate0Clothes,
		&battle.Teammate0ClothesMain, &battle.Teammate0ClothesSub0, &battle.Teammate0ClothesSub1,
		&battle.Teammate0ClothesSub2, &battle.Teammate0Shoes, &battle.Teammate0ShoesMain, &battle.Teammate0ShoesSub0,
		&battle.Teammate0ShoesSub1, &battle.Teammate0ShoesSub2, &battle.Teammate1SplatnetId, &battle.Teammate1Name,
		&battle.Teammate1Rank, &battle.Teammate1LevelStar, &battle.Teammate1Level, &battle.Teammate1Weapon,
		&battle.Teammate1Gender, &battle.Teammate1Species, &battle.Teammate1Assists, &battle.Teammate1Deaths,
		&battle.Teammate1GamePaintPoint, &battle.Teammate1Kills, &battle.Teammate1Specials, &battle.Teammate1Headgear,
		&battle.Teammate1HeadgearMain, &battle.Teammate1HeadgearSub0, &battle.Teammate1HeadgearSub1,
		&battle.Teammate1HeadgearSub2, &battle.Teammate1Clothes, &battle.Teammate1ClothesMain,
		&battle.Teammate1ClothesSub0, &battle.Teammate1ClothesSub1, &battle.Teammate1ClothesSub2,
		&battle.Teammate1Shoes, &battle.Teammate1ShoesMain, &battle.Teammate1ShoesSub0, &battle.Teammate1ShoesSub1,
		&battle.Teammate1ShoesSub2, &battle.Teammate2SplatnetId, &battle.Teammate2Name, &battle.Teammate2Rank,
		&battle.Teammate2LevelStar, &battle.Teammate2Level, &battle.Teammate2Weapon, &battle.Teammate2Gender,
		&battle.Teammate2Species, &battle.Teammate2Assists, &battle.Teammate2Deaths, &battle.Teammate2GamePaintPoint,
		&battle.Teammate2Kills, &battle.Teammate2Specials, &battle.Teammate2Headgear, &battle.Teammate2HeadgearMain,
		&battle.Teammate2HeadgearSub0, &battle.Teammate2HeadgearSub1, &battle.Teammate2HeadgearSub2,
		&battle.Teammate2Clothes, &battle.Teammate2ClothesMain, &battle.Teammate2ClothesSub0,
		&battle.Teammate2ClothesSub1, &battle.Teammate2ClothesSub2, &battle.Teammate2Shoes, &battle.Teammate2ShoesMain,
		&battle.Teammate2ShoesSub0, &battle.Teammate2ShoesSub1, &battle.Teammate2ShoesSub2,
		&battle.PlayerName, &battle.PlayerRank, &battle.PlayerLevelStar, &battle.PlayerLevel, &battle.PlayerWeapon,
		&battle.PlayerGender, &battle.PlayerSpecies, &battle.PlayerAssists, &battle.PlayerDeaths,
		&battle.PlayerGamePaintPoint, &battle.PlayerKills, &battle.PlayerSpecials, &battle.PlayerHeadgear,
		&battle.PlayerHeadgearMain, &battle.PlayerHeadgearSub0, &battle.PlayerHeadgearSub1, &battle.PlayerHeadgearSub2,
		&battle.PlayerClothes, &battle.PlayerClothesMain, &battle.PlayerClothesSub0, &battle.PlayerClothesSub1,
		&battle.PlayerClothesSub2, &battle.PlayerShoes, &battle.PlayerShoesMain, &battle.PlayerShoesSub0,
		&battle.PlayerShoesSub1, &battle.PlayerShoesSub2,
	); err != nil {
		return nil, err
	}
	var battleSplatnet *api_objects.BattleSplatnet
	if battle.SplatnetJson != nil {
		var err error
		battleSplatnet, err = readBattleSplatnet(*battle.SplatnetJson)
		if err != nil {
			return nil, err
		}
	}
	var battleStatInk *api_objects.BattleStatInk
	if battle.StatInkJson != nil {
		var err error
		battleStatInk, err = readBattleStatInk(*battle.StatInkJson)
		if err != nil {
			return nil, err
		}
	}
	hasDC := battle.HasDisconnectedPlayer != 0
	win := battle.Win != 0
	return &api_objects.Battle{
		UserId:                  battle.UserId,
		SplatnetJson:            battleSplatnet,
		StatInkJson:             battleStatInk,
		BattleNumber:            battle.SplatnetNumber,
		PlayerSplatnetId:        battle.PlayerSplatnetId,
		ElapsedTime:             battle.ElapsedTime,
		HasDisconnectedPlayer:   hasDC,
		LeaguePoint:             battle.LeaguePoint,
		MatchType:               battle.MatchType,
		Rule:                    battle.Rule,
		MyTeamCount:             battle.MyTeamCount,
		OtherTeamCount:          battle.OtherTeamCount,
		SplatfestPoint:          battle.SplatfestPoint,
		SplatfestTitle:          battle.SplatfestTitle,
		Stage:                   battle.Stage,
		TagId:                   battle.TagId,
		Time:                    battle.Time,
		Win:                     win,
		WinMeter:                battle.WinMeter,
		Opponent0SplatnetId:     battle.Opponent0SplatnetId,
		Opponent0Name:           battle.Opponent0Name,
		Opponent0Rank:           battle.Opponent0Rank,
		Opponent0LevelStar:      battle.Opponent0LevelStar,
		Opponent0Level:          battle.Opponent0Level,
		Opponent0Weapon:         battle.Opponent0Weapon,
		Opponent0Gender:         battle.Opponent0Gender,
		Opponent0Species:        battle.Opponent0Species,
		Opponent0Assists:        battle.Opponent0Assists,
		Opponent0Deaths:         battle.Opponent0Deaths,
		Opponent0GamePaintPoint: battle.Opponent0GamePaintPoint,
		Opponent0Kills:          battle.Opponent0Kills,
		Opponent0Specials:       battle.Opponent0Specials,
		Opponent0Headgear:       battle.Opponent0Headgear,
		Opponent0HeadgearMain:   battle.Opponent0HeadgearMain,
		Opponent0HeadgearSub0:   battle.Opponent0HeadgearSub0,
		Opponent0HeadgearSub1:   battle.Opponent0HeadgearSub1,
		Opponent0HeadgearSub2:   battle.Opponent0HeadgearSub2,
		Opponent0Clothes:        battle.Opponent0Clothes,
		Opponent0ClothesMain:    battle.Opponent0ClothesMain,
		Opponent0ClothesSub0:    battle.Opponent0ClothesSub0,
		Opponent0ClothesSub1:    battle.Opponent0ClothesSub1,
		Opponent0ClothesSub2:    battle.Opponent0ClothesSub2,
		Opponent0Shoes:          battle.Opponent0Shoes,
		Opponent0ShoesMain:      battle.Opponent0ShoesMain,
		Opponent0ShoesSub0:      battle.Opponent0ShoesSub0,
		Opponent0ShoesSub1:      battle.Opponent0ShoesSub1,
		Opponent0ShoesSub2:      battle.Opponent0ShoesSub2,
		Opponent1SplatnetId:     battle.Opponent1SplatnetId,
		Opponent1Name:           battle.Opponent1Name,
		Opponent1Rank:           battle.Opponent1Rank,
		Opponent1LevelStar:      battle.Opponent1LevelStar,
		Opponent1Level:          battle.Opponent1Level,
		Opponent1Weapon:         battle.Opponent1Weapon,
		Opponent1Gender:         battle.Opponent1Gender,
		Opponent1Species:        battle.Opponent1Species,
		Opponent1Assists:        battle.Opponent1Assists,
		Opponent1Deaths:         battle.Opponent1Deaths,
		Opponent1GamePaintPoint: battle.Opponent1GamePaintPoint,
		Opponent1Kills:          battle.Opponent1Kills,
		Opponent1Specials:       battle.Opponent1Specials,
		Opponent1Headgear:       battle.Opponent1Headgear,
		Opponent1HeadgearMain:   battle.Opponent1HeadgearMain,
		Opponent1HeadgearSub0:   battle.Opponent1HeadgearSub0,
		Opponent1HeadgearSub1:   battle.Opponent1HeadgearSub1,
		Opponent1HeadgearSub2:   battle.Opponent1HeadgearSub2,
		Opponent1Clothes:        battle.Opponent1Clothes,
		Opponent1ClothesMain:    battle.Opponent1ClothesMain,
		Opponent1ClothesSub0:    battle.Opponent1ClothesSub0,
		Opponent1ClothesSub1:    battle.Opponent1ClothesSub1,
		Opponent1ClothesSub2:    battle.Opponent1ClothesSub2,
		Opponent1Shoes:          battle.Opponent1Shoes,
		Opponent1ShoesMain:      battle.Opponent1ShoesMain,
		Opponent1ShoesSub0:      battle.Opponent1ShoesSub0,
		Opponent1ShoesSub1:      battle.Opponent1ShoesSub1,
		Opponent1ShoesSub2:      battle.Opponent1ShoesSub2,
		Opponent2SplatnetId:     battle.Opponent2SplatnetId,
		Opponent2Name:           battle.Opponent2Name,
		Opponent2Rank:           battle.Opponent2Rank,
		Opponent2LevelStar:      battle.Opponent2LevelStar,
		Opponent2Level:          battle.Opponent2Level,
		Opponent2Weapon:         battle.Opponent2Weapon,
		Opponent2Gender:         battle.Opponent2Gender,
		Opponent2Species:        battle.Opponent2Species,
		Opponent2Assists:        battle.Opponent2Assists,
		Opponent2Deaths:         battle.Opponent2Deaths,
		Opponent2GamePaintPoint: battle.Opponent2GamePaintPoint,
		Opponent2Kills:          battle.Opponent2Kills,
		Opponent2Specials:       battle.Opponent2Specials,
		Opponent2Headgear:       battle.Opponent2Headgear,
		Opponent2HeadgearMain:   battle.Opponent2HeadgearMain,
		Opponent2HeadgearSub0:   battle.Opponent2HeadgearSub0,
		Opponent2HeadgearSub1:   battle.Opponent2HeadgearSub1,
		Opponent2HeadgearSub2:   battle.Opponent2HeadgearSub2,
		Opponent2Clothes:        battle.Opponent2Clothes,
		Opponent2ClothesMain:    battle.Opponent2ClothesMain,
		Opponent2ClothesSub0:    battle.Opponent2ClothesSub0,
		Opponent2ClothesSub1:    battle.Opponent2ClothesSub1,
		Opponent2ClothesSub2:    battle.Opponent2ClothesSub2,
		Opponent2Shoes:          battle.Opponent2Shoes,
		Opponent2ShoesMain:      battle.Opponent2ShoesMain,
		Opponent2ShoesSub0:      battle.Opponent2ShoesSub0,
		Opponent2ShoesSub1:      battle.Opponent2ShoesSub1,
		Opponent2ShoesSub2:      battle.Opponent2ShoesSub2,
		Opponent3SplatnetId:     battle.Opponent3SplatnetId,
		Opponent3Name:           battle.Opponent3Name,
		Opponent3Rank:           battle.Opponent3Rank,
		Opponent3LevelStar:      battle.Opponent3LevelStar,
		Opponent3Level:          battle.Opponent3Level,
		Opponent3Weapon:         battle.Opponent3Weapon,
		Opponent3Gender:         battle.Opponent3Gender,
		Opponent3Species:        battle.Opponent3Species,
		Opponent3Assists:        battle.Opponent3Assists,
		Opponent3Deaths:         battle.Opponent3Deaths,
		Opponent3GamePaintPoint: battle.Opponent3GamePaintPoint,
		Opponent3Kills:          battle.Opponent3Kills,
		Opponent3Specials:       battle.Opponent3Specials,
		Opponent3Headgear:       battle.Opponent3Headgear,
		Opponent3HeadgearMain:   battle.Opponent3HeadgearMain,
		Opponent3HeadgearSub0:   battle.Opponent3HeadgearSub0,
		Opponent3HeadgearSub1:   battle.Opponent3HeadgearSub1,
		Opponent3HeadgearSub2:   battle.Opponent3HeadgearSub2,
		Opponent3Clothes:        battle.Opponent3Clothes,
		Opponent3ClothesMain:    battle.Opponent3ClothesMain,
		Opponent3ClothesSub0:    battle.Opponent3ClothesSub0,
		Opponent3ClothesSub1:    battle.Opponent3ClothesSub1,
		Opponent3ClothesSub2:    battle.Opponent3ClothesSub2,
		Opponent3Shoes:          battle.Opponent3Shoes,
		Opponent3ShoesMain:      battle.Opponent3ShoesMain,
		Opponent3ShoesSub0:      battle.Opponent3ShoesSub0,
		Opponent3ShoesSub1:      battle.Opponent3ShoesSub1,
		Opponent3ShoesSub2:      battle.Opponent3ShoesSub2,
		Teammate0SplatnetId:     battle.Teammate0SplatnetId,
		Teammate0Name:           battle.Teammate0Name,
		Teammate0Rank:           battle.Teammate0Rank,
		Teammate0LevelStar:      battle.Teammate0LevelStar,
		Teammate0Level:          battle.Teammate0Level,
		Teammate0Weapon:         battle.Teammate0Weapon,
		Teammate0Gender:         battle.Teammate0Gender,
		Teammate0Species:        battle.Teammate0Species,
		Teammate0Assists:        battle.Teammate0Assists,
		Teammate0Deaths:         battle.Teammate0Deaths,
		Teammate0GamePaintPoint: battle.Teammate0GamePaintPoint,
		Teammate0Kills:          battle.Teammate0Kills,
		Teammate0Specials:       battle.Teammate0Specials,
		Teammate0Headgear:       battle.Teammate0Headgear,
		Teammate0HeadgearMain:   battle.Teammate0HeadgearMain,
		Teammate0HeadgearSub0:   battle.Teammate0HeadgearSub0,
		Teammate0HeadgearSub1:   battle.Teammate0HeadgearSub1,
		Teammate0HeadgearSub2:   battle.Teammate0HeadgearSub2,
		Teammate0Clothes:        battle.Teammate0Clothes,
		Teammate0ClothesMain:    battle.Teammate0ClothesMain,
		Teammate0ClothesSub0:    battle.Teammate0ClothesSub0,
		Teammate0ClothesSub1:    battle.Teammate0ClothesSub1,
		Teammate0ClothesSub2:    battle.Teammate0ClothesSub2,
		Teammate0Shoes:          battle.Teammate0Shoes,
		Teammate0ShoesMain:      battle.Teammate0ShoesMain,
		Teammate0ShoesSub0:      battle.Teammate0ShoesSub0,
		Teammate0ShoesSub1:      battle.Teammate0ShoesSub1,
		Teammate0ShoesSub2:      battle.Teammate0ShoesSub2,
		Teammate1SplatnetId:     battle.Teammate1SplatnetId,
		Teammate1Name:           battle.Teammate1Name,
		Teammate1Rank:           battle.Teammate1Rank,
		Teammate1LevelStar:      battle.Teammate1LevelStar,
		Teammate1Level:          battle.Teammate1Level,
		Teammate1Weapon:         battle.Teammate1Weapon,
		Teammate1Gender:         battle.Teammate1Gender,
		Teammate1Species:        battle.Teammate1Species,
		Teammate1Assists:        battle.Teammate1Assists,
		Teammate1Deaths:         battle.Teammate1Deaths,
		Teammate1GamePaintPoint: battle.Teammate1GamePaintPoint,
		Teammate1Kills:          battle.Teammate1Kills,
		Teammate1Specials:       battle.Teammate1Specials,
		Teammate1Headgear:       battle.Teammate1Headgear,
		Teammate1HeadgearMain:   battle.Teammate1HeadgearMain,
		Teammate1HeadgearSub0:   battle.Teammate1HeadgearSub0,
		Teammate1HeadgearSub1:   battle.Teammate1HeadgearSub1,
		Teammate1HeadgearSub2:   battle.Teammate1HeadgearSub2,
		Teammate1Clothes:        battle.Teammate1Clothes,
		Teammate1ClothesMain:    battle.Teammate1ClothesMain,
		Teammate1ClothesSub0:    battle.Teammate1ClothesSub0,
		Teammate1ClothesSub1:    battle.Teammate1ClothesSub1,
		Teammate1ClothesSub2:    battle.Teammate1ClothesSub2,
		Teammate1Shoes:          battle.Teammate1Shoes,
		Teammate1ShoesMain:      battle.Teammate1ShoesMain,
		Teammate1ShoesSub0:      battle.Teammate1ShoesSub0,
		Teammate1ShoesSub1:      battle.Teammate1ShoesSub1,
		Teammate1ShoesSub2:      battle.Teammate1ShoesSub2,
		Teammate2SplatnetId:     battle.Teammate2SplatnetId,
		Teammate2Name:           battle.Teammate2Name,
		Teammate2Rank:           battle.Teammate2Rank,
		Teammate2LevelStar:      battle.Teammate2LevelStar,
		Teammate2Level:          battle.Teammate2Level,
		Teammate2Weapon:         battle.Teammate2Weapon,
		Teammate2Gender:         battle.Teammate2Gender,
		Teammate2Species:        battle.Teammate2Species,
		Teammate2Assists:        battle.Teammate2Assists,
		Teammate2Deaths:         battle.Teammate2Deaths,
		Teammate2GamePaintPoint: battle.Teammate2GamePaintPoint,
		Teammate2Kills:          battle.Teammate2Kills,
		Teammate2Specials:       battle.Teammate2Specials,
		Teammate2Headgear:       battle.Teammate2Headgear,
		Teammate2HeadgearMain:   battle.Teammate2HeadgearMain,
		Teammate2HeadgearSub0:   battle.Teammate2HeadgearSub0,
		Teammate2HeadgearSub1:   battle.Teammate2HeadgearSub1,
		Teammate2HeadgearSub2:   battle.Teammate2HeadgearSub2,
		Teammate2Clothes:        battle.Teammate2Clothes,
		Teammate2ClothesMain:    battle.Teammate2ClothesMain,
		Teammate2ClothesSub0:    battle.Teammate2ClothesSub0,
		Teammate2ClothesSub1:    battle.Teammate2ClothesSub1,
		Teammate2ClothesSub2:    battle.Teammate2ClothesSub2,
		Teammate2Shoes:          battle.Teammate2Shoes,
		Teammate2ShoesMain:      battle.Teammate2ShoesMain,
		Teammate2ShoesSub0:      battle.Teammate2ShoesSub0,
		Teammate2ShoesSub1:      battle.Teammate2ShoesSub1,
		Teammate2ShoesSub2:      battle.Teammate2ShoesSub2,
		PlayerName:              battle.PlayerName,
		PlayerRank:              battle.PlayerRank,
		PlayerLevelStar:         battle.PlayerLevelStar,
		PlayerLevel:             battle.PlayerLevel,
		PlayerWeapon:            battle.PlayerWeapon,
		PlayerGender:            battle.PlayerGender,
		PlayerSpecies:           battle.PlayerSpecies,
		PlayerAssists:           battle.PlayerAssists,
		PlayerDeaths:            battle.PlayerDeaths,
		PlayerGamePaintPoint:    battle.PlayerGamePaintPoint,
		PlayerKills:             battle.PlayerKills,
		PlayerSpecials:          battle.PlayerSpecials,
		PlayerHeadgear:          battle.PlayerHeadgear,
		PlayerHeadgearMain:      battle.PlayerHeadgearMain,
		PlayerHeadgearSub0:      battle.PlayerHeadgearSub0,
		PlayerHeadgearSub1:      battle.PlayerHeadgearSub1,
		PlayerHeadgearSub2:      battle.PlayerHeadgearSub2,
		PlayerClothes:           battle.PlayerClothes,
		PlayerClothesMain:       battle.PlayerClothesMain,
		PlayerClothesSub0:       battle.PlayerClothesSub0,
		PlayerClothesSub1:       battle.PlayerClothesSub1,
		PlayerClothesSub2:       battle.PlayerClothesSub2,
		PlayerShoes:             battle.PlayerShoes,
		PlayerShoesMain:         battle.PlayerShoesMain,
		PlayerShoesSub0:         battle.PlayerShoesSub0,
		PlayerShoesSub1:         battle.PlayerShoesSub1,
		PlayerShoesSub2:         battle.PlayerShoesSub2,
	}, nil
}

func GetBattleLean(userId int64, splatnetNumber int64) (*api_objects.Battle, error) {
	var battle db_objects.Battle
	var pk int
	if err := readObjWithUserSplatnet(userId, splatnetNumber, "two_battles_battle").Scan(
		&pk, &battle.UserId, &battle.SplatnetJson, &battle.SplatnetUpload, &battle.StatInkJson, &battle.StatInkUpload,
		&battle.SplatnetNumber, &battle.PlayerSplatnetId, &battle.ElapsedTime, &battle.HasDisconnectedPlayer,
		&battle.LeaguePoint, &battle.MatchType, &battle.Rule, &battle.MyTeamCount, &battle.OtherTeamCount,
		&battle.SplatfestPoint, &battle.SplatfestTitle, &battle.Stage, &battle.TagId, &battle.Time, &battle.Win,
		&battle.WinMeter, &battle.Opponent0SplatnetId, &battle.Opponent0Name, &battle.Opponent0Rank,
		&battle.Opponent0LevelStar, &battle.Opponent0Level, &battle.Opponent0Weapon, &battle.Opponent0Gender,
		&battle.Opponent0Species, &battle.Opponent0Assists, &battle.Opponent0Deaths, &battle.Opponent0GamePaintPoint,
		&battle.Opponent0Kills, &battle.Opponent0Specials, &battle.Opponent0Headgear, &battle.Opponent0HeadgearMain,
		&battle.Opponent0HeadgearSub0, &battle.Opponent0HeadgearSub1, &battle.Opponent0HeadgearSub2,
		&battle.Opponent0Clothes, &battle.Opponent0ClothesMain, &battle.Opponent0ClothesSub0,
		&battle.Opponent0ClothesSub1, &battle.Opponent0ClothesSub2, &battle.Opponent0Shoes, &battle.Opponent0ShoesMain,
		&battle.Opponent0ShoesSub0, &battle.Opponent0ShoesSub1, &battle.Opponent0ShoesSub2, &battle.Opponent1SplatnetId,
		&battle.Opponent1Name, &battle.Opponent1Rank, &battle.Opponent1LevelStar, &battle.Opponent1Level,
		&battle.Opponent1Weapon, &battle.Opponent1Gender, &battle.Opponent1Species, &battle.Opponent1Assists,
		&battle.Opponent1Deaths, &battle.Opponent1GamePaintPoint, &battle.Opponent1Kills, &battle.Opponent1Specials,
		&battle.Opponent1Headgear, &battle.Opponent1HeadgearMain, &battle.Opponent1HeadgearSub0,
		&battle.Opponent1HeadgearSub1, &battle.Opponent1HeadgearSub2, &battle.Opponent1Clothes,
		&battle.Opponent1ClothesMain, &battle.Opponent1ClothesSub0, &battle.Opponent1ClothesSub1,
		&battle.Opponent1ClothesSub2, &battle.Opponent1Shoes, &battle.Opponent1ShoesMain, &battle.Opponent1ShoesSub0,
		&battle.Opponent1ShoesSub1, &battle.Opponent1ShoesSub2, &battle.Opponent2SplatnetId, &battle.Opponent2Name,
		&battle.Opponent2Rank, &battle.Opponent2LevelStar, &battle.Opponent2Level, &battle.Opponent2Weapon,
		&battle.Opponent2Gender, &battle.Opponent2Species, &battle.Opponent2Assists, &battle.Opponent2Deaths,
		&battle.Opponent2GamePaintPoint, &battle.Opponent2Kills, &battle.Opponent2Specials, &battle.Opponent2Headgear,
		&battle.Opponent2HeadgearMain, &battle.Opponent2HeadgearSub0, &battle.Opponent2HeadgearSub1,
		&battle.Opponent2HeadgearSub2, &battle.Opponent2Clothes, &battle.Opponent2ClothesMain,
		&battle.Opponent2ClothesSub0, &battle.Opponent2ClothesSub1, &battle.Opponent2ClothesSub2,
		&battle.Opponent2Shoes, &battle.Opponent2ShoesMain, &battle.Opponent2ShoesSub0, &battle.Opponent2ShoesSub1,
		&battle.Opponent2ShoesSub2, &battle.Opponent3SplatnetId, &battle.Opponent3Name, &battle.Opponent3Rank,
		&battle.Opponent3LevelStar, &battle.Opponent3Level, &battle.Opponent3Weapon, &battle.Opponent3Gender,
		&battle.Opponent3Species, &battle.Opponent3Assists, &battle.Opponent3Deaths, &battle.Opponent3GamePaintPoint,
		&battle.Opponent3Kills, &battle.Opponent3Specials, &battle.Opponent3Headgear, &battle.Opponent3HeadgearMain,
		&battle.Opponent3HeadgearSub0, &battle.Opponent3HeadgearSub1, &battle.Opponent3HeadgearSub2,
		&battle.Opponent3Clothes, &battle.Opponent3ClothesMain, &battle.Opponent3ClothesSub0,
		&battle.Opponent3ClothesSub1, &battle.Opponent3ClothesSub2, &battle.Opponent3Shoes, &battle.Opponent3ShoesMain,
		&battle.Opponent3ShoesSub0, &battle.Opponent3ShoesSub1, &battle.Opponent3ShoesSub2, &battle.Teammate0SplatnetId,
		&battle.Teammate0Name, &battle.Teammate0Rank, &battle.Teammate0LevelStar, &battle.Teammate0Level,
		&battle.Teammate0Weapon, &battle.Teammate0Gender, &battle.Teammate0Species, &battle.Teammate0Assists,
		&battle.Teammate0Deaths, &battle.Teammate0GamePaintPoint, &battle.Teammate0Kills, &battle.Teammate0Specials,
		&battle.Teammate0Headgear, &battle.Teammate0HeadgearMain, &battle.Teammate0HeadgearSub0,
		&battle.Teammate0HeadgearSub1, &battle.Teammate0HeadgearSub2, &battle.Teammate0Clothes,
		&battle.Teammate0ClothesMain, &battle.Teammate0ClothesSub0, &battle.Teammate0ClothesSub1,
		&battle.Teammate0ClothesSub2, &battle.Teammate0Shoes, &battle.Teammate0ShoesMain, &battle.Teammate0ShoesSub0,
		&battle.Teammate0ShoesSub1, &battle.Teammate0ShoesSub2, &battle.Teammate1SplatnetId, &battle.Teammate1Name,
		&battle.Teammate1Rank, &battle.Teammate1LevelStar, &battle.Teammate1Level, &battle.Teammate1Weapon,
		&battle.Teammate1Gender, &battle.Teammate1Species, &battle.Teammate1Assists, &battle.Teammate1Deaths,
		&battle.Teammate1GamePaintPoint, &battle.Teammate1Kills, &battle.Teammate1Specials, &battle.Teammate1Headgear,
		&battle.Teammate1HeadgearMain, &battle.Teammate1HeadgearSub0, &battle.Teammate1HeadgearSub1,
		&battle.Teammate1HeadgearSub2, &battle.Teammate1Clothes, &battle.Teammate1ClothesMain,
		&battle.Teammate1ClothesSub0, &battle.Teammate1ClothesSub1, &battle.Teammate1ClothesSub2,
		&battle.Teammate1Shoes, &battle.Teammate1ShoesMain, &battle.Teammate1ShoesSub0, &battle.Teammate1ShoesSub1,
		&battle.Teammate1ShoesSub2, &battle.Teammate2SplatnetId, &battle.Teammate2Name, &battle.Teammate2Rank,
		&battle.Teammate2LevelStar, &battle.Teammate2Level, &battle.Teammate2Weapon, &battle.Teammate2Gender,
		&battle.Teammate2Species, &battle.Teammate2Assists, &battle.Teammate2Deaths, &battle.Teammate2GamePaintPoint,
		&battle.Teammate2Kills, &battle.Teammate2Specials, &battle.Teammate2Headgear, &battle.Teammate2HeadgearMain,
		&battle.Teammate2HeadgearSub0, &battle.Teammate2HeadgearSub1, &battle.Teammate2HeadgearSub2,
		&battle.Teammate2Clothes, &battle.Teammate2ClothesMain, &battle.Teammate2ClothesSub0,
		&battle.Teammate2ClothesSub1, &battle.Teammate2ClothesSub2, &battle.Teammate2Shoes, &battle.Teammate2ShoesMain,
		&battle.Teammate2ShoesSub0, &battle.Teammate2ShoesSub1, &battle.Teammate2ShoesSub2,
		&battle.PlayerName, &battle.PlayerRank, &battle.PlayerLevelStar, &battle.PlayerLevel, &battle.PlayerWeapon,
		&battle.PlayerGender, &battle.PlayerSpecies, &battle.PlayerAssists, &battle.PlayerDeaths,
		&battle.PlayerGamePaintPoint, &battle.PlayerKills, &battle.PlayerSpecials, &battle.PlayerHeadgear,
		&battle.PlayerHeadgearMain, &battle.PlayerHeadgearSub0, &battle.PlayerHeadgearSub1, &battle.PlayerHeadgearSub2,
		&battle.PlayerClothes, &battle.PlayerClothesMain, &battle.PlayerClothesSub0, &battle.PlayerClothesSub1,
		&battle.PlayerClothesSub2, &battle.PlayerShoes, &battle.PlayerShoesMain, &battle.PlayerShoesSub0,
		&battle.PlayerShoesSub1, &battle.PlayerShoesSub2,
	); err != nil {
		return nil, err
	}
	hasDC := battle.HasDisconnectedPlayer != 0
	win := battle.Win != 0
	return &api_objects.Battle{
		UserId:                  battle.UserId,
		SplatnetJson:            nil,
		StatInkJson:             nil,
		BattleNumber:            battle.SplatnetNumber,
		PlayerSplatnetId:        battle.PlayerSplatnetId,
		ElapsedTime:             battle.ElapsedTime,
		HasDisconnectedPlayer:   hasDC,
		LeaguePoint:             battle.LeaguePoint,
		MatchType:               battle.MatchType,
		Rule:                    battle.Rule,
		MyTeamCount:             battle.MyTeamCount,
		OtherTeamCount:          battle.OtherTeamCount,
		SplatfestPoint:          battle.SplatfestPoint,
		SplatfestTitle:          battle.SplatfestTitle,
		Stage:                   battle.Stage,
		TagId:                   battle.TagId,
		Time:                    battle.Time,
		Win:                     win,
		WinMeter:                battle.WinMeter,
		Opponent0SplatnetId:     battle.Opponent0SplatnetId,
		Opponent0Name:           battle.Opponent0Name,
		Opponent0Rank:           battle.Opponent0Rank,
		Opponent0LevelStar:      battle.Opponent0LevelStar,
		Opponent0Level:          battle.Opponent0Level,
		Opponent0Weapon:         battle.Opponent0Weapon,
		Opponent0Gender:         battle.Opponent0Gender,
		Opponent0Species:        battle.Opponent0Species,
		Opponent0Assists:        battle.Opponent0Assists,
		Opponent0Deaths:         battle.Opponent0Deaths,
		Opponent0GamePaintPoint: battle.Opponent0GamePaintPoint,
		Opponent0Kills:          battle.Opponent0Kills,
		Opponent0Specials:       battle.Opponent0Specials,
		Opponent0Headgear:       battle.Opponent0Headgear,
		Opponent0HeadgearMain:   battle.Opponent0HeadgearMain,
		Opponent0HeadgearSub0:   battle.Opponent0HeadgearSub0,
		Opponent0HeadgearSub1:   battle.Opponent0HeadgearSub1,
		Opponent0HeadgearSub2:   battle.Opponent0HeadgearSub2,
		Opponent0Clothes:        battle.Opponent0Clothes,
		Opponent0ClothesMain:    battle.Opponent0ClothesMain,
		Opponent0ClothesSub0:    battle.Opponent0ClothesSub0,
		Opponent0ClothesSub1:    battle.Opponent0ClothesSub1,
		Opponent0ClothesSub2:    battle.Opponent0ClothesSub2,
		Opponent0Shoes:          battle.Opponent0Shoes,
		Opponent0ShoesMain:      battle.Opponent0ShoesMain,
		Opponent0ShoesSub0:      battle.Opponent0ShoesSub0,
		Opponent0ShoesSub1:      battle.Opponent0ShoesSub1,
		Opponent0ShoesSub2:      battle.Opponent0ShoesSub2,
		Opponent1SplatnetId:     battle.Opponent1SplatnetId,
		Opponent1Name:           battle.Opponent1Name,
		Opponent1Rank:           battle.Opponent1Rank,
		Opponent1LevelStar:      battle.Opponent1LevelStar,
		Opponent1Level:          battle.Opponent1Level,
		Opponent1Weapon:         battle.Opponent1Weapon,
		Opponent1Gender:         battle.Opponent1Gender,
		Opponent1Species:        battle.Opponent1Species,
		Opponent1Assists:        battle.Opponent1Assists,
		Opponent1Deaths:         battle.Opponent1Deaths,
		Opponent1GamePaintPoint: battle.Opponent1GamePaintPoint,
		Opponent1Kills:          battle.Opponent1Kills,
		Opponent1Specials:       battle.Opponent1Specials,
		Opponent1Headgear:       battle.Opponent1Headgear,
		Opponent1HeadgearMain:   battle.Opponent1HeadgearMain,
		Opponent1HeadgearSub0:   battle.Opponent1HeadgearSub0,
		Opponent1HeadgearSub1:   battle.Opponent1HeadgearSub1,
		Opponent1HeadgearSub2:   battle.Opponent1HeadgearSub2,
		Opponent1Clothes:        battle.Opponent1Clothes,
		Opponent1ClothesMain:    battle.Opponent1ClothesMain,
		Opponent1ClothesSub0:    battle.Opponent1ClothesSub0,
		Opponent1ClothesSub1:    battle.Opponent1ClothesSub1,
		Opponent1ClothesSub2:    battle.Opponent1ClothesSub2,
		Opponent1Shoes:          battle.Opponent1Shoes,
		Opponent1ShoesMain:      battle.Opponent1ShoesMain,
		Opponent1ShoesSub0:      battle.Opponent1ShoesSub0,
		Opponent1ShoesSub1:      battle.Opponent1ShoesSub1,
		Opponent1ShoesSub2:      battle.Opponent1ShoesSub2,
		Opponent2SplatnetId:     battle.Opponent2SplatnetId,
		Opponent2Name:           battle.Opponent2Name,
		Opponent2Rank:           battle.Opponent2Rank,
		Opponent2LevelStar:      battle.Opponent2LevelStar,
		Opponent2Level:          battle.Opponent2Level,
		Opponent2Weapon:         battle.Opponent2Weapon,
		Opponent2Gender:         battle.Opponent2Gender,
		Opponent2Species:        battle.Opponent2Species,
		Opponent2Assists:        battle.Opponent2Assists,
		Opponent2Deaths:         battle.Opponent2Deaths,
		Opponent2GamePaintPoint: battle.Opponent2GamePaintPoint,
		Opponent2Kills:          battle.Opponent2Kills,
		Opponent2Specials:       battle.Opponent2Specials,
		Opponent2Headgear:       battle.Opponent2Headgear,
		Opponent2HeadgearMain:   battle.Opponent2HeadgearMain,
		Opponent2HeadgearSub0:   battle.Opponent2HeadgearSub0,
		Opponent2HeadgearSub1:   battle.Opponent2HeadgearSub1,
		Opponent2HeadgearSub2:   battle.Opponent2HeadgearSub2,
		Opponent2Clothes:        battle.Opponent2Clothes,
		Opponent2ClothesMain:    battle.Opponent2ClothesMain,
		Opponent2ClothesSub0:    battle.Opponent2ClothesSub0,
		Opponent2ClothesSub1:    battle.Opponent2ClothesSub1,
		Opponent2ClothesSub2:    battle.Opponent2ClothesSub2,
		Opponent2Shoes:          battle.Opponent2Shoes,
		Opponent2ShoesMain:      battle.Opponent2ShoesMain,
		Opponent2ShoesSub0:      battle.Opponent2ShoesSub0,
		Opponent2ShoesSub1:      battle.Opponent2ShoesSub1,
		Opponent2ShoesSub2:      battle.Opponent2ShoesSub2,
		Opponent3SplatnetId:     battle.Opponent3SplatnetId,
		Opponent3Name:           battle.Opponent3Name,
		Opponent3Rank:           battle.Opponent3Rank,
		Opponent3LevelStar:      battle.Opponent3LevelStar,
		Opponent3Level:          battle.Opponent3Level,
		Opponent3Weapon:         battle.Opponent3Weapon,
		Opponent3Gender:         battle.Opponent3Gender,
		Opponent3Species:        battle.Opponent3Species,
		Opponent3Assists:        battle.Opponent3Assists,
		Opponent3Deaths:         battle.Opponent3Deaths,
		Opponent3GamePaintPoint: battle.Opponent3GamePaintPoint,
		Opponent3Kills:          battle.Opponent3Kills,
		Opponent3Specials:       battle.Opponent3Specials,
		Opponent3Headgear:       battle.Opponent3Headgear,
		Opponent3HeadgearMain:   battle.Opponent3HeadgearMain,
		Opponent3HeadgearSub0:   battle.Opponent3HeadgearSub0,
		Opponent3HeadgearSub1:   battle.Opponent3HeadgearSub1,
		Opponent3HeadgearSub2:   battle.Opponent3HeadgearSub2,
		Opponent3Clothes:        battle.Opponent3Clothes,
		Opponent3ClothesMain:    battle.Opponent3ClothesMain,
		Opponent3ClothesSub0:    battle.Opponent3ClothesSub0,
		Opponent3ClothesSub1:    battle.Opponent3ClothesSub1,
		Opponent3ClothesSub2:    battle.Opponent3ClothesSub2,
		Opponent3Shoes:          battle.Opponent3Shoes,
		Opponent3ShoesMain:      battle.Opponent3ShoesMain,
		Opponent3ShoesSub0:      battle.Opponent3ShoesSub0,
		Opponent3ShoesSub1:      battle.Opponent3ShoesSub1,
		Opponent3ShoesSub2:      battle.Opponent3ShoesSub2,
		Teammate0SplatnetId:     battle.Teammate0SplatnetId,
		Teammate0Name:           battle.Teammate0Name,
		Teammate0Rank:           battle.Teammate0Rank,
		Teammate0LevelStar:      battle.Teammate0LevelStar,
		Teammate0Level:          battle.Teammate0Level,
		Teammate0Weapon:         battle.Teammate0Weapon,
		Teammate0Gender:         battle.Teammate0Gender,
		Teammate0Species:        battle.Teammate0Species,
		Teammate0Assists:        battle.Teammate0Assists,
		Teammate0Deaths:         battle.Teammate0Deaths,
		Teammate0GamePaintPoint: battle.Teammate0GamePaintPoint,
		Teammate0Kills:          battle.Teammate0Kills,
		Teammate0Specials:       battle.Teammate0Specials,
		Teammate0Headgear:       battle.Teammate0Headgear,
		Teammate0HeadgearMain:   battle.Teammate0HeadgearMain,
		Teammate0HeadgearSub0:   battle.Teammate0HeadgearSub0,
		Teammate0HeadgearSub1:   battle.Teammate0HeadgearSub1,
		Teammate0HeadgearSub2:   battle.Teammate0HeadgearSub2,
		Teammate0Clothes:        battle.Teammate0Clothes,
		Teammate0ClothesMain:    battle.Teammate0ClothesMain,
		Teammate0ClothesSub0:    battle.Teammate0ClothesSub0,
		Teammate0ClothesSub1:    battle.Teammate0ClothesSub1,
		Teammate0ClothesSub2:    battle.Teammate0ClothesSub2,
		Teammate0Shoes:          battle.Teammate0Shoes,
		Teammate0ShoesMain:      battle.Teammate0ShoesMain,
		Teammate0ShoesSub0:      battle.Teammate0ShoesSub0,
		Teammate0ShoesSub1:      battle.Teammate0ShoesSub1,
		Teammate0ShoesSub2:      battle.Teammate0ShoesSub2,
		Teammate1SplatnetId:     battle.Teammate1SplatnetId,
		Teammate1Name:           battle.Teammate1Name,
		Teammate1Rank:           battle.Teammate1Rank,
		Teammate1LevelStar:      battle.Teammate1LevelStar,
		Teammate1Level:          battle.Teammate1Level,
		Teammate1Weapon:         battle.Teammate1Weapon,
		Teammate1Gender:         battle.Teammate1Gender,
		Teammate1Species:        battle.Teammate1Species,
		Teammate1Assists:        battle.Teammate1Assists,
		Teammate1Deaths:         battle.Teammate1Deaths,
		Teammate1GamePaintPoint: battle.Teammate1GamePaintPoint,
		Teammate1Kills:          battle.Teammate1Kills,
		Teammate1Specials:       battle.Teammate1Specials,
		Teammate1Headgear:       battle.Teammate1Headgear,
		Teammate1HeadgearMain:   battle.Teammate1HeadgearMain,
		Teammate1HeadgearSub0:   battle.Teammate1HeadgearSub0,
		Teammate1HeadgearSub1:   battle.Teammate1HeadgearSub1,
		Teammate1HeadgearSub2:   battle.Teammate1HeadgearSub2,
		Teammate1Clothes:        battle.Teammate1Clothes,
		Teammate1ClothesMain:    battle.Teammate1ClothesMain,
		Teammate1ClothesSub0:    battle.Teammate1ClothesSub0,
		Teammate1ClothesSub1:    battle.Teammate1ClothesSub1,
		Teammate1ClothesSub2:    battle.Teammate1ClothesSub2,
		Teammate1Shoes:          battle.Teammate1Shoes,
		Teammate1ShoesMain:      battle.Teammate1ShoesMain,
		Teammate1ShoesSub0:      battle.Teammate1ShoesSub0,
		Teammate1ShoesSub1:      battle.Teammate1ShoesSub1,
		Teammate1ShoesSub2:      battle.Teammate1ShoesSub2,
		Teammate2SplatnetId:     battle.Teammate2SplatnetId,
		Teammate2Name:           battle.Teammate2Name,
		Teammate2Rank:           battle.Teammate2Rank,
		Teammate2LevelStar:      battle.Teammate2LevelStar,
		Teammate2Level:          battle.Teammate2Level,
		Teammate2Weapon:         battle.Teammate2Weapon,
		Teammate2Gender:         battle.Teammate2Gender,
		Teammate2Species:        battle.Teammate2Species,
		Teammate2Assists:        battle.Teammate2Assists,
		Teammate2Deaths:         battle.Teammate2Deaths,
		Teammate2GamePaintPoint: battle.Teammate2GamePaintPoint,
		Teammate2Kills:          battle.Teammate2Kills,
		Teammate2Specials:       battle.Teammate2Specials,
		Teammate2Headgear:       battle.Teammate2Headgear,
		Teammate2HeadgearMain:   battle.Teammate2HeadgearMain,
		Teammate2HeadgearSub0:   battle.Teammate2HeadgearSub0,
		Teammate2HeadgearSub1:   battle.Teammate2HeadgearSub1,
		Teammate2HeadgearSub2:   battle.Teammate2HeadgearSub2,
		Teammate2Clothes:        battle.Teammate2Clothes,
		Teammate2ClothesMain:    battle.Teammate2ClothesMain,
		Teammate2ClothesSub0:    battle.Teammate2ClothesSub0,
		Teammate2ClothesSub1:    battle.Teammate2ClothesSub1,
		Teammate2ClothesSub2:    battle.Teammate2ClothesSub2,
		Teammate2Shoes:          battle.Teammate2Shoes,
		Teammate2ShoesMain:      battle.Teammate2ShoesMain,
		Teammate2ShoesSub0:      battle.Teammate2ShoesSub0,
		Teammate2ShoesSub1:      battle.Teammate2ShoesSub1,
		Teammate2ShoesSub2:      battle.Teammate2ShoesSub2,
		PlayerName:              battle.PlayerName,
		PlayerRank:              battle.PlayerRank,
		PlayerLevelStar:         battle.PlayerLevelStar,
		PlayerLevel:             battle.PlayerLevel,
		PlayerWeapon:            battle.PlayerWeapon,
		PlayerGender:            battle.PlayerGender,
		PlayerSpecies:           battle.PlayerSpecies,
		PlayerAssists:           battle.PlayerAssists,
		PlayerDeaths:            battle.PlayerDeaths,
		PlayerGamePaintPoint:    battle.PlayerGamePaintPoint,
		PlayerKills:             battle.PlayerKills,
		PlayerSpecials:          battle.PlayerSpecials,
		PlayerHeadgear:          battle.PlayerHeadgear,
		PlayerHeadgearMain:      battle.PlayerHeadgearMain,
		PlayerHeadgearSub0:      battle.PlayerHeadgearSub0,
		PlayerHeadgearSub1:      battle.PlayerHeadgearSub1,
		PlayerHeadgearSub2:      battle.PlayerHeadgearSub2,
		PlayerClothes:           battle.PlayerClothes,
		PlayerClothesMain:       battle.PlayerClothesMain,
		PlayerClothesSub0:       battle.PlayerClothesSub0,
		PlayerClothesSub1:       battle.PlayerClothesSub1,
		PlayerClothesSub2:       battle.PlayerClothesSub2,
		PlayerShoes:             battle.PlayerShoes,
		PlayerShoesMain:         battle.PlayerShoesMain,
		PlayerShoesSub0:         battle.PlayerShoesSub0,
		PlayerShoesSub1:         battle.PlayerShoesSub1,
		PlayerShoesSub2:         battle.PlayerShoesSub2,
	}, nil
}

func GetBattlePk(pk int64) (*api_objects.Battle, error) {
	var battle db_objects.Battle
	if err := readObjWithId(pk, "two_battles_battle").Scan(
		&pk, &battle.UserId, &battle.SplatnetJson, &battle.SplatnetUpload, &battle.StatInkJson, &battle.StatInkUpload,
		&battle.SplatnetNumber, &battle.PlayerSplatnetId, &battle.ElapsedTime, &battle.HasDisconnectedPlayer,
		&battle.LeaguePoint, &battle.MatchType, &battle.Rule, &battle.MyTeamCount, &battle.OtherTeamCount,
		&battle.SplatfestPoint, &battle.SplatfestTitle, &battle.Stage, &battle.TagId, &battle.Time, &battle.Win,
		&battle.WinMeter, &battle.Opponent0SplatnetId, &battle.Opponent0Name, &battle.Opponent0Rank,
		&battle.Opponent0LevelStar, &battle.Opponent0Level, &battle.Opponent0Weapon, &battle.Opponent0Gender,
		&battle.Opponent0Species, &battle.Opponent0Assists, &battle.Opponent0Deaths, &battle.Opponent0GamePaintPoint,
		&battle.Opponent0Kills, &battle.Opponent0Specials, &battle.Opponent0Headgear, &battle.Opponent0HeadgearMain,
		&battle.Opponent0HeadgearSub0, &battle.Opponent0HeadgearSub1, &battle.Opponent0HeadgearSub2,
		&battle.Opponent0Clothes, &battle.Opponent0ClothesMain, &battle.Opponent0ClothesSub0,
		&battle.Opponent0ClothesSub1, &battle.Opponent0ClothesSub2, &battle.Opponent0Shoes, &battle.Opponent0ShoesMain,
		&battle.Opponent0ShoesSub0, &battle.Opponent0ShoesSub1, &battle.Opponent0ShoesSub2, &battle.Opponent1SplatnetId,
		&battle.Opponent1Name, &battle.Opponent1Rank, &battle.Opponent1LevelStar, &battle.Opponent1Level,
		&battle.Opponent1Weapon, &battle.Opponent1Gender, &battle.Opponent1Species, &battle.Opponent1Assists,
		&battle.Opponent1Deaths, &battle.Opponent1GamePaintPoint, &battle.Opponent1Kills, &battle.Opponent1Specials,
		&battle.Opponent1Headgear, &battle.Opponent1HeadgearMain, &battle.Opponent1HeadgearSub0,
		&battle.Opponent1HeadgearSub1, &battle.Opponent1HeadgearSub2, &battle.Opponent1Clothes,
		&battle.Opponent1ClothesMain, &battle.Opponent1ClothesSub0, &battle.Opponent1ClothesSub1,
		&battle.Opponent1ClothesSub2, &battle.Opponent1Shoes, &battle.Opponent1ShoesMain, &battle.Opponent1ShoesSub0,
		&battle.Opponent1ShoesSub1, &battle.Opponent1ShoesSub2, &battle.Opponent2SplatnetId, &battle.Opponent2Name,
		&battle.Opponent2Rank, &battle.Opponent2LevelStar, &battle.Opponent2Level, &battle.Opponent2Weapon,
		&battle.Opponent2Gender, &battle.Opponent2Species, &battle.Opponent2Assists, &battle.Opponent2Deaths,
		&battle.Opponent2GamePaintPoint, &battle.Opponent2Kills, &battle.Opponent2Specials, &battle.Opponent2Headgear,
		&battle.Opponent2HeadgearMain, &battle.Opponent2HeadgearSub0, &battle.Opponent2HeadgearSub1,
		&battle.Opponent2HeadgearSub2, &battle.Opponent2Clothes, &battle.Opponent2ClothesMain,
		&battle.Opponent2ClothesSub0, &battle.Opponent2ClothesSub1, &battle.Opponent2ClothesSub2,
		&battle.Opponent2Shoes, &battle.Opponent2ShoesMain, &battle.Opponent2ShoesSub0, &battle.Opponent2ShoesSub1,
		&battle.Opponent2ShoesSub2, &battle.Opponent3SplatnetId, &battle.Opponent3Name, &battle.Opponent3Rank,
		&battle.Opponent3LevelStar, &battle.Opponent3Level, &battle.Opponent3Weapon, &battle.Opponent3Gender,
		&battle.Opponent3Species, &battle.Opponent3Assists, &battle.Opponent3Deaths, &battle.Opponent3GamePaintPoint,
		&battle.Opponent3Kills, &battle.Opponent3Specials, &battle.Opponent3Headgear, &battle.Opponent3HeadgearMain,
		&battle.Opponent3HeadgearSub0, &battle.Opponent3HeadgearSub1, &battle.Opponent3HeadgearSub2,
		&battle.Opponent3Clothes, &battle.Opponent3ClothesMain, &battle.Opponent3ClothesSub0,
		&battle.Opponent3ClothesSub1, &battle.Opponent3ClothesSub2, &battle.Opponent3Shoes, &battle.Opponent3ShoesMain,
		&battle.Opponent3ShoesSub0, &battle.Opponent3ShoesSub1, &battle.Opponent3ShoesSub2, &battle.Teammate0SplatnetId,
		&battle.Teammate0Name, &battle.Teammate0Rank, &battle.Teammate0LevelStar, &battle.Teammate0Level,
		&battle.Teammate0Weapon, &battle.Teammate0Gender, &battle.Teammate0Species, &battle.Teammate0Assists,
		&battle.Teammate0Deaths, &battle.Teammate0GamePaintPoint, &battle.Teammate0Kills, &battle.Teammate0Specials,
		&battle.Teammate0Headgear, &battle.Teammate0HeadgearMain, &battle.Teammate0HeadgearSub0,
		&battle.Teammate0HeadgearSub1, &battle.Teammate0HeadgearSub2, &battle.Teammate0Clothes,
		&battle.Teammate0ClothesMain, &battle.Teammate0ClothesSub0, &battle.Teammate0ClothesSub1,
		&battle.Teammate0ClothesSub2, &battle.Teammate0Shoes, &battle.Teammate0ShoesMain, &battle.Teammate0ShoesSub0,
		&battle.Teammate0ShoesSub1, &battle.Teammate0ShoesSub2, &battle.Teammate1SplatnetId, &battle.Teammate1Name,
		&battle.Teammate1Rank, &battle.Teammate1LevelStar, &battle.Teammate1Level, &battle.Teammate1Weapon,
		&battle.Teammate1Gender, &battle.Teammate1Species, &battle.Teammate1Assists, &battle.Teammate1Deaths,
		&battle.Teammate1GamePaintPoint, &battle.Teammate1Kills, &battle.Teammate1Specials, &battle.Teammate1Headgear,
		&battle.Teammate1HeadgearMain, &battle.Teammate1HeadgearSub0, &battle.Teammate1HeadgearSub1,
		&battle.Teammate1HeadgearSub2, &battle.Teammate1Clothes, &battle.Teammate1ClothesMain,
		&battle.Teammate1ClothesSub0, &battle.Teammate1ClothesSub1, &battle.Teammate1ClothesSub2,
		&battle.Teammate1Shoes, &battle.Teammate1ShoesMain, &battle.Teammate1ShoesSub0, &battle.Teammate1ShoesSub1,
		&battle.Teammate1ShoesSub2, &battle.Teammate2SplatnetId, &battle.Teammate2Name, &battle.Teammate2Rank,
		&battle.Teammate2LevelStar, &battle.Teammate2Level, &battle.Teammate2Weapon, &battle.Teammate2Gender,
		&battle.Teammate2Species, &battle.Teammate2Assists, &battle.Teammate2Deaths, &battle.Teammate2GamePaintPoint,
		&battle.Teammate2Kills, &battle.Teammate2Specials, &battle.Teammate2Headgear, &battle.Teammate2HeadgearMain,
		&battle.Teammate2HeadgearSub0, &battle.Teammate2HeadgearSub1, &battle.Teammate2HeadgearSub2,
		&battle.Teammate2Clothes, &battle.Teammate2ClothesMain, &battle.Teammate2ClothesSub0,
		&battle.Teammate2ClothesSub1, &battle.Teammate2ClothesSub2, &battle.Teammate2Shoes, &battle.Teammate2ShoesMain,
		&battle.Teammate2ShoesSub0, &battle.Teammate2ShoesSub1, &battle.Teammate2ShoesSub2,
		&battle.PlayerName, &battle.PlayerRank, &battle.PlayerLevelStar, &battle.PlayerLevel, &battle.PlayerWeapon,
		&battle.PlayerGender, &battle.PlayerSpecies, &battle.PlayerAssists, &battle.PlayerDeaths,
		&battle.PlayerGamePaintPoint, &battle.PlayerKills, &battle.PlayerSpecials, &battle.PlayerHeadgear,
		&battle.PlayerHeadgearMain, &battle.PlayerHeadgearSub0, &battle.PlayerHeadgearSub1, &battle.PlayerHeadgearSub2,
		&battle.PlayerClothes, &battle.PlayerClothesMain, &battle.PlayerClothesSub0, &battle.PlayerClothesSub1,
		&battle.PlayerClothesSub2, &battle.PlayerShoes, &battle.PlayerShoesMain, &battle.PlayerShoesSub0,
		&battle.PlayerShoesSub1, &battle.PlayerShoesSub2,
	); err != nil {
		return nil, err
	}
	var battleSplatnet *api_objects.BattleSplatnet
	if battle.SplatnetJson != nil {
		var err error
		battleSplatnet, err = readBattleSplatnet(*battle.SplatnetJson)
		if err != nil {
			return nil, err
		}
	}
	var battleStatInk *api_objects.BattleStatInk
	if battle.StatInkJson != nil {
		var err error
		battleStatInk, err = readBattleStatInk(*battle.StatInkJson)
		if err != nil {
			return nil, err
		}
	}
	hasDC := battle.HasDisconnectedPlayer != 0
	win := battle.Win != 0
	return &api_objects.Battle{
		UserId:                  battle.UserId,
		SplatnetJson:            battleSplatnet,
		StatInkJson:             battleStatInk,
		BattleNumber:            battle.SplatnetNumber,
		PlayerSplatnetId:        battle.PlayerSplatnetId,
		ElapsedTime:             battle.ElapsedTime,
		HasDisconnectedPlayer:   hasDC,
		LeaguePoint:             battle.LeaguePoint,
		MatchType:               battle.MatchType,
		Rule:                    battle.Rule,
		MyTeamCount:             battle.MyTeamCount,
		OtherTeamCount:          battle.OtherTeamCount,
		SplatfestPoint:          battle.SplatfestPoint,
		SplatfestTitle:          battle.SplatfestTitle,
		Stage:                   battle.Stage,
		TagId:                   battle.TagId,
		Time:                    battle.Time,
		Win:                     win,
		WinMeter:                battle.WinMeter,
		Opponent0SplatnetId:     battle.Opponent0SplatnetId,
		Opponent0Name:           battle.Opponent0Name,
		Opponent0Rank:           battle.Opponent0Rank,
		Opponent0LevelStar:      battle.Opponent0LevelStar,
		Opponent0Level:          battle.Opponent0Level,
		Opponent0Weapon:         battle.Opponent0Weapon,
		Opponent0Gender:         battle.Opponent0Gender,
		Opponent0Species:        battle.Opponent0Species,
		Opponent0Assists:        battle.Opponent0Assists,
		Opponent0Deaths:         battle.Opponent0Deaths,
		Opponent0GamePaintPoint: battle.Opponent0GamePaintPoint,
		Opponent0Kills:          battle.Opponent0Kills,
		Opponent0Specials:       battle.Opponent0Specials,
		Opponent0Headgear:       battle.Opponent0Headgear,
		Opponent0HeadgearMain:   battle.Opponent0HeadgearMain,
		Opponent0HeadgearSub0:   battle.Opponent0HeadgearSub0,
		Opponent0HeadgearSub1:   battle.Opponent0HeadgearSub1,
		Opponent0HeadgearSub2:   battle.Opponent0HeadgearSub2,
		Opponent0Clothes:        battle.Opponent0Clothes,
		Opponent0ClothesMain:    battle.Opponent0ClothesMain,
		Opponent0ClothesSub0:    battle.Opponent0ClothesSub0,
		Opponent0ClothesSub1:    battle.Opponent0ClothesSub1,
		Opponent0ClothesSub2:    battle.Opponent0ClothesSub2,
		Opponent0Shoes:          battle.Opponent0Shoes,
		Opponent0ShoesMain:      battle.Opponent0ShoesMain,
		Opponent0ShoesSub0:      battle.Opponent0ShoesSub0,
		Opponent0ShoesSub1:      battle.Opponent0ShoesSub1,
		Opponent0ShoesSub2:      battle.Opponent0ShoesSub2,
		Opponent1SplatnetId:     battle.Opponent1SplatnetId,
		Opponent1Name:           battle.Opponent1Name,
		Opponent1Rank:           battle.Opponent1Rank,
		Opponent1LevelStar:      battle.Opponent1LevelStar,
		Opponent1Level:          battle.Opponent1Level,
		Opponent1Weapon:         battle.Opponent1Weapon,
		Opponent1Gender:         battle.Opponent1Gender,
		Opponent1Species:        battle.Opponent1Species,
		Opponent1Assists:        battle.Opponent1Assists,
		Opponent1Deaths:         battle.Opponent1Deaths,
		Opponent1GamePaintPoint: battle.Opponent1GamePaintPoint,
		Opponent1Kills:          battle.Opponent1Kills,
		Opponent1Specials:       battle.Opponent1Specials,
		Opponent1Headgear:       battle.Opponent1Headgear,
		Opponent1HeadgearMain:   battle.Opponent1HeadgearMain,
		Opponent1HeadgearSub0:   battle.Opponent1HeadgearSub0,
		Opponent1HeadgearSub1:   battle.Opponent1HeadgearSub1,
		Opponent1HeadgearSub2:   battle.Opponent1HeadgearSub2,
		Opponent1Clothes:        battle.Opponent1Clothes,
		Opponent1ClothesMain:    battle.Opponent1ClothesMain,
		Opponent1ClothesSub0:    battle.Opponent1ClothesSub0,
		Opponent1ClothesSub1:    battle.Opponent1ClothesSub1,
		Opponent1ClothesSub2:    battle.Opponent1ClothesSub2,
		Opponent1Shoes:          battle.Opponent1Shoes,
		Opponent1ShoesMain:      battle.Opponent1ShoesMain,
		Opponent1ShoesSub0:      battle.Opponent1ShoesSub0,
		Opponent1ShoesSub1:      battle.Opponent1ShoesSub1,
		Opponent1ShoesSub2:      battle.Opponent1ShoesSub2,
		Opponent2SplatnetId:     battle.Opponent2SplatnetId,
		Opponent2Name:           battle.Opponent2Name,
		Opponent2Rank:           battle.Opponent2Rank,
		Opponent2LevelStar:      battle.Opponent2LevelStar,
		Opponent2Level:          battle.Opponent2Level,
		Opponent2Weapon:         battle.Opponent2Weapon,
		Opponent2Gender:         battle.Opponent2Gender,
		Opponent2Species:        battle.Opponent2Species,
		Opponent2Assists:        battle.Opponent2Assists,
		Opponent2Deaths:         battle.Opponent2Deaths,
		Opponent2GamePaintPoint: battle.Opponent2GamePaintPoint,
		Opponent2Kills:          battle.Opponent2Kills,
		Opponent2Specials:       battle.Opponent2Specials,
		Opponent2Headgear:       battle.Opponent2Headgear,
		Opponent2HeadgearMain:   battle.Opponent2HeadgearMain,
		Opponent2HeadgearSub0:   battle.Opponent2HeadgearSub0,
		Opponent2HeadgearSub1:   battle.Opponent2HeadgearSub1,
		Opponent2HeadgearSub2:   battle.Opponent2HeadgearSub2,
		Opponent2Clothes:        battle.Opponent2Clothes,
		Opponent2ClothesMain:    battle.Opponent2ClothesMain,
		Opponent2ClothesSub0:    battle.Opponent2ClothesSub0,
		Opponent2ClothesSub1:    battle.Opponent2ClothesSub1,
		Opponent2ClothesSub2:    battle.Opponent2ClothesSub2,
		Opponent2Shoes:          battle.Opponent2Shoes,
		Opponent2ShoesMain:      battle.Opponent2ShoesMain,
		Opponent2ShoesSub0:      battle.Opponent2ShoesSub0,
		Opponent2ShoesSub1:      battle.Opponent2ShoesSub1,
		Opponent2ShoesSub2:      battle.Opponent2ShoesSub2,
		Opponent3SplatnetId:     battle.Opponent3SplatnetId,
		Opponent3Name:           battle.Opponent3Name,
		Opponent3Rank:           battle.Opponent3Rank,
		Opponent3LevelStar:      battle.Opponent3LevelStar,
		Opponent3Level:          battle.Opponent3Level,
		Opponent3Weapon:         battle.Opponent3Weapon,
		Opponent3Gender:         battle.Opponent3Gender,
		Opponent3Species:        battle.Opponent3Species,
		Opponent3Assists:        battle.Opponent3Assists,
		Opponent3Deaths:         battle.Opponent3Deaths,
		Opponent3GamePaintPoint: battle.Opponent3GamePaintPoint,
		Opponent3Kills:          battle.Opponent3Kills,
		Opponent3Specials:       battle.Opponent3Specials,
		Opponent3Headgear:       battle.Opponent3Headgear,
		Opponent3HeadgearMain:   battle.Opponent3HeadgearMain,
		Opponent3HeadgearSub0:   battle.Opponent3HeadgearSub0,
		Opponent3HeadgearSub1:   battle.Opponent3HeadgearSub1,
		Opponent3HeadgearSub2:   battle.Opponent3HeadgearSub2,
		Opponent3Clothes:        battle.Opponent3Clothes,
		Opponent3ClothesMain:    battle.Opponent3ClothesMain,
		Opponent3ClothesSub0:    battle.Opponent3ClothesSub0,
		Opponent3ClothesSub1:    battle.Opponent3ClothesSub1,
		Opponent3ClothesSub2:    battle.Opponent3ClothesSub2,
		Opponent3Shoes:          battle.Opponent3Shoes,
		Opponent3ShoesMain:      battle.Opponent3ShoesMain,
		Opponent3ShoesSub0:      battle.Opponent3ShoesSub0,
		Opponent3ShoesSub1:      battle.Opponent3ShoesSub1,
		Opponent3ShoesSub2:      battle.Opponent3ShoesSub2,
		Teammate0SplatnetId:     battle.Teammate0SplatnetId,
		Teammate0Name:           battle.Teammate0Name,
		Teammate0Rank:           battle.Teammate0Rank,
		Teammate0LevelStar:      battle.Teammate0LevelStar,
		Teammate0Level:          battle.Teammate0Level,
		Teammate0Weapon:         battle.Teammate0Weapon,
		Teammate0Gender:         battle.Teammate0Gender,
		Teammate0Species:        battle.Teammate0Species,
		Teammate0Assists:        battle.Teammate0Assists,
		Teammate0Deaths:         battle.Teammate0Deaths,
		Teammate0GamePaintPoint: battle.Teammate0GamePaintPoint,
		Teammate0Kills:          battle.Teammate0Kills,
		Teammate0Specials:       battle.Teammate0Specials,
		Teammate0Headgear:       battle.Teammate0Headgear,
		Teammate0HeadgearMain:   battle.Teammate0HeadgearMain,
		Teammate0HeadgearSub0:   battle.Teammate0HeadgearSub0,
		Teammate0HeadgearSub1:   battle.Teammate0HeadgearSub1,
		Teammate0HeadgearSub2:   battle.Teammate0HeadgearSub2,
		Teammate0Clothes:        battle.Teammate0Clothes,
		Teammate0ClothesMain:    battle.Teammate0ClothesMain,
		Teammate0ClothesSub0:    battle.Teammate0ClothesSub0,
		Teammate0ClothesSub1:    battle.Teammate0ClothesSub1,
		Teammate0ClothesSub2:    battle.Teammate0ClothesSub2,
		Teammate0Shoes:          battle.Teammate0Shoes,
		Teammate0ShoesMain:      battle.Teammate0ShoesMain,
		Teammate0ShoesSub0:      battle.Teammate0ShoesSub0,
		Teammate0ShoesSub1:      battle.Teammate0ShoesSub1,
		Teammate0ShoesSub2:      battle.Teammate0ShoesSub2,
		Teammate1SplatnetId:     battle.Teammate1SplatnetId,
		Teammate1Name:           battle.Teammate1Name,
		Teammate1Rank:           battle.Teammate1Rank,
		Teammate1LevelStar:      battle.Teammate1LevelStar,
		Teammate1Level:          battle.Teammate1Level,
		Teammate1Weapon:         battle.Teammate1Weapon,
		Teammate1Gender:         battle.Teammate1Gender,
		Teammate1Species:        battle.Teammate1Species,
		Teammate1Assists:        battle.Teammate1Assists,
		Teammate1Deaths:         battle.Teammate1Deaths,
		Teammate1GamePaintPoint: battle.Teammate1GamePaintPoint,
		Teammate1Kills:          battle.Teammate1Kills,
		Teammate1Specials:       battle.Teammate1Specials,
		Teammate1Headgear:       battle.Teammate1Headgear,
		Teammate1HeadgearMain:   battle.Teammate1HeadgearMain,
		Teammate1HeadgearSub0:   battle.Teammate1HeadgearSub0,
		Teammate1HeadgearSub1:   battle.Teammate1HeadgearSub1,
		Teammate1HeadgearSub2:   battle.Teammate1HeadgearSub2,
		Teammate1Clothes:        battle.Teammate1Clothes,
		Teammate1ClothesMain:    battle.Teammate1ClothesMain,
		Teammate1ClothesSub0:    battle.Teammate1ClothesSub0,
		Teammate1ClothesSub1:    battle.Teammate1ClothesSub1,
		Teammate1ClothesSub2:    battle.Teammate1ClothesSub2,
		Teammate1Shoes:          battle.Teammate1Shoes,
		Teammate1ShoesMain:      battle.Teammate1ShoesMain,
		Teammate1ShoesSub0:      battle.Teammate1ShoesSub0,
		Teammate1ShoesSub1:      battle.Teammate1ShoesSub1,
		Teammate1ShoesSub2:      battle.Teammate1ShoesSub2,
		Teammate2SplatnetId:     battle.Teammate2SplatnetId,
		Teammate2Name:           battle.Teammate2Name,
		Teammate2Rank:           battle.Teammate2Rank,
		Teammate2LevelStar:      battle.Teammate2LevelStar,
		Teammate2Level:          battle.Teammate2Level,
		Teammate2Weapon:         battle.Teammate2Weapon,
		Teammate2Gender:         battle.Teammate2Gender,
		Teammate2Species:        battle.Teammate2Species,
		Teammate2Assists:        battle.Teammate2Assists,
		Teammate2Deaths:         battle.Teammate2Deaths,
		Teammate2GamePaintPoint: battle.Teammate2GamePaintPoint,
		Teammate2Kills:          battle.Teammate2Kills,
		Teammate2Specials:       battle.Teammate2Specials,
		Teammate2Headgear:       battle.Teammate2Headgear,
		Teammate2HeadgearMain:   battle.Teammate2HeadgearMain,
		Teammate2HeadgearSub0:   battle.Teammate2HeadgearSub0,
		Teammate2HeadgearSub1:   battle.Teammate2HeadgearSub1,
		Teammate2HeadgearSub2:   battle.Teammate2HeadgearSub2,
		Teammate2Clothes:        battle.Teammate2Clothes,
		Teammate2ClothesMain:    battle.Teammate2ClothesMain,
		Teammate2ClothesSub0:    battle.Teammate2ClothesSub0,
		Teammate2ClothesSub1:    battle.Teammate2ClothesSub1,
		Teammate2ClothesSub2:    battle.Teammate2ClothesSub2,
		Teammate2Shoes:          battle.Teammate2Shoes,
		Teammate2ShoesMain:      battle.Teammate2ShoesMain,
		Teammate2ShoesSub0:      battle.Teammate2ShoesSub0,
		Teammate2ShoesSub1:      battle.Teammate2ShoesSub1,
		Teammate2ShoesSub2:      battle.Teammate2ShoesSub2,
		PlayerName:              battle.PlayerName,
		PlayerRank:              battle.PlayerRank,
		PlayerLevelStar:         battle.PlayerLevelStar,
		PlayerLevel:             battle.PlayerLevel,
		PlayerWeapon:            battle.PlayerWeapon,
		PlayerGender:            battle.PlayerGender,
		PlayerSpecies:           battle.PlayerSpecies,
		PlayerAssists:           battle.PlayerAssists,
		PlayerDeaths:            battle.PlayerDeaths,
		PlayerGamePaintPoint:    battle.PlayerGamePaintPoint,
		PlayerKills:             battle.PlayerKills,
		PlayerSpecials:          battle.PlayerSpecials,
		PlayerHeadgear:          battle.PlayerHeadgear,
		PlayerHeadgearMain:      battle.PlayerHeadgearMain,
		PlayerHeadgearSub0:      battle.PlayerHeadgearSub0,
		PlayerHeadgearSub1:      battle.PlayerHeadgearSub1,
		PlayerHeadgearSub2:      battle.PlayerHeadgearSub2,
		PlayerClothes:           battle.PlayerClothes,
		PlayerClothesMain:       battle.PlayerClothesMain,
		PlayerClothesSub0:       battle.PlayerClothesSub0,
		PlayerClothesSub1:       battle.PlayerClothesSub1,
		PlayerClothesSub2:       battle.PlayerClothesSub2,
		PlayerShoes:             battle.PlayerShoes,
		PlayerShoesMain:         battle.PlayerShoesMain,
		PlayerShoesSub0:         battle.PlayerShoesSub0,
		PlayerShoesSub1:         battle.PlayerShoesSub1,
		PlayerShoesSub2:         battle.PlayerShoesSub2,
	}, nil
}

func GetBattleLeanPk(pk int64) (*api_objects.Battle, error) {
	var battle db_objects.Battle
	if err := readObjWithId(pk, "two_battles_battle").Scan(
		&pk, &battle.UserId, &battle.SplatnetJson, &battle.SplatnetUpload, &battle.StatInkJson, &battle.StatInkUpload,
		&battle.SplatnetNumber, &battle.PlayerSplatnetId, &battle.ElapsedTime, &battle.HasDisconnectedPlayer,
		&battle.LeaguePoint, &battle.MatchType, &battle.Rule, &battle.MyTeamCount, &battle.OtherTeamCount,
		&battle.SplatfestPoint, &battle.SplatfestTitle, &battle.Stage, &battle.TagId, &battle.Time, &battle.Win,
		&battle.WinMeter, &battle.Opponent0SplatnetId, &battle.Opponent0Name, &battle.Opponent0Rank,
		&battle.Opponent0LevelStar, &battle.Opponent0Level, &battle.Opponent0Weapon, &battle.Opponent0Gender,
		&battle.Opponent0Species, &battle.Opponent0Assists, &battle.Opponent0Deaths, &battle.Opponent0GamePaintPoint,
		&battle.Opponent0Kills, &battle.Opponent0Specials, &battle.Opponent0Headgear, &battle.Opponent0HeadgearMain,
		&battle.Opponent0HeadgearSub0, &battle.Opponent0HeadgearSub1, &battle.Opponent0HeadgearSub2,
		&battle.Opponent0Clothes, &battle.Opponent0ClothesMain, &battle.Opponent0ClothesSub0,
		&battle.Opponent0ClothesSub1, &battle.Opponent0ClothesSub2, &battle.Opponent0Shoes, &battle.Opponent0ShoesMain,
		&battle.Opponent0ShoesSub0, &battle.Opponent0ShoesSub1, &battle.Opponent0ShoesSub2, &battle.Opponent1SplatnetId,
		&battle.Opponent1Name, &battle.Opponent1Rank, &battle.Opponent1LevelStar, &battle.Opponent1Level,
		&battle.Opponent1Weapon, &battle.Opponent1Gender, &battle.Opponent1Species, &battle.Opponent1Assists,
		&battle.Opponent1Deaths, &battle.Opponent1GamePaintPoint, &battle.Opponent1Kills, &battle.Opponent1Specials,
		&battle.Opponent1Headgear, &battle.Opponent1HeadgearMain, &battle.Opponent1HeadgearSub0,
		&battle.Opponent1HeadgearSub1, &battle.Opponent1HeadgearSub2, &battle.Opponent1Clothes,
		&battle.Opponent1ClothesMain, &battle.Opponent1ClothesSub0, &battle.Opponent1ClothesSub1,
		&battle.Opponent1ClothesSub2, &battle.Opponent1Shoes, &battle.Opponent1ShoesMain, &battle.Opponent1ShoesSub0,
		&battle.Opponent1ShoesSub1, &battle.Opponent1ShoesSub2, &battle.Opponent2SplatnetId, &battle.Opponent2Name,
		&battle.Opponent2Rank, &battle.Opponent2LevelStar, &battle.Opponent2Level, &battle.Opponent2Weapon,
		&battle.Opponent2Gender, &battle.Opponent2Species, &battle.Opponent2Assists, &battle.Opponent2Deaths,
		&battle.Opponent2GamePaintPoint, &battle.Opponent2Kills, &battle.Opponent2Specials, &battle.Opponent2Headgear,
		&battle.Opponent2HeadgearMain, &battle.Opponent2HeadgearSub0, &battle.Opponent2HeadgearSub1,
		&battle.Opponent2HeadgearSub2, &battle.Opponent2Clothes, &battle.Opponent2ClothesMain,
		&battle.Opponent2ClothesSub0, &battle.Opponent2ClothesSub1, &battle.Opponent2ClothesSub2,
		&battle.Opponent2Shoes, &battle.Opponent2ShoesMain, &battle.Opponent2ShoesSub0, &battle.Opponent2ShoesSub1,
		&battle.Opponent2ShoesSub2, &battle.Opponent3SplatnetId, &battle.Opponent3Name, &battle.Opponent3Rank,
		&battle.Opponent3LevelStar, &battle.Opponent3Level, &battle.Opponent3Weapon, &battle.Opponent3Gender,
		&battle.Opponent3Species, &battle.Opponent3Assists, &battle.Opponent3Deaths, &battle.Opponent3GamePaintPoint,
		&battle.Opponent3Kills, &battle.Opponent3Specials, &battle.Opponent3Headgear, &battle.Opponent3HeadgearMain,
		&battle.Opponent3HeadgearSub0, &battle.Opponent3HeadgearSub1, &battle.Opponent3HeadgearSub2,
		&battle.Opponent3Clothes, &battle.Opponent3ClothesMain, &battle.Opponent3ClothesSub0,
		&battle.Opponent3ClothesSub1, &battle.Opponent3ClothesSub2, &battle.Opponent3Shoes, &battle.Opponent3ShoesMain,
		&battle.Opponent3ShoesSub0, &battle.Opponent3ShoesSub1, &battle.Opponent3ShoesSub2, &battle.Teammate0SplatnetId,
		&battle.Teammate0Name, &battle.Teammate0Rank, &battle.Teammate0LevelStar, &battle.Teammate0Level,
		&battle.Teammate0Weapon, &battle.Teammate0Gender, &battle.Teammate0Species, &battle.Teammate0Assists,
		&battle.Teammate0Deaths, &battle.Teammate0GamePaintPoint, &battle.Teammate0Kills, &battle.Teammate0Specials,
		&battle.Teammate0Headgear, &battle.Teammate0HeadgearMain, &battle.Teammate0HeadgearSub0,
		&battle.Teammate0HeadgearSub1, &battle.Teammate0HeadgearSub2, &battle.Teammate0Clothes,
		&battle.Teammate0ClothesMain, &battle.Teammate0ClothesSub0, &battle.Teammate0ClothesSub1,
		&battle.Teammate0ClothesSub2, &battle.Teammate0Shoes, &battle.Teammate0ShoesMain, &battle.Teammate0ShoesSub0,
		&battle.Teammate0ShoesSub1, &battle.Teammate0ShoesSub2, &battle.Teammate1SplatnetId, &battle.Teammate1Name,
		&battle.Teammate1Rank, &battle.Teammate1LevelStar, &battle.Teammate1Level, &battle.Teammate1Weapon,
		&battle.Teammate1Gender, &battle.Teammate1Species, &battle.Teammate1Assists, &battle.Teammate1Deaths,
		&battle.Teammate1GamePaintPoint, &battle.Teammate1Kills, &battle.Teammate1Specials, &battle.Teammate1Headgear,
		&battle.Teammate1HeadgearMain, &battle.Teammate1HeadgearSub0, &battle.Teammate1HeadgearSub1,
		&battle.Teammate1HeadgearSub2, &battle.Teammate1Clothes, &battle.Teammate1ClothesMain,
		&battle.Teammate1ClothesSub0, &battle.Teammate1ClothesSub1, &battle.Teammate1ClothesSub2,
		&battle.Teammate1Shoes, &battle.Teammate1ShoesMain, &battle.Teammate1ShoesSub0, &battle.Teammate1ShoesSub1,
		&battle.Teammate1ShoesSub2, &battle.Teammate2SplatnetId, &battle.Teammate2Name, &battle.Teammate2Rank,
		&battle.Teammate2LevelStar, &battle.Teammate2Level, &battle.Teammate2Weapon, &battle.Teammate2Gender,
		&battle.Teammate2Species, &battle.Teammate2Assists, &battle.Teammate2Deaths, &battle.Teammate2GamePaintPoint,
		&battle.Teammate2Kills, &battle.Teammate2Specials, &battle.Teammate2Headgear, &battle.Teammate2HeadgearMain,
		&battle.Teammate2HeadgearSub0, &battle.Teammate2HeadgearSub1, &battle.Teammate2HeadgearSub2,
		&battle.Teammate2Clothes, &battle.Teammate2ClothesMain, &battle.Teammate2ClothesSub0,
		&battle.Teammate2ClothesSub1, &battle.Teammate2ClothesSub2, &battle.Teammate2Shoes, &battle.Teammate2ShoesMain,
		&battle.Teammate2ShoesSub0, &battle.Teammate2ShoesSub1, &battle.Teammate2ShoesSub2,
		&battle.PlayerName, &battle.PlayerRank, &battle.PlayerLevelStar, &battle.PlayerLevel, &battle.PlayerWeapon,
		&battle.PlayerGender, &battle.PlayerSpecies, &battle.PlayerAssists, &battle.PlayerDeaths,
		&battle.PlayerGamePaintPoint, &battle.PlayerKills, &battle.PlayerSpecials, &battle.PlayerHeadgear,
		&battle.PlayerHeadgearMain, &battle.PlayerHeadgearSub0, &battle.PlayerHeadgearSub1, &battle.PlayerHeadgearSub2,
		&battle.PlayerClothes, &battle.PlayerClothesMain, &battle.PlayerClothesSub0, &battle.PlayerClothesSub1,
		&battle.PlayerClothesSub2, &battle.PlayerShoes, &battle.PlayerShoesMain, &battle.PlayerShoesSub0,
		&battle.PlayerShoesSub1, &battle.PlayerShoesSub2,
	); err != nil {
		return nil, err
	}
	hasDC := battle.HasDisconnectedPlayer != 0
	win := battle.Win != 0
	return &api_objects.Battle{
		UserId:                  battle.UserId,
		BattleNumber:            battle.SplatnetNumber,
		PlayerSplatnetId:        battle.PlayerSplatnetId,
		ElapsedTime:             battle.ElapsedTime,
		HasDisconnectedPlayer:   hasDC,
		LeaguePoint:             battle.LeaguePoint,
		MatchType:               battle.MatchType,
		Rule:                    battle.Rule,
		MyTeamCount:             battle.MyTeamCount,
		OtherTeamCount:          battle.OtherTeamCount,
		SplatfestPoint:          battle.SplatfestPoint,
		SplatfestTitle:          battle.SplatfestTitle,
		Stage:                   battle.Stage,
		TagId:                   battle.TagId,
		Time:                    battle.Time,
		Win:                     win,
		WinMeter:                battle.WinMeter,
		Opponent0SplatnetId:     battle.Opponent0SplatnetId,
		Opponent0Name:           battle.Opponent0Name,
		Opponent0Rank:           battle.Opponent0Rank,
		Opponent0LevelStar:      battle.Opponent0LevelStar,
		Opponent0Level:          battle.Opponent0Level,
		Opponent0Weapon:         battle.Opponent0Weapon,
		Opponent0Gender:         battle.Opponent0Gender,
		Opponent0Species:        battle.Opponent0Species,
		Opponent0Assists:        battle.Opponent0Assists,
		Opponent0Deaths:         battle.Opponent0Deaths,
		Opponent0GamePaintPoint: battle.Opponent0GamePaintPoint,
		Opponent0Kills:          battle.Opponent0Kills,
		Opponent0Specials:       battle.Opponent0Specials,
		Opponent0Headgear:       battle.Opponent0Headgear,
		Opponent0HeadgearMain:   battle.Opponent0HeadgearMain,
		Opponent0HeadgearSub0:   battle.Opponent0HeadgearSub0,
		Opponent0HeadgearSub1:   battle.Opponent0HeadgearSub1,
		Opponent0HeadgearSub2:   battle.Opponent0HeadgearSub2,
		Opponent0Clothes:        battle.Opponent0Clothes,
		Opponent0ClothesMain:    battle.Opponent0ClothesMain,
		Opponent0ClothesSub0:    battle.Opponent0ClothesSub0,
		Opponent0ClothesSub1:    battle.Opponent0ClothesSub1,
		Opponent0ClothesSub2:    battle.Opponent0ClothesSub2,
		Opponent0Shoes:          battle.Opponent0Shoes,
		Opponent0ShoesMain:      battle.Opponent0ShoesMain,
		Opponent0ShoesSub0:      battle.Opponent0ShoesSub0,
		Opponent0ShoesSub1:      battle.Opponent0ShoesSub1,
		Opponent0ShoesSub2:      battle.Opponent0ShoesSub2,
		Opponent1SplatnetId:     battle.Opponent1SplatnetId,
		Opponent1Name:           battle.Opponent1Name,
		Opponent1Rank:           battle.Opponent1Rank,
		Opponent1LevelStar:      battle.Opponent1LevelStar,
		Opponent1Level:          battle.Opponent1Level,
		Opponent1Weapon:         battle.Opponent1Weapon,
		Opponent1Gender:         battle.Opponent1Gender,
		Opponent1Species:        battle.Opponent1Species,
		Opponent1Assists:        battle.Opponent1Assists,
		Opponent1Deaths:         battle.Opponent1Deaths,
		Opponent1GamePaintPoint: battle.Opponent1GamePaintPoint,
		Opponent1Kills:          battle.Opponent1Kills,
		Opponent1Specials:       battle.Opponent1Specials,
		Opponent1Headgear:       battle.Opponent1Headgear,
		Opponent1HeadgearMain:   battle.Opponent1HeadgearMain,
		Opponent1HeadgearSub0:   battle.Opponent1HeadgearSub0,
		Opponent1HeadgearSub1:   battle.Opponent1HeadgearSub1,
		Opponent1HeadgearSub2:   battle.Opponent1HeadgearSub2,
		Opponent1Clothes:        battle.Opponent1Clothes,
		Opponent1ClothesMain:    battle.Opponent1ClothesMain,
		Opponent1ClothesSub0:    battle.Opponent1ClothesSub0,
		Opponent1ClothesSub1:    battle.Opponent1ClothesSub1,
		Opponent1ClothesSub2:    battle.Opponent1ClothesSub2,
		Opponent1Shoes:          battle.Opponent1Shoes,
		Opponent1ShoesMain:      battle.Opponent1ShoesMain,
		Opponent1ShoesSub0:      battle.Opponent1ShoesSub0,
		Opponent1ShoesSub1:      battle.Opponent1ShoesSub1,
		Opponent1ShoesSub2:      battle.Opponent1ShoesSub2,
		Opponent2SplatnetId:     battle.Opponent2SplatnetId,
		Opponent2Name:           battle.Opponent2Name,
		Opponent2Rank:           battle.Opponent2Rank,
		Opponent2LevelStar:      battle.Opponent2LevelStar,
		Opponent2Level:          battle.Opponent2Level,
		Opponent2Weapon:         battle.Opponent2Weapon,
		Opponent2Gender:         battle.Opponent2Gender,
		Opponent2Species:        battle.Opponent2Species,
		Opponent2Assists:        battle.Opponent2Assists,
		Opponent2Deaths:         battle.Opponent2Deaths,
		Opponent2GamePaintPoint: battle.Opponent2GamePaintPoint,
		Opponent2Kills:          battle.Opponent2Kills,
		Opponent2Specials:       battle.Opponent2Specials,
		Opponent2Headgear:       battle.Opponent2Headgear,
		Opponent2HeadgearMain:   battle.Opponent2HeadgearMain,
		Opponent2HeadgearSub0:   battle.Opponent2HeadgearSub0,
		Opponent2HeadgearSub1:   battle.Opponent2HeadgearSub1,
		Opponent2HeadgearSub2:   battle.Opponent2HeadgearSub2,
		Opponent2Clothes:        battle.Opponent2Clothes,
		Opponent2ClothesMain:    battle.Opponent2ClothesMain,
		Opponent2ClothesSub0:    battle.Opponent2ClothesSub0,
		Opponent2ClothesSub1:    battle.Opponent2ClothesSub1,
		Opponent2ClothesSub2:    battle.Opponent2ClothesSub2,
		Opponent2Shoes:          battle.Opponent2Shoes,
		Opponent2ShoesMain:      battle.Opponent2ShoesMain,
		Opponent2ShoesSub0:      battle.Opponent2ShoesSub0,
		Opponent2ShoesSub1:      battle.Opponent2ShoesSub1,
		Opponent2ShoesSub2:      battle.Opponent2ShoesSub2,
		Opponent3SplatnetId:     battle.Opponent3SplatnetId,
		Opponent3Name:           battle.Opponent3Name,
		Opponent3Rank:           battle.Opponent3Rank,
		Opponent3LevelStar:      battle.Opponent3LevelStar,
		Opponent3Level:          battle.Opponent3Level,
		Opponent3Weapon:         battle.Opponent3Weapon,
		Opponent3Gender:         battle.Opponent3Gender,
		Opponent3Species:        battle.Opponent3Species,
		Opponent3Assists:        battle.Opponent3Assists,
		Opponent3Deaths:         battle.Opponent3Deaths,
		Opponent3GamePaintPoint: battle.Opponent3GamePaintPoint,
		Opponent3Kills:          battle.Opponent3Kills,
		Opponent3Specials:       battle.Opponent3Specials,
		Opponent3Headgear:       battle.Opponent3Headgear,
		Opponent3HeadgearMain:   battle.Opponent3HeadgearMain,
		Opponent3HeadgearSub0:   battle.Opponent3HeadgearSub0,
		Opponent3HeadgearSub1:   battle.Opponent3HeadgearSub1,
		Opponent3HeadgearSub2:   battle.Opponent3HeadgearSub2,
		Opponent3Clothes:        battle.Opponent3Clothes,
		Opponent3ClothesMain:    battle.Opponent3ClothesMain,
		Opponent3ClothesSub0:    battle.Opponent3ClothesSub0,
		Opponent3ClothesSub1:    battle.Opponent3ClothesSub1,
		Opponent3ClothesSub2:    battle.Opponent3ClothesSub2,
		Opponent3Shoes:          battle.Opponent3Shoes,
		Opponent3ShoesMain:      battle.Opponent3ShoesMain,
		Opponent3ShoesSub0:      battle.Opponent3ShoesSub0,
		Opponent3ShoesSub1:      battle.Opponent3ShoesSub1,
		Opponent3ShoesSub2:      battle.Opponent3ShoesSub2,
		Teammate0SplatnetId:     battle.Teammate0SplatnetId,
		Teammate0Name:           battle.Teammate0Name,
		Teammate0Rank:           battle.Teammate0Rank,
		Teammate0LevelStar:      battle.Teammate0LevelStar,
		Teammate0Level:          battle.Teammate0Level,
		Teammate0Weapon:         battle.Teammate0Weapon,
		Teammate0Gender:         battle.Teammate0Gender,
		Teammate0Species:        battle.Teammate0Species,
		Teammate0Assists:        battle.Teammate0Assists,
		Teammate0Deaths:         battle.Teammate0Deaths,
		Teammate0GamePaintPoint: battle.Teammate0GamePaintPoint,
		Teammate0Kills:          battle.Teammate0Kills,
		Teammate0Specials:       battle.Teammate0Specials,
		Teammate0Headgear:       battle.Teammate0Headgear,
		Teammate0HeadgearMain:   battle.Teammate0HeadgearMain,
		Teammate0HeadgearSub0:   battle.Teammate0HeadgearSub0,
		Teammate0HeadgearSub1:   battle.Teammate0HeadgearSub1,
		Teammate0HeadgearSub2:   battle.Teammate0HeadgearSub2,
		Teammate0Clothes:        battle.Teammate0Clothes,
		Teammate0ClothesMain:    battle.Teammate0ClothesMain,
		Teammate0ClothesSub0:    battle.Teammate0ClothesSub0,
		Teammate0ClothesSub1:    battle.Teammate0ClothesSub1,
		Teammate0ClothesSub2:    battle.Teammate0ClothesSub2,
		Teammate0Shoes:          battle.Teammate0Shoes,
		Teammate0ShoesMain:      battle.Teammate0ShoesMain,
		Teammate0ShoesSub0:      battle.Teammate0ShoesSub0,
		Teammate0ShoesSub1:      battle.Teammate0ShoesSub1,
		Teammate0ShoesSub2:      battle.Teammate0ShoesSub2,
		Teammate1SplatnetId:     battle.Teammate1SplatnetId,
		Teammate1Name:           battle.Teammate1Name,
		Teammate1Rank:           battle.Teammate1Rank,
		Teammate1LevelStar:      battle.Teammate1LevelStar,
		Teammate1Level:          battle.Teammate1Level,
		Teammate1Weapon:         battle.Teammate1Weapon,
		Teammate1Gender:         battle.Teammate1Gender,
		Teammate1Species:        battle.Teammate1Species,
		Teammate1Assists:        battle.Teammate1Assists,
		Teammate1Deaths:         battle.Teammate1Deaths,
		Teammate1GamePaintPoint: battle.Teammate1GamePaintPoint,
		Teammate1Kills:          battle.Teammate1Kills,
		Teammate1Specials:       battle.Teammate1Specials,
		Teammate1Headgear:       battle.Teammate1Headgear,
		Teammate1HeadgearMain:   battle.Teammate1HeadgearMain,
		Teammate1HeadgearSub0:   battle.Teammate1HeadgearSub0,
		Teammate1HeadgearSub1:   battle.Teammate1HeadgearSub1,
		Teammate1HeadgearSub2:   battle.Teammate1HeadgearSub2,
		Teammate1Clothes:        battle.Teammate1Clothes,
		Teammate1ClothesMain:    battle.Teammate1ClothesMain,
		Teammate1ClothesSub0:    battle.Teammate1ClothesSub0,
		Teammate1ClothesSub1:    battle.Teammate1ClothesSub1,
		Teammate1ClothesSub2:    battle.Teammate1ClothesSub2,
		Teammate1Shoes:          battle.Teammate1Shoes,
		Teammate1ShoesMain:      battle.Teammate1ShoesMain,
		Teammate1ShoesSub0:      battle.Teammate1ShoesSub0,
		Teammate1ShoesSub1:      battle.Teammate1ShoesSub1,
		Teammate1ShoesSub2:      battle.Teammate1ShoesSub2,
		Teammate2SplatnetId:     battle.Teammate2SplatnetId,
		Teammate2Name:           battle.Teammate2Name,
		Teammate2Rank:           battle.Teammate2Rank,
		Teammate2LevelStar:      battle.Teammate2LevelStar,
		Teammate2Level:          battle.Teammate2Level,
		Teammate2Weapon:         battle.Teammate2Weapon,
		Teammate2Gender:         battle.Teammate2Gender,
		Teammate2Species:        battle.Teammate2Species,
		Teammate2Assists:        battle.Teammate2Assists,
		Teammate2Deaths:         battle.Teammate2Deaths,
		Teammate2GamePaintPoint: battle.Teammate2GamePaintPoint,
		Teammate2Kills:          battle.Teammate2Kills,
		Teammate2Specials:       battle.Teammate2Specials,
		Teammate2Headgear:       battle.Teammate2Headgear,
		Teammate2HeadgearMain:   battle.Teammate2HeadgearMain,
		Teammate2HeadgearSub0:   battle.Teammate2HeadgearSub0,
		Teammate2HeadgearSub1:   battle.Teammate2HeadgearSub1,
		Teammate2HeadgearSub2:   battle.Teammate2HeadgearSub2,
		Teammate2Clothes:        battle.Teammate2Clothes,
		Teammate2ClothesMain:    battle.Teammate2ClothesMain,
		Teammate2ClothesSub0:    battle.Teammate2ClothesSub0,
		Teammate2ClothesSub1:    battle.Teammate2ClothesSub1,
		Teammate2ClothesSub2:    battle.Teammate2ClothesSub2,
		Teammate2Shoes:          battle.Teammate2Shoes,
		Teammate2ShoesMain:      battle.Teammate2ShoesMain,
		Teammate2ShoesSub0:      battle.Teammate2ShoesSub0,
		Teammate2ShoesSub1:      battle.Teammate2ShoesSub1,
		Teammate2ShoesSub2:      battle.Teammate2ShoesSub2,
		PlayerName:              battle.PlayerName,
		PlayerRank:              battle.PlayerRank,
		PlayerLevelStar:         battle.PlayerLevelStar,
		PlayerLevel:             battle.PlayerLevel,
		PlayerWeapon:            battle.PlayerWeapon,
		PlayerGender:            battle.PlayerGender,
		PlayerSpecies:           battle.PlayerSpecies,
		PlayerAssists:           battle.PlayerAssists,
		PlayerDeaths:            battle.PlayerDeaths,
		PlayerGamePaintPoint:    battle.PlayerGamePaintPoint,
		PlayerKills:             battle.PlayerKills,
		PlayerSpecials:          battle.PlayerSpecials,
		PlayerHeadgear:          battle.PlayerHeadgear,
		PlayerHeadgearMain:      battle.PlayerHeadgearMain,
		PlayerHeadgearSub0:      battle.PlayerHeadgearSub0,
		PlayerHeadgearSub1:      battle.PlayerHeadgearSub1,
		PlayerHeadgearSub2:      battle.PlayerHeadgearSub2,
		PlayerClothes:           battle.PlayerClothes,
		PlayerClothesMain:       battle.PlayerClothesMain,
		PlayerClothesSub0:       battle.PlayerClothesSub0,
		PlayerClothesSub1:       battle.PlayerClothesSub1,
		PlayerClothesSub2:       battle.PlayerClothesSub2,
		PlayerShoes:             battle.PlayerShoes,
		PlayerShoesMain:         battle.PlayerShoesMain,
		PlayerShoesSub0:         battle.PlayerShoesSub0,
		PlayerShoesSub1:         battle.PlayerShoesSub1,
		PlayerShoesSub2:         battle.PlayerShoesSub2,
	}, nil
}

func readBattleSplatnet(pk int64) (*api_objects.BattleSplatnet, error) {
	battle := db_objects.BattleSplatnet{}
	if err := readObjWithId(pk, "two_battles_battle_splatnet").Scan(
		&pk, &battle.Udemae, &battle.Stage,
		&battle.OtherTeamCount, &battle.MyTeamCount, &battle.StarRank, &battle.Rule, &battle.PlayerResult,
		&battle.EstimateGachiPower, &battle.ElapsedTime, &battle.StartTime, &battle.GameMode,
		//&battle.XPower,
		&battle.BattleNumber, &battle.Type, &battle.PlayerRank,
		//&battle.CrownPlayers,
		&battle.WeaponPaintPoint,
		//&battle.Rank,
		&battle.MyTeamResult,
		//&battle.EstimateXPower,
		&battle.OtherTeamResult, &battle.LeaguePoint, &battle.WinMeter, &battle.MyTeamPercentage,
		&battle.OtherTeamPercentage, &battle.TagId,
	); err != nil {
		log.Printf("Called by readBattleSplatnet(%d)\n", pk)
		return nil, err
	}
	var udemae *api_objects.BattleSplatnetUdemae
	if battle.Udemae != nil {
		var err error
		udemae, err = readBattleSplatnetUdemae(*battle.Udemae)
		if err != nil {
			log.Printf("Called by readBattleSplatnet(%d)\n", pk)
			return nil, err
		}
	}
	teammatesKeys, err := readKeyArrayWithKeyCondition(pk, true, "parent", "my_team", "player_result", "two_battles_battle_splatnet_team_member")
	if err != nil {
		log.Printf("Called by readBattleSplatnet(%d)\n", pk)
		return nil, err
	}
	teammates := make([]api_objects.BattleSplatnetPlayerResult, len(teammatesKeys))
	for i := range teammatesKeys {
		teammateTemp, err := readBattleSplatnetPlayerResult(teammatesKeys[i])
		if err != nil {
			log.Printf("Called by readBattleSplatnet(%d)\n", pk)
			return nil, err
		}
		teammates[i] = *teammateTemp
	}
	opponentsKeys, err := readKeyArrayWithKeyCondition(pk, false, "parent", "my_team", "player_result", "two_battles_battle_splatnet_team_member")
	if err != nil {
		log.Printf("Called by readBattleSplatnet(%d)\n", pk)
		return nil, err
	}
	opponents := make([]api_objects.BattleSplatnetPlayerResult, len(opponentsKeys))
	for i := range opponentsKeys {
		opponentTemp, err := readBattleSplatnetPlayerResult(opponentsKeys[i])
		if err != nil {
			log.Printf("Called by readBattleSplatnet(%d)\n", pk)
			return nil, err
		}
		opponents[i] = *opponentTemp
	}
	playerResult, err := readBattleSplatnetPlayerResult(battle.PlayerResult)
	if err != nil {
		log.Printf("Called by readBattleSplatnet(%d)\n", pk)
		return nil, err
	}
	rule, err := readBattleSplatnetRule(battle.Rule)
	if err != nil {
		log.Printf("Called by readBattleSplatnet(%d)\n", pk)
		return nil, err
	}
	stage, err := readSplatnetTriple(battle.Stage)
	if err != nil {
		log.Printf("Called by readBattleSplatnet(%d)\n", pk)
		return nil, err
	}
	gameMode, err := getSplatnetDouble(battle.GameMode)
	if err != nil {
		log.Printf("Called by readBattleSplatnet(%d)\n", pk)
		return nil, err
	}
	myTeamResult, err := getSplatnetDouble(battle.MyTeamResult)
	if err != nil {
		log.Printf("Called by readBattleSplatnet(%d)\n", pk)
		return nil, err
	}
	otherTeamResult, err := getSplatnetDouble(battle.OtherTeamResult)
	if err != nil {
		log.Printf("Called by readBattleSplatnet(%d)\n", pk)
		return nil, err
	}
	return &api_objects.BattleSplatnet{
		Udemae:             udemae,
		Stage:              *stage,
		OtherTeamCount:     battle.OtherTeamCount,
		MyTeamCount:        battle.MyTeamCount,
		StarRank:           battle.StarRank,
		Rule:               *rule,
		PlayerResult:       *playerResult,
		EstimateGachiPower: battle.EstimateGachiPower,
		ElapsedTime:        battle.ElapsedTime,
		StartTime:          battle.StartTime,
		GameMode:           *gameMode,
		//XPower:              battle.XPower,
		BattleNumber: battle.BattleNumber,
		Type:         battle.Type,
		PlayerRank:   battle.PlayerRank,
		//CrownPlayers:        battle.CrownPlayers,
		MyTeamMembers:    teammates,
		OtherTeamMembers: opponents,
		WeaponPaintPoint: battle.WeaponPaintPoint,
		//Rank:                battle.Rank,
		MyTeamResult: *myTeamResult,
		//EstimateXPower:      battle.EstimateXPower,
		OtherTeamResult:     *otherTeamResult,
		LeaguePoint:         battle.LeaguePoint,
		WinMeter:            battle.WinMeter,
		MyTeamPercentage:    battle.MyTeamPercentage,
		OtherTeamPercentage: battle.OtherTeamPercentage,
		TagId:               battle.TagId,
	}, nil
}

func readBattleSplatnetUdemae(pk int64) (*api_objects.BattleSplatnetUdemae, error) {
	udemae := api_objects.BattleSplatnetUdemae{}
	var junk int
	if err := readObjWithId(pk, "two_battles_battle_splatnet_udemae").Scan(&junk, &udemae.Name, &udemae.IsX,
		&udemae.IsNumberReached, &udemae.Number, &udemae.SPlusNumber); err != nil {
		return nil, err
	}
	return &udemae, nil
}

func readSplatnetTriple(pk int64) (*api_objects.BattleSplatnetTriple, error) {
	triple := api_objects.BattleSplatnetTriple{}
	var junk int
	if err := readObjWithId(pk, "splatnet_triple").Scan(
		&junk, &triple.Id, &triple.Image, &triple.Name,
	); err != nil {
		return nil, err
	}
	return &triple, nil
}

func readBattleSplatnetRule(pk int64) (*api_objects.BattleSplatnetRule, error) {
	rule := api_objects.BattleSplatnetRule{}
	var junk int
	if err := readObjWithId(pk, "two_battles_battle_splatnet_rule").Scan(
		&junk, &rule.Key, &rule.Name, &rule.MultilineName,
	); err != nil {
		return nil, err
	}
	return &rule, nil
}

func readBattleSplatnetPlayerResult(pk int64) (*api_objects.BattleSplatnetPlayerResult, error) {
	result := db_objects.BattleSplatnetPlayerResult{}
	if err := readObjWithId(pk, "two_battles_battle_splatnet_player_result").Scan(
		&pk, &result.DeathCount, &result.GamePaintPoint, &result.KillCount, &result.SpecialCount,
		&result.AssistCount, &result.SortScore, &result.Player,
	); err != nil {
		return nil, err
	}
	player, err := readBattleSplatnetPlayerResultPlayer(result.Player)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleSplatnetPlayerResult{
		DeathCount:     result.DeathCount,
		GamePaintPoint: result.GamePaintPoint,
		KillCount:      result.KillCount,
		SpecialCount:   result.SpecialCount,
		AssistCount:    result.AssistCount,
		SortScore:      result.SortScore,
		Player:         *player,
	}, nil
}

func readBattleSplatnetPlayerResultPlayer(pk int64) (*api_objects.BattleSplatnetPlayerResultPlayer, error) {
	player := db_objects.BattleSplatnetPlayerResultPlayer{}
	if err := readObjWithId(pk, "two_battles_battle_splatnet_player_result_player").Scan(
		&pk, &player.HeadSkills, &player.ShoesSkills, &player.ClothesSkills, &player.PlayerRank,
		&player.StarRank, &player.Nickname, &player.PlayerType, &player.PrincipalId, &player.Head, &player.Clothes,
		&player.Shoes, &player.Udemae, &player.Weapon,
	); err != nil {
		return nil, err
	}
	var udemae *api_objects.BattleSplatnetUdemae
	if player.Udemae != nil {
		var err error
		udemae, err = readBattleSplatnetUdemae(*player.Udemae)
		if err != nil {
			return nil, err
		}
	}
	head, err := readBattleSplatnetPlayerResultPlayerClothing(player.Head)
	if err != nil {
		return nil, err
	}
	clothes, err := readBattleSplatnetPlayerResultPlayerClothing(player.Clothes)
	if err != nil {
		return nil, err
	}
	shoes, err := readBattleSplatnetPlayerResultPlayerClothing(player.Shoes)
	if err != nil {
		return nil, err
	}
	weapon, err := readBattleSplatnetPlayerResultPlayerWeapon(player.Weapon)
	if err != nil {
		return nil, err
	}
	playerType, err := getSplatnetPlayerType(player.PlayerType)
	if err != nil {
		return nil, err
	}
	headSkills, err := readBattleSplatnetPlayerResultPlayerSkills(player.HeadSkills)
	if err != nil {
		return nil, err
	}
	shoesSkills, err := readBattleSplatnetPlayerResultPlayerSkills(player.ShoesSkills)
	if err != nil {
		return nil, err
	}
	clothesSkills, err := readBattleSplatnetPlayerResultPlayerSkills(player.ClothesSkills)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleSplatnetPlayerResultPlayer{
		HeadSkills:    *headSkills,
		ShoesSkills:   *shoesSkills,
		ClothesSkills: *clothesSkills,
		PlayerRank:    player.PlayerRank,
		StarRank:      player.StarRank,
		Nickname:      player.Nickname,
		PlayerType:    *playerType,
		PrincipalId:   player.PrincipalId,
		Head:          *head,
		Clothes:       *clothes,
		Shoes:         *shoes,
		Udemae:        udemae,
		Weapon:        *weapon,
	}, nil
}

func readBattleSplatnetPlayerResultPlayerSkills(pk int64) (*api_objects.BattleSplatnetPlayerResultPlayerSkills, error) {
	skills := db_objects.BattleSplatnetPlayerResultPlayerSkills{}
	if err := readObjWithId(pk, "two_battles_battle_splatnet_player_result_player_skills").Scan(
		&pk, &skills.Main,
	); err != nil {
		return nil, err
	}
	subKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "sub", "two_battles_battle_splatnet_player_result_player_skills_sub", "asc", "pk")
	if err != nil {
		return nil, err
	}
	subs := make([]api_objects.BattleSplatnetTriple, len(subKeys))
	for i := range subKeys {
		subTemp, err := readSplatnetTriple(subKeys[i])
		if err != nil {
			return nil, err
		}
		subs[i] = *subTemp
	}
	main, err := readSplatnetTriple(skills.Main)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleSplatnetPlayerResultPlayerSkills{
		Main: *main,
		Subs: subs,
	}, nil
}

func getSplatnetPlayerType(pk int64) (*api_objects.SplatnetPlayerType, error) {
	playerType := api_objects.SplatnetPlayerType{}
	var junk int
	if err := readObjWithId(pk, "splatnet_player_type").Scan(
		&junk, &playerType.Gender, &playerType.Species,
	); err != nil {
		log.Printf("Called by getSplatnetPlayerType(%d)\n", pk)
		return nil, err
	}
	return &playerType, nil
}

func readBattleSplatnetPlayerResultPlayerClothing(pk int64) (*api_objects.BattleSplatnetPlayerResultPlayerClothing, error) {
	clothing := db_objects.BattleSplatnetPlayerResultPlayerClothing{}
	if err := readObjWithId(pk, "two_battles_battle_splatnet_player_result_player_clothing").Scan(
		&pk, &clothing.Id, &clothing.Image, &clothing.Name, &clothing.Thumbnail, &clothing.Kind,
		&clothing.Rarity, &clothing.Brand,
	); err != nil {
		return nil, err
	}
	brand, err := readBattleSplatnetPlayerResultPlayerClothingBrand(clothing.Brand)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleSplatnetPlayerResultPlayerClothing{
		Id:        clothing.Id,
		Image:     clothing.Image,
		Name:      clothing.Name,
		Thumbnail: clothing.Thumbnail,
		Kind:      clothing.Kind,
		Rarity:    clothing.Rarity,
		Brand:     *brand,
	}, nil
}

func readBattleSplatnetPlayerResultPlayerClothingBrand(pk int64) (*api_objects.BattleSplatnetPlayerResultPlayerClothingBrand, error) {
	brand := db_objects.BattleSplatnetPlayerResultPlayerClothingBrand{}
	if err := readObjWithId(pk, "two_battles_battle_splatnet_player_result_player_clothing_brand").Scan(
		&pk, &brand.Id, &brand.Image, &brand.Name, &brand.FrequentSkill,
	); err != nil {
		return nil, err
	}
	frequentSkill, err := readSplatnetTriple(brand.FrequentSkill)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleSplatnetPlayerResultPlayerClothingBrand{
		Id:            brand.Id,
		Image:         brand.Image,
		Name:          brand.Name,
		FrequentSkill: *frequentSkill,
	}, nil
}

func readBattleSplatnetPlayerResultPlayerWeapon(pk int64) (*api_objects.BattleSplatnetPlayerResultPlayerWeapon, error) {
	weapon := db_objects.BattleSplatnetPlayerResultPlayerWeapon{}
	if err := readObjWithId(pk, "two_battles_battle_splatnet_player_result_player_weapon").Scan(
		&pk, &weapon.Id, &weapon.Image, &weapon.Name, &weapon.Thumbnail, &weapon.Sub, &weapon.Special,
	); err != nil {
		return nil, err
	}
	sub, err := getSplatnetQuad(weapon.Sub)
	if err != nil {
		return nil, err
	}
	special, err := getSplatnetQuad(weapon.Special)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleSplatnetPlayerResultPlayerWeapon{
		Id:        weapon.Id,
		Image:     weapon.Image,
		Name:      weapon.Name,
		Thumbnail: weapon.Thumbnail,
		Sub:       *sub,
		Special:   *special,
	}, nil
}

func getSplatnetQuad(pk int64) (*api_objects.SplatnetQuad, error) {
	quad := api_objects.SplatnetQuad{}
	var junk int
	if err := readObjWithId(pk, "splatnet_quad").Scan(
		&junk, &quad.Id, &quad.ImageA, &quad.ImageB, &quad.Name,
	); err != nil {
		return nil, err
	}
	return &quad, nil
}

func getSplatnetDouble(pk int64) (*api_objects.SplatnetDouble, error) {
	double := api_objects.SplatnetDouble{}
	var junk int
	if err := readObjWithId(pk, "splatnet_double").Scan(
		&junk, &double.Key, &double.Name,
	); err != nil {
		return nil, err
	}
	return &double, nil
}

func readBattleStatInk(pk int64) (*api_objects.BattleStatInk, error) {
	battle := db_objects.BattleStatInk{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink").Scan(
		&pk, &battle.Id, &battle.SplatnetNumber, &battle.Url, &battle.User, &battle.Lobby,
		&battle.Mode, &battle.Rule, &battle.Map, &battle.Weapon, &battle.Freshness, &battle.Rank, &battle.RankExp,
		&battle.RankAfter,
		//&battle.XPower, &battle.XPowerAfter, &battle.EstimateXPower,
		&battle.Level,
		&battle.LevelAfter, &battle.StarRank, &battle.Result, &battle.KnockOut, &battle.RankInTeam, &battle.Kill,
		&battle.Death, &battle.KillOrAssist, &battle.Special, &battle.KillRatio, &battle.KillRate,
		//&battle.MaxKillCombo, &battle.MaxKillStreak, &battle.DeathReasons,
		&battle.MyPoint, &battle.EstimateGachiPower, &battle.LeaguePoint,
		&battle.MyTeamEstimateLeaguePoint, &battle.HisTeamEstimateLeaguePoint,
		//&battle.MyTeamPoint, &battle.HisTeamPoint,
		&battle.MyTeamPercent, &battle.HisTeamPercent, &battle.MyTeamId, &battle.HisTeamId,
		&battle.Species, &battle.Gender, &battle.FestTitle, &battle.FestExp, &battle.FestTitleAfter,
		&battle.FestExpAfter, &battle.FestPower, &battle.MyTeamEstimateFestPower,
		&battle.HisTeamMyTeamEstimateFestPower, &battle.MyTeamFestTheme, &battle.MyTeamNickname,
		&battle.HisTeamNickname, &battle.Clout, &battle.TotalClout, &battle.TotalCloutAfter, &battle.MyTeamWinStreak,
		&battle.HisTeamWinStreak, &battle.SynergyBonus, &battle.SpecialBattle, &battle.ImageResult, &battle.ImageGear,
		&battle.Gears, &battle.Period, &battle.PeriodRange,
		//&battle.Events, &battle.SplatnetJson,
		&battle.Agent, &battle.Automated,
		//&battle.Environment,
		&battle.LinkUrl,
		//&battle.Note,
		&battle.GameVersion,
		&battle.NawabariBonus, &battle.StartAt, &battle.EndAt, &battle.RegisterAt,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	var freshness *api_objects.BattleStatInkFreshness
	if battle.Freshness != nil {
		var err error
		freshness, err = readBattleStatInkFreshness(*battle.Freshness)
		if err != nil {

			return nil, err
		}
	}
	var rank *api_objects.BattleStatInkRank
	if battle.Rank != nil {
		var err error
		rank, err = readBattleStatInkRank(*battle.Rank)
		if err != nil {

			return nil, err
		}
	}
	var rankAfter *api_objects.BattleStatInkRank
	if battle.RankAfter != nil {
		var err error
		rankAfter, err = readBattleStatInkRank(*battle.RankAfter)
		if err != nil {

			return nil, err
		}
	}
	var festTitle *api_objects.StatInkKeyName
	if battle.FestTitle != nil {
		var err error
		festTitle, err = readStatInkKeyName(*battle.FestTitle)
		if err != nil {

			return nil, err
		}
	}
	var festTitleAfter *api_objects.StatInkKeyName
	if battle.FestTitleAfter != nil {
		var err error
		festTitleAfter, err = readStatInkKeyName(*battle.FestTitleAfter)
		if err != nil {

			return nil, err
		}
	}
	var specialBattle *api_objects.StatInkKeyName
	if battle.SpecialBattle != nil {
		var err error
		specialBattle, err = readStatInkKeyName(*battle.SpecialBattle)
		if err != nil {

			return nil, err
		}
	}
	playerKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "pk", "two_battles_battle_stat_ink_player", "asc", "pk")
	if err != nil {

		return nil, err
	}
	players := make([]api_objects.BattleStatInkPlayer, len(playerKeys))
	for i := range playerKeys {
		playerTemp, err := readBattleStatInkPlayer(playerKeys[i])
		if err != nil {

			return nil, err
		}
		players[i] = *playerTemp
	}
	agent, err := readBattleStatInkAgent(battle.Agent)
	if err != nil {

		return nil, err
	}
	user, err := readBattleStatInkUser(battle.User)
	if err != nil {

		return nil, err
	}
	startAt, err := readStatInkTime(battle.StartAt)
	if err != nil {

		return nil, err
	}
	endAt, err := readStatInkTime(battle.EndAt)
	if err != nil {

		return nil, err
	}
	registerAt, err := readStatInkTime(battle.RegisterAt)
	if err != nil {

		return nil, err
	}
	stage, err := readBattleStatInkMap(battle.Map)
	if err != nil {

		return nil, err
	}
	species, err := readStatInkKeyName(battle.Species)
	if err != nil {

		return nil, err
	}
	lobby, err := readStatInkKeyName(battle.Lobby)
	if err != nil {

		return nil, err
	}
	mode, err := readStatInkKeyName(battle.Mode)
	if err != nil {

		return nil, err
	}
	rule, err := readStatInkKeyName(battle.Rule)
	if err != nil {

		return nil, err
	}
	weapon, err := readBattleStatInkWeapon(battle.Weapon)
	if err != nil {

		return nil, err
	}
	gender, err := readStatInkGender(battle.Gender)
	if err != nil {

		return nil, err
	}
	gears, err := readBattleStatInkGears(battle.Gears)
	if err != nil {

		return nil, err
	}
	return &api_objects.BattleStatInk{
		Id:             battle.Id,
		SplatnetNumber: battle.SplatnetNumber,
		Url:            battle.Url,
		User:           *user,
		Lobby:          *lobby,
		Mode:           *mode,
		Rule:           *rule,
		Map:            *stage,
		Weapon:         *weapon,
		Freshness:      freshness,
		Rank:           rank,
		RankExp:        battle.RankExp,
		RankAfter:      rankAfter,
		//XPower:                         battle.XPower,
		//XPowerAfter:                    battle.XPowerAfter,
		//EstimateXPower:                 battle.EstimateXPower,
		Level:        battle.Level,
		LevelAfter:   battle.LevelAfter,
		StarRank:     battle.StarRank,
		Result:       battle.Result,
		KnockOut:     battle.KnockOut,
		RankInTeam:   battle.RankInTeam,
		Kill:         battle.Kill,
		Death:        battle.Death,
		KillOrAssist: battle.KillOrAssist,
		Special:      battle.Special,
		KillRatio:    battle.KillRatio,
		KillRate:     battle.KillRate,
		//MaxKillCombo:                   battle.MaxKillCombo,
		//MaxKillStreak:                  battle.MaxKillStreak,
		//DeathReasons:                   battle.DeathReasons,
		MyPoint:                    battle.MyPoint,
		EstimateGachiPower:         battle.EstimateGachiPower,
		LeaguePoint:                battle.LeaguePoint,
		MyTeamEstimateLeaguePoint:  battle.MyTeamEstimateLeaguePoint,
		HisTeamEstimateLeaguePoint: battle.HisTeamEstimateLeaguePoint,
		//MyTeamPoint:                    battle.MyTeamPoint,
		//HisTeamPoint:                   battle.HisTeamPoint,
		MyTeamPercent:                  battle.MyTeamPercent,
		HisTeamPercent:                 battle.HisTeamPercent,
		MyTeamId:                       battle.MyTeamId,
		HisTeamId:                      battle.HisTeamId,
		Species:                        *species,
		Gender:                         *gender,
		FestTitle:                      festTitle,
		FestExp:                        battle.FestExp,
		FestTitleAfter:                 festTitleAfter,
		FestExpAfter:                   battle.FestExpAfter,
		FestPower:                      battle.FestPower,
		MyTeamEstimateFestPower:        battle.MyTeamEstimateFestPower,
		HisTeamMyTeamEstimateFestPower: battle.HisTeamMyTeamEstimateFestPower,
		MyTeamFestTheme:                battle.MyTeamFestTheme,
		MyTeamNickname:                 battle.MyTeamNickname,
		HisTeamNickname:                battle.HisTeamNickname,
		Clout:                          battle.Clout,
		TotalClout:                     battle.TotalClout,
		TotalCloutAfter:                battle.TotalCloutAfter,
		MyTeamWinStreak:                battle.MyTeamWinStreak,
		HisTeamWinStreak:               battle.HisTeamWinStreak,
		SynergyBonus:                   battle.SynergyBonus,
		SpecialBattle:                  specialBattle,
		ImageResult:                    battle.ImageResult,
		ImageGear:                      battle.ImageGear,
		Gears:                          *gears,
		Period:                         battle.Period,
		PeriodRange:                    battle.PeriodRange,
		Players:                        players,
		//Events:                         battle.Events,
		//SplatnetJson:                   battle.SplatnetJson,
		Agent:     *agent,
		Automated: battle.Automated,
		//Environment:                    battle.Environment,
		LinkUrl: battle.LinkUrl,
		//Note:                           battle.Note,
		GameVersion:   battle.GameVersion,
		NawabariBonus: battle.NawabariBonus,
		StartAt:       *startAt,
		EndAt:         *endAt,
		RegisterAt:    *registerAt,
	}, nil
}

func readBattleStatInkUser(pk int64) (*api_objects.BattleStatInkUser, error) {
	var user db_objects.BattleStatInkUser
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_user").Scan(
		&pk, &user.Id, &user.Name, &user.ScreenName, &user.Url, &user.JoinAt, &user.Profile,
		//&user.Stat,
		&user.Stats,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	profile, err := readStatInkProfile(user.Profile)
	if err != nil {

		return nil, err
	}
	stats, err := readBattleStatInkUserStats(user.Stats)
	if err != nil {

		return nil, err
	}
	joinAt, err := readStatInkTime(user.JoinAt)
	if err != nil {

		return nil, err
	}
	return &api_objects.BattleStatInkUser{
		Id:         user.Id,
		Name:       user.Name,
		ScreenName: user.ScreenName,
		Url:        user.Url,
		JoinAt:     *joinAt,
		Profile:    *profile,
		//Stat:       user.Stat,
		Stats: *stats,
	}, nil
}

func readStatInkProfile(pk int64) (*api_objects.StatInkProfile, error) {
	profile := api_objects.StatInkProfile{}
	var junk int
	if err := readObjWithId(pk, "stat_ink_profile").Scan(
		&junk,
		//&profile.Nnid,
		&profile.FriendCode, &profile.Twitter,
		//&profile.Ikanakama, &profile.Ikanakama2, &profile.Environment,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &profile, nil
}

func readBattleStatInkUserStats(pk int64) (*api_objects.BattleStatInkUserStats, error) {
	var stats db_objects.BattleStatInkUserStats
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_user_stats").Scan(
		&pk,
		//&stats.V1,
		&stats.V2,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	var v2 *api_objects.BattleStatInkUserStatsV2 = nil
	if stats.V2 != nil {
		var err error
		v2, err = readBattleStatInkUserStatsV2(*stats.V2)
		if err != nil {

			return nil, err
		}
	}
	return &api_objects.BattleStatInkUserStats{
		//V1: stats.V1,
		V2: v2,
	}, nil
}

func readStatInkTime(pk int64) (*api_objects.StatInkTime, error) {
	timeObj := api_objects.StatInkTime{}
	row := readObjWithId(pk, "stat_ink_time")
	err := (*row).Scan(&pk, &timeObj.Time, &timeObj.Iso8601)
	if err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &timeObj, nil
}

func readBattleStatInkUserStatsV2(pk int64) (*api_objects.BattleStatInkUserStatsV2, error) {
	v2 := db_objects.BattleStatInkUserStatsV2{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_user_stats_v2").Scan(
		&pk, &v2.UpdatedAt, &v2.Entire, &v2.Nawabari, &v2.Gachi,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	updatedAt, err := readStatInkTime(v2.UpdatedAt)
	if err != nil {
		return nil, err
	}
	entire, err := readBattleStatInkUserStatsV2Entire(v2.Entire)
	if err != nil {
		return nil, err
	}
	nawabari, err := readBattleStatInkUserStatsV2Nawabari(v2.Nawabari)
	if err != nil {
		return nil, err
	}
	gachi, err := readBattleStatInkUserStatsV2Gachi(v2.Gachi)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkUserStatsV2{
		UpdatedAt: *updatedAt,
		Entire:    *entire,
		Nawabari:  *nawabari,
		Gachi:     *gachi,
	}, nil
}

func readBattleStatInkUserStatsV2Entire(pk int64) (*api_objects.BattleStatInkUserStatsV2Entire, error) {
	entire := api_objects.BattleStatInkUserStatsV2Entire{}
	var junk int
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_user_stats_v2_entire").Scan(
		&junk, &entire.Battles, &entire.WinPct, &entire.KillRatio, &entire.KillTotal,
		&entire.KillAvg, &entire.KillPerMin, &entire.DeathTotal, &entire.DeathAvg, &entire.DeathPerMin,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &entire, nil
}

func readBattleStatInkUserStatsV2Nawabari(pk int64) (*api_objects.BattleStatInkUserStatsV2Nawabari, error) {
	nawabari := api_objects.BattleStatInkUserStatsV2Nawabari{}
	var junk int
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_user_stats_v2_nawabari").Scan(
		&junk, &nawabari.Battles, &nawabari.WinPct, &nawabari.KillRatio, &nawabari.KillTotal,
		&nawabari.KillAvg, &nawabari.KillPerMin, &nawabari.DeathTotal, &nawabari.DeathAvg, &nawabari.DeathPerMin,
		&nawabari.TotalInked, &nawabari.MaxInked, &nawabari.AvgInked,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &nawabari, nil
}

func readBattleStatInkUserStatsV2Gachi(pk int64) (*api_objects.BattleStatInkUserStatsV2Gachi, error) {
	gachi := db_objects.BattleStatInkUserStatsV2Gachi{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_user_stats_v2_gachi").Scan(
		&pk, &gachi.Battles, &gachi.WinPct, &gachi.KillRatio, &gachi.KillTotal, &gachi.KillAvg, &gachi.KillPerMin,
		&gachi.DeathTotal, &gachi.DeathAvg, &gachi.DeathPerMin, &gachi.Rules,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	rules, err := readBattleStatInkUserStatsV2GachiRules(gachi.Rules)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkUserStatsV2Gachi{
		Battles:     gachi.Battles,
		WinPct:      gachi.WinPct,
		KillRatio:   gachi.KillRatio,
		KillTotal:   gachi.KillTotal,
		KillAvg:     gachi.KillAvg,
		KillPerMin:  gachi.KillPerMin,
		DeathTotal:  gachi.DeathTotal,
		DeathAvg:    gachi.DeathAvg,
		DeathPerMin: gachi.DeathPerMin,
		Rules:       *rules,
	}, nil
}

func readStatInkKeyName(pk int64) (*api_objects.StatInkKeyName, error) {
	keyName := db_objects.StatInkKeyName{}
	if err := readObjWithId(pk, "stat_ink_key_name").Scan(
		&pk, &keyName.Key, &keyName.Name,
	); err != nil {
		debug.PrintStack()
		log.Printf("pk: %d\n", pk)
		return nil, err
	}
	name, err := readStatInkName(keyName.Name)
	if err != nil {
		return nil, err
	}
	return &api_objects.StatInkKeyName{
		Key:  keyName.Key,
		Name: *name,
	}, nil
}

func readShiftStatInkFailReason(pk int64) (*api_objects.ShiftStatInkFailReason, error) {
	keyName := db_objects.StatInkKeyName{}
	if err := readObjWithId(pk, "stat_ink_key_name").Scan(
		&pk, &keyName.Key, &keyName.Name,
	); err != nil {
		debug.PrintStack()
		log.Printf("pk: %d\n", pk)
		return nil, err
	}
	name, err := readStatInkName(keyName.Name)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInkFailReason{
		Key:  enums.FailureReasonEnum(keyName.Key),
		Name: *name,
	}, nil
}

func readStatInkName(pk int64) (*api_objects.StatInkName, error) {
	name := api_objects.StatInkName{}
	var junk int
	if err := readObjWithId(pk, "stat_ink_name").Scan(
		&junk, &name.DeDE, &name.EnGB, &name.EnUS, &name.EsES, &name.EsMX, &name.FrCA, &name.FrFR, &name.ItIT,
		&name.JaJP, &name.NlNL, &name.RuRU, &name.ZhCN, &name.ZhTW,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &name, nil
}

func readShiftStatInkStageName(pk int64) (*api_objects.ShiftStatInkStageName, error) {
	name := api_objects.ShiftStatInkStageName{}
	var junk int
	if err := readObjWithId(pk, "stat_ink_name").Scan(
		&junk, &name.DeDE, &name.EnGB, &name.EnUS, &name.EsES, &name.EsMX, &name.FrCA, &name.FrFR, &name.ItIT,
		&name.JaJP, &name.NlNL, &name.RuRU, &name.ZhCN, &name.ZhTW,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &name, nil
}

func readBattleStatInkUserStatsV2GachiRules(pk int64) (*api_objects.BattleStatInkUserStatsV2GachiRules, error) {
	rules := db_objects.BattleStatInkUserStatsV2GachiRules{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_user_stats_v2_gachi_rules").Scan(
		&pk, &rules.Area, &rules.Yagura, &rules.Hoko, &rules.Asari,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	area, err := readBattleStatInkUserStatsV2GachiRulesSub(rules.Area)
	if err != nil {
		return nil, err
	}
	yagura, err := readBattleStatInkUserStatsV2GachiRulesSub(rules.Yagura)
	if err != nil {
		return nil, err
	}
	hoko, err := readBattleStatInkUserStatsV2GachiRulesSub(rules.Hoko)
	if err != nil {
		return nil, err
	}
	asari, err := readBattleStatInkUserStatsV2GachiRulesSub(rules.Asari)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkUserStatsV2GachiRules{
		Area:   *area,
		Yagura: *yagura,
		Hoko:   *hoko,
		Asari:  *asari,
	}, nil
}

func readBattleStatInkUserStatsV2GachiRulesSub(pk int64) (*api_objects.BattleStatInkUserStatsV2GachiRulesSub, error) {
	rules := api_objects.BattleStatInkUserStatsV2GachiRulesSub{}
	var junk int
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_user_stats_v2_gachi_rules_sub").Scan(
		&junk, &rules.RankPeak, &rules.RankCurrent, //&rules.XPowerPeak, &rules.XPowerCurrent,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &rules, nil
}

func readBattleStatInkMap(pk int64) (*api_objects.BattleStatInkMap, error) {
	stage := db_objects.BattleStatInkMap{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_map").Scan(
		&pk, &stage.Key, &stage.Name, &stage.Splatnet, &stage.Area, &stage.ReleaseAt, &stage.ShortName,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	releaseAt, err := readStatInkTime(stage.ReleaseAt)
	if err != nil {
		return nil, err
	}
	name, err := readStatInkName(stage.Name)
	if err != nil {
		return nil, err
	}
	shortName, err := readStatInkName(stage.ShortName)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkMap{
		Key:       stage.Key,
		Name:      *name,
		Splatnet:  stage.Splatnet,
		Area:      stage.Area,
		ReleaseAt: *releaseAt,
		ShortName: *shortName,
	}, nil
}

func readBattleStatInkWeapon(pk int64) (*api_objects.BattleStatInkWeapon, error) {
	weapon := db_objects.BattleStatInkWeapon{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_weapon").Scan(
		&pk, &weapon.Key, &weapon.Name, &weapon.Splatnet, &weapon.Type, &weapon.ReskinOf, &weapon.MainRef,
		&weapon.Sub, &weapon.Special, &weapon.MainPowerUp,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	name, err := readStatInkName(weapon.Name)
	if err != nil {
		return nil, err
	}
	sub, err := readStatInkKeyName(weapon.Sub)
	if err != nil {
		return nil, err
	}
	special, err := readStatInkKeyName(weapon.Special)
	if err != nil {
		return nil, err
	}
	mainPowerUp, err := readStatInkKeyName(weapon.MainPowerUp)
	if err != nil {
		return nil, err
	}
	weaponType, err := readBattleStatInkWeaponType(weapon.Type)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkWeapon{
		Key:         weapon.Key,
		Name:        *name,
		Splatnet:    weapon.Splatnet,
		Type:        *weaponType,
		ReskinOf:    weapon.ReskinOf,
		MainRef:     weapon.MainRef,
		Sub:         *sub,
		Special:     *special,
		MainPowerUp: *mainPowerUp,
	}, nil
}

func readBattleStatInkWeaponType(pk int64) (*api_objects.BattleStatInkWeaponType, error) {
	weaponType := db_objects.BattleStatInkWeaponType{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_weapon_type").Scan(
		&pk, &weaponType.Key, &weaponType.Name, &weaponType.Category,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	name, err := readStatInkName(weaponType.Name)
	if err != nil {
		return nil, err
	}
	category, err := readStatInkKeyName(weaponType.Category)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkWeaponType{
		Key:      weaponType.Key,
		Name:     *name,
		Category: *category,
	}, nil
}

func readBattleStatInkFreshness(pk int64) (*api_objects.BattleStatInkFreshness, error) {
	freshness := db_objects.BattleStatInkFreshness{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_freshness").Scan(
		&pk, &freshness.Freshness, &freshness.Title,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	title, err := readStatInkName(freshness.Title)
	if err != nil {

		return nil, err
	}
	return &api_objects.BattleStatInkFreshness{
		Freshness: freshness.Freshness,
		Title:     *title,
	}, nil
}

func readBattleStatInkRank(pk int64) (*api_objects.BattleStatInkRank, error) {
	rank := db_objects.BattleStatInkRank{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_rank").Scan(
		&pk, &rank.Key, &rank.Name, &rank.Zone,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	name, err := readStatInkName(rank.Name)
	if err != nil {

		return nil, err
	}
	zone, err := readStatInkKeyName(rank.Zone)
	if err != nil {

		return nil, err
	}
	return &api_objects.BattleStatInkRank{
		Key:  rank.Key,
		Name: *name,
		Zone: *zone,
	}, nil
}

func readStatInkGender(pk int64) (*api_objects.StatInkGender, error) {
	gender := db_objects.StatInkGender{}
	if err := readObjWithId(pk, "stat_ink_gender").Scan(
		&pk, &gender.Key, &gender.Name, &gender.Iso5218,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	name, err := readStatInkName(gender.Name)
	if err != nil {

		return nil, err
	}
	return &api_objects.StatInkGender{
		Key:     gender.Key,
		Name:    *name,
		Iso5218: gender.Iso5218,
	}, nil
}

func readBattleStatInkGears(pk int64) (*api_objects.BattleStatInkGears, error) {
	gears := db_objects.BattleStatInkGears{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_gears").Scan(
		&pk, &gears.Headgear, &gears.Clothing, &gears.Shoes,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	headgear, err := readBattleStatInkGearsClothes(gears.Headgear)
	if err != nil {
		return nil, err
	}
	clothing, err := readBattleStatInkGearsClothes(gears.Clothing)
	if err != nil {
		return nil, err
	}
	shoes, err := readBattleStatInkGearsClothes(gears.Shoes)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkGears{
		Headgear: *headgear,
		Clothing: *clothing,
		Shoes:    *shoes,
	}, nil
}

func readBattleStatInkGearsClothes(pk int64) (*api_objects.BattleStatInkGearsClothes, error) {
	clothes := db_objects.BattleStatInkGearsClothes{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_gears_clothes").Scan(
		&pk, &clothes.Gear, &clothes.PrimaryAbility,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	secondaryAbilityKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "secondary_ability", "two_battles_battle_stat_ink_gears_clothes_sa_container", "asc", "pk")
	if err != nil {
		return nil, err
	}
	secondaryAbilities := make([]api_objects.StatInkKeyName, len(secondaryAbilityKeys))
	for i := range secondaryAbilityKeys {
		secondaryAbilityTemp, err := readStatInkKeyName(secondaryAbilityKeys[i])
		if err != nil {
			return nil, err
		}
		secondaryAbilities[i] = *secondaryAbilityTemp
	}
	primaryAbility, err := readStatInkKeyName(clothes.PrimaryAbility)
	if err != nil {
		return nil, err
	}
	gear, err := readBattleStatInkGearsClothesGear(clothes.Gear)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkGearsClothes{
		Gear:               *gear,
		PrimaryAbility:     *primaryAbility,
		SecondaryAbilities: secondaryAbilities,
	}, nil
}

func readBattleStatInkGearsClothesGear(pk int64) (*api_objects.BattleStatInkGearsClothesGear, error) {
	gear := db_objects.BattleStatInkGearsClothesGear{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_gears_clothes_gear").Scan(
		&pk, &gear.Key, &gear.Name, &gear.Splatnet, &gear.Type, &gear.Brand, &gear.PrimaryAbility,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	name, err := readStatInkName(gear.Name)
	if err != nil {
		return nil, err
	}
	gearType, err := readStatInkKeyName(gear.Type)
	if err != nil {
		return nil, err
	}
	brand, err := readStatInkKeyName(gear.Brand)
	if err != nil {
		return nil, err
	}
	primaryAbility, err := readStatInkKeyName(gear.PrimaryAbility)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkGearsClothesGear{
		Key:            gear.Key,
		Name:           *name,
		Splatnet:       gear.Splatnet,
		Type:           *gearType,
		Brand:          *brand,
		PrimaryAbility: *primaryAbility,
	}, nil
}

func readBattleStatInkPlayer(pk int64) (*api_objects.BattleStatInkPlayer, error) {
	player := db_objects.BattleStatInkPlayer{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_player").Scan(
		&pk, &player.Parent, &player.Team, &player.IsMe, &player.Weapon, &player.Level, &player.Rank, &player.StarRank,
		&player.RankInTeam, &player.Kill, &player.Death, &player.KillOrAssist, &player.Special, //&player.MyKill,
		&player.Point, &player.Name, &player.Special, &player.Gender, &player.FestTitle, &player.SplatnetId,
		&player.Top500, &player.Icon,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	var rank *api_objects.BattleStatInkRank
	if player.Rank != nil {
		var err error
		rank, err = readBattleStatInkRank(*player.Rank)
		if err != nil {
			return nil, err
		}
	}
	var festTitle *api_objects.StatInkKeyName
	if player.FestTitle != nil {
		var err error
		festTitle, err = readStatInkKeyName(*player.FestTitle)
		if err != nil {
			return nil, err
		}
	}
	weapon, err := readBattleStatInkWeapon(player.Weapon)
	if err != nil {
		return nil, err
	}
	var species *api_objects.StatInkKeyName
	if player.Species != nil {
		species, err = readStatInkKeyName(*player.Species)
		if err != nil {
			return nil, err
		}
	}
	gender, err := readStatInkGender(player.Gender)
	if err != nil {
		return nil, err
	}
	return &api_objects.BattleStatInkPlayer{
		Team:         player.Team,
		IsMe:         player.IsMe,
		Weapon:       *weapon,
		Level:        player.Level,
		Rank:         rank,
		StarRank:     player.StarRank,
		RankInTeam:   player.RankInTeam,
		Kill:         player.Kill,
		Death:        player.Death,
		KillOrAssist: player.KillOrAssist,
		Special:      player.Special,
		//MyKill:       player.MyKill,
		Point:      player.Point,
		Name:       player.Name,
		Species:    species,
		Gender:     *gender,
		FestTitle:  festTitle,
		SplatnetId: player.SplatnetId,
		Top500:     player.Top500,
		Icon:       player.Icon,
	}, nil
}

func readBattleStatInkAgent(pk int64) (*api_objects.BattleStatInkAgent, error) {
	agent := db_objects.BattleStatInkAgent{}
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_agent").Scan(
		&pk, &agent.Name, &agent.Version, //&agent.GameVersion, &agent.GameVersionDate, &agent.Custom,
		&agent.Variables,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	var variables *api_objects.BattleStatInkAgentVariables
	if agent.Variables != nil {
		var err error
		variables, err = readBattleStatInkAgentVariables(*agent.Variables)
		if err != nil {
			return nil, err
		}
	}
	return &api_objects.BattleStatInkAgent{
		Name:    agent.Name,
		Version: agent.Version,
		//GameVersion:     agent.GameVersion,
		//GameVersionDate: agent.GameVersionDate,
		//Custom:          agent.Custom,
		Variables: variables,
	}, nil
}

func readBattleStatInkAgentVariables(pk int64) (*api_objects.BattleStatInkAgentVariables, error) {
	variables := api_objects.BattleStatInkAgentVariables{}
	var junk int
	if err := readObjWithId(pk, "two_battles_battle_stat_ink_agent_variables").Scan(
		&junk, &variables.UploadMode,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &variables, nil
}

func GetShift(userId, splatnetNumber int64) (*api_objects.Shift, error) {
	shift := db_objects.Shift{}
	var junk int64
	if err := readObjWithUserSplatnet(userId, splatnetNumber, "two_salmon_shift").Scan(
		&junk,
		&shift.UserId, &shift.PlayerSplatnetId, &shift.JobId,
		&shift.SplatnetJson, &shift.StatInkJson, &shift.StartTime, &shift.PlayTime, &shift.EndTime, &shift.DangerRate,
		&shift.IsClear, &shift.JobFailureReason, &shift.FailureWave, &shift.GradePoint, &shift.GradePointDelta,
		&shift.JobScore,
		&shift.DrizzlerCount, &shift.FlyfishCount, &shift.GoldieCount, &shift.GrillerCount,
		&shift.MawsCount, &shift.ScrapperCount, &shift.SteelEelCount, &shift.SteelheadCount, &shift.StingerCount,
		&shift.Stage,
		&shift.PlayerName, &shift.PlayerDeathCount, &shift.PlayerReviveCount, &shift.PlayerGoldenEggs,
		&shift.PlayerPowerEggs, &shift.PlayerSpecial, &shift.PlayerTitle, &shift.PlayerSpecies, &shift.PlayerGender,
		&shift.PlayerW1Specials, &shift.PlayerW2Specials, &shift.PlayerW3Specials, &shift.PlayerW1Weapon,
		&shift.PlayerW2Weapon, &shift.PlayerW3Weapon, &shift.PlayerDrizzlerKills, &shift.PlayerFlyfishKills,
		&shift.PlayerGoldieKills, &shift.PlayerGrillerKills, &shift.PlayerMawsKills, &shift.PlayerScrapperKills,
		&shift.PlayerSteelEelKills, &shift.PlayerSteelheadKills, &shift.PlayerStingerKills,
		&shift.Teammate0SplatnetId, &shift.Teammate0Name, &shift.Teammate0DeathCount, &shift.Teammate0ReviveCount,
		&shift.Teammate0GoldenEggs, &shift.Teammate0PowerEggs, &shift.Teammate0Special, &shift.Teammate0Species,
		&shift.Teammate0Gender, &shift.Teammate0W1Specials, &shift.Teammate0W2Specials, &shift.Teammate0W3Specials,
		&shift.Teammate0W1Weapon, &shift.Teammate0W2Weapon, &shift.Teammate0W3Weapon, &shift.Teammate0DrizzlerKills,
		&shift.Teammate0FlyfishKills, &shift.Teammate0GoldieKills, &shift.Teammate0GrillerKills,
		&shift.Teammate0MawsKills, &shift.Teammate0ScrapperKills, &shift.Teammate0SteelEelKills,
		&shift.Teammate0SteelheadKills, &shift.Teammate0StingerKills,
		&shift.Teammate1SplatnetId, &shift.Teammate1Name, &shift.Teammate1DeathCount, &shift.Teammate1ReviveCount,
		&shift.Teammate1GoldenEggs, &shift.Teammate1PowerEggs, &shift.Teammate1Special, &shift.Teammate1Species,
		&shift.Teammate1Gender, &shift.Teammate1W1Specials, &shift.Teammate1W2Specials, &shift.Teammate1W3Specials,
		&shift.Teammate1W1Weapon, &shift.Teammate1W2Weapon, &shift.Teammate1W3Weapon, &shift.Teammate1DrizzlerKills,
		&shift.Teammate1FlyfishKills, &shift.Teammate1GoldieKills, &shift.Teammate1GrillerKills,
		&shift.Teammate1MawsKills, &shift.Teammate1ScrapperKills, &shift.Teammate1SteelEelKills,
		&shift.Teammate1SteelheadKills, &shift.Teammate1StingerKills,
		&shift.Teammate2SplatnetId, &shift.Teammate2Name, &shift.Teammate2DeathCount, &shift.Teammate2ReviveCount,
		&shift.Teammate2GoldenEggs, &shift.Teammate2PowerEggs, &shift.Teammate2Special, &shift.Teammate2Species,
		&shift.Teammate2Gender, &shift.Teammate2W1Specials, &shift.Teammate2W2Specials, &shift.Teammate2W3Specials,
		&shift.Teammate2W1Weapon, &shift.Teammate2W2Weapon, &shift.Teammate2W3Weapon, &shift.Teammate2DrizzlerKills,
		&shift.Teammate2FlyfishKills, &shift.Teammate2GoldieKills, &shift.Teammate2GrillerKills,
		&shift.Teammate2MawsKills, &shift.Teammate2ScrapperKills, &shift.Teammate2SteelEelKills,
		&shift.Teammate2SteelheadKills, &shift.Teammate2StingerKills,
		&shift.ScheduleEndTime, &shift.ScheduleStartTime, &shift.ScheduleWeapon0, &shift.ScheduleWeapon1,
		&shift.ScheduleWeapon2, &shift.ScheduleWeapon3,
		&shift.Wave1WaterLevel, &shift.Wave1EventType, &shift.Wave1GoldenIkuraNum, &shift.Wave1GoldenIkuraPopNum,
		&shift.Wave1IkuraNum, &shift.Wave1QuotaNum, &shift.Wave2WaterLevel, &shift.Wave2EventType,
		&shift.Wave2GoldenIkuraNum, &shift.Wave2GoldenIkuraPopNum, &shift.Wave2IkuraNum, &shift.Wave2QuotaNum,
		&shift.Wave3WaterLevel, &shift.Wave3EventType, &shift.Wave3GoldenIkuraNum, &shift.Wave3GoldenIkuraPopNum,
		&shift.Wave3IkuraNum, &shift.Wave3QuotaNum,
	); err != nil {
		return nil, err
	}
	var splatnet *api_objects.ShiftSplatnet
	if shift.SplatnetJson != nil {
		var err error
		splatnet, err = getShiftSplatnet(*shift.SplatnetJson)
		if err != nil {
			return nil, err
		}
	}
	var statink *api_objects.ShiftStatInk
	if shift.StatInkJson != nil {
		var err error
		statink, err = getShiftStatInk(*shift.StatInkJson)
		if err != nil {
			return nil, err
		}
	}
	return &api_objects.Shift{
		UserId:                  shift.UserId,
		PlayerSplatnetId:        shift.PlayerSplatnetId,
		JobId:                   shift.JobId,
		SplatnetJson:            splatnet,
		StatInkJson:             statink,
		StartTime:               shift.StartTime,
		PlayTime:                shift.PlayTime,
		EndTime:                 shift.EndTime,
		DangerRate:              shift.DangerRate,
		IsClear:                 shift.IsClear,
		JobFailureReason:        shift.JobFailureReason,
		FailureWave:             shift.FailureWave,
		GradePoint:              shift.GradePoint,
		GradePointDelta:         shift.GradePointDelta,
		JobScore:                shift.JobScore,
		DrizzlerCount:           shift.DrizzlerCount,
		FlyfishCount:            shift.FlyfishCount,
		GoldieCount:             shift.GoldieCount,
		GrillerCount:            shift.GrillerCount,
		MawsCount:               shift.MawsCount,
		ScrapperCount:           shift.ScrapperCount,
		SteelEelCount:           shift.SteelEelCount,
		SteelheadCount:          shift.SteelheadCount,
		StingerCount:            shift.StingerCount,
		Stage:                   shift.Stage,
		PlayerName:              shift.PlayerName,
		PlayerDeathCount:        shift.PlayerDeathCount,
		PlayerReviveCount:       shift.PlayerReviveCount,
		PlayerGoldenEggs:        shift.PlayerGoldenEggs,
		PlayerPowerEggs:         shift.PlayerPowerEggs,
		PlayerSpecial:           shift.PlayerSpecial,
		PlayerTitle:             shift.PlayerTitle,
		PlayerSpecies:           shift.PlayerSpecies,
		PlayerGender:            shift.PlayerGender,
		PlayerW1Specials:        shift.PlayerW1Specials,
		PlayerW2Specials:        shift.PlayerW2Specials,
		PlayerW3Specials:        shift.PlayerW3Specials,
		PlayerW1Weapon:          shift.PlayerW1Weapon,
		PlayerW2Weapon:          shift.PlayerW2Weapon,
		PlayerW3Weapon:          shift.PlayerW3Weapon,
		PlayerDrizzlerKills:     shift.PlayerDrizzlerKills,
		PlayerFlyfishKills:      shift.PlayerFlyfishKills,
		PlayerGoldieKills:       shift.PlayerGoldieKills,
		PlayerGrillerKills:      shift.PlayerGrillerKills,
		PlayerMawsKills:         shift.PlayerMawsKills,
		PlayerScrapperKills:     shift.PlayerScrapperKills,
		PlayerSteelEelKills:     shift.PlayerSteelEelKills,
		PlayerSteelheadKills:    shift.PlayerSteelheadKills,
		PlayerStingerKills:      shift.PlayerStingerKills,
		Teammate0SplatnetId:     shift.Teammate0SplatnetId,
		Teammate0Name:           shift.Teammate0Name,
		Teammate0DeathCount:     shift.Teammate0DeathCount,
		Teammate0ReviveCount:    shift.Teammate0ReviveCount,
		Teammate0GoldenEggs:     shift.Teammate0GoldenEggs,
		Teammate0PowerEggs:      shift.Teammate0PowerEggs,
		Teammate0Special:        shift.Teammate0Special,
		Teammate0Species:        shift.Teammate0Species,
		Teammate0Gender:         shift.Teammate0Gender,
		Teammate0W1Specials:     shift.Teammate0W1Specials,
		Teammate0W2Specials:     shift.Teammate0W2Specials,
		Teammate0W3Specials:     shift.Teammate0W3Specials,
		Teammate0W1Weapon:       shift.Teammate0W1Weapon,
		Teammate0W2Weapon:       shift.Teammate0W2Weapon,
		Teammate0W3Weapon:       shift.Teammate0W3Weapon,
		Teammate0DrizzlerKills:  shift.Teammate0DrizzlerKills,
		Teammate0FlyfishKills:   shift.Teammate0FlyfishKills,
		Teammate0GoldieKills:    shift.Teammate0GoldieKills,
		Teammate0GrillerKills:   shift.Teammate0GrillerKills,
		Teammate0MawsKills:      shift.Teammate0MawsKills,
		Teammate0ScrapperKills:  shift.Teammate0ScrapperKills,
		Teammate0SteelEelKills:  shift.Teammate0SteelEelKills,
		Teammate0SteelheadKills: shift.Teammate0SteelheadKills,
		Teammate0StingerKills:   shift.Teammate0StingerKills,
		Teammate1SplatnetId:     shift.Teammate1SplatnetId,
		Teammate1Name:           shift.Teammate1Name,
		Teammate1DeathCount:     shift.Teammate1DeathCount,
		Teammate1ReviveCount:    shift.Teammate1ReviveCount,
		Teammate1GoldenEggs:     shift.Teammate1GoldenEggs,
		Teammate1PowerEggs:      shift.Teammate1PowerEggs,
		Teammate1Special:        shift.Teammate1Special,
		Teammate1Species:        shift.Teammate1Species,
		Teammate1Gender:         shift.Teammate1Gender,
		Teammate1W1Specials:     shift.Teammate1W1Specials,
		Teammate1W2Specials:     shift.Teammate1W2Specials,
		Teammate1W3Specials:     shift.Teammate1W3Specials,
		Teammate1W1Weapon:       shift.Teammate1W1Weapon,
		Teammate1W2Weapon:       shift.Teammate1W2Weapon,
		Teammate1W3Weapon:       shift.Teammate1W3Weapon,
		Teammate1DrizzlerKills:  shift.Teammate1DrizzlerKills,
		Teammate1FlyfishKills:   shift.Teammate1FlyfishKills,
		Teammate1GoldieKills:    shift.Teammate1GoldieKills,
		Teammate1GrillerKills:   shift.Teammate1GrillerKills,
		Teammate1MawsKills:      shift.Teammate1MawsKills,
		Teammate1ScrapperKills:  shift.Teammate1ScrapperKills,
		Teammate1SteelEelKills:  shift.Teammate1SteelEelKills,
		Teammate1SteelheadKills: shift.Teammate1SteelheadKills,
		Teammate1StingerKills:   shift.Teammate1StingerKills,
		Teammate2SplatnetId:     shift.Teammate2SplatnetId,
		Teammate2Name:           shift.Teammate2Name,
		Teammate2DeathCount:     shift.Teammate2DeathCount,
		Teammate2ReviveCount:    shift.Teammate2ReviveCount,
		Teammate2GoldenEggs:     shift.Teammate2GoldenEggs,
		Teammate2PowerEggs:      shift.Teammate2PowerEggs,
		Teammate2Special:        shift.Teammate2Special,
		Teammate2Species:        shift.Teammate2Species,
		Teammate2Gender:         shift.Teammate2Gender,
		Teammate2W1Specials:     shift.Teammate2W1Specials,
		Teammate2W2Specials:     shift.Teammate2W2Specials,
		Teammate2W3Specials:     shift.Teammate2W3Specials,
		Teammate2W1Weapon:       shift.Teammate2W1Weapon,
		Teammate2W2Weapon:       shift.Teammate2W2Weapon,
		Teammate2W3Weapon:       shift.Teammate2W3Weapon,
		Teammate2DrizzlerKills:  shift.Teammate2DrizzlerKills,
		Teammate2FlyfishKills:   shift.Teammate2FlyfishKills,
		Teammate2GoldieKills:    shift.Teammate2GoldieKills,
		Teammate2GrillerKills:   shift.Teammate2GrillerKills,
		Teammate2MawsKills:      shift.Teammate2MawsKills,
		Teammate2ScrapperKills:  shift.Teammate2ScrapperKills,
		Teammate2SteelEelKills:  shift.Teammate2SteelEelKills,
		Teammate2SteelheadKills: shift.Teammate2SteelheadKills,
		Teammate2StingerKills:   shift.Teammate2StingerKills,
		ScheduleEndTime:         shift.ScheduleEndTime,
		ScheduleStartTime:       *shift.ScheduleStartTime,
		ScheduleWeapon0:         shift.ScheduleWeapon0,
		ScheduleWeapon1:         shift.ScheduleWeapon1,
		ScheduleWeapon2:         shift.ScheduleWeapon2,
		ScheduleWeapon3:         shift.ScheduleWeapon3,
		Wave1WaterLevel:         shift.Wave1WaterLevel,
		Wave1EventType:          shift.Wave1EventType,
		Wave1GoldenDelivered:    shift.Wave1GoldenIkuraNum,
		Wave1GoldenAppear:       shift.Wave1GoldenIkuraPopNum,
		Wave1PowerEggs:          shift.Wave1IkuraNum,
		Wave1Quota:              shift.Wave1QuotaNum,
		Wave2WaterLevel:         shift.Wave2WaterLevel,
		Wave2EventType:          shift.Wave2EventType,
		Wave2GoldenDelivered:    shift.Wave2GoldenIkuraNum,
		Wave2GoldenAppear:       shift.Wave2GoldenIkuraPopNum,
		Wave2PowerEggs:          shift.Wave2IkuraNum,
		Wave2Quota:              shift.Wave2QuotaNum,
		Wave3WaterLevel:         shift.Wave3WaterLevel,
		Wave3EventType:          shift.Wave3EventType,
		Wave3GoldenDelivered:    shift.Wave3GoldenIkuraNum,
		Wave3GoldenAppear:       shift.Wave3GoldenIkuraPopNum,
		Wave3PowerEggs:          shift.Wave3IkuraNum,
		Wave3Quota:              shift.Wave3QuotaNum,
	}, nil
}

func GetShiftLean(userId, splatnetNumber int64) (*api_objects.Shift, error) {
	shift := db_objects.Shift{}
	var junk int64
	if err := readObjWithUserSplatnet(userId, splatnetNumber, "two_salmon_shift").Scan(
		&junk,
		&shift.UserId, &shift.PlayerSplatnetId, &shift.JobId,
		&shift.SplatnetJson, &shift.StatInkJson, &shift.StartTime, &shift.PlayTime, &shift.EndTime, &shift.DangerRate,
		&shift.IsClear, &shift.JobFailureReason, &shift.FailureWave, &shift.GradePoint, &shift.GradePointDelta,
		&shift.JobScore,
		&shift.DrizzlerCount, &shift.FlyfishCount, &shift.GoldieCount, &shift.GrillerCount,
		&shift.MawsCount, &shift.ScrapperCount, &shift.SteelEelCount, &shift.SteelheadCount, &shift.StingerCount,
		&shift.Stage,
		&shift.PlayerName, &shift.PlayerDeathCount, &shift.PlayerReviveCount, &shift.PlayerGoldenEggs,
		&shift.PlayerPowerEggs, &shift.PlayerSpecial, &shift.PlayerTitle, &shift.PlayerSpecies, &shift.PlayerGender,
		&shift.PlayerW1Specials, &shift.PlayerW2Specials, &shift.PlayerW3Specials, &shift.PlayerW1Weapon,
		&shift.PlayerW2Weapon, &shift.PlayerW3Weapon, &shift.PlayerDrizzlerKills, &shift.PlayerFlyfishKills,
		&shift.PlayerGoldieKills, &shift.PlayerGrillerKills, &shift.PlayerMawsKills, &shift.PlayerScrapperKills,
		&shift.PlayerSteelEelKills, &shift.PlayerSteelheadKills, &shift.PlayerStingerKills,
		&shift.Teammate0SplatnetId, &shift.Teammate0Name, &shift.Teammate0DeathCount, &shift.Teammate0ReviveCount,
		&shift.Teammate0GoldenEggs, &shift.Teammate0PowerEggs, &shift.Teammate0Special, &shift.Teammate0Species,
		&shift.Teammate0Gender, &shift.Teammate0W1Specials, &shift.Teammate0W2Specials, &shift.Teammate0W3Specials,
		&shift.Teammate0W1Weapon, &shift.Teammate0W2Weapon, &shift.Teammate0W3Weapon, &shift.Teammate0DrizzlerKills,
		&shift.Teammate0FlyfishKills, &shift.Teammate0GoldieKills, &shift.Teammate0GrillerKills,
		&shift.Teammate0MawsKills, &shift.Teammate0ScrapperKills, &shift.Teammate0SteelEelKills,
		&shift.Teammate0SteelheadKills, &shift.Teammate0StingerKills,
		&shift.Teammate1SplatnetId, &shift.Teammate1Name, &shift.Teammate1DeathCount, &shift.Teammate1ReviveCount,
		&shift.Teammate1GoldenEggs, &shift.Teammate1PowerEggs, &shift.Teammate1Special, &shift.Teammate1Species,
		&shift.Teammate1Gender, &shift.Teammate1W1Specials, &shift.Teammate1W2Specials, &shift.Teammate1W3Specials,
		&shift.Teammate1W1Weapon, &shift.Teammate1W2Weapon, &shift.Teammate1W3Weapon, &shift.Teammate1DrizzlerKills,
		&shift.Teammate1FlyfishKills, &shift.Teammate1GoldieKills, &shift.Teammate1GrillerKills,
		&shift.Teammate1MawsKills, &shift.Teammate1ScrapperKills, &shift.Teammate1SteelEelKills,
		&shift.Teammate1SteelheadKills, &shift.Teammate1StingerKills,
		&shift.Teammate2SplatnetId, &shift.Teammate2Name, &shift.Teammate2DeathCount, &shift.Teammate2ReviveCount,
		&shift.Teammate2GoldenEggs, &shift.Teammate2PowerEggs, &shift.Teammate2Special, &shift.Teammate2Species,
		&shift.Teammate2Gender, &shift.Teammate2W1Specials, &shift.Teammate2W2Specials, &shift.Teammate2W3Specials,
		&shift.Teammate2W1Weapon, &shift.Teammate2W2Weapon, &shift.Teammate2W3Weapon, &shift.Teammate2DrizzlerKills,
		&shift.Teammate2FlyfishKills, &shift.Teammate2GoldieKills, &shift.Teammate2GrillerKills,
		&shift.Teammate2MawsKills, &shift.Teammate2ScrapperKills, &shift.Teammate2SteelEelKills,
		&shift.Teammate2SteelheadKills, &shift.Teammate2StingerKills,
		&shift.ScheduleEndTime, &shift.ScheduleStartTime, &shift.ScheduleWeapon0, &shift.ScheduleWeapon1,
		&shift.ScheduleWeapon2, &shift.ScheduleWeapon3,
		&shift.Wave1WaterLevel, &shift.Wave1EventType, &shift.Wave1GoldenIkuraNum, &shift.Wave1GoldenIkuraPopNum,
		&shift.Wave1IkuraNum, &shift.Wave1QuotaNum, &shift.Wave2WaterLevel, &shift.Wave2EventType,
		&shift.Wave2GoldenIkuraNum, &shift.Wave2GoldenIkuraPopNum, &shift.Wave2IkuraNum, &shift.Wave2QuotaNum,
		&shift.Wave3WaterLevel, &shift.Wave3EventType, &shift.Wave3GoldenIkuraNum, &shift.Wave3GoldenIkuraPopNum,
		&shift.Wave3IkuraNum, &shift.Wave3QuotaNum,
	); err != nil {
		return nil, err
	}
	return &api_objects.Shift{
		UserId:                  shift.UserId,
		PlayerSplatnetId:        shift.PlayerSplatnetId,
		JobId:                   shift.JobId,
		SplatnetJson:            nil,
		StatInkJson:             nil,
		StartTime:               shift.StartTime,
		PlayTime:                shift.PlayTime,
		EndTime:                 shift.EndTime,
		DangerRate:              shift.DangerRate,
		IsClear:                 shift.IsClear,
		JobFailureReason:        shift.JobFailureReason,
		FailureWave:             shift.FailureWave,
		GradePoint:              shift.GradePoint,
		GradePointDelta:         shift.GradePointDelta,
		JobScore:                shift.JobScore,
		DrizzlerCount:           shift.DrizzlerCount,
		FlyfishCount:            shift.FlyfishCount,
		GoldieCount:             shift.GoldieCount,
		GrillerCount:            shift.GrillerCount,
		MawsCount:               shift.MawsCount,
		ScrapperCount:           shift.ScrapperCount,
		SteelEelCount:           shift.SteelEelCount,
		SteelheadCount:          shift.SteelheadCount,
		StingerCount:            shift.StingerCount,
		Stage:                   shift.Stage,
		PlayerName:              shift.PlayerName,
		PlayerDeathCount:        shift.PlayerDeathCount,
		PlayerReviveCount:       shift.PlayerReviveCount,
		PlayerGoldenEggs:        shift.PlayerGoldenEggs,
		PlayerPowerEggs:         shift.PlayerPowerEggs,
		PlayerSpecial:           shift.PlayerSpecial,
		PlayerTitle:             shift.PlayerTitle,
		PlayerSpecies:           shift.PlayerSpecies,
		PlayerGender:            shift.PlayerGender,
		PlayerW1Specials:        shift.PlayerW1Specials,
		PlayerW2Specials:        shift.PlayerW2Specials,
		PlayerW3Specials:        shift.PlayerW3Specials,
		PlayerW1Weapon:          shift.PlayerW1Weapon,
		PlayerW2Weapon:          shift.PlayerW2Weapon,
		PlayerW3Weapon:          shift.PlayerW3Weapon,
		PlayerDrizzlerKills:     shift.PlayerDrizzlerKills,
		PlayerFlyfishKills:      shift.PlayerFlyfishKills,
		PlayerGoldieKills:       shift.PlayerGoldieKills,
		PlayerGrillerKills:      shift.PlayerGrillerKills,
		PlayerMawsKills:         shift.PlayerMawsKills,
		PlayerScrapperKills:     shift.PlayerScrapperKills,
		PlayerSteelEelKills:     shift.PlayerSteelEelKills,
		PlayerSteelheadKills:    shift.PlayerSteelheadKills,
		PlayerStingerKills:      shift.PlayerStingerKills,
		Teammate0SplatnetId:     shift.Teammate0SplatnetId,
		Teammate0Name:           shift.Teammate0Name,
		Teammate0DeathCount:     shift.Teammate0DeathCount,
		Teammate0ReviveCount:    shift.Teammate0ReviveCount,
		Teammate0GoldenEggs:     shift.Teammate0GoldenEggs,
		Teammate0PowerEggs:      shift.Teammate0PowerEggs,
		Teammate0Special:        shift.Teammate0Special,
		Teammate0Species:        shift.Teammate0Species,
		Teammate0Gender:         shift.Teammate0Gender,
		Teammate0W1Specials:     shift.Teammate0W1Specials,
		Teammate0W2Specials:     shift.Teammate0W2Specials,
		Teammate0W3Specials:     shift.Teammate0W3Specials,
		Teammate0W1Weapon:       shift.Teammate0W1Weapon,
		Teammate0W2Weapon:       shift.Teammate0W2Weapon,
		Teammate0W3Weapon:       shift.Teammate0W3Weapon,
		Teammate0DrizzlerKills:  shift.Teammate0DrizzlerKills,
		Teammate0FlyfishKills:   shift.Teammate0FlyfishKills,
		Teammate0GoldieKills:    shift.Teammate0GoldieKills,
		Teammate0GrillerKills:   shift.Teammate0GrillerKills,
		Teammate0MawsKills:      shift.Teammate0MawsKills,
		Teammate0ScrapperKills:  shift.Teammate0ScrapperKills,
		Teammate0SteelEelKills:  shift.Teammate0SteelEelKills,
		Teammate0SteelheadKills: shift.Teammate0SteelheadKills,
		Teammate0StingerKills:   shift.Teammate0StingerKills,
		Teammate1SplatnetId:     shift.Teammate1SplatnetId,
		Teammate1Name:           shift.Teammate1Name,
		Teammate1DeathCount:     shift.Teammate1DeathCount,
		Teammate1ReviveCount:    shift.Teammate1ReviveCount,
		Teammate1GoldenEggs:     shift.Teammate1GoldenEggs,
		Teammate1PowerEggs:      shift.Teammate1PowerEggs,
		Teammate1Special:        shift.Teammate1Special,
		Teammate1Species:        shift.Teammate1Species,
		Teammate1Gender:         shift.Teammate1Gender,
		Teammate1W1Specials:     shift.Teammate1W1Specials,
		Teammate1W2Specials:     shift.Teammate1W2Specials,
		Teammate1W3Specials:     shift.Teammate1W3Specials,
		Teammate1W1Weapon:       shift.Teammate1W1Weapon,
		Teammate1W2Weapon:       shift.Teammate1W2Weapon,
		Teammate1W3Weapon:       shift.Teammate1W3Weapon,
		Teammate1DrizzlerKills:  shift.Teammate1DrizzlerKills,
		Teammate1FlyfishKills:   shift.Teammate1FlyfishKills,
		Teammate1GoldieKills:    shift.Teammate1GoldieKills,
		Teammate1GrillerKills:   shift.Teammate1GrillerKills,
		Teammate1MawsKills:      shift.Teammate1MawsKills,
		Teammate1ScrapperKills:  shift.Teammate1ScrapperKills,
		Teammate1SteelEelKills:  shift.Teammate1SteelEelKills,
		Teammate1SteelheadKills: shift.Teammate1SteelheadKills,
		Teammate1StingerKills:   shift.Teammate1StingerKills,
		Teammate2SplatnetId:     shift.Teammate2SplatnetId,
		Teammate2Name:           shift.Teammate2Name,
		Teammate2DeathCount:     shift.Teammate2DeathCount,
		Teammate2ReviveCount:    shift.Teammate2ReviveCount,
		Teammate2GoldenEggs:     shift.Teammate2GoldenEggs,
		Teammate2PowerEggs:      shift.Teammate2PowerEggs,
		Teammate2Special:        shift.Teammate2Special,
		Teammate2Species:        shift.Teammate2Species,
		Teammate2Gender:         shift.Teammate2Gender,
		Teammate2W1Specials:     shift.Teammate2W1Specials,
		Teammate2W2Specials:     shift.Teammate2W2Specials,
		Teammate2W3Specials:     shift.Teammate2W3Specials,
		Teammate2W1Weapon:       shift.Teammate2W1Weapon,
		Teammate2W2Weapon:       shift.Teammate2W2Weapon,
		Teammate2W3Weapon:       shift.Teammate2W3Weapon,
		Teammate2DrizzlerKills:  shift.Teammate2DrizzlerKills,
		Teammate2FlyfishKills:   shift.Teammate2FlyfishKills,
		Teammate2GoldieKills:    shift.Teammate2GoldieKills,
		Teammate2GrillerKills:   shift.Teammate2GrillerKills,
		Teammate2MawsKills:      shift.Teammate2MawsKills,
		Teammate2ScrapperKills:  shift.Teammate2ScrapperKills,
		Teammate2SteelEelKills:  shift.Teammate2SteelEelKills,
		Teammate2SteelheadKills: shift.Teammate2SteelheadKills,
		Teammate2StingerKills:   shift.Teammate2StingerKills,
		ScheduleEndTime:         shift.ScheduleEndTime,
		ScheduleStartTime:       *shift.ScheduleStartTime,
		ScheduleWeapon0:         shift.ScheduleWeapon0,
		ScheduleWeapon1:         shift.ScheduleWeapon1,
		ScheduleWeapon2:         shift.ScheduleWeapon2,
		ScheduleWeapon3:         shift.ScheduleWeapon3,
		Wave1WaterLevel:         shift.Wave1WaterLevel,
		Wave1EventType:          shift.Wave1EventType,
		Wave1GoldenDelivered:    shift.Wave1GoldenIkuraNum,
		Wave1GoldenAppear:       shift.Wave1GoldenIkuraPopNum,
		Wave1PowerEggs:          shift.Wave1IkuraNum,
		Wave1Quota:              shift.Wave1QuotaNum,
		Wave2WaterLevel:         shift.Wave2WaterLevel,
		Wave2EventType:          shift.Wave2EventType,
		Wave2GoldenDelivered:    shift.Wave2GoldenIkuraNum,
		Wave2GoldenAppear:       shift.Wave2GoldenIkuraPopNum,
		Wave2PowerEggs:          shift.Wave2IkuraNum,
		Wave2Quota:              shift.Wave2QuotaNum,
		Wave3WaterLevel:         shift.Wave3WaterLevel,
		Wave3EventType:          shift.Wave3EventType,
		Wave3GoldenDelivered:    shift.Wave3GoldenIkuraNum,
		Wave3GoldenAppear:       shift.Wave3GoldenIkuraPopNum,
		Wave3PowerEggs:          shift.Wave3IkuraNum,
		Wave3Quota:              shift.Wave3QuotaNum,
	}, nil
}

func GetShiftPk(pk int64) (*api_objects.Shift, error) {
	shift := db_objects.Shift{}
	if err := readObjWithId(pk, "two_salmon_shift").Scan(
		&pk,
		&shift.UserId, &shift.PlayerSplatnetId, &shift.JobId,
		&shift.SplatnetJson, &shift.StatInkJson, &shift.StartTime, &shift.PlayTime, &shift.EndTime, &shift.DangerRate,
		&shift.IsClear, &shift.JobFailureReason, &shift.FailureWave, &shift.GradePoint, &shift.GradePointDelta,
		&shift.JobScore,
		&shift.DrizzlerCount, &shift.FlyfishCount, &shift.GoldieCount, &shift.GrillerCount,
		&shift.MawsCount, &shift.ScrapperCount, &shift.SteelEelCount, &shift.SteelheadCount, &shift.StingerCount,
		&shift.Stage,
		&shift.PlayerName, &shift.PlayerDeathCount, &shift.PlayerReviveCount, &shift.PlayerGoldenEggs,
		&shift.PlayerPowerEggs, &shift.PlayerSpecial, &shift.PlayerTitle, &shift.PlayerSpecies, &shift.PlayerGender,
		&shift.PlayerW1Specials, &shift.PlayerW2Specials, &shift.PlayerW3Specials, &shift.PlayerW1Weapon,
		&shift.PlayerW2Weapon, &shift.PlayerW3Weapon, &shift.PlayerDrizzlerKills, &shift.PlayerFlyfishKills,
		&shift.PlayerGoldieKills, &shift.PlayerGrillerKills, &shift.PlayerMawsKills, &shift.PlayerScrapperKills,
		&shift.PlayerSteelEelKills, &shift.PlayerSteelheadKills, &shift.PlayerStingerKills,
		&shift.Teammate0SplatnetId, &shift.Teammate0Name, &shift.Teammate0DeathCount, &shift.Teammate0ReviveCount,
		&shift.Teammate0GoldenEggs, &shift.Teammate0PowerEggs, &shift.Teammate0Special, &shift.Teammate0Species,
		&shift.Teammate0Gender, &shift.Teammate0W1Specials, &shift.Teammate0W2Specials, &shift.Teammate0W3Specials,
		&shift.Teammate0W1Weapon, &shift.Teammate0W2Weapon, &shift.Teammate0W3Weapon, &shift.Teammate0DrizzlerKills,
		&shift.Teammate0FlyfishKills, &shift.Teammate0GoldieKills, &shift.Teammate0GrillerKills,
		&shift.Teammate0MawsKills, &shift.Teammate0ScrapperKills, &shift.Teammate0SteelEelKills,
		&shift.Teammate0SteelheadKills, &shift.Teammate0StingerKills,
		&shift.Teammate1SplatnetId, &shift.Teammate1Name, &shift.Teammate1DeathCount, &shift.Teammate1ReviveCount,
		&shift.Teammate1GoldenEggs, &shift.Teammate1PowerEggs, &shift.Teammate1Special, &shift.Teammate1Species,
		&shift.Teammate1Gender, &shift.Teammate1W1Specials, &shift.Teammate1W2Specials, &shift.Teammate1W3Specials,
		&shift.Teammate1W1Weapon, &shift.Teammate1W2Weapon, &shift.Teammate1W3Weapon, &shift.Teammate1DrizzlerKills,
		&shift.Teammate1FlyfishKills, &shift.Teammate1GoldieKills, &shift.Teammate1GrillerKills,
		&shift.Teammate1MawsKills, &shift.Teammate1ScrapperKills, &shift.Teammate1SteelEelKills,
		&shift.Teammate1SteelheadKills, &shift.Teammate1StingerKills,
		&shift.Teammate2SplatnetId, &shift.Teammate2Name, &shift.Teammate2DeathCount, &shift.Teammate2ReviveCount,
		&shift.Teammate2GoldenEggs, &shift.Teammate2PowerEggs, &shift.Teammate2Special, &shift.Teammate2Species,
		&shift.Teammate2Gender, &shift.Teammate2W1Specials, &shift.Teammate2W2Specials, &shift.Teammate2W3Specials,
		&shift.Teammate2W1Weapon, &shift.Teammate2W2Weapon, &shift.Teammate2W3Weapon, &shift.Teammate2DrizzlerKills,
		&shift.Teammate2FlyfishKills, &shift.Teammate2GoldieKills, &shift.Teammate2GrillerKills,
		&shift.Teammate2MawsKills, &shift.Teammate2ScrapperKills, &shift.Teammate2SteelEelKills,
		&shift.Teammate2SteelheadKills, &shift.Teammate2StingerKills,
		&shift.ScheduleEndTime, &shift.ScheduleStartTime, &shift.ScheduleWeapon0, &shift.ScheduleWeapon1,
		&shift.ScheduleWeapon2, &shift.ScheduleWeapon3,
		&shift.Wave1WaterLevel, &shift.Wave1EventType, &shift.Wave1GoldenIkuraNum, &shift.Wave1GoldenIkuraPopNum,
		&shift.Wave1IkuraNum, &shift.Wave1QuotaNum, &shift.Wave2WaterLevel, &shift.Wave2EventType,
		&shift.Wave2GoldenIkuraNum, &shift.Wave2GoldenIkuraPopNum, &shift.Wave2IkuraNum, &shift.Wave2QuotaNum,
		&shift.Wave3WaterLevel, &shift.Wave3EventType, &shift.Wave3GoldenIkuraNum, &shift.Wave3GoldenIkuraPopNum,
		&shift.Wave3IkuraNum, &shift.Wave3QuotaNum,
	); err != nil {
		return nil, err
	}
	var splatnet *api_objects.ShiftSplatnet
	if shift.SplatnetJson != nil {
		var err error
		splatnet, err = getShiftSplatnet(*shift.SplatnetJson)
		if err != nil {
			return nil, err
		}
	}
	var statink *api_objects.ShiftStatInk
	if shift.StatInkJson != nil {
		var err error
		statink, err = getShiftStatInk(*shift.StatInkJson)
		if err != nil {
			return nil, err
		}
	}
	return &api_objects.Shift{
		UserId:                  shift.UserId,
		PlayerSplatnetId:        shift.PlayerSplatnetId,
		JobId:                   shift.JobId,
		SplatnetJson:            splatnet,
		StatInkJson:             statink,
		StartTime:               shift.StartTime,
		PlayTime:                shift.PlayTime,
		EndTime:                 shift.EndTime,
		DangerRate:              shift.DangerRate,
		IsClear:                 shift.IsClear,
		JobFailureReason:        shift.JobFailureReason,
		FailureWave:             shift.FailureWave,
		GradePoint:              shift.GradePoint,
		GradePointDelta:         shift.GradePointDelta,
		JobScore:                shift.JobScore,
		DrizzlerCount:           shift.DrizzlerCount,
		FlyfishCount:            shift.FlyfishCount,
		GoldieCount:             shift.GoldieCount,
		GrillerCount:            shift.GrillerCount,
		MawsCount:               shift.MawsCount,
		ScrapperCount:           shift.ScrapperCount,
		SteelEelCount:           shift.SteelEelCount,
		SteelheadCount:          shift.SteelheadCount,
		StingerCount:            shift.StingerCount,
		Stage:                   shift.Stage,
		PlayerName:              shift.PlayerName,
		PlayerDeathCount:        shift.PlayerDeathCount,
		PlayerReviveCount:       shift.PlayerReviveCount,
		PlayerGoldenEggs:        shift.PlayerGoldenEggs,
		PlayerPowerEggs:         shift.PlayerPowerEggs,
		PlayerSpecial:           shift.PlayerSpecial,
		PlayerTitle:             shift.PlayerTitle,
		PlayerSpecies:           shift.PlayerSpecies,
		PlayerGender:            shift.PlayerGender,
		PlayerW1Specials:        shift.PlayerW1Specials,
		PlayerW2Specials:        shift.PlayerW2Specials,
		PlayerW3Specials:        shift.PlayerW3Specials,
		PlayerW1Weapon:          shift.PlayerW1Weapon,
		PlayerW2Weapon:          shift.PlayerW2Weapon,
		PlayerW3Weapon:          shift.PlayerW3Weapon,
		PlayerDrizzlerKills:     shift.PlayerDrizzlerKills,
		PlayerFlyfishKills:      shift.PlayerFlyfishKills,
		PlayerGoldieKills:       shift.PlayerGoldieKills,
		PlayerGrillerKills:      shift.PlayerGrillerKills,
		PlayerMawsKills:         shift.PlayerMawsKills,
		PlayerScrapperKills:     shift.PlayerScrapperKills,
		PlayerSteelEelKills:     shift.PlayerSteelEelKills,
		PlayerSteelheadKills:    shift.PlayerSteelheadKills,
		PlayerStingerKills:      shift.PlayerStingerKills,
		Teammate0SplatnetId:     shift.Teammate0SplatnetId,
		Teammate0Name:           shift.Teammate0Name,
		Teammate0DeathCount:     shift.Teammate0DeathCount,
		Teammate0ReviveCount:    shift.Teammate0ReviveCount,
		Teammate0GoldenEggs:     shift.Teammate0GoldenEggs,
		Teammate0PowerEggs:      shift.Teammate0PowerEggs,
		Teammate0Special:        shift.Teammate0Special,
		Teammate0Species:        shift.Teammate0Species,
		Teammate0Gender:         shift.Teammate0Gender,
		Teammate0W1Specials:     shift.Teammate0W1Specials,
		Teammate0W2Specials:     shift.Teammate0W2Specials,
		Teammate0W3Specials:     shift.Teammate0W3Specials,
		Teammate0W1Weapon:       shift.Teammate0W1Weapon,
		Teammate0W2Weapon:       shift.Teammate0W2Weapon,
		Teammate0W3Weapon:       shift.Teammate0W3Weapon,
		Teammate0DrizzlerKills:  shift.Teammate0DrizzlerKills,
		Teammate0FlyfishKills:   shift.Teammate0FlyfishKills,
		Teammate0GoldieKills:    shift.Teammate0GoldieKills,
		Teammate0GrillerKills:   shift.Teammate0GrillerKills,
		Teammate0MawsKills:      shift.Teammate0MawsKills,
		Teammate0ScrapperKills:  shift.Teammate0ScrapperKills,
		Teammate0SteelEelKills:  shift.Teammate0SteelEelKills,
		Teammate0SteelheadKills: shift.Teammate0SteelheadKills,
		Teammate0StingerKills:   shift.Teammate0StingerKills,
		Teammate1SplatnetId:     shift.Teammate1SplatnetId,
		Teammate1Name:           shift.Teammate1Name,
		Teammate1DeathCount:     shift.Teammate1DeathCount,
		Teammate1ReviveCount:    shift.Teammate1ReviveCount,
		Teammate1GoldenEggs:     shift.Teammate1GoldenEggs,
		Teammate1PowerEggs:      shift.Teammate1PowerEggs,
		Teammate1Special:        shift.Teammate1Special,
		Teammate1Species:        shift.Teammate1Species,
		Teammate1Gender:         shift.Teammate1Gender,
		Teammate1W1Specials:     shift.Teammate1W1Specials,
		Teammate1W2Specials:     shift.Teammate1W2Specials,
		Teammate1W3Specials:     shift.Teammate1W3Specials,
		Teammate1W1Weapon:       shift.Teammate1W1Weapon,
		Teammate1W2Weapon:       shift.Teammate1W2Weapon,
		Teammate1W3Weapon:       shift.Teammate1W3Weapon,
		Teammate1DrizzlerKills:  shift.Teammate1DrizzlerKills,
		Teammate1FlyfishKills:   shift.Teammate1FlyfishKills,
		Teammate1GoldieKills:    shift.Teammate1GoldieKills,
		Teammate1GrillerKills:   shift.Teammate1GrillerKills,
		Teammate1MawsKills:      shift.Teammate1MawsKills,
		Teammate1ScrapperKills:  shift.Teammate1ScrapperKills,
		Teammate1SteelEelKills:  shift.Teammate1SteelEelKills,
		Teammate1SteelheadKills: shift.Teammate1SteelheadKills,
		Teammate1StingerKills:   shift.Teammate1StingerKills,
		Teammate2SplatnetId:     shift.Teammate2SplatnetId,
		Teammate2Name:           shift.Teammate2Name,
		Teammate2DeathCount:     shift.Teammate2DeathCount,
		Teammate2ReviveCount:    shift.Teammate2ReviveCount,
		Teammate2GoldenEggs:     shift.Teammate2GoldenEggs,
		Teammate2PowerEggs:      shift.Teammate2PowerEggs,
		Teammate2Special:        shift.Teammate2Special,
		Teammate2Species:        shift.Teammate2Species,
		Teammate2Gender:         shift.Teammate2Gender,
		Teammate2W1Specials:     shift.Teammate2W1Specials,
		Teammate2W2Specials:     shift.Teammate2W2Specials,
		Teammate2W3Specials:     shift.Teammate2W3Specials,
		Teammate2W1Weapon:       shift.Teammate2W1Weapon,
		Teammate2W2Weapon:       shift.Teammate2W2Weapon,
		Teammate2W3Weapon:       shift.Teammate2W3Weapon,
		Teammate2DrizzlerKills:  shift.Teammate2DrizzlerKills,
		Teammate2FlyfishKills:   shift.Teammate2FlyfishKills,
		Teammate2GoldieKills:    shift.Teammate2GoldieKills,
		Teammate2GrillerKills:   shift.Teammate2GrillerKills,
		Teammate2MawsKills:      shift.Teammate2MawsKills,
		Teammate2ScrapperKills:  shift.Teammate2ScrapperKills,
		Teammate2SteelEelKills:  shift.Teammate2SteelEelKills,
		Teammate2SteelheadKills: shift.Teammate2SteelheadKills,
		Teammate2StingerKills:   shift.Teammate2StingerKills,
		ScheduleEndTime:         shift.ScheduleEndTime,
		ScheduleStartTime:       *shift.ScheduleStartTime,
		ScheduleWeapon0:         shift.ScheduleWeapon0,
		ScheduleWeapon1:         shift.ScheduleWeapon1,
		ScheduleWeapon2:         shift.ScheduleWeapon2,
		ScheduleWeapon3:         shift.ScheduleWeapon3,
		Wave1WaterLevel:         shift.Wave1WaterLevel,
		Wave1EventType:          shift.Wave1EventType,
		Wave1GoldenDelivered:    shift.Wave1GoldenIkuraNum,
		Wave1GoldenAppear:       shift.Wave1GoldenIkuraPopNum,
		Wave1PowerEggs:          shift.Wave1IkuraNum,
		Wave1Quota:              shift.Wave1QuotaNum,
		Wave2WaterLevel:         shift.Wave2WaterLevel,
		Wave2EventType:          shift.Wave2EventType,
		Wave2GoldenDelivered:    shift.Wave2GoldenIkuraNum,
		Wave2GoldenAppear:       shift.Wave2GoldenIkuraPopNum,
		Wave2PowerEggs:          shift.Wave2IkuraNum,
		Wave2Quota:              shift.Wave2QuotaNum,
		Wave3WaterLevel:         shift.Wave3WaterLevel,
		Wave3EventType:          shift.Wave3EventType,
		Wave3GoldenDelivered:    shift.Wave3GoldenIkuraNum,
		Wave3GoldenAppear:       shift.Wave3GoldenIkuraPopNum,
		Wave3PowerEggs:          shift.Wave3IkuraNum,
		Wave3Quota:              shift.Wave3QuotaNum,
	}, nil
}

func GetShiftLeanPk(pk int64) (*api_objects.Shift, error) {
	shift := db_objects.Shift{}
	if err := readObjWithId(pk, "two_salmon_shift").Scan(
		&pk,
		&shift.UserId, &shift.PlayerSplatnetId, &shift.JobId,
		&shift.SplatnetJson, &shift.StatInkJson, &shift.StartTime, &shift.PlayTime, &shift.EndTime, &shift.DangerRate,
		&shift.IsClear, &shift.JobFailureReason, &shift.FailureWave, &shift.GradePoint, &shift.GradePointDelta,
		&shift.JobScore,
		&shift.DrizzlerCount, &shift.FlyfishCount, &shift.GoldieCount, &shift.GrillerCount,
		&shift.MawsCount, &shift.ScrapperCount, &shift.SteelEelCount, &shift.SteelheadCount, &shift.StingerCount,
		&shift.Stage,
		&shift.PlayerName, &shift.PlayerDeathCount, &shift.PlayerReviveCount, &shift.PlayerGoldenEggs,
		&shift.PlayerPowerEggs, &shift.PlayerSpecial, &shift.PlayerTitle, &shift.PlayerSpecies, &shift.PlayerGender,
		&shift.PlayerW1Specials, &shift.PlayerW2Specials, &shift.PlayerW3Specials, &shift.PlayerW1Weapon,
		&shift.PlayerW2Weapon, &shift.PlayerW3Weapon, &shift.PlayerDrizzlerKills, &shift.PlayerFlyfishKills,
		&shift.PlayerGoldieKills, &shift.PlayerGrillerKills, &shift.PlayerMawsKills, &shift.PlayerScrapperKills,
		&shift.PlayerSteelEelKills, &shift.PlayerSteelheadKills, &shift.PlayerStingerKills,
		&shift.Teammate0SplatnetId, &shift.Teammate0Name, &shift.Teammate0DeathCount, &shift.Teammate0ReviveCount,
		&shift.Teammate0GoldenEggs, &shift.Teammate0PowerEggs, &shift.Teammate0Special, &shift.Teammate0Species,
		&shift.Teammate0Gender, &shift.Teammate0W1Specials, &shift.Teammate0W2Specials, &shift.Teammate0W3Specials,
		&shift.Teammate0W1Weapon, &shift.Teammate0W2Weapon, &shift.Teammate0W3Weapon, &shift.Teammate0DrizzlerKills,
		&shift.Teammate0FlyfishKills, &shift.Teammate0GoldieKills, &shift.Teammate0GrillerKills,
		&shift.Teammate0MawsKills, &shift.Teammate0ScrapperKills, &shift.Teammate0SteelEelKills,
		&shift.Teammate0SteelheadKills, &shift.Teammate0StingerKills,
		&shift.Teammate1SplatnetId, &shift.Teammate1Name, &shift.Teammate1DeathCount, &shift.Teammate1ReviveCount,
		&shift.Teammate1GoldenEggs, &shift.Teammate1PowerEggs, &shift.Teammate1Special, &shift.Teammate1Species,
		&shift.Teammate1Gender, &shift.Teammate1W1Specials, &shift.Teammate1W2Specials, &shift.Teammate1W3Specials,
		&shift.Teammate1W1Weapon, &shift.Teammate1W2Weapon, &shift.Teammate1W3Weapon, &shift.Teammate1DrizzlerKills,
		&shift.Teammate1FlyfishKills, &shift.Teammate1GoldieKills, &shift.Teammate1GrillerKills,
		&shift.Teammate1MawsKills, &shift.Teammate1ScrapperKills, &shift.Teammate1SteelEelKills,
		&shift.Teammate1SteelheadKills, &shift.Teammate1StingerKills,
		&shift.Teammate2SplatnetId, &shift.Teammate2Name, &shift.Teammate2DeathCount, &shift.Teammate2ReviveCount,
		&shift.Teammate2GoldenEggs, &shift.Teammate2PowerEggs, &shift.Teammate2Special, &shift.Teammate2Species,
		&shift.Teammate2Gender, &shift.Teammate2W1Specials, &shift.Teammate2W2Specials, &shift.Teammate2W3Specials,
		&shift.Teammate2W1Weapon, &shift.Teammate2W2Weapon, &shift.Teammate2W3Weapon, &shift.Teammate2DrizzlerKills,
		&shift.Teammate2FlyfishKills, &shift.Teammate2GoldieKills, &shift.Teammate2GrillerKills,
		&shift.Teammate2MawsKills, &shift.Teammate2ScrapperKills, &shift.Teammate2SteelEelKills,
		&shift.Teammate2SteelheadKills, &shift.Teammate2StingerKills,
		&shift.ScheduleEndTime, &shift.ScheduleStartTime, &shift.ScheduleWeapon0, &shift.ScheduleWeapon1,
		&shift.ScheduleWeapon2, &shift.ScheduleWeapon3,
		&shift.Wave1WaterLevel, &shift.Wave1EventType, &shift.Wave1GoldenIkuraNum, &shift.Wave1GoldenIkuraPopNum,
		&shift.Wave1IkuraNum, &shift.Wave1QuotaNum, &shift.Wave2WaterLevel, &shift.Wave2EventType,
		&shift.Wave2GoldenIkuraNum, &shift.Wave2GoldenIkuraPopNum, &shift.Wave2IkuraNum, &shift.Wave2QuotaNum,
		&shift.Wave3WaterLevel, &shift.Wave3EventType, &shift.Wave3GoldenIkuraNum, &shift.Wave3GoldenIkuraPopNum,
		&shift.Wave3IkuraNum, &shift.Wave3QuotaNum,
	); err != nil {
		return nil, err
	}
	return &api_objects.Shift{
		UserId:                  shift.UserId,
		PlayerSplatnetId:        shift.PlayerSplatnetId,
		JobId:                   shift.JobId,
		SplatnetJson:            nil,
		StatInkJson:             nil,
		StartTime:               shift.StartTime,
		PlayTime:                shift.PlayTime,
		EndTime:                 shift.EndTime,
		DangerRate:              shift.DangerRate,
		IsClear:                 shift.IsClear,
		JobFailureReason:        shift.JobFailureReason,
		FailureWave:             shift.FailureWave,
		GradePoint:              shift.GradePoint,
		GradePointDelta:         shift.GradePointDelta,
		JobScore:                shift.JobScore,
		DrizzlerCount:           shift.DrizzlerCount,
		FlyfishCount:            shift.FlyfishCount,
		GoldieCount:             shift.GoldieCount,
		GrillerCount:            shift.GrillerCount,
		MawsCount:               shift.MawsCount,
		ScrapperCount:           shift.ScrapperCount,
		SteelEelCount:           shift.SteelEelCount,
		SteelheadCount:          shift.SteelheadCount,
		StingerCount:            shift.StingerCount,
		Stage:                   shift.Stage,
		PlayerName:              shift.PlayerName,
		PlayerDeathCount:        shift.PlayerDeathCount,
		PlayerReviveCount:       shift.PlayerReviveCount,
		PlayerGoldenEggs:        shift.PlayerGoldenEggs,
		PlayerPowerEggs:         shift.PlayerPowerEggs,
		PlayerSpecial:           shift.PlayerSpecial,
		PlayerTitle:             shift.PlayerTitle,
		PlayerSpecies:           shift.PlayerSpecies,
		PlayerGender:            shift.PlayerGender,
		PlayerW1Specials:        shift.PlayerW1Specials,
		PlayerW2Specials:        shift.PlayerW2Specials,
		PlayerW3Specials:        shift.PlayerW3Specials,
		PlayerW1Weapon:          shift.PlayerW1Weapon,
		PlayerW2Weapon:          shift.PlayerW2Weapon,
		PlayerW3Weapon:          shift.PlayerW3Weapon,
		PlayerDrizzlerKills:     shift.PlayerDrizzlerKills,
		PlayerFlyfishKills:      shift.PlayerFlyfishKills,
		PlayerGoldieKills:       shift.PlayerGoldieKills,
		PlayerGrillerKills:      shift.PlayerGrillerKills,
		PlayerMawsKills:         shift.PlayerMawsKills,
		PlayerScrapperKills:     shift.PlayerScrapperKills,
		PlayerSteelEelKills:     shift.PlayerSteelEelKills,
		PlayerSteelheadKills:    shift.PlayerSteelheadKills,
		PlayerStingerKills:      shift.PlayerStingerKills,
		Teammate0SplatnetId:     shift.Teammate0SplatnetId,
		Teammate0Name:           shift.Teammate0Name,
		Teammate0DeathCount:     shift.Teammate0DeathCount,
		Teammate0ReviveCount:    shift.Teammate0ReviveCount,
		Teammate0GoldenEggs:     shift.Teammate0GoldenEggs,
		Teammate0PowerEggs:      shift.Teammate0PowerEggs,
		Teammate0Special:        shift.Teammate0Special,
		Teammate0Species:        shift.Teammate0Species,
		Teammate0Gender:         shift.Teammate0Gender,
		Teammate0W1Specials:     shift.Teammate0W1Specials,
		Teammate0W2Specials:     shift.Teammate0W2Specials,
		Teammate0W3Specials:     shift.Teammate0W3Specials,
		Teammate0W1Weapon:       shift.Teammate0W1Weapon,
		Teammate0W2Weapon:       shift.Teammate0W2Weapon,
		Teammate0W3Weapon:       shift.Teammate0W3Weapon,
		Teammate0DrizzlerKills:  shift.Teammate0DrizzlerKills,
		Teammate0FlyfishKills:   shift.Teammate0FlyfishKills,
		Teammate0GoldieKills:    shift.Teammate0GoldieKills,
		Teammate0GrillerKills:   shift.Teammate0GrillerKills,
		Teammate0MawsKills:      shift.Teammate0MawsKills,
		Teammate0ScrapperKills:  shift.Teammate0ScrapperKills,
		Teammate0SteelEelKills:  shift.Teammate0SteelEelKills,
		Teammate0SteelheadKills: shift.Teammate0SteelheadKills,
		Teammate0StingerKills:   shift.Teammate0StingerKills,
		Teammate1SplatnetId:     shift.Teammate1SplatnetId,
		Teammate1Name:           shift.Teammate1Name,
		Teammate1DeathCount:     shift.Teammate1DeathCount,
		Teammate1ReviveCount:    shift.Teammate1ReviveCount,
		Teammate1GoldenEggs:     shift.Teammate1GoldenEggs,
		Teammate1PowerEggs:      shift.Teammate1PowerEggs,
		Teammate1Special:        shift.Teammate1Special,
		Teammate1Species:        shift.Teammate1Species,
		Teammate1Gender:         shift.Teammate1Gender,
		Teammate1W1Specials:     shift.Teammate1W1Specials,
		Teammate1W2Specials:     shift.Teammate1W2Specials,
		Teammate1W3Specials:     shift.Teammate1W3Specials,
		Teammate1W1Weapon:       shift.Teammate1W1Weapon,
		Teammate1W2Weapon:       shift.Teammate1W2Weapon,
		Teammate1W3Weapon:       shift.Teammate1W3Weapon,
		Teammate1DrizzlerKills:  shift.Teammate1DrizzlerKills,
		Teammate1FlyfishKills:   shift.Teammate1FlyfishKills,
		Teammate1GoldieKills:    shift.Teammate1GoldieKills,
		Teammate1GrillerKills:   shift.Teammate1GrillerKills,
		Teammate1MawsKills:      shift.Teammate1MawsKills,
		Teammate1ScrapperKills:  shift.Teammate1ScrapperKills,
		Teammate1SteelEelKills:  shift.Teammate1SteelEelKills,
		Teammate1SteelheadKills: shift.Teammate1SteelheadKills,
		Teammate1StingerKills:   shift.Teammate1StingerKills,
		Teammate2SplatnetId:     shift.Teammate2SplatnetId,
		Teammate2Name:           shift.Teammate2Name,
		Teammate2DeathCount:     shift.Teammate2DeathCount,
		Teammate2ReviveCount:    shift.Teammate2ReviveCount,
		Teammate2GoldenEggs:     shift.Teammate2GoldenEggs,
		Teammate2PowerEggs:      shift.Teammate2PowerEggs,
		Teammate2Special:        shift.Teammate2Special,
		Teammate2Species:        shift.Teammate2Species,
		Teammate2Gender:         shift.Teammate2Gender,
		Teammate2W1Specials:     shift.Teammate2W1Specials,
		Teammate2W2Specials:     shift.Teammate2W2Specials,
		Teammate2W3Specials:     shift.Teammate2W3Specials,
		Teammate2W1Weapon:       shift.Teammate2W1Weapon,
		Teammate2W2Weapon:       shift.Teammate2W2Weapon,
		Teammate2W3Weapon:       shift.Teammate2W3Weapon,
		Teammate2DrizzlerKills:  shift.Teammate2DrizzlerKills,
		Teammate2FlyfishKills:   shift.Teammate2FlyfishKills,
		Teammate2GoldieKills:    shift.Teammate2GoldieKills,
		Teammate2GrillerKills:   shift.Teammate2GrillerKills,
		Teammate2MawsKills:      shift.Teammate2MawsKills,
		Teammate2ScrapperKills:  shift.Teammate2ScrapperKills,
		Teammate2SteelEelKills:  shift.Teammate2SteelEelKills,
		Teammate2SteelheadKills: shift.Teammate2SteelheadKills,
		Teammate2StingerKills:   shift.Teammate2StingerKills,
		ScheduleEndTime:         shift.ScheduleEndTime,
		ScheduleStartTime:       *shift.ScheduleStartTime,
		ScheduleWeapon0:         shift.ScheduleWeapon0,
		ScheduleWeapon1:         shift.ScheduleWeapon1,
		ScheduleWeapon2:         shift.ScheduleWeapon2,
		ScheduleWeapon3:         shift.ScheduleWeapon3,
		Wave1WaterLevel:         shift.Wave1WaterLevel,
		Wave1EventType:          shift.Wave1EventType,
		Wave1GoldenDelivered:    shift.Wave1GoldenIkuraNum,
		Wave1GoldenAppear:       shift.Wave1GoldenIkuraPopNum,
		Wave1PowerEggs:          shift.Wave1IkuraNum,
		Wave1Quota:              shift.Wave1QuotaNum,
		Wave2WaterLevel:         shift.Wave2WaterLevel,
		Wave2EventType:          shift.Wave2EventType,
		Wave2GoldenDelivered:    shift.Wave2GoldenIkuraNum,
		Wave2GoldenAppear:       shift.Wave2GoldenIkuraPopNum,
		Wave2PowerEggs:          shift.Wave2IkuraNum,
		Wave2Quota:              shift.Wave2QuotaNum,
		Wave3WaterLevel:         shift.Wave3WaterLevel,
		Wave3EventType:          shift.Wave3EventType,
		Wave3GoldenDelivered:    shift.Wave3GoldenIkuraNum,
		Wave3GoldenAppear:       shift.Wave3GoldenIkuraPopNum,
		Wave3PowerEggs:          shift.Wave3IkuraNum,
		Wave3Quota:              shift.Wave3QuotaNum,
	}, nil
}

func getShiftSplatnet(pk int64) (*api_objects.ShiftSplatnet, error) {
	shift := db_objects.ShiftSplatnet{}
	if err := readObjWithId(pk, "two_salmon_shift_splatnet").Scan(
		&pk, &shift.JobId, &shift.DangerRate, &shift.JobResult, &shift.JobScore, &shift.JobRate, &shift.GradePoint,
		&shift.GradePointDelta, &shift.KumaPoint, &shift.StartTime, &shift.PlayerType,
		&shift.PlayTime, &shift.BossCounts, &shift.EndTime, &shift.MyResult, &shift.Grade,
		&shift.Schedule,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	otherResultsKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "player", "two_salmon_shift_splatnet_player_container", "asc", "pk")
	if err != nil {

		return nil, err
	}
	otherResults := make([]api_objects.ShiftSplatnetPlayer, len(otherResultsKeys))
	for i := range otherResultsKeys {
		otherResultsTemp, err := getShiftSplatnetPlayer(otherResultsKeys[i])
		if err != nil {

			return nil, err
		}
		otherResults[i] = *otherResultsTemp
	}
	waveDetailsKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "pk", "two_salmon_shift_splatnet_wave", "asc", "pk")
	if err != nil {

		return nil, err
	}
	waveDetails := make([]api_objects.ShiftSplatnetWave, len(waveDetailsKeys))
	for i := range waveDetailsKeys {
		waveDetailTemp, err := getShiftSplatnetWave(waveDetailsKeys[i])
		if err != nil {

			return nil, err
		}
		waveDetails[i] = *waveDetailTemp
	}
	myResult, err := getShiftSplatnetPlayer(shift.MyResult)
	if err != nil {

		return nil, err
	}
	playerType, err := getSplatnetPlayerType(shift.PlayerType)
	if err != nil {

		return nil, err
	}
	bossCounts, err := getShiftSplatnetBossCounts(shift.BossCounts)
	if err != nil {

		return nil, err
	}
	grade, err := getShiftSplatnetGrade(shift.Grade)
	if err != nil {

		return nil, err
	}
	schedule, err := getShiftSplatnetSchedule(shift.Schedule)
	if err != nil {

		return nil, err
	}
	jobResult, err := getShiftSplatnetJobResult(shift.JobResult)
	if err != nil {

		return nil, err
	}
	return &api_objects.ShiftSplatnet{
		JobId:           shift.JobId,
		DangerRate:      shift.DangerRate,
		JobResult:       *jobResult,
		JobScore:        shift.JobScore,
		JobRate:         shift.JobRate,
		GradePoint:      shift.GradePoint,
		GradePointDelta: shift.GradePointDelta,
		OtherResults:    otherResults,
		KumaPoint:       shift.KumaPoint,
		StartTime:       shift.StartTime,
		PlayerType:      *playerType,
		PlayTime:        shift.PlayTime,
		BossCounts:      *bossCounts,
		EndTime:         shift.EndTime,
		MyResult:        *myResult,
		WaveDetails:     waveDetails,
		Grade:           *grade,
		Schedule:        *schedule,
	}, nil
}

func getShiftSplatnetPlayer(pk int64) (*api_objects.ShiftSplatnetPlayer, error) {
	player := db_objects.ShiftSplatnetPlayer{}
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_player").Scan(
		&pk, &player.Special, &player.Pid, &player.PlayerType,
		&player.Name, &player.DeadCount, &player.GoldenIkuraNum, &player.BossKillCounts, &player.IkuraNum,
		&player.HelpCount,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	weaponListKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "pk", "two_salmon_shift_splatnet_player_weapon_list", "asc", "pk")
	if err != nil {

		return nil, err
	}
	weaponList := make([]api_objects.ShiftSplatnetPlayerWeaponList, len(weaponListKeys))
	for i := range weaponListKeys {
		weaponListTemp, err := getShiftSplatnetPlayerWeaponList(weaponListKeys[i])
		if err != nil {

			return nil, err
		}
		weaponList[i] = *weaponListTemp
	}
	specialCountKeys, err := ReadKeyArrayWithKeyTable(fmt.Sprint(pk), "parent", "two_salmon_shift_splatnet_player_weapon_list", "parent_table", "pk", "int_container")
	if err != nil {

		return nil, err
	}
	specialCounts := make([]int, len(specialCountKeys))
	for i := range specialCountKeys {
		specialCountsTemp, err := getIntContainer(specialCountKeys[i])
		if err != nil {

			return nil, err
		}
		specialCounts[i] = *specialCountsTemp
	}
	playerType, err := getSplatnetPlayerType(player.PlayerType)
	if err != nil {

		return nil, err
	}
	bossKillCounts, err := getShiftSplatnetBossCounts(player.BossKillCounts)
	if err != nil {

		return nil, err
	}
	special, err := getSplatnetQuad(player.Special)
	if err != nil {

		return nil, err
	}
	return &api_objects.ShiftSplatnetPlayer{
		SpecialCounts:  specialCounts,
		Special:        *special,
		Pid:            player.Pid,
		PlayerType:     *playerType,
		WeaponList:     weaponList,
		Name:           player.Name,
		DeadCount:      player.DeadCount,
		GoldenEggs:     player.GoldenIkuraNum,
		BossKillCounts: *bossKillCounts,
		PowerEggs:      player.IkuraNum,
		HelpCount:      player.HelpCount,
	}, nil
}

func getShiftSplatnetPlayerWeaponList(pk int64) (*api_objects.ShiftSplatnetPlayerWeaponList, error) {
	weaponList := db_objects.ShiftSplatnetPlayerWeaponList{}
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_player_weapon_list").Scan(
		&pk, &weaponList.Parent, &weaponList.Id, &weaponList.Weapon,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	weapon, err := getShiftSplatnetPlayerWeaponListWeapon(weaponList.Weapon)
	if err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &api_objects.ShiftSplatnetPlayerWeaponList{
		Id:     weaponList.Id,
		Weapon: *weapon,
	}, nil
}

func getShiftSplatnetPlayerWeaponListWeapon(pk int64) (*api_objects.ShiftSplatnetPlayerWeaponListWeapon, error) {
	weapon := api_objects.ShiftSplatnetPlayerWeaponListWeapon{}
	var junk int
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_player_weapon_list_weapon").Scan(
		&junk, &weapon.Id, &weapon.Image, &weapon.Name, &weapon.Thumbnail,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &weapon, nil
}

func getShiftSplatnetBossCounts(pk int64) (*api_objects.ShiftSplatnetBossCounts, error) {
	bossCounts := db_objects.ShiftSplatnetBossCounts{}
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_boss_counts").Scan(
		&pk, &bossCounts.Var3, &bossCounts.Var6, &bossCounts.Var9, &bossCounts.Var12, &bossCounts.Var13,
		&bossCounts.Var14, &bossCounts.Var15, &bossCounts.Var16, &bossCounts.Var21,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	var3, err := getShiftSplatnetBossCountsBoss(bossCounts.Var3)
	if err != nil {

		return nil, err
	}
	var6, err := getShiftSplatnetBossCountsBoss(bossCounts.Var6)
	if err != nil {

		return nil, err
	}
	var9, err := getShiftSplatnetBossCountsBoss(bossCounts.Var9)
	if err != nil {

		return nil, err
	}
	var12, err := getShiftSplatnetBossCountsBoss(bossCounts.Var12)
	if err != nil {

		return nil, err
	}
	var13, err := getShiftSplatnetBossCountsBoss(bossCounts.Var13)
	if err != nil {

		return nil, err
	}
	var14, err := getShiftSplatnetBossCountsBoss(bossCounts.Var14)
	if err != nil {

		return nil, err
	}
	var15, err := getShiftSplatnetBossCountsBoss(bossCounts.Var15)
	if err != nil {

		return nil, err
	}
	var16, err := getShiftSplatnetBossCountsBoss(bossCounts.Var16)
	if err != nil {

		return nil, err
	}
	var21, err := getShiftSplatnetBossCountsBoss(bossCounts.Var21)
	if err != nil {

		return nil, err
	}
	return &api_objects.ShiftSplatnetBossCounts{
		Goldie:    *var3,
		Steelhead: *var6,
		Flyfish:   *var9,
		Scrapper:  *var12,
		SteelEel:  *var13,
		Stinger:   *var14,
		Maws:      *var15,
		Griller:   *var16,
		Drizzler:  *var21,
	}, nil
}

func getShiftSplatnetBossCountsBoss(pk int64) (*api_objects.ShiftSplatnetBossCountsBoss, error) {
	boss := db_objects.ShiftSplatnetBossCountsBoss{}
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_boss_counts_boss").Scan(
		&pk, &boss.Boss, &boss.Count,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	bossInside, err := getSplatnetDouble(boss.Boss)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftSplatnetBossCountsBoss{
		Boss:  splatnetDoubleToShiftSplatnetBossCountsBossDouble(*bossInside),
		Count: boss.Count,
	}, nil
}

func splatnetDoubleToShiftSplatnetBossCountsBossDouble(double api_objects.SplatnetDouble) api_objects.ShiftSplatnetBossCountsBossDouble {
	return api_objects.ShiftSplatnetBossCountsBossDouble{
		Key:  enums.ShiftSplatnetBossCountsBossDoubleKeyEnum(double.Key),
		Name: double.Name,
	}
}

func getShiftSplatnetWave(pk int64) (*api_objects.ShiftSplatnetWave, error) {
	wave := db_objects.ShiftSplatnetWave{}
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_wave").Scan(
		&pk, &wave.Parent, &wave.WaterLevel, &wave.EventType, &wave.GoldenIkuraNum, &wave.GoldenIkuraPopNum,
		&wave.IkuraNum, &wave.QuotaNum,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	waterLevel, err := getSplatnetDouble(wave.WaterLevel)
	if err != nil {

		return nil, err
	}
	eventType, err := getSplatnetDouble(wave.EventType)
	if err != nil {

		return nil, err
	}
	return &api_objects.ShiftSplatnetWave{
		WaterLevel:   *waterLevel,
		EventType:    *eventType,
		GoldenEggs:   wave.GoldenIkuraNum,
		GoldenAppear: wave.GoldenIkuraPopNum,
		PowerEggs:    wave.IkuraNum,
		QuotaNum:     wave.QuotaNum,
	}, nil
}

func getShiftSplatnetJobResult(pk int64) (*api_objects.ShiftSplatnetJobResult, error) {
	jobResult := api_objects.ShiftSplatnetJobResult{}
	var junk int
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_job_result").Scan(
		&junk, &jobResult.IsClear, &jobResult.FailureReason, &jobResult.FailureWave,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &jobResult, nil
}

func getShiftSplatnetGrade(pk int64) (*api_objects.ShiftSplatnetGrade, error) {
	grade := api_objects.ShiftSplatnetGrade{}
	var junk int
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_grade").Scan(
		&junk, &grade.Id, &grade.ShortName, &grade.LongName, &grade.Name,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &grade, nil
}

func getShiftSplatnetSchedule(pk int64) (*api_objects.ShiftSplatnetSchedule, error) {
	schedule := db_objects.ShiftSplatnetSchedule{}
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_schedule").Scan(
		&pk, &schedule.StartTime, &schedule.EndTime, &schedule.Stage,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	weaponKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "pk", "two_salmon_shift_splatnet_schedule_weapon", "asc", "pk")
	if err != nil {

		return nil, err
	}
	weapons := make([]api_objects.ShiftSplatnetScheduleWeapon, len(weaponKeys))
	for i := range weaponKeys {
		weaponTemp, err := getShiftSplatnetScheduleWeapon(weaponKeys[i])
		if err != nil {

			return nil, err
		}
		weapons[i] = *weaponTemp
	}
	stage, err := getShiftSplatnetScheduleStage(schedule.Stage)
	if err != nil {

		return nil, err
	}
	return &api_objects.ShiftSplatnetSchedule{
		StartTime: schedule.StartTime,
		Weapons:   weapons,
		EndTime:   schedule.EndTime,
		Stage:     *stage,
	}, nil
}

func getShiftSplatnetScheduleWeapon(pk int64) (*api_objects.ShiftSplatnetScheduleWeapon, error) {
	weapon := db_objects.ShiftSplatnetScheduleWeapon{}
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_schedule_weapon").Scan(
		&pk, &weapon.Parent, &weapon.Id, &weapon.Weapon, &weapon.CoopSpecialWeapon,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	var weaponInside *api_objects.ShiftSplatnetScheduleWeaponWeapon
	if weapon.Weapon != nil {
		var err error
		weaponInside, err = getShiftSplatnetScheduleWeaponWeapon(*weapon.Weapon)
		if err != nil {

			return nil, err
		}
	}
	var special *api_objects.ShiftSplatnetScheduleWeaponSpecialWeapon
	if weapon.CoopSpecialWeapon != nil {
		var err error
		special, err = getShiftSplatnetScheduleWeaponSpecialWeapon(*weapon.CoopSpecialWeapon)
		if err != nil {

			return nil, err
		}
	}
	return &api_objects.ShiftSplatnetScheduleWeapon{
		Id:                weapon.Id,
		Weapon:            weaponInside,
		CoopSpecialWeapon: special,
	}, nil
}

func getShiftSplatnetScheduleWeaponWeapon(pk int64) (*api_objects.ShiftSplatnetScheduleWeaponWeapon, error) {
	weapon := api_objects.ShiftSplatnetScheduleWeaponWeapon{}
	var junk int
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_schedule_weapon_weapon").Scan(
		&junk, &weapon.Id, &weapon.Image, &weapon.Name, &weapon.Thumbnail,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &weapon, nil
}

func getShiftSplatnetScheduleWeaponSpecialWeapon(pk int64) (*api_objects.ShiftSplatnetScheduleWeaponSpecialWeapon, error) {
	weapon := api_objects.ShiftSplatnetScheduleWeaponSpecialWeapon{}
	var junk int
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_schedule_weapon_special_weapon").Scan(
		&junk, &weapon.Image, &weapon.Name,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &weapon, nil
}

func getShiftSplatnetScheduleStage(pk int64) (*api_objects.ShiftSplatnetScheduleStage, error) {
	stage := api_objects.ShiftSplatnetScheduleStage{}
	var junk int
	if err := readObjWithId(pk, "two_salmon_shift_splatnet_schedule_stage").Scan(
		&junk, &stage.Image, &stage.Name,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &stage, nil
}

func getShiftStatInk(pk int64) (*api_objects.ShiftStatInk, error) {
	shift := db_objects.ShiftStatInk{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink").Scan(
		&pk, &shift.Id, &shift.Uuid, &shift.SplatnetNumber, &shift.Url, &shift.ApiEndpoint, &shift.User,
		&shift.Stage, &shift.IsCleared, &shift.FailReason, &shift.ClearWaves, &shift.DangerRate,
		&shift.Title, &shift.TitleExp, &shift.TitleAfter, &shift.TitleExpAfter,
		&shift.MyData, &shift.Agent, &shift.Automated,
		//&shift.Note, &shift.LinkUrl,
		&shift.ShiftStartAt, &shift.StartAt,
		//&shift.EndAt,
		&shift.RegisterAt,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	var failReason *api_objects.ShiftStatInkFailReason
	if shift.FailReason != nil {
		var err error
		failReason, err = readShiftStatInkFailReason(*shift.FailReason)
		if err != nil {
			return nil, err
		}
	}
	bossAppearanceKeys, err := ReadKeyArrayWithKeyTable(fmt.Sprint(pk), "parent", "two_salmon_shift_stat_ink", "parent_table", "pk", "two_salmon_shift_stat_ink_boss_data")
	if err != nil {
		return nil, err
	}
	bossAppearances := make([]api_objects.ShiftStatInkBossData, len(bossAppearanceKeys))
	for i := range bossAppearanceKeys {
		bossAppearanceTemp, err := getShiftStatInkBossData(bossAppearanceKeys[i])
		if err != nil {
			return nil, err
		}
		bossAppearances[i] = *bossAppearanceTemp
	}
	waveKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "pk", "two_salmon_shift_stat_ink_wave", "asc", "pk")
	if err != nil {
		return nil, err
	}
	waves := make([]api_objects.ShiftStatInkWave, len(waveKeys))
	for i := range waveKeys {
		waveTemp, err := getShiftStatInkWave(waveKeys[i])
		if err != nil {
			return nil, err
		}
		waves[i] = *waveTemp
	}
	teammateKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "pk", "two_salmon_shift_stat_ink_player_container", "asc", "pk")
	if err != nil {
		return nil, err
	}
	teammates := make([]api_objects.ShiftStatInkPlayer, len(teammateKeys))
	for i := range teammateKeys {
		teammateTemp, err := getShiftStatInkPlayer(teammateKeys[i])
		if err != nil {
			return nil, err
		}
		teammates[i] = *teammateTemp
	}
	quotaKeys, err := ReadKeyArrayWithKeyTable(fmt.Sprint(pk), "parent", "two_salmon_shift_stat_ink", "parent_table", "pk", "int_container")
	if err != nil {
		return nil, err
	}
	quota := make([]int, len(quotaKeys))
	for i := range quotaKeys {
		quotaTemp, err := getIntContainer(quotaKeys[i])
		if err != nil {
			return nil, err
		}
		quota[i] = *quotaTemp
	}
	myData, err := getShiftStatInkPlayer(shift.MyData)
	if err != nil {
		return nil, err
	}
	agent, err := getShiftStatInkAgent(shift.Agent)
	if err != nil {
		return nil, err
	}
	title, err := getShiftStatInkTitle(shift.Title)
	if err != nil {
		return nil, err
	}
	titleAfter, err := getShiftStatInkTitle(shift.TitleAfter)
	if err != nil {
		return nil, err
	}
	user, err := getShiftStatInkUser(shift.User)
	if err != nil {
		return nil, err
	}
	stage, err := getShiftStatInkStage(shift.Stage)
	if err != nil {
		return nil, err
	}
	shiftStartAt, err := readStatInkTime(shift.ShiftStartAt)
	if err != nil {
		return nil, err
	}
	startAt, err := readStatInkTime(shift.StartAt)
	if err != nil {
		return nil, err
	}
	registerAt, err := readStatInkTime(shift.RegisterAt)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInk{
		Id:              shift.Id,
		Uuid:            shift.Uuid,
		SplatnetNumber:  shift.SplatnetNumber,
		Url:             shift.Url,
		ApiEndpoint:     shift.ApiEndpoint,
		User:            *user,
		Stage:           *stage,
		IsCleared:       shift.IsCleared,
		FailReason:      failReason,
		ClearWaves:      shift.ClearWaves,
		DangerRate:      shift.DangerRate,
		Quota:           quota,
		Title:           *title,
		TitleExp:        shift.TitleExp,
		TitleAfter:      *titleAfter,
		TitleExpAfter:   shift.TitleExpAfter,
		BossAppearances: bossAppearances,
		Waves:           waves,
		MyData:          *myData,
		Teammates:       teammates,
		Agent:           *agent,
		Automated:       shift.Automated,
		ShiftStartAt:    *shiftStartAt,
		StartAt:         *startAt,
		RegisterAt:      *registerAt,
	}, nil
}

func getShiftStatInkBossData(pk int64) (*api_objects.ShiftStatInkBossData, error) {
	bossData := db_objects.ShiftStatInkBossData{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_boss_data").Scan(
		&pk, &bossData.Parent, &bossData.ParentTable, &bossData.Boss, &bossData.Count,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	boss, err := getShiftStatInkBossDataBoss(bossData.Boss)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInkBossData{
		Boss:  *boss,
		Count: bossData.Count,
	}, nil
}

func getShiftStatInkBossDataBoss(pk int64) (*api_objects.ShiftStatInkBossDataBoss, error) {
	boss := api_objects.ShiftStatInkBossDataBoss{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_boss_data_boss").Scan(
		&pk, &boss.Splatnet, &boss.SplatnetStr,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &boss, nil
}

func getShiftStatInkWave(pk int64) (*api_objects.ShiftStatInkWave, error) {
	wave := db_objects.ShiftStatInkWave{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_wave").Scan(
		&pk, &wave.Parent, &wave.KnownOccurrence, &wave.WaterLevel, &wave.GoldenEggQuota, &wave.GoldenEggAppearances,
		&wave.GoldenEggDelivered, &wave.PowerEggCollected,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	var knownOccurrence *api_objects.ShiftStatInkTripleString
	if wave.KnownOccurrence != nil {
		var err error
		knownOccurrence, err = getShiftStatInkTripleString(*wave.KnownOccurrence)
		if err != nil {
			return nil, err
		}
	}
	var waterLevel *api_objects.ShiftStatInkTripleString
	if wave.WaterLevel != nil {
		var err error
		waterLevel, err = getShiftStatInkTripleString(*wave.WaterLevel)
		if err != nil {
			return nil, err
		}
	}
	return &api_objects.ShiftStatInkWave{
		KnownOccurrence:      knownOccurrence,
		WaterLevel:           waterLevel,
		GoldenEggQuota:       wave.GoldenEggQuota,
		GoldenEggAppearances: wave.GoldenEggAppearances,
		GoldenEggDelivered:   wave.GoldenEggDelivered,
		PowerEggCollected:    wave.PowerEggCollected,
	}, nil
}

func getShiftStatInkTriple(pk int64) (*api_objects.ShiftStatInkTripleInt, error) {
	triple := db_objects.ShiftStatInkTriple{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_triple").Scan(
		&pk, &triple.Key, &triple.Name, &triple.Splatnet,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	name, err := readStatInkName(triple.Name)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInkTripleInt{
		Key:      triple.Key,
		Name:     *name,
		Splatnet: triple.Splatnet,
	}, nil
}

func getShiftStatInkTripleString(pk int64) (*api_objects.ShiftStatInkTripleString, error) {
	triple := db_objects.ShiftStatInkTripleString{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_triple_string").Scan(
		&pk, &triple.Key, &triple.Name, &triple.Splatnet,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	name, err := readStatInkName(triple.Name)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInkTripleString{
		Key:      triple.Key,
		Name:     *name,
		Splatnet: triple.Splatnet,
	}, nil
}

func getShiftStatInkStage(pk int64) (*api_objects.ShiftStatInkStage, error) {
	triple := db_objects.ShiftStatInkTripleString{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_triple_string").Scan(
		&pk, &triple.Key, &triple.Name, &triple.Splatnet,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	name, err := readShiftStatInkStageName(triple.Name)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInkStage{
		Key:      triple.Key,
		Name:     *name,
		Splatnet: triple.Splatnet,
	}, nil
}

func getShiftStatInkPlayer(pk int64) (*api_objects.ShiftStatInkPlayer, error) {
	player := db_objects.ShiftStatInkPlayer{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_player").Scan(
		&pk, &player.SplatnetId, &player.Name, &player.Special, &player.Rescue, &player.Death,
		&player.GoldenEggDelivered, &player.PowerEggCollected, &player.Species, &player.Gender,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	specialUseKeys, err := ReadKeyArrayWithKeyTable(fmt.Sprint(pk), "parent", "two_salmon_shift_stat_ink_player", "parent_table", "pk", "int_container")
	if err != nil {
		return nil, err
	}
	specialUses := make([]int, len(specialUseKeys))
	for i := range specialUseKeys {
		specialUseTemp, err := getIntContainer(specialUseKeys[i])
		if err != nil {
			return nil, err
		}
		specialUses[i] = *specialUseTemp
	}
	weaponKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "triple", "two_salmon_shift_stat_ink_triple_container", "asc", "pk")
	if err != nil {
		return nil, err
	}
	weapons := make([]api_objects.ShiftStatInkTripleInt, len(weaponKeys))
	for i := range weaponKeys {
		weaponTemp, err := getShiftStatInkTriple(weaponKeys[i])
		if err != nil {
			return nil, err
		}
		weapons[i] = *weaponTemp
	}
	bossKillKeys, err := ReadKeyArrayWithKey(fmt.Sprint(pk), "parent", "pk", "two_salmon_shift_stat_ink_boss_data", "asc", "pk")
	if err != nil {
		return nil, err
	}
	bossKills := make([]api_objects.ShiftStatInkBossData, len(bossKillKeys))
	for i := range bossKillKeys {
		bossKillTemp, err := getShiftStatInkBossData(bossKillKeys[i])
		if err != nil {
			return nil, err
		}
		bossKills[i] = *bossKillTemp
	}
	special, err := getShiftStatInkTriple(player.Special)
	if err != nil {
		return nil, err
	}
	species, err := readStatInkKeyName(player.Species)
	if err != nil {
		return nil, err
	}
	gender, err := readStatInkGender(player.Gender)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInkPlayer{
		SplatnetId:         player.SplatnetId,
		Name:               player.Name,
		Special:            *special,
		Rescue:             player.Rescue,
		Death:              player.Death,
		GoldenEggDelivered: player.GoldenEggDelivered,
		PowerEggCollected:  player.PowerEggCollected,
		Species:            *species,
		Gender:             *gender,
		SpecialUses:        specialUses,
		Weapons:            weapons,
		BossKills:          bossKills,
	}, nil
}

func getShiftStatInkUser(pk int64) (*api_objects.ShiftStatInkUser, error) {
	user := db_objects.ShiftStatInkUser{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_user").Scan(
		&pk, &user.Id, &user.Name, &user.ScreenName, &user.Url, &user.SalmonUrl, &user.BattleUrl, &user.JoinAt,
		&user.Profile, &user.Stats,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	stats, err := getShiftStatInkUserStats(user.Stats)
	if err != nil {
		return nil, err
	}
	profile, err := readStatInkProfile(user.Profile)
	if err != nil {
		return nil, err
	}
	joinAt, err := readStatInkTime(user.JoinAt)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInkUser{
		Id:         user.Id,
		Name:       user.Name,
		ScreenName: user.ScreenName,
		Url:        user.Url,
		SalmonUrl:  user.SalmonUrl,
		BattleUrl:  user.BattleUrl,
		JoinAt:     *joinAt,
		Profile:    *profile,
		Stats:      *stats,
	}, nil
}

func getShiftStatInkUserStats(pk int64) (*api_objects.ShiftStatInkUserStats, error) {
	stats := db_objects.ShiftStatInkUserStats{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_user_stats").Scan(
		&pk, &stats.WorkCount, &stats.TotalGoldenEggs, &stats.TotalEggs, &stats.TotalRescued, &stats.TotalPoint,
		&stats.AsOf, &stats.RegisteredAt,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	asOf, err := readStatInkTime(stats.AsOf)
	if err != nil {
		return nil, err
	}
	registeredAt, err := readStatInkTime(stats.RegisteredAt)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInkUserStats{
		WorkCount:       stats.WorkCount,
		TotalGoldenEggs: stats.TotalGoldenEggs,
		TotalEggs:       stats.TotalEggs,
		TotalRescued:    stats.TotalRescued,
		TotalPoint:      stats.TotalPoint,
		AsOf:            *asOf,
		RegisteredAt:    *registeredAt,
	}, nil
}

func getShiftStatInkTitle(pk int64) (*api_objects.ShiftStatInkTitle, error) {
	title := db_objects.ShiftStatInkTitle{}
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_title").Scan(
		&pk, &title.Splatnet, &title.GenericName,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	genericName, err := readStatInkName(title.GenericName)
	if err != nil {
		return nil, err
	}
	return &api_objects.ShiftStatInkTitle{
		Splatnet:    title.Splatnet,
		GenericName: *genericName,
	}, nil
}

func getShiftStatInkAgent(pk int64) (*api_objects.ShiftStatInkAgent, error) {
	agent := api_objects.ShiftStatInkAgent{}
	var junk int
	if err := readObjWithId(pk, "two_salmon_shift_stat_ink_agent").Scan(
		&junk, &agent.Name, &agent.Version,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &agent, nil
}

func ReadUser(username string) (*db_objects.User, error) {
	user := db_objects.User{}
	pk, err := ReadKeyArrayWithKey(username, "username", "pk", "auth_user", "asc", "pk")
	if err != nil {
		return nil, err
	}
	if err := readObjWithId(pk[0], "auth_user").Scan(
		&user.Pk, &user.Username, &user.Password, &user.Email, &user.EmailVerified,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &user, nil
}

func ReadUserById(pk int64) (*db_objects.User, error) {
	user := db_objects.User{}
	if err := readObjWithId(pk, "auth_user").Scan(
		&pk, &user.Username, &user.Password, &user.Email, &user.EmailVerified,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &user, nil
}

func CheckSessionToken(sessionToken string) (*int64, error) {
	token := db_objects.SessionToken{}
	if err := ReadValuesWithKey(sessionToken, "session_token", "auth_user_session_token", []string{
		"parent",
	}).Scan(
		&token.Parent,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &token.Parent, nil
}

func getIntContainer(pk int64) (*int, error) {
	var junk int64
	var parent int64
	var parentTable string
	var result int
	if err := readObjWithId(pk, "int_container").Scan(
		&junk, &parent, &parentTable, &result,
	); err != nil {
		debug.PrintStack()
		return nil, err
	}
	return &result, nil
}
