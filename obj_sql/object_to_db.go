package obj_sql

import (
	"errors"
	"fmt"
	"github.com/cass-dlcm/SplatStatsGo/api_objects"
	"github.com/cass-dlcm/SplatStatsGo/db_objects"
	"log"
	"reflect"
)

func WriteNewUser(user *db_objects.User) error {
	values := make([]interface{}, 4)
	values[0] = (*user).Username
	values[1] = (*user).Password
	values[2] = (*user).Email
	values[3] = false
	return writeIntoTable("auth_user", []string{
		"username", "password", "email", "email_verified",
	}, values)
}

func btoi(b bool) int8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func WriteNewBattle(battle *api_objects.Battle) error {
	_, err := GetBattleLean((*battle).UserId, (*battle).BattleNumber)
	if fmt.Sprint(err) != "sql: no rows in result set" {
		return errors.New("battle already exists")
	}
	var splatnet *int64
	if (*battle).SplatnetJson != nil {
		splatnet, err = writeNewBattleSplatnet((*battle).SplatnetJson)
		if err != nil {
			return err
		}
	}
	var statInk *int64
	if (*battle).StatInkJson != nil {
		statInk, err = writeNewBattleStatInk((*battle).StatInkJson)
		if err != nil {
			return err
		}
	}
	v := reflect.ValueOf(db_objects.Battle{
		UserId:                  (*battle).UserId,
		SplatnetJson:            splatnet,
		StatInkJson:             statInk,
		SplatnetNumber:          (*battle).BattleNumber,
		PlayerSplatnetId:        (*battle).PlayerSplatnetId,
		ElapsedTime:             (*battle).ElapsedTime,
		HasDisconnectedPlayer:   btoi((*battle).HasDisconnectedPlayer),
		LeaguePoint:             (*battle).LeaguePoint,
		MatchType:               (*battle).MatchType,
		Rule:                    (*battle).Rule,
		MyTeamCount:             (*battle).MyTeamCount,
		OtherTeamCount:          (*battle).OtherTeamCount,
		SplatfestPoint:          (*battle).SplatfestPoint,
		SplatfestTitle:          (*battle).SplatfestTitle,
		Stage:                   (*battle).Stage,
		TagId:                   (*battle).TagId,
		Time:                    (*battle).Time,
		Win:                     btoi((*battle).Win),
		WinMeter:                (*battle).WinMeter,
		Opponent0SplatnetId:     (*battle).Opponent0SplatnetId,
		Opponent0Name:           (*battle).Opponent0Name,
		Opponent0Rank:           (*battle).Opponent0Rank,
		Opponent0LevelStar:      (*battle).Opponent0LevelStar,
		Opponent0Level:          (*battle).Opponent0Level,
		Opponent0Weapon:         (*battle).Opponent0Weapon,
		Opponent0Gender:         (*battle).Opponent0Gender,
		Opponent0Species:        (*battle).Opponent0Species,
		Opponent0Assists:        (*battle).Opponent0Assists,
		Opponent0Deaths:         (*battle).Opponent0Deaths,
		Opponent0GamePaintPoint: (*battle).Opponent0GamePaintPoint,
		Opponent0Kills:          (*battle).Opponent0Kills,
		Opponent0Specials:       (*battle).Opponent0Specials,
		Opponent0Headgear:       (*battle).Opponent0Headgear,
		Opponent0HeadgearMain:   (*battle).Opponent0HeadgearMain,
		Opponent0HeadgearSub0:   (*battle).Opponent0HeadgearSub0,
		Opponent0HeadgearSub1:   (*battle).Opponent0HeadgearSub1,
		Opponent0HeadgearSub2:   (*battle).Opponent0HeadgearSub2,
		Opponent0Clothes:        (*battle).Opponent0Clothes,
		Opponent0ClothesMain:    (*battle).Opponent0ClothesMain,
		Opponent0ClothesSub0:    (*battle).Opponent0ClothesSub0,
		Opponent0ClothesSub1:    (*battle).Opponent0ClothesSub1,
		Opponent0ClothesSub2:    (*battle).Opponent0ClothesSub2,
		Opponent0Shoes:          (*battle).Opponent0Shoes,
		Opponent0ShoesMain:      (*battle).Opponent0ShoesMain,
		Opponent0ShoesSub0:      (*battle).Opponent0ShoesSub0,
		Opponent0ShoesSub1:      (*battle).Opponent0ShoesSub1,
		Opponent0ShoesSub2:      (*battle).Opponent0ShoesSub2,
		Opponent1SplatnetId:     (*battle).Opponent1SplatnetId,
		Opponent1Name:           (*battle).Opponent1Name,
		Opponent1Rank:           (*battle).Opponent1Rank,
		Opponent1LevelStar:      (*battle).Opponent1LevelStar,
		Opponent1Level:          (*battle).Opponent1Level,
		Opponent1Weapon:         (*battle).Opponent1Weapon,
		Opponent1Gender:         (*battle).Opponent1Gender,
		Opponent1Species:        (*battle).Opponent1Species,
		Opponent1Assists:        (*battle).Opponent1Assists,
		Opponent1Deaths:         (*battle).Opponent1Deaths,
		Opponent1GamePaintPoint: (*battle).Opponent1GamePaintPoint,
		Opponent1Kills:          (*battle).Opponent1Kills,
		Opponent1Specials:       (*battle).Opponent1Specials,
		Opponent1Headgear:       (*battle).Opponent1Headgear,
		Opponent1HeadgearMain:   (*battle).Opponent1HeadgearMain,
		Opponent1HeadgearSub0:   (*battle).Opponent1HeadgearSub0,
		Opponent1HeadgearSub1:   (*battle).Opponent1HeadgearSub1,
		Opponent1HeadgearSub2:   (*battle).Opponent1HeadgearSub2,
		Opponent1Clothes:        (*battle).Opponent1Clothes,
		Opponent1ClothesMain:    (*battle).Opponent1ClothesMain,
		Opponent1ClothesSub0:    (*battle).Opponent1ClothesSub0,
		Opponent1ClothesSub1:    (*battle).Opponent1ClothesSub1,
		Opponent1ClothesSub2:    (*battle).Opponent1ClothesSub2,
		Opponent1Shoes:          (*battle).Opponent1Shoes,
		Opponent1ShoesMain:      (*battle).Opponent1ShoesMain,
		Opponent1ShoesSub0:      (*battle).Opponent1ShoesSub0,
		Opponent1ShoesSub1:      (*battle).Opponent1ShoesSub1,
		Opponent1ShoesSub2:      (*battle).Opponent1ShoesSub2,
		Opponent2SplatnetId:     (*battle).Opponent2SplatnetId,
		Opponent2Name:           (*battle).Opponent2Name,
		Opponent2Rank:           (*battle).Opponent2Rank,
		Opponent2LevelStar:      (*battle).Opponent2LevelStar,
		Opponent2Level:          (*battle).Opponent2Level,
		Opponent2Weapon:         (*battle).Opponent2Weapon,
		Opponent2Gender:         (*battle).Opponent2Gender,
		Opponent2Species:        (*battle).Opponent2Species,
		Opponent2Assists:        (*battle).Opponent2Assists,
		Opponent2Deaths:         (*battle).Opponent2Deaths,
		Opponent2GamePaintPoint: (*battle).Opponent2GamePaintPoint,
		Opponent2Kills:          (*battle).Opponent2Kills,
		Opponent2Specials:       (*battle).Opponent2Specials,
		Opponent2Headgear:       (*battle).Opponent2Headgear,
		Opponent2HeadgearMain:   (*battle).Opponent2HeadgearMain,
		Opponent2HeadgearSub0:   (*battle).Opponent2HeadgearSub0,
		Opponent2HeadgearSub1:   (*battle).Opponent2HeadgearSub1,
		Opponent2HeadgearSub2:   (*battle).Opponent2HeadgearSub2,
		Opponent2Clothes:        (*battle).Opponent2Clothes,
		Opponent2ClothesMain:    (*battle).Opponent2ClothesMain,
		Opponent2ClothesSub0:    (*battle).Opponent2ClothesSub0,
		Opponent2ClothesSub1:    (*battle).Opponent2ClothesSub1,
		Opponent2ClothesSub2:    (*battle).Opponent2ClothesSub2,
		Opponent2Shoes:          (*battle).Opponent2Shoes,
		Opponent2ShoesMain:      (*battle).Opponent2ShoesMain,
		Opponent2ShoesSub0:      (*battle).Opponent2ShoesSub0,
		Opponent2ShoesSub1:      (*battle).Opponent2ShoesSub1,
		Opponent2ShoesSub2:      (*battle).Opponent2ShoesSub2,
		Opponent3SplatnetId:     (*battle).Opponent3SplatnetId,
		Opponent3Name:           (*battle).Opponent3Name,
		Opponent3Rank:           (*battle).Opponent3Rank,
		Opponent3LevelStar:      (*battle).Opponent3LevelStar,
		Opponent3Level:          (*battle).Opponent3Level,
		Opponent3Weapon:         (*battle).Opponent3Weapon,
		Opponent3Gender:         (*battle).Opponent3Gender,
		Opponent3Species:        (*battle).Opponent3Species,
		Opponent3Assists:        (*battle).Opponent3Assists,
		Opponent3Deaths:         (*battle).Opponent3Deaths,
		Opponent3GamePaintPoint: (*battle).Opponent3GamePaintPoint,
		Opponent3Kills:          (*battle).Opponent3Kills,
		Opponent3Specials:       (*battle).Opponent3Specials,
		Opponent3Headgear:       (*battle).Opponent3Headgear,
		Opponent3HeadgearMain:   (*battle).Opponent3HeadgearMain,
		Opponent3HeadgearSub0:   (*battle).Opponent3HeadgearSub0,
		Opponent3HeadgearSub1:   (*battle).Opponent3HeadgearSub1,
		Opponent3HeadgearSub2:   (*battle).Opponent3HeadgearSub2,
		Opponent3Clothes:        (*battle).Opponent3Clothes,
		Opponent3ClothesMain:    (*battle).Opponent3ClothesMain,
		Opponent3ClothesSub0:    (*battle).Opponent3ClothesSub0,
		Opponent3ClothesSub1:    (*battle).Opponent3ClothesSub1,
		Opponent3ClothesSub2:    (*battle).Opponent3ClothesSub2,
		Opponent3Shoes:          (*battle).Opponent3Shoes,
		Opponent3ShoesMain:      (*battle).Opponent3ShoesMain,
		Opponent3ShoesSub0:      (*battle).Opponent3ShoesSub0,
		Opponent3ShoesSub1:      (*battle).Opponent3ShoesSub1,
		Opponent3ShoesSub2:      (*battle).Opponent3ShoesSub2,
		Teammate0SplatnetId:     (*battle).Teammate0SplatnetId,
		Teammate0Name:           (*battle).Teammate0Name,
		Teammate0Rank:           (*battle).Teammate0Rank,
		Teammate0LevelStar:      (*battle).Teammate0LevelStar,
		Teammate0Level:          (*battle).Teammate0Level,
		Teammate0Weapon:         (*battle).Teammate0Weapon,
		Teammate0Gender:         (*battle).Teammate0Gender,
		Teammate0Species:        (*battle).Teammate0Species,
		Teammate0Assists:        (*battle).Teammate0Assists,
		Teammate0Deaths:         (*battle).Teammate0Deaths,
		Teammate0GamePaintPoint: (*battle).Teammate0GamePaintPoint,
		Teammate0Kills:          (*battle).Teammate0Kills,
		Teammate0Specials:       (*battle).Teammate0Specials,
		Teammate0Headgear:       (*battle).Teammate0Headgear,
		Teammate0HeadgearMain:   (*battle).Teammate0HeadgearMain,
		Teammate0HeadgearSub0:   (*battle).Teammate0HeadgearSub0,
		Teammate0HeadgearSub1:   (*battle).Teammate0HeadgearSub1,
		Teammate0HeadgearSub2:   (*battle).Teammate0HeadgearSub2,
		Teammate0Clothes:        (*battle).Teammate0Clothes,
		Teammate0ClothesMain:    (*battle).Teammate0ClothesMain,
		Teammate0ClothesSub0:    (*battle).Teammate0ClothesSub0,
		Teammate0ClothesSub1:    (*battle).Teammate0ClothesSub1,
		Teammate0ClothesSub2:    (*battle).Teammate0ClothesSub2,
		Teammate0Shoes:          (*battle).Teammate0Shoes,
		Teammate0ShoesMain:      (*battle).Teammate0ShoesMain,
		Teammate0ShoesSub0:      (*battle).Teammate0ShoesSub0,
		Teammate0ShoesSub1:      (*battle).Teammate0ShoesSub1,
		Teammate0ShoesSub2:      (*battle).Teammate0ShoesSub2,
		Teammate1SplatnetId:     (*battle).Teammate1SplatnetId,
		Teammate1Name:           (*battle).Teammate1Name,
		Teammate1Rank:           (*battle).Teammate1Rank,
		Teammate1LevelStar:      (*battle).Teammate1LevelStar,
		Teammate1Level:          (*battle).Teammate1Level,
		Teammate1Weapon:         (*battle).Teammate1Weapon,
		Teammate1Gender:         (*battle).Teammate1Gender,
		Teammate1Species:        (*battle).Teammate1Species,
		Teammate1Assists:        (*battle).Teammate1Assists,
		Teammate1Deaths:         (*battle).Teammate1Deaths,
		Teammate1GamePaintPoint: (*battle).Teammate1GamePaintPoint,
		Teammate1Kills:          (*battle).Teammate1Kills,
		Teammate1Specials:       (*battle).Teammate1Specials,
		Teammate1Headgear:       (*battle).Teammate1Headgear,
		Teammate1HeadgearMain:   (*battle).Teammate1HeadgearMain,
		Teammate1HeadgearSub0:   (*battle).Teammate1HeadgearSub0,
		Teammate1HeadgearSub1:   (*battle).Teammate1HeadgearSub1,
		Teammate1HeadgearSub2:   (*battle).Teammate1HeadgearSub2,
		Teammate1Clothes:        (*battle).Teammate1Clothes,
		Teammate1ClothesMain:    (*battle).Teammate1ClothesMain,
		Teammate1ClothesSub0:    (*battle).Teammate1ClothesSub0,
		Teammate1ClothesSub1:    (*battle).Teammate1ClothesSub1,
		Teammate1ClothesSub2:    (*battle).Teammate1ClothesSub2,
		Teammate1Shoes:          (*battle).Teammate1Shoes,
		Teammate1ShoesMain:      (*battle).Teammate1ShoesMain,
		Teammate1ShoesSub0:      (*battle).Teammate1ShoesSub0,
		Teammate1ShoesSub1:      (*battle).Teammate1ShoesSub1,
		Teammate1ShoesSub2:      (*battle).Teammate1ShoesSub2,
		Teammate2SplatnetId:     (*battle).Teammate2SplatnetId,
		Teammate2Name:           (*battle).Teammate2Name,
		Teammate2Rank:           (*battle).Teammate2Rank,
		Teammate2LevelStar:      (*battle).Teammate2LevelStar,
		Teammate2Level:          (*battle).Teammate2Level,
		Teammate2Weapon:         (*battle).Teammate2Weapon,
		Teammate2Gender:         (*battle).Teammate2Gender,
		Teammate2Species:        (*battle).Teammate2Species,
		Teammate2Assists:        (*battle).Teammate2Assists,
		Teammate2Deaths:         (*battle).Teammate2Deaths,
		Teammate2GamePaintPoint: (*battle).Teammate2GamePaintPoint,
		Teammate2Kills:          (*battle).Teammate2Kills,
		Teammate2Specials:       (*battle).Teammate2Specials,
		Teammate2Headgear:       (*battle).Teammate2Headgear,
		Teammate2HeadgearMain:   (*battle).Teammate2HeadgearMain,
		Teammate2HeadgearSub0:   (*battle).Teammate2HeadgearSub0,
		Teammate2HeadgearSub1:   (*battle).Teammate2HeadgearSub1,
		Teammate2HeadgearSub2:   (*battle).Teammate2HeadgearSub2,
		Teammate2Clothes:        (*battle).Teammate2Clothes,
		Teammate2ClothesMain:    (*battle).Teammate2ClothesMain,
		Teammate2ClothesSub0:    (*battle).Teammate2ClothesSub0,
		Teammate2ClothesSub1:    (*battle).Teammate2ClothesSub1,
		Teammate2ClothesSub2:    (*battle).Teammate2ClothesSub2,
		Teammate2Shoes:          (*battle).Teammate2Shoes,
		Teammate2ShoesMain:      (*battle).Teammate2ShoesMain,
		Teammate2ShoesSub0:      (*battle).Teammate2ShoesSub0,
		Teammate2ShoesSub1:      (*battle).Teammate2ShoesSub1,
		Teammate2ShoesSub2:      (*battle).Teammate2ShoesSub2,
		PlayerName:              (*battle).PlayerName,
		PlayerRank:              (*battle).PlayerRank,
		PlayerLevelStar:         (*battle).PlayerLevelStar,
		PlayerLevel:             (*battle).PlayerLevel,
		PlayerWeapon:            (*battle).PlayerWeapon,
		PlayerGender:            (*battle).PlayerGender,
		PlayerSpecies:           (*battle).PlayerSpecies,
		PlayerAssists:           (*battle).PlayerAssists,
		PlayerDeaths:            (*battle).PlayerDeaths,
		PlayerGamePaintPoint:    (*battle).PlayerGamePaintPoint,
		PlayerKills:             (*battle).PlayerKills,
		PlayerSpecials:          (*battle).PlayerSpecials,
		PlayerHeadgear:          (*battle).PlayerHeadgear,
		PlayerHeadgearMain:      (*battle).PlayerHeadgearMain,
		PlayerHeadgearSub0:      (*battle).PlayerHeadgearSub0,
		PlayerHeadgearSub1:      (*battle).PlayerHeadgearSub1,
		PlayerHeadgearSub2:      (*battle).PlayerHeadgearSub2,
		PlayerClothes:           (*battle).PlayerClothes,
		PlayerClothesMain:       (*battle).PlayerClothesMain,
		PlayerClothesSub0:       (*battle).PlayerClothesSub0,
		PlayerClothesSub1:       (*battle).PlayerClothesSub1,
		PlayerClothesSub2:       (*battle).PlayerClothesSub2,
		PlayerShoes:             (*battle).PlayerShoes,
		PlayerShoesMain:         (*battle).PlayerShoesMain,
		PlayerShoesSub0:         (*battle).PlayerShoesSub0,
		PlayerShoesSub1:         (*battle).PlayerShoesSub1,
		PlayerShoesSub2:         (*battle).PlayerShoesSub2,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTable("two_battles_battle", []string{
		"user_id",
		"splatnet_json", "splatnet_upload", "stat_ink_json", "stat_ink_upload", "splatnet_number", "player_splatnet_id",
		"elapsed_time", "has_disconnected_player", "league_point", "match_type", "rule", "my_team_count",
		"other_team_count", "splatfest_point", "splatfest_title", "stage", "tag_id", "time", "win", "win_meter",
		"opponent0_splatnet_id", "opponent0_name", "opponent0_rank", "opponent0_level_star", "opponent0_level",
		"opponent0_weapon", "opponent0_gender", "opponent0_species", "opponent0_assists", "opponent0_deaths",
		"opponent0_game_paint_point", "opponent0_kills", "opponent0_specials", "opponent0_headgear",
		"opponent0_headgear_main", "opponent0_headgear_sub0", "opponent0_headgear_sub1", "opponent0_headgear_sub2",
		"opponent0_clothes", "opponent0_clothes_main", "opponent0_clothes_sub0", "opponent0_clothes_sub1",
		"opponent0_clothes_sub2", "opponent0_shoes", "opponent0_shoes_main", "opponent0_shoes_sub0",
		"opponent0_shoes_sub1", "opponent0_shoes_sub2",
		"opponent1_splatnet_id", "opponent1_name", "opponent1_rank", "opponent1_level_star", "opponent1_level",
		"opponent1_weapon", "opponent1_gender", "opponent1_species", "opponent1_assists", "opponent1_deaths",
		"opponent1_game_paint_point", "opponent1_kills", "opponent1_specials", "opponent1_headgear",
		"opponent1_headgear_main", "opponent1_headgear_sub0", "opponent1_headgear_sub1", "opponent1_headgear_sub2",
		"opponent1_clothes", "opponent1_clothes_main", "opponent1_clothes_sub0", "opponent1_clothes_sub1",
		"opponent1_clothes_sub2", "opponent1_shoes", "opponent1_shoes_main", "opponent1_shoes_sub0",
		"opponent1_shoes_sub1", "opponent1_shoes_sub2",
		"opponent2_splatnet_id", "opponent2_name", "opponent2_rank", "opponent2_level_star", "opponent2_level",
		"opponent2_weapon", "opponent2_gender", "opponent2_species", "opponent2_assists", "opponent2_deaths",
		"opponent2_game_paint_point", "opponent2_kills", "opponent2_specials", "opponent2_headgear",
		"opponent2_headgear_main", "opponent2_headgear_sub0", "opponent2_headgear_sub1", "opponent2_headgear_sub2",
		"opponent2_clothes", "opponent2_clothes_main", "opponent2_clothes_sub0", "opponent2_clothes_sub1",
		"opponent2_clothes_sub2", "opponent2_shoes", "opponent2_shoes_main", "opponent2_shoes_sub0",
		"opponent2_shoes_sub1", "opponent2_shoes_sub2",
		"opponent3_splatnet_id", "opponent3_name", "opponent3_rank", "opponent3_level_star", "opponent3_level",
		"opponent3_weapon", "opponent3_gender", "opponent3_species", "opponent3_assists", "opponent3_deaths",
		"opponent3_game_paint_point", "opponent3_kills", "opponent3_specials", "opponent3_headgear",
		"opponent3_headgear_main", "opponent3_headgear_sub0", "opponent3_headgear_sub1", "opponent3_headgear_sub2",
		"opponent3_clothes", "opponent3_clothes_main", "opponent3_clothes_sub0", "opponent3_clothes_sub1",
		"opponent3_clothes_sub2", "opponent3_shoes", "opponent3_shoes_main", "opponent3_shoes_sub0",
		"opponent3_shoes_sub1", "opponent3_shoes_sub2",
		"teammate0_splatnet_id", "teammate0_name", "teammate0_rank", "teammate0_level_star", "teammate0_level",
		"teammate0_weapon", "teammate0_gender", "teammate0_species", "teammate0_assists", "teammate0_deaths",
		"teammate0_game_paint_point", "teammate0_kills", "teammate0_specials", "teammate0_headgear",
		"teammate0_headgear_main", "teammate0_headgear_sub0", "teammate0_headgear_sub1", "teammate0_headgear_sub2",
		"teammate0_clothes", "teammate0_clothes_main", "teammate0_clothes_sub0", "teammate0_clothes_sub1",
		"teammate0_clothes_sub2", "teammate0_shoes", "teammate0_shoes_main", "teammate0_shoes_sub0",
		"teammate0_shoes_sub1", "teammate0_shoes_sub2",
		"teammate1_splatnet_id", "teammate1_name", "teammate1_rank", "teammate1_level_star", "teammate1_level",
		"teammate1_weapon", "teammate1_gender", "teammate1_species", "teammate1_assists", "teammate1_deaths",
		"teammate1_game_paint_point", "teammate1_kills", "teammate1_specials", "teammate1_headgear",
		"teammate1_headgear_main", "teammate1_headgear_sub0", "teammate1_headgear_sub1", "teammate1_headgear_sub2",
		"teammate1_clothes", "teammate1_clothes_main", "teammate1_clothes_sub0", "teammate1_clothes_sub1",
		"teammate1_clothes_sub2", "teammate1_shoes", "teammate1_shoes_main", "teammate1_shoes_sub0",
		"teammate1_shoes_sub1", "teammate1_shoes_sub2",
		"teammate2_splatnet_id", "teammate2_name", "teammate2_rank", "teammate2_level_star", "teammate2_level",
		"teammate2_weapon", "teammate2_gender", "teammate2_species", "teammate2_assists", "teammate2_deaths",
		"teammate2_game_paint_point", "teammate2_kills", "teammate2_specials", "teammate2_headgear",
		"teammate2_headgear_main", "teammate2_headgear_sub0", "teammate2_headgear_sub1", "teammate2_headgear_sub2",
		"teammate2_clothes", "teammate2_clothes_main", "teammate2_clothes_sub0", "teammate2_clothes_sub1",
		"teammate2_clothes_sub2", "teammate2_shoes", "teammate2_shoes_main", "teammate2_shoes_sub0",
		"teammate2_shoes_sub1", "teammate2_shoes_sub2",
		"player_name", "player_rank", "player_level_star", "player_level",
		"player_weapon", "player_gender", "player_species", "player_assists", "player_deaths",
		"player_game_paint_point", "player_kills", "player_specials", "player_headgear",
		"player_headgear_main", "player_headgear_sub0", "player_headgear_sub1", "player_headgear_sub2",
		"player_clothes", "player_clothes_main", "player_clothes_sub0", "player_clothes_sub1",
		"player_clothes_sub2", "player_shoes", "player_shoes_main", "player_shoes_sub0",
		"player_shoes_sub1", "player_shoes_sub2",
	}, values)
}

func writeNewBattleSplatnet(battle *api_objects.BattleSplatnet) (*int64, error) {
	var udemae *int64
	if (*battle).Udemae != nil {
		var err error
		udemae, err = writeNewBattleSplatnetUdemae((*battle).Udemae)
		if err != nil {
			log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
			return nil, err
		}
	}
	playerResult, err := writeNewBattleSplatnetPlayerResult(&(*battle).PlayerResult)
	if err != nil {
		log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
		return nil, err
	}
	stage, err := writeNewSplatnetTriple(&(*battle).Stage)
	if err != nil {
		log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
		return nil, err
	}
	rule, err := writeNewBattleSplatnetRule(&(*battle).Rule)
	if err != nil {
		log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
		return nil, err
	}
	gameMode, err := writeNewSplatnetDouble(&(*battle).GameMode)
	if err != nil {
		log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
		return nil, err
	}
	myTeamResult, err := writeNewSplatnetDouble(&(*battle).MyTeamResult)
	if err != nil {
		log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
		return nil, err
	}
	otherTeamResult, err := writeNewSplatnetDouble(&(*battle).OtherTeamResult)
	if err != nil {
		log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleSplatnet{
		Udemae:             udemae,
		Stage:              *stage,
		OtherTeamCount:     (*battle).OtherTeamCount,
		MyTeamCount:        (*battle).MyTeamCount,
		StarRank:           (*battle).StarRank,
		Rule:               *rule,
		PlayerResult:       *playerResult,
		EstimateGachiPower: (*battle).EstimateGachiPower,
		ElapsedTime:        (*battle).ElapsedTime,
		StartTime:          (*battle).StartTime,
		GameMode:           *gameMode,
		//XPower:              (*battle).XPower,
		BattleNumber: (*battle).BattleNumber,
		Type:         (*battle).Type,
		PlayerRank:   (*battle).PlayerRank,
		//CrownPlayers:        (*battle).CrownPlayers,
		WeaponPaintPoint: (*battle).WeaponPaintPoint,
		//Rank:                (*battle).Rank,
		MyTeamResult: *myTeamResult,
		//EstimateXPower:      (*battle).EstimateXPower,
		OtherTeamResult:     *otherTeamResult,
		LeaguePoint:         (*battle).LeaguePoint,
		WinMeter:            (*battle).WinMeter,
		MyTeamPercentage:    (*battle).MyTeamPercentage,
		OtherTeamPercentage: (*battle).OtherTeamPercentage,
		TagId:               (*battle).TagId,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	pk, err := writeIntoTableGetPk("two_battles_battle_splatnet", []string{
		"udemae", "stage", "other_team_count", "my_team_count", "star_rank", "rule", "player_result",
		"estimate_gachi_power", "elapsed_time", "start_time", "game_mode",
		// "x_power",
		"battle_number", "type", "player_rank",
		// "crown_players",
		"weapon_paint_point",
		// "rank",
		"my_team_result",
		// "estimate_x_power",
		"other_team_result", "league_point", "win_meter", "my_team_percentage",
		"other_team_percentage", "tag_id",
	}, values)
	if err != nil {
		log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
		return nil, err
	}
	for i := range (*battle).MyTeamMembers {
		if _, err := writeNewBattleSplatnetTeamMember(*pk, true, &(*battle).MyTeamMembers[i]); err != nil {
			log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
			return nil, err
		}
	}
	for i := range (*battle).OtherTeamMembers {
		if _, err := writeNewBattleSplatnetTeamMember(*pk, false, &(*battle).OtherTeamMembers[i]); err != nil {
			log.Printf("Called by writeNewBattleSplatnet(%+v)\n", *battle)
			return nil, err
		}
	}
	return pk, nil
}

func writeNewBattleSplatnetTeamMember(parent int64, myTeam bool, playerResult *api_objects.BattleSplatnetPlayerResult) (*int64, error) {
	playerResultKey, err := writeNewBattleSplatnetPlayerResult(playerResult)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleSplatnetTeamMember{
		MyTeam:       myTeam,
		Parent:       parent,
		PlayerResult: *playerResultKey,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_splatnet_team_member", []string{
		"my_team", "parent", "player_result",
	}, values)
}

func writeNewBattleSplatnetUdemae(udemae *api_objects.BattleSplatnetUdemae) (*int64, error) {
	v := reflect.ValueOf(*udemae)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_splatnet_udemae", []string{
		"name", "is_x", "is_number_reached", "number", "s_plus_number",
	}, values)
}

func writeNewBattleSplatnetPlayerResult(playerResult *api_objects.BattleSplatnetPlayerResult) (*int64, error) {
	player, err := writeNewBattleSplatnetPlayerResultPlayer(&(*playerResult).Player)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleSplatnetPlayerResult{
		DeathCount:     (*playerResult).DeathCount,
		GamePaintPoint: (*playerResult).GamePaintPoint,
		KillCount:      (*playerResult).KillCount,
		SpecialCount:   (*playerResult).SpecialCount,
		AssistCount:    (*playerResult).AssistCount,
		SortScore:      (*playerResult).SortScore,
		Player:         *player,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_splatnet_player_result", []string{
		"death_count", "game_paint_point", "kill_count", "special_count", "assist_count", "sort_score", "player",
	}, values)
}

