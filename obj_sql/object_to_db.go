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
	_, err := db.Exec("insert into auth_user (username, password, email, email_verified) values ($1, $2, $3, false)",
		user.Username, user.Password, user.Email)
	return err
}

func btoi(b bool) int8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func WriteNewBattle(battle *api_objects.Battle2) error {
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
	_, err = db.Exec("insert into auth_user_session_token (parent, session_token) values ($1, $2)",
		user.Pk, sessionToken)
	return err
}

func WriteNewShift3(shift *api_objects.Shift3, userId int64) error {
	if err := shift.DecodeId(); err != nil {
		return err
	}
	exists, err := checkIfShift3Exists(userId, shift.Data.CoopHistoryDetail.ID)
	if err != nil {
		return err
	}
	if *exists {
		return errors.New("shift already exists")
	}
	if err := shift.Data.CoopHistoryDetail.AfterGrade.DecodeId(); err != nil {
		return err
	}
	if err := shift.Data.CoopHistoryDetail.MyResult.Player.Nameplate.Background.DecodeId(); err != nil {
		return err
	}
	if err := shift.Data.CoopHistoryDetail.MyResult.Player.Uniform.DecodeId(); err != nil {
		return err
	}
	if err := shift.Data.CoopHistoryDetail.MyResult.DecodeId(); err != nil {
		return err
	}
	if err := shift.Data.CoopHistoryDetail.CoopStage.DecodeId(); err != nil {
		return err
	}
	if shift.Data.CoopHistoryDetail.BossResult != nil {
		if err := shift.Data.CoopHistoryDetail.BossResult.Boss.DecodeId(); err != nil {
			return err
		}
		if _, err := db.Exec("insert into three_salmon_shift (userId, typename, id, afterGrade, playerByname, playerBackground, playerName,\n                                playerNameId, playerUniform, playerId, playerSpecies, playerSpecialWeapon,\n                                playerDefeatEnemyCount, playerDeliverCount, playerGoldenAssistCount,\n                                playerGoldenDeliverCount, playerRescueCount, playerRescuedCount, resultWave, playedTime,\n                                rule, stage, dangerRate, scenarioCode, smellMeter, afterGradePoint, scaleBronze,\n                                scaleSilver, scaleGold, jobPoint, jobScore, jobRate, jobBonus, hasDefeatBoss, boss,\n                                weapon0, weapon1, weapon2, weapon3)\nvalues ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24,\n        $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39);",
			userId, shift.Data.CoopHistoryDetail.Typename, shift.Data.CoopHistoryDetail.ID, shift.Data.CoopHistoryDetail.AfterGrade.ID,
			shift.Data.CoopHistoryDetail.MyResult.Player.Byname, shift.Data.CoopHistoryDetail.MyResult.Player.Nameplate.Background.ID,
			shift.Data.CoopHistoryDetail.MyResult.Player.Name, shift.Data.CoopHistoryDetail.MyResult.Player.NameID,
			shift.Data.CoopHistoryDetail.MyResult.Player.Uniform.ID, shift.Data.CoopHistoryDetail.MyResult.Player.ID,
			shift.Data.CoopHistoryDetail.MyResult.Player.Species,
			shift.Data.CoopHistoryDetail.MyResult.SpecialWeapon.WeaponID,
			shift.Data.CoopHistoryDetail.MyResult.DefeatEnemyCount,
			shift.Data.CoopHistoryDetail.MyResult.DeliverCount,
			shift.Data.CoopHistoryDetail.MyResult.GoldenAssistCount,
			shift.Data.CoopHistoryDetail.MyResult.GoldenDeliverCount,
			shift.Data.CoopHistoryDetail.MyResult.RescueCount,
			shift.Data.CoopHistoryDetail.MyResult.RescuedCount, shift.Data.CoopHistoryDetail.ResultWave,
			shift.Data.CoopHistoryDetail.PlayedTime, shift.Data.CoopHistoryDetail.Rule, shift.Data.CoopHistoryDetail.CoopStage.ID,
			shift.Data.CoopHistoryDetail.DangerRate, shift.Data.CoopHistoryDetail.ScenarioCode,
			shift.Data.CoopHistoryDetail.SmellMeter, shift.Data.CoopHistoryDetail.AfterGradePoint,
			shift.Data.CoopHistoryDetail.Scale.Bronze, shift.Data.CoopHistoryDetail.Scale.Silver,
			shift.Data.CoopHistoryDetail.Scale.Gold, shift.Data.CoopHistoryDetail.JobPoint,
			shift.Data.CoopHistoryDetail.JobScore, shift.Data.CoopHistoryDetail.JobRate,
			shift.Data.CoopHistoryDetail.JobBonus, shift.Data.CoopHistoryDetail.BossResult.HasDefeatBoss, shift.Data.CoopHistoryDetail.BossResult.Boss.ID,
			shift.Data.CoopHistoryDetail.Weapons[0].Name, shift.Data.CoopHistoryDetail.Weapons[1].Name, shift.Data.CoopHistoryDetail.Weapons[2].Name, shift.Data.CoopHistoryDetail.Weapons[3].Name,
		); err != nil {
			return err
		}
	} else {
		if _, err := db.Exec("insert into three_salmon_shift (userId, typename, id, afterGrade, playerByname, playerBackground, playerName,\n                                playerNameId, playerUniform, playerId, playerSpecies, playerSpecialWeapon,\n                                playerDefeatEnemyCount, playerDeliverCount, playerGoldenAssistCount,\n                                playerGoldenDeliverCount, playerRescueCount, playerRescuedCount, resultWave, playedTime,\n                                rule, stage, dangerRate, scenarioCode, smellMeter, afterGradePoint, jobPoint, jobScore,\n                                jobRate, jobBonus, weapon0, weapon1, weapon2, weapon3)\nvalues ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24,\n        $25, $26, $27, $28, $29, $30, $31, $32, $33, $34);",
			userId, shift.Data.CoopHistoryDetail.Typename, shift.Data.CoopHistoryDetail.ID, shift.Data.CoopHistoryDetail.AfterGrade.ID,
			shift.Data.CoopHistoryDetail.MyResult.Player.Byname, shift.Data.CoopHistoryDetail.MyResult.Player.Nameplate.Background.ID,
			shift.Data.CoopHistoryDetail.MyResult.Player.Name,
			shift.Data.CoopHistoryDetail.MyResult.Player.NameID,
			shift.Data.CoopHistoryDetail.MyResult.Player.Uniform.ID, shift.Data.CoopHistoryDetail.MyResult.Player.ID,
			shift.Data.CoopHistoryDetail.MyResult.Player.Species,
			shift.Data.CoopHistoryDetail.MyResult.SpecialWeapon.WeaponID,
			shift.Data.CoopHistoryDetail.MyResult.DefeatEnemyCount,
			shift.Data.CoopHistoryDetail.MyResult.DeliverCount,
			shift.Data.CoopHistoryDetail.MyResult.GoldenAssistCount,
			shift.Data.CoopHistoryDetail.MyResult.GoldenDeliverCount,
			shift.Data.CoopHistoryDetail.MyResult.RescueCount,
			shift.Data.CoopHistoryDetail.MyResult.RescuedCount, shift.Data.CoopHistoryDetail.ResultWave,
			shift.Data.CoopHistoryDetail.PlayedTime, shift.Data.CoopHistoryDetail.Rule, shift.Data.CoopHistoryDetail.CoopStage.ID,
			shift.Data.CoopHistoryDetail.DangerRate, shift.Data.CoopHistoryDetail.ScenarioCode,
			shift.Data.CoopHistoryDetail.SmellMeter, shift.Data.CoopHistoryDetail.AfterGradePoint,
			shift.Data.CoopHistoryDetail.JobPoint, shift.Data.CoopHistoryDetail.JobScore,
			shift.Data.CoopHistoryDetail.JobRate, shift.Data.CoopHistoryDetail.JobBonus,
			shift.Data.CoopHistoryDetail.Weapons[0].Name, shift.Data.CoopHistoryDetail.Weapons[1].Name, shift.Data.CoopHistoryDetail.Weapons[2].Name, shift.Data.CoopHistoryDetail.Weapons[3].Name,
		); err != nil {
			return err
		}
	}
	for i, b := range shift.Data.CoopHistoryDetail.MyResult.Player.Nameplate.Badges {
		if err := b.DecodeId(); err != nil {
			return err
		}
		if _, err := db.Exec("insert into three_salmon_user_badge (userId, shiftId, badgeSlot, badge) values ($1, $2, $3, $4);",
			userId, shift.Data.CoopHistoryDetail.ID, i, b.ID); err != nil {
			return err
		}
	}
	for i, w := range shift.Data.CoopHistoryDetail.MyResult.Weapons {
		if _, err := db.Exec("insert into three_salmon_user_weapon (userId, shiftId, wave, weapon) values ($1, $2, $3, $4);",
			userId, shift.Data.CoopHistoryDetail.ID, i+1, w.Name); err != nil {
			return err
		}
	}
	for _, w := range shift.Data.CoopHistoryDetail.WaveResults {
		if w.EventWave != nil {
			if err := w.EventWave.DecodeId(); err != nil {
				return err
			}
			if _, err := db.Exec("insert into three_salmon_wave (userId, shiftId, waveNumber, waterLevel, eventWave, deliverNorm, goldenPopCount,\n                               teamDeliverCount)\nVALUES ($1, $2, $3, $4, $5, $6, $7, $8);",
				userId, shift.Data.CoopHistoryDetail.ID, w.WaveNumber, w.WaterLevel, w.EventWave.ID, w.DeliverNorm, w.GoldenPopCount, w.TeamDeliverCount); err != nil {
				return err
			}
		} else {
			if _, err := db.Exec("insert into three_salmon_wave (userId, shiftId, waveNumber, waterLevel, deliverNorm, goldenPopCount,\n                               teamDeliverCount)\nVALUES ($1, $2, $3, $4, $5, $6, $7);",
				userId, shift.Data.CoopHistoryDetail.ID, w.WaveNumber, w.WaterLevel, w.DeliverNorm, w.GoldenPopCount, w.TeamDeliverCount); err != nil {
				return err
			}
		}
		for _, we := range w.SpecialWeapons {
			if err := we.DecodeId(); err != nil {
				return err
			}
			if _, err := db.Exec("insert into three_salmon_wave_special (userId, shiftId, waveNumber, special) values ($1, $2, $3, $4);",
				userId, shift.Data.CoopHistoryDetail.ID, w.WaveNumber, we.ID); err != nil {
				return err
			}
		}
	}
	for _, e := range shift.Data.CoopHistoryDetail.EnemyResults {
		if err := e.Enemy.DecodeId(); err != nil {
			return err
		}
		if _, err := db.Exec("insert into three_salmon_enemy_result (userId, shiftID, defeatCount, teamDefeatCount, popCount, enemy) values ($1, $2, $3, $4, $5, $6);",
			userId, shift.Data.CoopHistoryDetail.ID, e.DefeatCount, e.TeamDefeatCount, e.PopCount, e.Enemy.ID); err != nil {
			return err
		}
	}
	for _, p := range shift.Data.CoopHistoryDetail.MemberResults {
		if err := p.DecodeId(); err != nil {
			return err
		}
		if err := p.Player.Nameplate.Background.DecodeId(); err != nil {
			return err
		}
		if err := p.Player.Uniform.DecodeId(); err != nil {
			return err
		}
		if _, err := db.Exec("insert into three_salmon_player (userId, shiftId, byname, name, nameId, background, uniform, id,\n                                  species, special, defeatEnemyCount, deliverCount, goldenAssistCount,\n                                  goldenDeliverCount, rescueCount, rescuedCount)\nvalues ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);",
			userId, shift.Data.CoopHistoryDetail.ID, p.Player.Byname, p.Player.Name, p.Player.NameID,
			p.Player.Nameplate.Background.ID, p.Player.Uniform.ID, p.Player.ID, p.Player.Species, p.SpecialWeapon.WeaponID,
			p.DefeatEnemyCount, p.DeliverCount, p.GoldenAssistCount, p.GoldenDeliverCount, p.RescueCount,
			p.RescuedCount); err != nil {
			return err
		}
		for i, w := range p.Weapons {
			if _, err := db.Exec("insert into three_salmon_player_weapon (userId, shiftId, playerId, wave, weapon) values ($1, $2, $3, $4, $5);",
				userId, shift.Data.CoopHistoryDetail.ID, p.Player.ID, i+1, w.Name); err != nil {
				return err
			}
		}
		for i, b := range p.Player.Nameplate.Badges {
			if err := b.DecodeId(); err != nil {
				return err
			}
			if _, err := db.Exec("insert into three_salmon_player_badge (userId, shiftId, playerId, badgeSlot, badge) values ($1, $2, $3, $4, $5);",
				userId, shift.Data.CoopHistoryDetail.ID, p.Player.ID, i, b.ID); err != nil {

			}
		}
	}
	return nil
}

func WriteNewBattle3(battle *api_objects.Battle3, userId int64) error {
	return nil
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