func writeNewBattleSplatnetPlayerResultPlayer(player *api_objects.BattleSplatnetPlayerResultPlayer) (*int64, error) {
	var udemae *int64
	var err error
	if (*player).Udemae != nil {
		udemae, err = writeNewBattleSplatnetUdemae((*player).Udemae)
		if err != nil {
			return nil, err
		}
	}
	headSkills, err := writeNewBattleSplatnetPlayerResultPlayerSkills(&(*player).HeadSkills)
	if err != nil {
		return nil, err
	}
	shoesSkills, err := writeNewBattleSplatnetPlayerResultPlayerSkills(&(*player).ShoesSkills)
	if err != nil {
		return nil, err
	}
	clothesSkills, err := writeNewBattleSplatnetPlayerResultPlayerSkills(&(*player).ClothesSkills)
	if err != nil {
		return nil, err
	}
	playerType, err := writeNewSplatnetPlayerType(&(*player).PlayerType)
	if err != nil {
		return nil, err
	}
	head, err := writeNewBattleSplatnetPlayerResultPlayerClothing(&(*player).Head)
	if err != nil {
		return nil, err
	}
	clothes, err := writeNewBattleSplatnetPlayerResultPlayerClothing(&(*player).Clothes)
	if err != nil {
		return nil, err
	}
	shoes, err := writeNewBattleSplatnetPlayerResultPlayerClothing(&(*player).Shoes)
	if err != nil {
		return nil, err
	}
	weapon, err := writeNewBattleSplatnetPlayerResultPlayerWeapon(&(*player).Weapon)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleSplatnetPlayerResultPlayer{
		HeadSkills:    *headSkills,
		ShoesSkills:   *shoesSkills,
		ClothesSkills: *clothesSkills,
		PlayerRank:    (*player).PlayerRank,
		StarRank:      (*player).StarRank,
		Nickname:      (*player).Nickname,
		PlayerType:    *playerType,
		PrincipalId:   (*player).PrincipalId,
		Head:          *head,
		Clothes:       *clothes,
		Shoes:         *shoes,
		Udemae:        udemae,
		Weapon:        *weapon,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_splatnet_player_result_player", []string{
		"head_skills", "shoes_skills", "clothes_skills", "player_rank", "star_rank", "nickname", "player_type",
		"principal_id", "head", "clothes", "shoes", "udemae", "weapon",
	}, values)
}

func writeNewBattleSplatnetPlayerResultPlayerSkills(skills *api_objects.BattleSplatnetPlayerResultPlayerSkills) (*int64, error) {
	main, err := writeNewSplatnetTriple(&(*skills).Main)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleSplatnetPlayerResultPlayerSkills{
		Main: *main,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	pk, err := writeIntoTableGetPk("two_battles_battle_splatnet_player_result_player_skills", []string{
		"main",
	}, values)
	if err != nil {
		return nil, err
	}
	for i := range (*skills).Subs {
		if _, err := writeNewBattleSplatnetResultPlayerSkillsSub(*pk, &(*skills).Subs[i]); err != nil {
			return nil, err
		}
	}
	return pk, nil
}

func writeNewBattleSplatnetResultPlayerSkillsSub(pk int64, sub *api_objects.BattleSplatnetTriple) (*int64, error) {
	subKey, err := writeNewSplatnetTriple(sub)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleSplatnetPlayerResultPlayerSkillsSubContainer{
		Parent: pk,
		Sub:    *subKey,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_splatnet_player_result_player_skills_sub", []string{
		"parent", "sub",
	}, values)
}

func writeNewSplatnetTriple(triple *api_objects.BattleSplatnetTriple) (*int64, error) {
	v := reflect.ValueOf(*triple)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("splatnet_triple", []string{
		"id", "image", "name",
	}, values)
}

func writeNewSplatnetPlayerType(playerType *api_objects.SplatnetPlayerType) (*int64, error) {
	v := reflect.ValueOf(*playerType)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("splatnet_player_type", []string{
		"gender", "species",
	}, values)
}

func writeNewBattleSplatnetPlayerResultPlayerClothing(clothing *api_objects.BattleSplatnetPlayerResultPlayerClothing) (*int64, error) {
	brand, err := writeNewBattleSplatnetPlayerResultPlayerClothingBrand(&(*clothing).Brand)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleSplatnetPlayerResultPlayerClothing{
		Id:        (*clothing).Id,
		Image:     (*clothing).Image,
		Name:      (*clothing).Name,
		Thumbnail: (*clothing).Thumbnail,
		Kind:      (*clothing).Kind,
		Rarity:    (*clothing).Rarity,
		Brand:     *brand,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_splatnet_player_result_player_clothing", []string{
		"id", "image", "name", "thumbnail", "kind", "rarity", "brand",
	}, values)
}

func writeNewBattleSplatnetPlayerResultPlayerClothingBrand(brand *api_objects.BattleSplatnetPlayerResultPlayerClothingBrand) (*int64, error) {
	frequentSkill, err := writeNewSplatnetTriple(&(*brand).FrequentSkill)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleSplatnetPlayerResultPlayerClothingBrand{
		Id:            (*brand).Id,
		Image:         (*brand).Image,
		Name:          (*brand).Name,
		FrequentSkill: *frequentSkill,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_splatnet_player_result_player_clothing_brand", []string{
		"id", "image", "name", "frequent_skill",
	}, values)
}

func writeNewBattleSplatnetPlayerResultPlayerWeapon(weapon *api_objects.BattleSplatnetPlayerResultPlayerWeapon) (*int64, error) {
	sub, err := writeNewSplatnetQuad(&(*weapon).Sub)
	if err != nil {
		return nil, err
	}
	special, err := writeNewSplatnetQuad(&(*weapon).Special)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleSplatnetPlayerResultPlayerWeapon{
		Id:        (*weapon).Id,
		Image:     (*weapon).Image,
		Name:      (*weapon).Name,
		Thumbnail: (*weapon).Thumbnail,
		Sub:       *sub,
		Special:   *special,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_splatnet_player_result_player_weapon", []string{
		"id", "image", "name", "thumbnail", "sub", "special",
	}, values)
}

func writeNewSplatnetQuad(quad *api_objects.SplatnetQuad) (*int64, error) {
	v := reflect.ValueOf(*quad)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("splatnet_quad", []string{
		"id", "image_a", "image_b", "name",
	}, values)
}

func writeNewBattleSplatnetRule(rule *api_objects.BattleSplatnetRule) (*int64, error) {
	v := reflect.ValueOf(*rule)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_splatnet_rule", []string{
		"key", "name", "multiline_name",
	}, values)
}

func writeNewSplatnetDouble(double *api_objects.SplatnetDouble) (*int64, error) {
	v := reflect.ValueOf(*double)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("splatnet_double", []string{
		"key", "name",
	}, values)
}

func writeNewBattleStatInk(battle *api_objects.BattleStatInk) (*int64, error) {
	var err error
	var freshness *int64
	if (*battle).Freshness != nil {
		freshness, err = writeNewBattleStatInkFreshness((*battle).Freshness)
		if err != nil {
			return nil, err
		}
	}
	var rank *int64
	if (*battle).Rank != nil {
		rank, err = writeNewBattleStatInkRank((*battle).Rank)
		if err != nil {
			return nil, err
		}
	}
	var rankAfter *int64
	if (*battle).Rank != nil {
		rankAfter, err = writeNewBattleStatInkRank((*battle).RankAfter)
		if err != nil {
			return nil, err
		}
	}
	var festTitle *int64
	if (*battle).FestTitle != nil {
		festTitle, err = writeNewStatInkKeyName((*battle).FestTitle)
		if err != nil {
			return nil, err
		}
	}
	var festTitleAfter *int64
	if (*battle).FestTitleAfter != nil {
		festTitleAfter, err = writeNewStatInkKeyName((*battle).FestTitleAfter)
		if err != nil {
			return nil, err
		}
	}
	var specialBattle *int64
	if (*battle).SpecialBattle != nil {
		specialBattle, err = writeNewStatInkKeyName((*battle).SpecialBattle)
		if err != nil {
			return nil, err
		}
	}
	lobby, err := writeNewStatInkKeyName(&(*battle).Lobby)
	if err != nil {
		return nil, err
	}
	mode, err := writeNewStatInkKeyName(&(*battle).Mode)
	if err != nil {
		return nil, err
	}
	rule, err := writeNewStatInkKeyName(&(*battle).Rule)
	if err != nil {
		return nil, err
	}
	species, err := writeNewStatInkKeyName(&(*battle).Species)
	if err != nil {
		return nil, err
	}
	weapon, err := writeNewBattleStatInkWeapon(&(*battle).Weapon)
	if err != nil {
		return nil, err
	}
	user, err := writeNewBattleStatInkUser(&(*battle).User)
	if err != nil {
		return nil, err
	}
	gender, err := writeNewStatInkGender(&(*battle).Gender)
	if err != nil {
		return nil, err
	}
	startAt, err := writeNewStatInkTime(&(*battle).StartAt)
	if err != nil {
		return nil, err
	}
	endAt, err := writeNewStatInkTime(&(*battle).EndAt)
	if err != nil {
		return nil, err
	}
	registerAt, err := writeNewStatInkTime(&(*battle).RegisterAt)
	if err != nil {
		return nil, err
	}
	stage, err := writeNewBattleStatInkMap(&(*battle).Map)
	if err != nil {
		return nil, err
	}
	gears, err := writeNewBattleStatInkGears(&(*battle).Gears)
	if err != nil {
		return nil, err
	}
	agent, err := writeNewBattleStatInkAgent(&(*battle).Agent)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInk{
		Id:             (*battle).Id,
		SplatnetNumber: (*battle).SplatnetNumber,
		Url:            (*battle).Url,
		User:           *user,
		Lobby:          *lobby,
		Mode:           *mode,
		Rule:           *rule,
		Map:            *stage,
		Weapon:         *weapon,
		Freshness:      freshness,
		Rank:           rank,
		RankExp:        (*battle).RankExp,
		RankAfter:      rankAfter,
		//XPower:                         (*battle).XPower,
		//XPowerAfter:                    (*battle).XPowerAfter,
		//EstimateXPower:                 (*battle).EstimateXPower,
		Level:        (*battle).Level,
		LevelAfter:   (*battle).LevelAfter,
		StarRank:     (*battle).StarRank,
		Result:       (*battle).Result,
		KnockOut:     (*battle).KnockOut,
		RankInTeam:   (*battle).RankInTeam,
		Kill:         (*battle).Kill,
		Death:        (*battle).Death,
		KillOrAssist: (*battle).KillOrAssist,
		Special:      (*battle).Special,
		KillRatio:    (*battle).KillRatio,
		KillRate:     (*battle).KillRate,
		//MaxKillCombo:                   (*battle).MaxKillCombo,
		//MaxKillStreak:                  (*battle).MaxKillStreak,
		//DeathReasons:                   (*battle).DeathReasons,
		MyPoint:                    (*battle).MyPoint,
		EstimateGachiPower:         (*battle).EstimateGachiPower,
		LeaguePoint:                (*battle).LeaguePoint,
		MyTeamEstimateLeaguePoint:  (*battle).MyTeamEstimateLeaguePoint,
		HisTeamEstimateLeaguePoint: (*battle).HisTeamEstimateLeaguePoint,
		//MyTeamPoint:                    (*battle).MyTeamPoint,
		//HisTeamPoint:                   (*battle).HisTeamPoint,
		MyTeamPercent:                  (*battle).MyTeamPercent,
		HisTeamPercent:                 (*battle).HisTeamPercent,
		MyTeamId:                       (*battle).HisTeamId,
		Species:                        *species,
		Gender:                         *gender,
		FestTitle:                      festTitle,
		FestExp:                        (*battle).FestExp,
		FestTitleAfter:                 festTitleAfter,
		FestExpAfter:                   (*battle).FestExpAfter,
		FestPower:                      (*battle).FestPower,
		MyTeamEstimateFestPower:        (*battle).MyTeamEstimateFestPower,
		HisTeamMyTeamEstimateFestPower: (*battle).HisTeamMyTeamEstimateFestPower,
		MyTeamFestTheme:                (*battle).MyTeamFestTheme,
		MyTeamNickname:                 (*battle).MyTeamNickname,
		HisTeamNickname:                (*battle).HisTeamNickname,
		Clout:                          (*battle).Clout,
		TotalClout:                     (*battle).TotalClout,
		TotalCloutAfter:                (*battle).TotalCloutAfter,
		MyTeamWinStreak:                (*battle).MyTeamWinStreak,
		HisTeamWinStreak:               (*battle).HisTeamWinStreak,
		SynergyBonus:                   (*battle).SynergyBonus,
		SpecialBattle:                  specialBattle,
		ImageResult:                    (*battle).ImageResult,
		ImageGear:                      (*battle).ImageGear,
		Gears:                          *gears,
		Period:                         (*battle).Period,
		PeriodRange:                    (*battle).PeriodRange,
		//Events:                         (*battle).Events,
		//SplatnetJson:                   (*battle).SplatnetJson,
		Agent:     *agent,
		Automated: (*battle).Automated,
		//Environment:                    (*battle).Environment,
		LinkUrl: (*battle).LinkUrl,
		//Note:                           (*battle).Note,
		GameVersion:   (*battle).GameVersion,
		NawabariBonus: (*battle).NawabariBonus,
		StartAt:       *startAt,
		EndAt:         *endAt,
		RegisterAt:    *registerAt,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	pk, err := writeIntoTableGetPk("two_battles_battle_stat_ink", []string{
		"id", "splatnet_number", "url", "user", "lobby", "mode", "rule", "map", "weapon", "freshness", "rank",
		"rank_exp", "rank_after",
		//"x_power", "x_power_after", "estimate_x_power",
		"level", "level_after", "star_rank",
		"result", "knock_out", "rank_in_team", "kill", "death", "kill_or_assist", "special", "kill_ratio", "kill_rate",
		//"max_kill_combo", "max_kill_streak", "death_reasons",
		"my_point", "estimate_gachi_power", "league_point",
		"my_team_estimate_league_point", "his_team_estimate_league_point",
		//"my_team_point", "his_team_point",
		"my_team_percent", "his_team_percent", "my_team_id", "his_team_id", "species", "gender", "fest_title",
		"fest_exp", "fest_title_after", "fest_exp_after", "fest_power", "my_team_estimate_fest_power",
		"his_team_my_team_estimate_fest_power", "my_team_fest_theme", "my_team_nickname", "his_team_nickname", "clout",
		"total_clout", "total_clout_after", "my_team_win_streak", "his_team_win_streak", "synergy_bonus",
		"special_battle", "image_result", "image_gear", "gears", "period", "period_range",
		//"events", "splatnet_json",
		"agent", "automated",
		//"environment",
		"link_url",
		//"note",
		"game_version", "nawabari_bonus", "start_at", "end_at",
		"register_at",
	}, values)
	if err != nil {
		return nil, err
	}
	for i := range (*battle).Players {
		if _, err := writeNewBattleStatInkPlayer(*pk, &(*battle).Players[i]); err != nil {
			return nil, err
		}
	}
	return pk, err
}

func writeNewBattleStatInkPlayer(parent int64, player *api_objects.BattleStatInkPlayer) (*int64, error) {
	var rank *int64
	var err error
	if (*player).Rank != nil {
		rank, err = writeNewBattleStatInkRank((*player).Rank)
		if err != nil {
			return nil, err
		}
	}
	var festTitle *int64
	if (*player).FestTitle != nil {
		festTitle, err = writeNewStatInkKeyName((*player).FestTitle)
		if err != nil {
			return nil, err
		}
	}
	var species *int64
	if (*player).Species != nil {
		species, err = writeNewStatInkKeyName((*player).Species)
	}
	if err != nil {
		return nil, err
	}
	weapon, err := writeNewBattleStatInkWeapon(&(*player).Weapon)
	if err != nil {
		return nil, err
	}
	gender, err := writeNewStatInkGender(&(*player).Gender)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkPlayer{
		Parent:       parent,
		Team:         (*player).Team,
		IsMe:         (*player).IsMe,
		Weapon:       *weapon,
		Level:        (*player).Level,
		Rank:         rank,
		StarRank:     (*player).StarRank,
		RankInTeam:   (*player).RankInTeam,
		Kill:         (*player).Kill,
		Death:        (*player).Death,
		KillOrAssist: (*player).KillOrAssist,
		Special:      (*player).Special,
		// MyKill:       (*player).MyKill,
		Point:      (*player).Point,
		Name:       (*player).Name,
		Species:    species,
		Gender:     *gender,
		FestTitle:  festTitle,
		SplatnetId: (*player).SplatnetId,
		Top500:     (*player).Top500,
		Icon:       (*player).Icon,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_player", []string{
		"parent", "team", "is_me", "weapon", "level", "rank", "star_rank", "rank_in_team", "kill", "death",
		"kill_or_assist", "special",
		//"my_kill",
		"point", "name", "species", "gender", "fest_title", "splatnet_id",
		"top_500", "icon",
	}, values)
}

func writeNewBattleStatInkRank(rank *api_objects.BattleStatInkRank) (*int64, error) {
	name, err := writeNewStatInkName(&(*rank).Name)
	if err != nil {
		return nil, err
	}
	zone, err := writeNewStatInkKeyName(&(*rank).Zone)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkRank{
		Key:  (*rank).Key,
		Name: *name,
		Zone: *zone,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_rank", []string{
		"key", "name", "zone",
	}, values)
}

func writeNewStatInkName(name *api_objects.StatInkName) (*int64, error) {
	v := reflect.ValueOf(*name)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("stat_ink_name", []string{
		"de_DE", "en_GB", "en_US", "es_ES", "es_MX", "fr_CA", "fr_FR", "it_IT", "ja_JP", "nl_NL", "ru_RU", "zh_CN",
		"zh_TW",
	}, values)
}

func writeNewShiftStatInkStageName(name *api_objects.ShiftStatInkStageName) (*int64, error) {
	v := reflect.ValueOf(*name)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("stat_ink_name", []string{
		"de_DE", "en_GB", "en_US", "es_ES", "es_MX", "fr_CA", "fr_FR", "it_IT", "ja_JP", "nl_NL", "ru_RU", "zh_CN",
		"zh_TW",
	}, values)
}

func writeNewStatInkKeyName(keyName *api_objects.StatInkKeyName) (*int64, error) {
	name, err := writeNewStatInkName(&(*keyName).Name)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.StatInkKeyName{
		Key:  (*keyName).Key,
		Name: *name,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("stat_ink_key_name", []string{
		"key", "name",
	}, values)
}

func writeNewShiftStatInkFailReason(reason *api_objects.ShiftStatInkFailReason) (*int64, error) {
	name, err := writeNewStatInkName(&(*reason).Name)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkFailReason{
		Key:  (*reason).Key,
		Name: *name,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("stat_ink_key_name", []string{
		"key", "name",
	}, values)
}

func writeNewBattleStatInkWeapon(weapon *api_objects.BattleStatInkWeapon) (*int64, error) {
	name, err := writeNewStatInkName(&(*weapon).Name)
	if err != nil {
		return nil, err
	}
	sub, err := writeNewStatInkKeyName(&(*weapon).Sub)
	if err != nil {
		return nil, err
	}
	special, err := writeNewStatInkKeyName(&(*weapon).Special)
	if err != nil {
		return nil, err
	}
	mainPowerUp, err := writeNewStatInkKeyName(&(*weapon).MainPowerUp)
	if err != nil {
		return nil, err
	}
	typeKey, err := writeNewBattleStatInkWeaponType(&(*weapon).Type)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkWeapon{
		Key:         (*weapon).Key,
		Name:        *name,
		Splatnet:    (*weapon).Splatnet,
		Type:        *typeKey,
		ReskinOf:    (*weapon).ReskinOf,
		MainRef:     (*weapon).MainRef,
		Sub:         *sub,
		Special:     *special,
		MainPowerUp: *mainPowerUp,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_weapon", []string{
		"key", "name", "splatnet", "type", "reskin_of", "main_ref", "sub", "special", "main_power_up",
	}, values)
}

func writeNewBattleStatInkWeaponType(weaponType *api_objects.BattleStatInkWeaponType) (*int64, error) {
	name, err := writeNewStatInkName(&(*weaponType).Name)
	if err != nil {
		return nil, err
	}
	category, err := writeNewStatInkKeyName(&(*weaponType).Category)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkWeaponType{
		Key:      (*weaponType).Key,
		Name:     *name,
		Category: *category,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_weapon_type", []string{
		"key", "name", "category",
	}, values)
}

func writeNewStatInkGender(gender *api_objects.StatInkGender) (*int64, error) {
	name, err := writeNewStatInkName(&(*gender).Name)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.StatInkGender{
		Key:     (*gender).Key,
		Name:    *name,
		Iso5218: (*gender).Iso5218,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("stat_ink_gender", []string{
		"key", "name", "iso5218",
	}, values)
}

func writeNewBattleStatInkUser(user *api_objects.BattleStatInkUser) (*int64, error) {
	joinAt, err := writeNewStatInkTime(&(*user).JoinAt)
	if err != nil {
		return nil, err
	}
	profile, err := writeNewStatInkProfile(&(*user).Profile)
	if err != nil {
		return nil, err
	}
	stats, err := writeNewBattleStatInkUserStats(&(*user).Stats)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkUser{
		Id:         (*user).Id,
		Name:       (*user).Name,
		ScreenName: (*user).ScreenName,
		Url:        (*user).Url,
		JoinAt:     *joinAt,
		Profile:    *profile,
		//Stat:       (*user).Stat,
		Stats: *stats,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_user", []string{
		"id", "name", "screen_name", "url", "join_at", "profile",
		//"stat",
		"stats",
	}, values)
}

func writeNewStatInkTime(timeObj *api_objects.StatInkTime) (*int64, error) {
	v := reflect.ValueOf(*timeObj)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("stat_ink_time", []string{
		"time", "iso8601",
	}, values)
}

func writeNewStatInkProfile(profile *api_objects.StatInkProfile) (*int64, error) {
	v := reflect.ValueOf(*profile)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("stat_ink_profile", []string{
		// "nnid", "friend_code", "twitter", "ikanakama", "ikanakama2", "environment",
		"friend_code", "twitter",
	}, values)
}

func writeNewBattleStatInkUserStats(stats *api_objects.BattleStatInkUserStats) (*int64, error) {
	var v2 *int64
	var err error
	if (*stats).V2 != nil {
		v2, err = writeNewBattleStatInkUserStatsV2((*stats).V2)
		if err != nil {
			return nil, err
		}
	}
	v := reflect.ValueOf(db_objects.BattleStatInkUserStats{
		// V1: (*stats).V1,
		V2: v2,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_user_stats", []string{
		//"v1",
		"v2",
	}, values)
}

func writeNewBattleStatInkUserStatsV2(v2 *api_objects.BattleStatInkUserStatsV2) (*int64, error) {
	updatedAt, err := writeNewStatInkTime(&(*v2).UpdatedAt)
	if err != nil {
		return nil, err
	}
	entire, err := writeNewBattleStatInkUserStatsV2Entire(&(*v2).Entire)
	if err != nil {
		return nil, err
	}
	nawabari, err := writeNewBattleStatInkUserStatsV2Nawabari(&(*v2).Nawabari)
	if err != nil {
		return nil, err
	}
	gachi, err := writeNewBattleStatInkUserStatsV2Gachi(&(*v2).Gachi)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkUserStatsV2{
		UpdatedAt: *updatedAt,
		Entire:    *entire,
		Nawabari:  *nawabari,
		Gachi:     *gachi,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_user_stats_v2", []string{
		"updated_at", "entire", "nawabari", "gachi",
	}, values)
}

func writeNewBattleStatInkUserStatsV2Entire(entire *api_objects.BattleStatInkUserStatsV2Entire) (*int64, error) {
	v := reflect.ValueOf(*entire)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_user_stats_v2_entire", []string{
		"battles", "win_pct", "kill_ratio", "kill_total", "kill_avg", "kill_per_min", "death_total", "death_avg",
		"death_per_min",
	}, values)
}

func writeNewBattleStatInkUserStatsV2Nawabari(nawabari *api_objects.BattleStatInkUserStatsV2Nawabari) (*int64, error) {
	v := reflect.ValueOf(*nawabari)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_user_stats_v2_nawabari", []string{
		"battles", "win_pct", "kill_ratio", "kill_total", "kill_avg", "kill_per_min", "death_total", "death_avg",
		"death_per_min", "total_inked", "max_inked", "avg_inked",
	}, values)
}

func writeNewBattleStatInkUserStatsV2Gachi(gachi *api_objects.BattleStatInkUserStatsV2Gachi) (*int64, error) {
	rules, err := writeNewBattleStatInkUserStatsV2GachiRules(&(*gachi).Rules)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkUserStatsV2Gachi{
		Battles:     (*gachi).Battles,
		WinPct:      (*gachi).WinPct,
		KillRatio:   (*gachi).KillRatio,
		KillTotal:   (*gachi).KillTotal,
		KillAvg:     (*gachi).KillAvg,
		KillPerMin:  (*gachi).KillPerMin,
		DeathTotal:  (*gachi).DeathTotal,
		DeathAvg:    (*gachi).DeathAvg,
		DeathPerMin: (*gachi).DeathPerMin,
		Rules:       *rules,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_user_stats_v2_gachi", []string{
		"battles", "win_pct", "kill_ratio", "kill_total", "kill_avg", "kill_per_min", "death_total", "death_avg",
		"death_per_min", "rules",
	}, values)
}

func writeNewBattleStatInkUserStatsV2GachiRules(rules *api_objects.BattleStatInkUserStatsV2GachiRules) (*int64, error) {
	area, err := writeNewBattleStatInkUserStatsV2GachiRulesSub(&(*rules).Area)
	if err != nil {
		return nil, err
	}
	yagura, err := writeNewBattleStatInkUserStatsV2GachiRulesSub(&(*rules).Yagura)
	if err != nil {
		return nil, err
	}
	hoko, err := writeNewBattleStatInkUserStatsV2GachiRulesSub(&(*rules).Hoko)
	if err != nil {
		return nil, err
	}
	asari, err := writeNewBattleStatInkUserStatsV2GachiRulesSub(&(*rules).Asari)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkUserStatsV2GachiRules{
		Area:   *area,
		Yagura: *yagura,
		Hoko:   *hoko,
		Asari:  *asari,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_user_stats_v2_gachi_rules", []string{
		"area", "yagura", "hoko", "asari",
	}, values)
}

func writeNewBattleStatInkUserStatsV2GachiRulesSub(sub *api_objects.BattleStatInkUserStatsV2GachiRulesSub) (*int64, error) {
	v := reflect.ValueOf(*sub)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_user_stats_v2_gachi_rules_sub", []string{
		"rank_peak", "rank_current", // "x_power_peak", "x_power_current",
	}, values)
}

func writeNewBattleStatInkMap(stage *api_objects.BattleStatInkMap) (*int64, error) {
	name, err := writeNewStatInkName(&(*stage).Name)
	if err != nil {
		return nil, err
	}
	releaseAt, err := writeNewStatInkTime(&(*stage).ReleaseAt)
	if err != nil {
		return nil, err
	}
	shortName, err := writeNewStatInkName(&(*stage).ShortName)
	if err != nil {
		return nil, err
	}
	dbStage := db_objects.BattleStatInkMap{
		Key:       (*stage).Key,
		Name:      *name,
		Splatnet:  (*stage).Splatnet,
		Area:      (*stage).Area,
		ReleaseAt: *releaseAt,
		ShortName: *shortName,
	}
	v := reflect.ValueOf(dbStage)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_map", []string{
		"key", "name", "splatnet", "area", "release_at", "short_name",
	}, values)
}

func writeNewBattleStatInkFreshness(freshness *api_objects.BattleStatInkFreshness) (*int64, error) {
	title, err := writeNewStatInkName(&(*freshness).Title)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkFreshness{
		Freshness: (*freshness).Freshness,
		Title:     *title,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_freshness", []string{
		"freshness", "title",
	}, values)
}

func writeNewBattleStatInkGears(gears *api_objects.BattleStatInkGears) (*int64, error) {
	headgear, err := writeNewBattleStatInkGearsClothes(&(*gears).Headgear)
	if err != nil {
		return nil, err
	}
	clothing, err := writeNewBattleStatInkGearsClothes(&(*gears).Clothing)
	if err != nil {
		return nil, err
	}
	shoes, err := writeNewBattleStatInkGearsClothes(&(*gears).Shoes)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkGears{
		Headgear: *headgear,
		Clothing: *clothing,
		Shoes:    *shoes,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_gears", []string{
		"headgear", "clothing", "shoes",
	}, values)
}

func writeNewBattleStatInkGearsClothes(clothes *api_objects.BattleStatInkGearsClothes) (*int64, error) {
	gear, err := writeNewBattleStatInkGearsClothesGear(&(*clothes).Gear)
	if err != nil {
		return nil, err
	}
	primaryAbility, err := writeNewStatInkKeyName(&(*clothes).PrimaryAbility)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkGearsClothes{
		Gear:           *gear,
		PrimaryAbility: *primaryAbility,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	pk, err := writeIntoTableGetPk("two_battles_battle_stat_ink_gears_clothes", []string{
		"gear", "primary_ability",
	}, values)
	if err != nil {
		return nil, err
	}
	for i := range (*clothes).SecondaryAbilities {
		if _, err := writeNewBattleStatInkGearsClothesSecondaryAbility(*pk, &(*clothes).SecondaryAbilities[i]); err != nil {
			return nil, err
		}
	}
	return pk, nil
}

func writeNewBattleStatInkGearsClothesGear(gear *api_objects.BattleStatInkGearsClothesGear) (*int64, error) {
	name, err := writeNewStatInkName(&(*gear).Name)
	if err != nil {
		return nil, err
	}
	typeKey, err := writeNewStatInkKeyName(&(*gear).Type)
	if err != nil {
		return nil, err
	}
	brand, err := writeNewStatInkKeyName(&(*gear).Brand)
	if err != nil {
		return nil, err
	}
	primaryAbility, err := writeNewStatInkKeyName(&(*gear).PrimaryAbility)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkGearsClothesGear{
		Key:            (*gear).Key,
		Name:           *name,
		Splatnet:       (*gear).Splatnet,
		Type:           *typeKey,
		Brand:          *brand,
		PrimaryAbility: *primaryAbility,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_gears_clothes_gear", []string{
		"key", "name", "splatnet", "type", "brand", "primary_ability",
	}, values)
}

func writeNewBattleStatInkGearsClothesSecondaryAbility(parent int64, secondaryAbility *api_objects.StatInkKeyName) (*int64, error) {
	secondaryAbilityKey, err := writeNewStatInkKeyName(secondaryAbility)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.BattleStatInkGearsClothesSecondaryAbilityContainer{
		Parent:           parent,
		SecondaryAbility: *secondaryAbilityKey,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_gears_clothes_sa_container", []string{
		"parent", "secondary_ability",
	}, values)
}

func writeNewBattleStatInkAgent(agent *api_objects.BattleStatInkAgent) (*int64, error) {
	var variables *int64
	if (*agent).Variables != nil {
		var err error
		variables, err = writeNewBattleStatInkVariables((*agent).Variables)
		if err != nil {
			return nil, err
		}
	}
	v := reflect.ValueOf(db_objects.BattleStatInkAgent{
		Name:    (*agent).Name,
		Version: (*agent).Version,
		// GameVersion:     (*agent).GameVersion,
		// GameVersionDate: (*agent).GameVersionDate,
		// Custom:          (*agent).Custom,
		Variables: variables,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_agent", []string{
		"name", "version",
		//"game_version", "game_version_data", "custom",
		"variables",
	}, values)
}

func writeNewBattleStatInkVariables(variables *api_objects.BattleStatInkAgentVariables) (*int64, error) {
	v := reflect.ValueOf(*variables)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_battles_battle_stat_ink_agent_variables", []string{
		"upload_mode",
	}, values)
}

func AddSessionToken(sessionToken, username string) error {
	user, err := ReadUser(username)
	if err != nil {
		return err
	}
	values := make([]interface{}, 2)
	values[0] = (*user).Pk
	values[1] = sessionToken
	fmt.Println(len(sessionToken))
	return writeIntoTable("auth_user_session_token", []string{
		"parent", "session_token",
	}, values)
}

func WriteNewShift(shift *api_objects.Shift) error {
	_, err := GetShift((*shift).UserId, (*shift).JobId)
	if fmt.Sprint(err) != "sql: no rows in result set" {
		return errors.New("shift already exists")
	}
	var splatnet *int64
	if (*shift).SplatnetJson != nil {
		splatnet, err = writeNewShiftSplatnet((*shift).SplatnetJson)
		if err != nil {
			return err
		}
	}
	var statInk *int64
	if (*shift).StatInkJson != nil {
		statInk, err = writeNewShiftStatInk((*shift).StatInkJson)
		if err != nil {
			return err
		}
	}
	v := reflect.ValueOf(db_objects.Shift{
		UserId:                  (*shift).UserId,
		PlayerSplatnetId:        (*shift).PlayerSplatnetId,
		JobId:                   (*shift).JobId,
		SplatnetJson:            splatnet,
		StatInkJson:             statInk,
		StartTime:               (*shift).StartTime,
		PlayTime:                (*shift).PlayTime,
		EndTime:                 (*shift).EndTime,
		DangerRate:              (*shift).DangerRate,
		IsClear:                 (*shift).IsClear,
		JobFailureReason:        (*shift).JobFailureReason,
		FailureWave:             (*shift).FailureWave,
		GradePoint:              (*shift).GradePoint,
		GradePointDelta:         (*shift).GradePointDelta,
		JobScore:                (*shift).JobScore,
		DrizzlerCount:           (*shift).DrizzlerCount,
		FlyfishCount:            (*shift).FlyfishCount,
		GoldieCount:             (*shift).GoldieCount,
		GrillerCount:            (*shift).GrillerCount,
		MawsCount:               (*shift).MawsCount,
		ScrapperCount:           (*shift).ScrapperCount,
		SteelEelCount:           (*shift).SteelEelCount,
		SteelheadCount:          (*shift).SteelheadCount,
		StingerCount:            (*shift).StingerCount,
		Stage:                   (*shift).Stage,
		PlayerName:              (*shift).PlayerName,
		PlayerDeathCount:        (*shift).PlayerDeathCount,
		PlayerReviveCount:       (*shift).PlayerReviveCount,
		PlayerGoldenEggs:        (*shift).PlayerGoldenEggs,
		PlayerPowerEggs:         (*shift).PlayerPowerEggs,
		PlayerSpecial:           (*shift).PlayerSpecial,
		PlayerTitle:             (*shift).PlayerTitle,
		PlayerSpecies:           (*shift).PlayerSpecies,
		PlayerGender:            (*shift).PlayerGender,
		PlayerW1Specials:        (*shift).PlayerW1Specials,
		PlayerW2Specials:        (*shift).PlayerW2Specials,
		PlayerW3Specials:        (*shift).PlayerW3Specials,
		PlayerW1Weapon:          (*shift).PlayerW1Weapon,
		PlayerW2Weapon:          (*shift).PlayerW2Weapon,
		PlayerW3Weapon:          (*shift).PlayerW3Weapon,
		PlayerDrizzlerKills:     (*shift).PlayerDrizzlerKills,
		PlayerFlyfishKills:      (*shift).PlayerFlyfishKills,
		PlayerGoldieKills:       (*shift).PlayerGoldieKills,
		PlayerGrillerKills:      (*shift).PlayerGrillerKills,
		PlayerMawsKills:         (*shift).PlayerMawsKills,
		PlayerScrapperKills:     (*shift).PlayerScrapperKills,
		PlayerSteelEelKills:     (*shift).PlayerSteelEelKills,
		PlayerSteelheadKills:    (*shift).PlayerSteelheadKills,
		PlayerStingerKills:      (*shift).PlayerStingerKills,
		Teammate0SplatnetId:     (*shift).Teammate0SplatnetId,
		Teammate0Name:           (*shift).Teammate0Name,
		Teammate0DeathCount:     (*shift).Teammate0DeathCount,
		Teammate0ReviveCount:    (*shift).Teammate0ReviveCount,
		Teammate0GoldenEggs:     (*shift).Teammate0GoldenEggs,
		Teammate0PowerEggs:      (*shift).Teammate0PowerEggs,
		Teammate0Special:        (*shift).Teammate0Special,
		Teammate0Species:        (*shift).Teammate0Species,
		Teammate0Gender:         (*shift).Teammate0Gender,
		Teammate0W1Specials:     (*shift).Teammate0W1Specials,
		Teammate0W2Specials:     (*shift).Teammate0W2Specials,
		Teammate0W3Specials:     (*shift).Teammate0W3Specials,
		Teammate0W1Weapon:       (*shift).Teammate0W1Weapon,
		Teammate0W2Weapon:       (*shift).Teammate0W2Weapon,
		Teammate0W3Weapon:       (*shift).Teammate0W3Weapon,
		Teammate0DrizzlerKills:  (*shift).Teammate0DrizzlerKills,
		Teammate0FlyfishKills:   (*shift).Teammate0FlyfishKills,
		Teammate0GoldieKills:    (*shift).Teammate0GoldieKills,
		Teammate0GrillerKills:   (*shift).Teammate0GrillerKills,
		Teammate0MawsKills:      (*shift).Teammate0MawsKills,
		Teammate0ScrapperKills:  (*shift).Teammate0ScrapperKills,
		Teammate0SteelEelKills:  (*shift).Teammate0SteelEelKills,
		Teammate0SteelheadKills: (*shift).Teammate0SteelheadKills,
		Teammate0StingerKills:   (*shift).Teammate0StingerKills,
		Teammate1SplatnetId:     (*shift).Teammate1SplatnetId,
		Teammate1Name:           (*shift).Teammate1Name,
		Teammate1DeathCount:     (*shift).Teammate1DeathCount,
		Teammate1ReviveCount:    (*shift).Teammate1ReviveCount,
		Teammate1GoldenEggs:     (*shift).Teammate1GoldenEggs,
		Teammate1PowerEggs:      (*shift).Teammate1PowerEggs,
		Teammate1Special:        (*shift).Teammate1Special,
		Teammate1Species:        (*shift).Teammate1Species,
		Teammate1Gender:         (*shift).Teammate1Gender,
		Teammate1W1Specials:     (*shift).Teammate1W1Specials,
		Teammate1W2Specials:     (*shift).Teammate1W2Specials,
		Teammate1W3Specials:     (*shift).Teammate1W3Specials,
		Teammate1W1Weapon:       (*shift).Teammate1W1Weapon,
		Teammate1W2Weapon:       (*shift).Teammate1W2Weapon,
		Teammate1W3Weapon:       (*shift).Teammate1W3Weapon,
		Teammate1DrizzlerKills:  (*shift).Teammate1DrizzlerKills,
		Teammate1FlyfishKills:   (*shift).Teammate1FlyfishKills,
		Teammate1GoldieKills:    (*shift).Teammate1GoldieKills,
		Teammate1GrillerKills:   (*shift).Teammate1GrillerKills,
		Teammate1MawsKills:      (*shift).Teammate1MawsKills,
		Teammate1ScrapperKills:  (*shift).Teammate1ScrapperKills,
		Teammate1SteelEelKills:  (*shift).Teammate1SteelEelKills,
		Teammate1SteelheadKills: (*shift).Teammate1SteelheadKills,
		Teammate1StingerKills:   (*shift).Teammate1StingerKills,
		Teammate2SplatnetId:     (*shift).Teammate2SplatnetId,
		Teammate2Name:           (*shift).Teammate2Name,
		Teammate2DeathCount:     (*shift).Teammate2DeathCount,
		Teammate2ReviveCount:    (*shift).Teammate2ReviveCount,
		Teammate2GoldenEggs:     (*shift).Teammate2GoldenEggs,
		Teammate2PowerEggs:      (*shift).Teammate2PowerEggs,
		Teammate2Special:        (*shift).Teammate2Special,
		Teammate2Species:        (*shift).Teammate2Species,
		Teammate2Gender:         (*shift).Teammate2Gender,
		Teammate2W1Specials:     (*shift).Teammate2W1Specials,
		Teammate2W2Specials:     (*shift).Teammate2W2Specials,
		Teammate2W3Specials:     (*shift).Teammate2W3Specials,
		Teammate2W1Weapon:       (*shift).Teammate2W1Weapon,
		Teammate2W2Weapon:       (*shift).Teammate2W2Weapon,
		Teammate2W3Weapon:       (*shift).Teammate2W3Weapon,
		Teammate2DrizzlerKills:  (*shift).Teammate2DrizzlerKills,
		Teammate2FlyfishKills:   (*shift).Teammate2FlyfishKills,
		Teammate2GoldieKills:    (*shift).Teammate2GoldieKills,
		Teammate2GrillerKills:   (*shift).Teammate2GrillerKills,
		Teammate2MawsKills:      (*shift).Teammate2MawsKills,
		Teammate2ScrapperKills:  (*shift).Teammate2ScrapperKills,
		Teammate2SteelEelKills:  (*shift).Teammate2SteelEelKills,
		Teammate2SteelheadKills: (*shift).Teammate2SteelheadKills,
		Teammate2StingerKills:   (*shift).Teammate2StingerKills,
		ScheduleEndTime:         (*shift).ScheduleEndTime,
		ScheduleStartTime:       &(*shift).ScheduleStartTime,
		ScheduleWeapon0:         (*shift).ScheduleWeapon0,
		ScheduleWeapon1:         (*shift).ScheduleWeapon1,
		ScheduleWeapon2:         (*shift).ScheduleWeapon2,
		ScheduleWeapon3:         (*shift).ScheduleWeapon3,
		Wave1WaterLevel:         (*shift).Wave1WaterLevel,
		Wave1EventType:          (*shift).Wave1EventType,
		Wave1GoldenIkuraNum:     (*shift).Wave1GoldenDelivered,
		Wave1GoldenIkuraPopNum:  (*shift).Wave1GoldenAppear,
		Wave1IkuraNum:           (*shift).Wave1PowerEggs,
		Wave1QuotaNum:           (*shift).Wave1Quota,
		Wave2WaterLevel:         (*shift).Wave2WaterLevel,
		Wave2EventType:          (*shift).Wave2EventType,
		Wave2GoldenIkuraNum:     (*shift).Wave2GoldenDelivered,
		Wave2GoldenIkuraPopNum:  (*shift).Wave2GoldenAppear,
		Wave2IkuraNum:           (*shift).Wave2PowerEggs,
		Wave2QuotaNum:           (*shift).Wave2Quota,
		Wave3WaterLevel:         (*shift).Wave3WaterLevel,
		Wave3EventType:          (*shift).Wave3EventType,
		Wave3GoldenIkuraNum:     (*shift).Wave3GoldenDelivered,
		Wave3GoldenIkuraPopNum:  (*shift).Wave3GoldenAppear,
		Wave3IkuraNum:           (*shift).Wave3PowerEggs,
		Wave3QuotaNum:           (*shift).Wave3Quota,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTable("two_salmon_shift", []string{
		"user_id", "player_splatnet_id", "splatnet_number", "splatnet_json",
		"stat_ink_json", "start_time", "play_time", "end_time", "danger_rate", "is_clear", "job_failure_reason",
		"failure_wave", "grade_point", "grade_point_delta", "job_score", "drizzler_count", "flyfish_count",
		"goldie_count", "griller_count", "maws_count", "scrapper_count", "steel_eel_count", "steelhead_count",
		"stinger_count", "stage",
		"player_name", "player_death_count", "player_revive_count", "player_golden_eggs", "player_power_eggs",
		"player_special", "player_title", "player_species", "player_gender", "player_w1_specials", "player_w2_specials",
		"player_w3_specials", "player_w1_weapon", "player_w2_weapon", "player_w3_weapon", "player_drizzler_kills",
		"player_flyfish_kills", "player_goldie_kills", "player_griller_kills", "player_maws_kills",
		"player_scrapper_kills", "player_steel_eel_kills", "player_steelhead_kills", "player_stinger_kills",
		"teammate0_splatnet_id", "teammate0_name", "teammate0_death_count", "teammate0_revive_count",
		"teammate0_golden_eggs", "teammate0_power_eggs", "teammate0_special", "teammate0_species", "teammate0_gender",
		"teammate0_w1_specials", "teammate0_w2_specials", "teammate0_w3_specials", "teammate0_w1_weapon",
		"teammate0_w2_weapon", "teammate0_w3_weapon", "teammate0_drizzler_kills", "teammate0_flyfish_kills",
		"teammate0_goldie_kills", "teammate0_griller_kills", "teammate0_maws_kills", "teammate0_scrapper_kills",
		"teammate0_steel_eel_kills", "teammate0_steelhead_kills", "teammate0_stinger_kills",
		"teammate1_splatnet_id", "teammate1_name", "teammate1_death_count", "teammate1_revive_count",
		"teammate1_golden_eggs", "teammate1_power_eggs", "teammate1_special", "teammate1_species", "teammate1_gender",
		"teammate1_w1_specials", "teammate1_w2_specials", "teammate1_w3_specials", "teammate1_w1_weapon",
		"teammate1_w2_weapon", "teammate1_w3_weapon", "teammate1_drizzler_kills", "teammate1_flyfish_kills",
		"teammate1_goldie_kills", "teammate1_griller_kills", "teammate1_maws_kills", "teammate1_scrapper_kills",
		"teammate1_steel_eel_kills", "teammate1_steelhead_kills", "teammate1_stinger_kills",
		"teammate2_splatnet_id", "teammate2_name", "teammate2_death_count", "teammate2_revive_count",
		"teammate2_golden_eggs", "teammate2_power_eggs", "teammate2_special", "teammate2_species", "teammate2_gender",
		"teammate2_w1_specials", "teammate2_w2_specials", "teammate2_w3_specials", "teammate2_w1_weapon",
		"teammate2_w2_weapon", "teammate2_w3_weapon", "teammate2_drizzler_kills", "teammate2_flyfish_kills",
		"teammate2_goldie_kills", "teammate2_griller_kills", "teammate2_maws_kills", "teammate2_scrapper_kills",
		"teammate2_steel_eel_kills", "teammate2_steelhead_kills", "teammate2_stinger_kills",
		"schedule_end_time", "schedule_start_time", "schedule_weapon0", "schedule_weapon1", "schedule_weapon2",
		"schedule_weapon3",
		"wave1_water_level", "wave1_event_type", "wave1_golden_ikura_num", "wave1_golden_ikura_pop_num",
		"wave1_ikura_num", "wave1_quota_num",
		"wave2_water_level", "wave2_event_type", "wave2_golden_ikura_num", "wave2_golden_ikura_pop_num",
		"wave2_ikura_num", "wave2_quota_num",
		"wave3_water_level", "wave3_event_type", "wave3_golden_ikura_num", "wave3_golden_ikura_pop_num",
		"wave3_ikura_num", "wave3_quota_num",
	}, values)
}

func writeNewShiftSplatnet(shift *api_objects.ShiftSplatnet) (*int64, error) {
	jobResult, err := writeNewShiftsplatnetJobResult(&(*shift).JobResult)
	if err != nil {
		return nil, err
	}
	playerType, err := writeNewSplatnetPlayerType(&(*shift).PlayerType)
	if err != nil {
		return nil, err
	}
	bossCounts, err := writeNewShiftSplatnetBossCounts(&(*shift).BossCounts)
	if err != nil {
		return nil, err
	}
	myResult, err := writeNewShiftSplatnetPlayer(&(*shift).MyResult)
	if err != nil {
		return nil, err
	}
	grade, err := writeNewShiftSplatnetGrade(&(*shift).Grade)
	if err != nil {
		return nil, err
	}
	schedule, err := writeNewShiftSplatnetSchedule(&(*shift).Schedule)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftSplatnet{
		JobId:           (*shift).JobId,
		DangerRate:      (*shift).DangerRate,
		JobResult:       *jobResult,
		JobScore:        (*shift).JobScore,
		JobRate:         (*shift).JobRate,
		GradePoint:      (*shift).GradePoint,
		GradePointDelta: (*shift).GradePointDelta,
		KumaPoint:       (*shift).KumaPoint,
		StartTime:       (*shift).StartTime,
		PlayerType:      *playerType,
		PlayTime:        (*shift).PlayTime,
		BossCounts:      *bossCounts,
		EndTime:         (*shift).EndTime,
		MyResult:        *myResult,
		Grade:           *grade,
		Schedule:        *schedule,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	pk, err := writeIntoTableGetPk("two_salmon_shift_splatnet", []string{
		"job_id",
		"danger_rate", "job_result", "job_score", "job_rate", "grade_point", "grade_point_delta", "kuma_point",
		"start_time", "player_type", "play_time", "boss_counts", "end_time", "my_result", "grade", "schedule",
	}, values)
	if err != nil {
		return nil, err
	}
	for i := range (*shift).OtherResults {
		if _, err := writeNewShiftSplatnetPlayerContainer(*pk, &(*shift).OtherResults[i]); err != nil {
			return nil, err
		}
	}
	for i := range (*shift).WaveDetails {
		if _, err := writeNewShiftSplatnetWave(*pk, &(*shift).WaveDetails[i]); err != nil {
			return nil, err
		}
	}
	return pk, nil
}

func writeNewShiftStatInk(shift *api_objects.ShiftStatInk) (*int64, error) {
	user, err := writeNewShiftStatInkUser(&(*shift).User)
	if err != nil {
		return nil, err
	}
	stage, err := writeNewShiftStatInkStage(&(*shift).Stage)
	if err != nil {
		return nil, err
	}
	var failReason *int64
	if (*shift).FailReason != nil {
		failReason, err = writeNewShiftStatInkFailReason((*shift).FailReason)
		if err != nil {
			return nil, err
		}
	}
	title, err := writeNewShiftStatInkTitle(&(*shift).Title)
	if err != nil {
		return nil, err
	}
	titleAfter, err := writeNewShiftStatInkTitle(&(*shift).TitleAfter)
	if err != nil {
		return nil, err
	}
	myData, err := writeNewShiftStatInkPlayer(&(*shift).MyData)
	if err != nil {
		return nil, err
	}
	agent, err := writeNewShiftStatInkAgent(&(*shift).Agent)
	if err != nil {
		return nil, err
	}
	shiftStartAt, err := writeNewStatInkTime(&(*shift).ShiftStartAt)
	if err != nil {
		return nil, err
	}
	startAt, err := writeNewStatInkTime(&(*shift).StartAt)
	if err != nil {
		return nil, err
	}
	registerAt, err := writeNewStatInkTime(&(*shift).RegisterAt)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInk{
		Id:             (*shift).Id,
		Uuid:           (*shift).Uuid,
		SplatnetNumber: (*shift).SplatnetNumber,
		Url:            (*shift).Url,
		ApiEndpoint:    (*shift).ApiEndpoint,
		User:           *user,
		Stage:          *stage,
		IsCleared:      (*shift).IsCleared,
		FailReason:     failReason,
		ClearWaves:     (*shift).ClearWaves,
		DangerRate:     (*shift).DangerRate,
		Title:          *title,
		TitleExp:       (*shift).TitleExp,
		TitleAfter:     *titleAfter,
		TitleExpAfter:  (*shift).TitleExpAfter,
		MyData:         *myData,
		Agent:          *agent,
		Automated:      (*shift).Automated,
		//Note:           (*shift).Note,
		//LinkUrl:        (*shift).LinkUrl,
		ShiftStartAt: *shiftStartAt,
		StartAt:      *startAt,
		//EndAt:          (*shift).EndAt,
		RegisterAt: *registerAt,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	pk, err := writeIntoTableGetPk("two_salmon_shift_stat_ink", []string{
		"id", "uuid", "splatnet_number", "url", "api_endpoint", "user", "stage", "is_cleared", "fail_reason",
		"clear_waves", "danger_rate", "title", "title_exp", "title_after", "title_exp_after", "my_data", "agent",
		"automated",
		//"note", "link_url",
		"shift_start_at", "start_at",
		//"end_at",
		"register_at",
	}, values)
	if err != nil {
		return nil, err
	}
	for i := range (*shift).Quota {
		if _, err := writeNewIntContainer(*pk, "two_salmon_shift_stat_ink", (*shift).Quota[i]); err != nil {
			return nil, err
		}
	}
	for i := range (*shift).BossAppearances {
		if _, err := writeNewShiftStatInkBossData(*pk, "two_salmon_shift_stat_ink", &(*shift).BossAppearances[i]); err != nil {
			return nil, err
		}
	}
	for i := range (*shift).Waves {
		if _, err := writeNewShiftStatInkWave(*pk, &(*shift).Waves[i]); err != nil {
			return nil, err
		}
	}
	for i := range (*shift).Teammates {
		if _, err := writeNewShiftStatInkPlayerContainer(*pk, &(*shift).Teammates[i]); err != nil {
			return nil, err
		}
	}
	return pk, nil
}

func writeNewShiftsplatnetJobResult(result *api_objects.ShiftSplatnetJobResult) (*int64, error) {
	v := reflect.ValueOf(*result)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_job_result", []string{
		"is_clear", "failure_reason", "failure_wave",
	}, values)
}

func writeNewShiftSplatnetBossCounts(counts *api_objects.ShiftSplatnetBossCounts) (*int64, error) {
	var3, err := writeNewShiftSplatnetBossCountsBoss(&(*counts).Goldie)
	if err != nil {
		return nil, err
	}
	var6, err := writeNewShiftSplatnetBossCountsBoss(&(*counts).Steelhead)
	if err != nil {
		return nil, err
	}
	var9, err := writeNewShiftSplatnetBossCountsBoss(&(*counts).Flyfish)
	if err != nil {
		return nil, err
	}
	var12, err := writeNewShiftSplatnetBossCountsBoss(&(*counts).Scrapper)
	if err != nil {
		return nil, err
	}
	var13, err := writeNewShiftSplatnetBossCountsBoss(&(*counts).SteelEel)
	if err != nil {
		return nil, err
	}
	var14, err := writeNewShiftSplatnetBossCountsBoss(&(*counts).Stinger)
	if err != nil {
		return nil, err
	}
	var15, err := writeNewShiftSplatnetBossCountsBoss(&(*counts).Maws)
	if err != nil {
		return nil, err
	}
	var16, err := writeNewShiftSplatnetBossCountsBoss(&(*counts).Griller)
	if err != nil {
		return nil, err
	}
	var21, err := writeNewShiftSplatnetBossCountsBoss(&(*counts).Drizzler)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftSplatnetBossCounts{
		Var3:  *var3,
		Var6:  *var6,
		Var9:  *var9,
		Var12: *var12,
		Var13: *var13,
		Var14: *var14,
		Var15: *var15,
		Var16: *var16,
		Var21: *var21,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_boss_counts", []string{
		"var3", "var6", "var9", "var12", "var13", "var14", "var15", "var16", "var21",
	}, values)
}

func writeNewShiftSplatnetPlayer(player *api_objects.ShiftSplatnetPlayer) (*int64, error) {
	special, err := writeNewSplatnetQuad(&(*player).Special)
	if err != nil {
		return nil, err
	}
	playerType, err := writeNewSplatnetPlayerType(&(*player).PlayerType)
	if err != nil {
		return nil, err
	}
	bossKillCounts, err := writeNewShiftSplatnetBossCounts(&(*player).BossKillCounts)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftSplatnetPlayer{
		Special:        *special,
		Pid:            (*player).Pid,
		PlayerType:     *playerType,
		Name:           (*player).Name,
		DeadCount:      (*player).DeadCount,
		GoldenIkuraNum: (*player).GoldenEggs,
		BossKillCounts: *bossKillCounts,
		IkuraNum:       (*player).PowerEggs,
		HelpCount:      (*player).HelpCount,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	pk, err := writeIntoTableGetPk("two_salmon_shift_splatnet_player", []string{
		"special", "pid", "player_type", "name", "dead_count", "golden_ikura_num", "boss_kill_counts", "ikura_num",
		"help_count",
	}, values)
	if err != nil {
		return nil, err
	}
	for i := range (*player).SpecialCounts {
		if _, err := writeNewIntContainer(*pk, "two_salmon_shift_splatnet_boss_counts", (*player).SpecialCounts[i]); err != nil {
			return nil, err
		}
	}
	for i := range (*player).WeaponList {
		if _, err := writeNewShiftSplatnetPlayerWeaponList(*pk, &(*player).WeaponList[i]); err != nil {
			return nil, err
		}
	}
	return pk, nil
}

func writeNewShiftSplatnetGrade(grade *api_objects.ShiftSplatnetGrade) (*int64, error) {
	v := reflect.ValueOf(*grade)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_grade", []string{
		"id", "short_name", "long_name", "name",
	}, values)
}

func writeNewShiftSplatnetSchedule(schedule *api_objects.ShiftSplatnetSchedule) (*int64, error) {
	stage, err := writeNewShiftSplatnetScheduleStage(&(*schedule).Stage)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftSplatnetSchedule{
		StartTime: (*schedule).StartTime,
		EndTime:   (*schedule).EndTime,
		Stage:     *stage,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	pk, err := writeIntoTableGetPk("two_salmon_shift_splatnet_schedule", []string{
		"start_time", "end_time", "stage",
	}, values)
	if err != nil {
		return nil, err
	}
	for i := range (*schedule).Weapons {
		if _, err := writeNewShiftSplatnetScheduleWeapon(*pk, &(*schedule).Weapons[i]); err != nil {
			return nil, err
		}
	}
	return pk, nil
}

func writeNewShiftSplatnetPlayerContainer(parent int64, player *api_objects.ShiftSplatnetPlayer) (*int64, error) {
	playerKey, err := writeNewShiftSplatnetPlayer(player)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftSplatnetPlayerContainer{
		Parent: parent,
		Player: *playerKey,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_player_container", []string{
		"parent", "player",
	}, values)
}

func writeNewShiftSplatnetWave(parent int64, wave *api_objects.ShiftSplatnetWave) (*int64, error) {
	waterLevel, err := writeNewSplatnetDouble(&(*wave).WaterLevel)
	if err != nil {
		return nil, err
	}
	eventType, err := writeNewSplatnetDouble(&(*wave).EventType)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftSplatnetWave{
		Parent:            parent,
		WaterLevel:        *waterLevel,
		EventType:         *eventType,
		GoldenIkuraNum:    (*wave).GoldenEggs,
		GoldenIkuraPopNum: (*wave).GoldenAppear,
		IkuraNum:          (*wave).PowerEggs,
		QuotaNum:          (*wave).QuotaNum,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_wave", []string{
		"parent", "water_level", "event_type", "golden_ikura_num", "golden_ikura_pop_num", "ikura_num", "quota_num",
	}, values)
}

func writeNewShiftStatInkUser(user *api_objects.ShiftStatInkUser) (*int64, error) {
	joinAt, err := writeNewStatInkTime(&(*user).JoinAt)
	if err != nil {
		return nil, err
	}
	profile, err := writeNewStatInkProfile(&(*user).Profile)
	if err != nil {
		return nil, err
	}
	stats, err := writeNewShiftStatInkUserStats(&(*user).Stats)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkUser{
		Id:         (*user).Id,
		Name:       (*user).Name,
		ScreenName: (*user).ScreenName,
		Url:        (*user).Url,
		SalmonUrl:  (*user).SalmonUrl,
		BattleUrl:  (*user).BattleUrl,
		JoinAt:     *joinAt,
		Profile:    *profile,
		Stats:      *stats,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_user", []string{
		"id", "name", "screen_name", "url", "salmon_url", "battle_url", "join_at", "profile", "stats",
	}, values)
}

func writeNewShiftStatInkTriple(triple *api_objects.ShiftStatInkTripleInt) (*int64, error) {
	name, err := writeNewStatInkName(&(*triple).Name)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkTriple{
		Key:      (*triple).Key,
		Name:     *name,
		Splatnet: (*triple).Splatnet,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_triple", []string{
		"key", "name", "splatnet",
	}, values)
}

func writeNewShiftStatInkTitle(title *api_objects.ShiftStatInkTitle) (*int64, error) {
	genericName, err := writeNewStatInkName(&(*title).GenericName)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkTitle{
		Splatnet:    (*title).Splatnet,
		GenericName: *genericName,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_title", []string{
		"splatnet", "generic_name",
	}, values)
}

func writeNewShiftStatInkPlayer(player *api_objects.ShiftStatInkPlayer) (*int64, error) {
	special, err := writeNewShiftStatInkTriple(&(*player).Special)
	if err != nil {
		return nil, err
	}
	species, err := writeNewStatInkKeyName(&(*player).Species)
	if err != nil {
		return nil, err
	}
	gender, err := writeNewStatInkGender(&(*player).Gender)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkPlayer{
		SplatnetId:         (*player).SplatnetId,
		Name:               (*player).Name,
		Special:            *special,
		Rescue:             (*player).Rescue,
		Death:              (*player).Death,
		GoldenEggDelivered: (*player).GoldenEggDelivered,
		PowerEggCollected:  (*player).PowerEggCollected,
		Species:            *species,
		Gender:             *gender,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	pk, err := writeIntoTableGetPk("two_salmon_shift_stat_ink_player", []string{
		"splatnet_id", "name", "special", "rescue", "death", "golden_egg_delivered", "power_egg_collected", "species", "gender",
	}, values)
	if err != nil {
		return nil, err
	}
	for i := range (*player).SpecialUses {
		if _, err := writeNewIntContainer(*pk, "two_salmon_shift_stat_ink_player", (*player).SpecialUses[i]); err != nil {
			return nil, err
		}
	}
	for i := range (*player).Weapons {
		if _, err := writeNewShiftStatInkTripleContainer(*pk, &(*player).Weapons[i]); err != nil {
			return nil, err
		}
	}
	for i := range (*player).BossKills {
		if _, err := writeNewShiftStatInkBossData(*pk, "two_salmon_shift_stat_ink_player", &(*player).BossKills[i]); err != nil {
			return nil, err
		}
	}
	return pk, nil
}

func writeNewShiftStatInkAgent(agent *api_objects.ShiftStatInkAgent) (*int64, error) {
	v := reflect.ValueOf(*agent)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_agent", []string{
		"name", "version",
	}, values)
}

func writeNewIntContainer(parent int64, parentTable string, value int) (*int64, error) {
	v := reflect.ValueOf(db_objects.IntContainer{
		Parent:      parent,
		ParentTable: parentTable,
		Value:       value,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("int_container", []string{
		"parent", "parent_table", "value",
	}, values)
}

func writeNewShiftStatInkBossData(parent int64, parentTable string, bossData *api_objects.ShiftStatInkBossData) (*int64, error) {
	boss, err := writeNewShiftStatInkBossDataBoss(&(*bossData).Boss)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkBossData{
		Parent:      parent,
		ParentTable: parentTable,
		Boss:        *boss,
		Count:       (*bossData).Count,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_boss_data", []string{
		"parent", "parent_table", "boss", "count",
	}, values)
}

func writeNewShiftStatInkWave(parent int64, wave *api_objects.ShiftStatInkWave) (*int64, error) {
	var err error
	var knownOccurrence *int64
	if (*wave).KnownOccurrence != nil {
		knownOccurrence, err = writeNewShiftStatInkTripleString((*wave).KnownOccurrence)
		if err != nil {
			return nil, err
		}
	}
	var waterLevel *int64
	if (*wave).WaterLevel != nil {
		waterLevel, err = writeNewShiftStatInkTripleString((*wave).WaterLevel)
		if err != nil {
			return nil, err
		}
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkWave{
		Parent:               parent,
		KnownOccurrence:      knownOccurrence,
		WaterLevel:           waterLevel,
		GoldenEggQuota:       (*wave).GoldenEggQuota,
		GoldenEggAppearances: (*wave).GoldenEggAppearances,
		GoldenEggDelivered:   (*wave).GoldenEggDelivered,
		PowerEggCollected:    (*wave).PowerEggCollected,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_wave", []string{
		"parent", "known_occurrence", "water_level", "golden_egg_quota", "golden_egg_appearances",
		"golden_egg_delivered", "power_egg_collected",
	}, values)
}

func writeNewShiftStatInkPlayerContainer(parent int64, player *api_objects.ShiftStatInkPlayer) (*int64, error) {
	playerKey, err := writeNewShiftStatInkPlayer(player)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkPlayerContainer{
		Parent: parent,
		Player: *playerKey,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_player_container", []string{
		"parent", "player",
	}, values)
}

func writeNewShiftSplatnetBossCountsBoss(boss *api_objects.ShiftSplatnetBossCountsBoss) (*int64, error) {
	bossInside := shiftSplatnetBossCountsBossDoubleToSplatnetDouble((*boss).Boss)
	bossSub, err := writeNewSplatnetDouble(&bossInside)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftSplatnetBossCountsBoss{
		Boss:  *bossSub,
		Count: (*boss).Count,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_boss_counts_boss", []string{
		"boss", "count",
	}, values)
}

func shiftSplatnetBossCountsBossDoubleToSplatnetDouble(double api_objects.ShiftSplatnetBossCountsBossDouble) api_objects.SplatnetDouble {
	return api_objects.SplatnetDouble{
		Key:  fmt.Sprint(double.Key),
		Name: double.Name,
	}
}

func writeNewShiftSplatnetPlayerWeaponList(parent int64, weaponList *api_objects.ShiftSplatnetPlayerWeaponList) (*int64, error) {
	weapon, err := writeNewShiftSplatnetPlayerWeaponListWeapon(&(*weaponList).Weapon)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftSplatnetPlayerWeaponList{
		Parent: parent,
		Id:     (*weaponList).Id,
		Weapon: *weapon,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_player_weapon_list", []string{
		"parent", "id", "weapon",
	}, values)
}

func writeNewShiftSplatnetScheduleStage(stage *api_objects.ShiftSplatnetScheduleStage) (*int64, error) {
	v := reflect.ValueOf(*stage)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_schedule_stage", []string{
		"image", "name",
	}, values)
}

func writeNewShiftSplatnetScheduleWeapon(parent int64, weapon *api_objects.ShiftSplatnetScheduleWeapon) (*int64, error) {
	var weaponSub *int64
	if (*weapon).Weapon != nil {
		var err error
		if weaponSub, err = writeNewShiftSplatnetScheduleWeaponWeapon((*weapon).Weapon); err != nil {
			return nil, err
		}
	}
	var special *int64
	if (*weapon).CoopSpecialWeapon != nil {
		var err error
		if special, err = writeNewShiftSplatnetScheduleWeaponSpecialWeapon((*weapon).CoopSpecialWeapon); err != nil {
			return nil, err
		}
	}
	v := reflect.ValueOf(db_objects.ShiftSplatnetScheduleWeapon{
		Parent:            parent,
		Id:                (*weapon).Id,
		Weapon:            weaponSub,
		CoopSpecialWeapon: special,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_schedule_weapon", []string{
		"parent", "id", "weapon", "coop_special_weapon",
	}, values)
}

func writeNewShiftStatInkUserStats(stats *api_objects.ShiftStatInkUserStats) (*int64, error) {
	asOf, err := writeNewStatInkTime(&(*stats).AsOf)
	if err != nil {
		return nil, err
	}
	registeredAt, err := writeNewStatInkTime(&(*stats).RegisteredAt)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkUserStats{
		WorkCount:       (*stats).WorkCount,
		TotalGoldenEggs: (*stats).TotalGoldenEggs,
		TotalEggs:       (*stats).TotalEggs,
		TotalRescued:    (*stats).TotalRescued,
		TotalPoint:      (*stats).TotalPoint,
		AsOf:            *asOf,
		RegisteredAt:    *registeredAt,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_user_stats", []string{
		"work_count", "total_golden_eggs", "total_eggs", "total_rescued", "total_point", "as_of", "registered_at",
	}, values)
}

func writeNewShiftStatInkTripleContainer(parent int64, triple *api_objects.ShiftStatInkTripleInt) (*int64, error) {
	tripleKey, err := writeNewShiftStatInkTriple(triple)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkTripleContainer{
		Parent: parent,
		Triple: *tripleKey,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_triple_container", []string{
		"parent", "triple",
	}, values)
}

func writeNewShiftStatInkTripleString(triple *api_objects.ShiftStatInkTripleString) (*int64, error) {
	name, err := writeNewStatInkName(&(*triple).Name)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkTripleString{
		Key:      (*triple).Key,
		Name:     *name,
		Splatnet: (*triple).Splatnet,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_triple_string", []string{
		"key", "name", "splatnet",
	}, values)
}

func writeNewShiftStatInkStage(triple *api_objects.ShiftStatInkStage) (*int64, error) {
	name, err := writeNewShiftStatInkStageName(&(*triple).Name)
	if err != nil {
		return nil, err
	}
	v := reflect.ValueOf(db_objects.ShiftStatInkTripleString{
		Key:      (*triple).Key,
		Name:     *name,
		Splatnet: (*triple).Splatnet,
	})
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_triple_string", []string{
		"key", "name", "splatnet",
	}, values)
}

func writeNewShiftStatInkBossDataBoss(boss *api_objects.ShiftStatInkBossDataBoss) (*int64, error) {
	v := reflect.ValueOf(*boss)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_stat_ink_boss_data_boss", []string{
		"splatnet", "splatnet_str",
	}, values)
}

func writeNewShiftSplatnetPlayerWeaponListWeapon(weapon *api_objects.ShiftSplatnetPlayerWeaponListWeapon) (*int64, error) {
	v := reflect.ValueOf(*weapon)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_player_weapon_list_weapon", []string{
		"id", "image", "name", "thumbnail",
	}, values)
}

func writeNewShiftSplatnetScheduleWeaponWeapon(weapon *api_objects.ShiftSplatnetScheduleWeaponWeapon) (*int64, error) {
	v := reflect.ValueOf(*weapon)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_schedule_weapon_weapon", []string{
		"id", "image", "name", "thumbnail",
	}, values)
}

func writeNewShiftSplatnetScheduleWeaponSpecialWeapon(weapon *api_objects.ShiftSplatnetScheduleWeaponSpecialWeapon) (*int64, error) {
	v := reflect.ValueOf(*weapon)
	values := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}
	return writeIntoTableGetPk("two_salmon_shift_splatnet_schedule_weapon_special_weapon", []string{
		"image", "name",
	}, values)
}

func VerifyEmailUser(email string) error {
	userPk, err := ReadKeyArrayWithKey(email, "email", "pk", "auth_user", "asc", "pk")
	if err != nil {
		return err
	}
	return updateTable("auth_user", []string{
		"email_verified",
	}, []interface{}{1}, userPk[0])
}
