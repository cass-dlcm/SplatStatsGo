CREATE TABLE `auth_user` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`username` varchar(128) NOT NULL,
	`password` char(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
	`email` varchar(128) NOT NULL,
	`email_verified` tinyint unsigned NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `username_UNIQUE` (`username`),
	UNIQUE KEY `email_UNIQUE` (`email`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `auth_user_session_token` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`session_token` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`),
	UNIQUE KEY `session_token_UNIQUE` (`session_token`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `int_container` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`parent_table` varchar(64) NOT NULL,
	`value` int NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=20439 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `splatnet_double` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=104693 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `splatnet_player_type` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`gender` varchar(4) NOT NULL,
	`species` varchar(9) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=13475 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `splatnet_quad` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` varchar(45) NOT NULL,
	`image_a` varchar(60) NOT NULL,
	`image_b` varchar(60) NOT NULL,
	`name` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=14539 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `splatnet_triple` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` varchar(45) NOT NULL,
	`image` char(58) NOT NULL,
	`name` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=47663 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `stat_ink_gender` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` bigint NOT NULL,
	`iso5218` int NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3392 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `stat_ink_key_name` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` bigint NOT NULL,
	PRIMARY KEY (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=29221 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `stat_ink_name` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`de_DE` varchar(45) NOT NULL,
	`en_GB` varchar(45) NOT NULL,
	`en_US` varchar(45) NOT NULL,
	`es_ES` varchar(45) NOT NULL,
	`es_MX` varchar(45) NOT NULL,
	`fr_CA` varchar(45) NOT NULL,
	`fr_FR` varchar(45) NOT NULL,
	`it_IT` varchar(45) NOT NULL,
	`ja_JP` varchar(45) NOT NULL,
	`nl_NL` varchar(45) NOT NULL,
	`ru_RU` varchar(45) NOT NULL,
	`zh_CN` varchar(45) NOT NULL,
	`zh_TW` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=44436 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `stat_ink_profile` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`friend_code` varchar(45) DEFAULT NULL,
	`twitter` varchar(45) DEFAULT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=386 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `stat_ink_time` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`time` int NOT NULL,
	`iso8601` datetime NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=2304 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`user_id` bigint NOT NULL,
	`splatnet_json` bigint DEFAULT NULL,
	`splatnet_upload` tinyint(1) NOT NULL,
	`stat_ink_json` bigint DEFAULT NULL,
	`stat_ink_upload` tinyint(1) NOT NULL,
	`splatnet_number` int NOT NULL,
	`player_splatnet_id` char(16) NOT NULL,
	`elapsed_time` int NOT NULL,
	`has_disconnected_player` tinyint(1) NOT NULL,
	`league_point` decimal(5,1) DEFAULT NULL,
	`match_type` varchar(45) NOT NULL,
	`rule` varchar(45) NOT NULL,
	`my_team_count` decimal(4,1) NOT NULL,
	`other_team_count` decimal(4,1) NOT NULL,
	`splatfest_point` decimal(5,1) DEFAULT NULL,
	`splatfest_title` varchar(45) DEFAULT NULL,
	`stage` varchar(45) NOT NULL,
	`tag_id` varchar(45) DEFAULT NULL,
	`time` int NOT NULL,
	`win` tinyint(1) NOT NULL,
	`win_meter` decimal(3,1) DEFAULT NULL,
	`opponent0_splatnet_id` char(16) DEFAULT NULL,
	`opponent0_name` varchar(10) DEFAULT NULL,
	`opponent0_rank` varchar(45) DEFAULT NULL,
	`opponent0_level_star` int DEFAULT NULL,
	`opponent0_level` int DEFAULT NULL,
	`opponent0_weapon` varchar(45) DEFAULT NULL,
	`opponent0_gender` varchar(4) DEFAULT NULL,
	`opponent0_species` varchar(9) DEFAULT NULL,
	`opponent0_assists` int DEFAULT NULL,
	`opponent0_deaths` int DEFAULT NULL,
	`opponent0_game_paint_point` int DEFAULT NULL,
	`opponent0_kills` int DEFAULT NULL,
	`opponent0_specials` int DEFAULT NULL,
	`opponent0_headgear` varchar(45) DEFAULT NULL,
	`opponent0_headgear_main` varchar(45) DEFAULT NULL,
	`opponent0_headgear_sub0` varchar(45) DEFAULT NULL,
	`opponent0_headgear_sub1` varchar(45) DEFAULT NULL,
	`opponent0_headgear_sub2` varchar(45) DEFAULT NULL,
	`opponent0_clothes` varchar(45) DEFAULT NULL,
	`opponent0_clothes_main` varchar(45) DEFAULT NULL,
	`opponent0_clothes_sub0` varchar(45) DEFAULT NULL,
	`opponent0_clothes_sub1` varchar(45) DEFAULT NULL,
	`opponent0_clothes_sub2` varchar(45) DEFAULT NULL,
	`opponent0_shoes` varchar(45) DEFAULT NULL,
	`opponent0_shoes_main` varchar(45) DEFAULT NULL,
	`opponent0_shoes_sub0` varchar(45) DEFAULT NULL,
	`opponent0_shoes_sub1` varchar(45) DEFAULT NULL,
	`opponent0_shoes_sub2` varchar(45) DEFAULT NULL,
	`opponent1_splatnet_id` char(16) DEFAULT NULL,
	`opponent1_name` varchar(10) DEFAULT NULL,
	`opponent1_rank` varchar(45) DEFAULT NULL,
	`opponent1_level_star` int DEFAULT NULL,
	`opponent1_level` int DEFAULT NULL,
	`opponent1_weapon` varchar(45) DEFAULT NULL,
	`opponent1_gender` varchar(4) DEFAULT NULL,
	`opponent1_species` varchar(9) DEFAULT NULL,
	`opponent1_assists` int DEFAULT NULL,
	`opponent1_deaths` int DEFAULT NULL,
	`opponent1_game_paint_point` int DEFAULT NULL,
	`opponent1_kills` int DEFAULT NULL,
	`opponent1_specials` int DEFAULT NULL,
	`opponent1_headgear` varchar(45) DEFAULT NULL,
	`opponent1_headgear_main` varchar(45) DEFAULT NULL,
	`opponent1_headgear_sub0` varchar(45) DEFAULT NULL,
	`opponent1_headgear_sub1` varchar(45) DEFAULT NULL,
	`opponent1_headgear_sub2` varchar(45) DEFAULT NULL,
	`opponent1_clothes` varchar(45) DEFAULT NULL,
	`opponent1_clothes_main` varchar(45) DEFAULT NULL,
	`opponent1_clothes_sub0` varchar(45) DEFAULT NULL,
	`opponent1_clothes_sub1` varchar(45) DEFAULT NULL,
	`opponent1_clothes_sub2` varchar(45) DEFAULT NULL,
	`opponent1_shoes` varchar(45) DEFAULT NULL,
	`opponent1_shoes_main` varchar(45) DEFAULT NULL,
	`opponent1_shoes_sub0` varchar(45) DEFAULT NULL,
	`opponent1_shoes_sub1` varchar(45) DEFAULT NULL,
	`opponent1_shoes_sub2` varchar(45) DEFAULT NULL,
	`opponent2_splatnet_id` char(16) DEFAULT NULL,
	`opponent2_name` varchar(10) DEFAULT NULL,
	`opponent2_rank` varchar(45) DEFAULT NULL,
	`opponent2_level_star` int DEFAULT NULL,
	`opponent2_level` int DEFAULT NULL,
	`opponent2_weapon` varchar(45) DEFAULT NULL,
	`opponent2_gender` varchar(4) DEFAULT NULL,
	`opponent2_species` varchar(9) DEFAULT NULL,
	`opponent2_assists` int DEFAULT NULL,
	`opponent2_deaths` int DEFAULT NULL,
	`opponent2_game_paint_point` int DEFAULT NULL,
	`opponent2_kills` int DEFAULT NULL,
	`opponent2_specials` int DEFAULT NULL,
	`opponent2_headgear` varchar(45) DEFAULT NULL,
	`opponent2_headgear_main` varchar(45) DEFAULT NULL,
	`opponent2_headgear_sub0` varchar(45) DEFAULT NULL,
	`opponent2_headgear_sub1` varchar(45) DEFAULT NULL,
	`opponent2_headgear_sub2` varchar(45) DEFAULT NULL,
	`opponent2_clothes` varchar(45) DEFAULT NULL,
	`opponent2_clothes_main` varchar(45) DEFAULT NULL,
	`opponent2_clothes_sub0` varchar(45) DEFAULT NULL,
	`opponent2_clothes_sub1` varchar(45) DEFAULT NULL,
	`opponent2_clothes_sub2` varchar(45) DEFAULT NULL,
	`opponent2_shoes` varchar(45) DEFAULT NULL,
	`opponent2_shoes_main` varchar(45) DEFAULT NULL,
	`opponent2_shoes_sub0` varchar(45) DEFAULT NULL,
	`opponent2_shoes_sub1` varchar(45) DEFAULT NULL,
	`opponent2_shoes_sub2` varchar(45) DEFAULT NULL,
	`opponent3_splatnet_id` char(16) DEFAULT NULL,
	`opponent3_name` varchar(10) DEFAULT NULL,
	`opponent3_rank` varchar(45) DEFAULT NULL,
	`opponent3_level_star` int DEFAULT NULL,
	`opponent3_level` int DEFAULT NULL,
	`opponent3_weapon` varchar(45) DEFAULT NULL,
	`opponent3_gender` varchar(4) DEFAULT NULL,
	`opponent3_species` varchar(9) DEFAULT NULL,
	`opponent3_assists` int DEFAULT NULL,
	`opponent3_deaths` int DEFAULT NULL,
	`opponent3_game_paint_point` int DEFAULT NULL,
	`opponent3_kills` int DEFAULT NULL,
	`opponent3_specials` int DEFAULT NULL,
	`opponent3_headgear` varchar(45) DEFAULT NULL,
	`opponent3_headgear_main` varchar(45) DEFAULT NULL,
	`opponent3_headgear_sub0` varchar(45) DEFAULT NULL,
	`opponent3_headgear_sub1` varchar(45) DEFAULT NULL,
	`opponent3_headgear_sub2` varchar(45) DEFAULT NULL,
	`opponent3_clothes` varchar(45) DEFAULT NULL,
	`opponent3_clothes_main` varchar(45) DEFAULT NULL,
	`opponent3_clothes_sub0` varchar(45) DEFAULT NULL,
	`opponent3_clothes_sub1` varchar(45) DEFAULT NULL,
	`opponent3_clothes_sub2` varchar(45) DEFAULT NULL,
	`opponent3_shoes` varchar(45) DEFAULT NULL,
	`opponent3_shoes_main` varchar(45) DEFAULT NULL,
	`opponent3_shoes_sub0` varchar(45) DEFAULT NULL,
	`opponent3_shoes_sub1` varchar(45) DEFAULT NULL,
	`opponent3_shoes_sub2` varchar(45) DEFAULT NULL,
	`teammate0_splatnet_id` char(16) DEFAULT NULL,
	`teammate0_name` varchar(10) DEFAULT NULL,
	`teammate0_rank` varchar(45) DEFAULT NULL,
	`teammate0_level_star` int DEFAULT NULL,
	`teammate0_level` int DEFAULT NULL,
	`teammate0_weapon` varchar(45) DEFAULT NULL,
	`teammate0_gender` varchar(4) DEFAULT NULL,
	`teammate0_species` varchar(9) DEFAULT NULL,
	`teammate0_assists` int DEFAULT NULL,
	`teammate0_deaths` int DEFAULT NULL,
	`teammate0_game_paint_point` int DEFAULT NULL,
	`teammate0_kills` int DEFAULT NULL,
	`teammate0_specials` int DEFAULT NULL,
	`teammate0_headgear` varchar(45) DEFAULT NULL,
	`teammate0_headgear_main` varchar(45) DEFAULT NULL,
	`teammate0_headgear_sub0` varchar(45) DEFAULT NULL,
	`teammate0_headgear_sub1` varchar(45) DEFAULT NULL,
	`teammate0_headgear_sub2` varchar(45) DEFAULT NULL,
	`teammate0_clothes` varchar(45) DEFAULT NULL,
	`teammate0_clothes_main` varchar(45) DEFAULT NULL,
	`teammate0_clothes_sub0` varchar(45) DEFAULT NULL,
	`teammate0_clothes_sub1` varchar(45) DEFAULT NULL,
	`teammate0_clothes_sub2` varchar(45) DEFAULT NULL,
	`teammate0_shoes` varchar(45) DEFAULT NULL,
	`teammate0_shoes_main` varchar(45) DEFAULT NULL,
	`teammate0_shoes_sub0` varchar(45) DEFAULT NULL,
	`teammate0_shoes_sub1` varchar(45) DEFAULT NULL,
	`teammate0_shoes_sub2` varchar(45) DEFAULT NULL,
	`teammate1_splatnet_id` char(16) DEFAULT NULL,
	`teammate1_name` varchar(10) DEFAULT NULL,
	`teammate1_rank` varchar(45) DEFAULT NULL,
	`teammate1_level_star` int DEFAULT NULL,
	`teammate1_level` int DEFAULT NULL,
	`teammate1_weapon` varchar(45) DEFAULT NULL,
	`teammate1_gender` varchar(4) DEFAULT NULL,
	`teammate1_species` varchar(9) DEFAULT NULL,
	`teammate1_assists` int DEFAULT NULL,
	`teammate1_deaths` int DEFAULT NULL,
	`teammate1_game_paint_point` int DEFAULT NULL,
	`teammate1_kills` int DEFAULT NULL,
	`teammate1_specials` int DEFAULT NULL,
	`teammate1_headgear` varchar(45) DEFAULT NULL,
	`teammate1_headgear_main` varchar(45) DEFAULT NULL,
	`teammate1_headgear_sub0` varchar(45) DEFAULT NULL,
	`teammate1_headgear_sub1` varchar(45) DEFAULT NULL,
	`teammate1_headgear_sub2` varchar(45) DEFAULT NULL,
	`teammate1_clothes` varchar(45) DEFAULT NULL,
	`teammate1_clothes_main` varchar(45) DEFAULT NULL,
	`teammate1_clothes_sub0` varchar(45) DEFAULT NULL,
	`teammate1_clothes_sub1` varchar(45) DEFAULT NULL,
	`teammate1_clothes_sub2` varchar(45) DEFAULT NULL,
	`teammate1_shoes` varchar(45) DEFAULT NULL,
	`teammate1_shoes_main` varchar(45) DEFAULT NULL,
	`teammate1_shoes_sub0` varchar(45) DEFAULT NULL,
	`teammate1_shoes_sub1` varchar(45) DEFAULT NULL,
	`teammate1_shoes_sub2` varchar(45) DEFAULT NULL,
	`teammate2_splatnet_id` char(16) DEFAULT NULL,
	`teammate2_name` varchar(10) DEFAULT NULL,
	`teammate2_rank` varchar(45) DEFAULT NULL,
	`teammate2_level_star` int DEFAULT NULL,
	`teammate2_level` int DEFAULT NULL,
	`teammate2_weapon` varchar(45) DEFAULT NULL,
	`teammate2_gender` varchar(4) DEFAULT NULL,
	`teammate2_species` varchar(9) DEFAULT NULL,
	`teammate2_assists` int DEFAULT NULL,
	`teammate2_deaths` int DEFAULT NULL,
	`teammate2_game_paint_point` int DEFAULT NULL,
	`teammate2_kills` int DEFAULT NULL,
	`teammate2_specials` int DEFAULT NULL,
	`teammate2_headgear` varchar(45) DEFAULT NULL,
	`teammate2_headgear_main` varchar(45) DEFAULT NULL,
	`teammate2_headgear_sub0` varchar(45) DEFAULT NULL,
	`teammate2_headgear_sub1` varchar(45) DEFAULT NULL,
	`teammate2_headgear_sub2` varchar(45) DEFAULT NULL,
	`teammate2_clothes` varchar(45) DEFAULT NULL,
	`teammate2_clothes_main` varchar(45) DEFAULT NULL,
	`teammate2_clothes_sub0` varchar(45) DEFAULT NULL,
	`teammate2_clothes_sub1` varchar(45) DEFAULT NULL,
	`teammate2_clothes_sub2` varchar(45) DEFAULT NULL,
	`teammate2_shoes` varchar(45) DEFAULT NULL,
	`teammate2_shoes_main` varchar(45) DEFAULT NULL,
	`teammate2_shoes_sub0` varchar(45) DEFAULT NULL,
	`teammate2_shoes_sub1` varchar(45) DEFAULT NULL,
	`teammate2_shoes_sub2` varchar(45) DEFAULT NULL,
	`player_name` varchar(10) NOT NULL,
	`player_rank` varchar(45) DEFAULT NULL,
	`player_level_star` int NOT NULL,
	`player_level` int NOT NULL,
	`player_weapon` varchar(45) NOT NULL,
	`player_gender` varchar(4) NOT NULL,
	`player_species` varchar(9) NOT NULL,
	`player_assists` int NOT NULL,
	`player_deaths` int NOT NULL,
	`player_game_paint_point` int NOT NULL,
	`player_kills` int NOT NULL,
	`player_specials` int NOT NULL,
	`player_headgear` varchar(45) NOT NULL,
	`player_headgear_main` varchar(45) NOT NULL,
	`player_headgear_sub0` varchar(45) NOT NULL,
	`player_headgear_sub1` varchar(45) NOT NULL,
	`player_headgear_sub2` varchar(45) NOT NULL,
	`player_clothes` varchar(45) NOT NULL,
	`player_clothes_main` varchar(45) NOT NULL,
	`player_clothes_sub0` varchar(45) NOT NULL,
	`player_clothes_sub1` varchar(45) NOT NULL,
	`player_clothes_sub2` varchar(45) NOT NULL,
	`player_shoes` varchar(45) NOT NULL,
	`player_shoes_main` varchar(45) NOT NULL,
	`player_shoes_sub0` varchar(45) NOT NULL,
	`player_shoes_sub1` varchar(45) NOT NULL,
	`player_shoes_sub2` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=857 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`udemae` bigint DEFAULT NULL,
	`stage` bigint NOT NULL,
	`other_team_count` int DEFAULT NULL,
	`my_team_count` int DEFAULT NULL,
	`star_rank` int NOT NULL,
	`rule` bigint NOT NULL,
	`player_result` bigint NOT NULL,
	`estimate_gachi_power` int DEFAULT NULL,
	`elapsed_time` int NOT NULL,
	`start_time` int NOT NULL,
	`game_mode` bigint NOT NULL,
	`battle_number` varchar(45) NOT NULL,
	`type` varchar(45) NOT NULL,
	`player_rank` int NOT NULL,
	`weapon_paint_point` int NOT NULL,
	`my_team_result` bigint NOT NULL,
	`other_team_result` bigint NOT NULL,
	`league_point` decimal(5,1) DEFAULT NULL,
	`win_meter` decimal(3,1) DEFAULT NULL,
	`my_team_percentage` decimal(4,1) DEFAULT NULL,
	`other_team_percentage` decimal(4,1) DEFAULT NULL,
	`tag_id` varchar(45) DEFAULT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=489 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_player_result` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`death_count` int NOT NULL,
	`game_paint_point` int NOT NULL,
	`kill_count` int NOT NULL,
	`special_count` int NOT NULL,
	`assist_count` int NOT NULL,
	`sort_score` int NOT NULL,
	`player` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3137 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_player_result_player` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`head_skills` bigint NOT NULL,
	`shoes_skills` bigint NOT NULL,
	`clothes_skills` bigint NOT NULL,
	`player_rank` int NOT NULL,
	`star_rank` int NOT NULL,
	`nickname` varchar(10) NOT NULL,
	`player_type` bigint NOT NULL,
	`principal_id` char(16) NOT NULL,
	`head` bigint NOT NULL,
	`clothes` bigint NOT NULL,
	`shoes` bigint NOT NULL,
	`udemae` bigint DEFAULT NULL,
	`weapon` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3137 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_player_result_player_clothing` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` varchar(45) NOT NULL,
	`image` char(60) NOT NULL,
	`name` varchar(45) NOT NULL,
	`thumbnail` char(60) NOT NULL,
	`kind` varchar(45) NOT NULL,
	`rarity` varchar(45) NOT NULL,
	`brand` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=9418 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_player_result_player_clothing_brand` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` varchar(45) NOT NULL,
	`image` char(58) NOT NULL,
	`name` varchar(45) NOT NULL,
	`frequent_skill` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=9418 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_player_result_player_skills` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`main` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=9439 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_player_result_player_skills_sub` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`sub` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=28313 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_player_result_player_weapon` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` varchar(45) NOT NULL,
	`image` char(59) NOT NULL,
	`name` varchar(45) NOT NULL,
	`thumbnail` char(59) NOT NULL,
	`sub` bigint NOT NULL,
	`special` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3137 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_rule` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` varchar(45) NOT NULL,
	`multiline_name` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=494 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_team_member` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`my_team` tinyint(1) NOT NULL,
	`parent` bigint NOT NULL,
	`player_result` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=2643 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_splatnet_udemae` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`name` varchar(45) NOT NULL,
	`is_x` tinyint(1) NOT NULL,
	`is_number_reached` tinyint(1) NOT NULL,
	`number` int NOT NULL,
	`s_plus_number` int DEFAULT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3530 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` int NOT NULL,
	`splatnet_number` int NOT NULL,
	`url` varchar(45) NOT NULL,
	`user` bigint NOT NULL,
	`lobby` bigint NOT NULL,
	`mode` bigint NOT NULL,
	`rule` bigint NOT NULL,
	`map` bigint NOT NULL,
	`weapon` bigint NOT NULL,
	`freshness` bigint DEFAULT NULL,
	`rank` bigint DEFAULT NULL,
	`rank_exp` int DEFAULT NULL,
	`rank_after` bigint DEFAULT NULL,
	`level` int NOT NULL,
	`level_after` int NOT NULL,
	`star_rank` int NOT NULL,
	`result` varchar(45) NOT NULL,
	`knock_out` tinyint(1) NOT NULL,
	`rank_in_team` int NOT NULL,
	`kill` int NOT NULL,
	`death` int NOT NULL,
	`kill_or_assist` int NOT NULL,
	`special` int NOT NULL,
	`kill_ratio` float NOT NULL,
	`kill_rate` float NOT NULL,
	`my_point` int NOT NULL,
	`estimate_gachi_power` int DEFAULT NULL,
	`league_point` varchar(45) DEFAULT NULL,
	`my_team_estimate_league_point` int DEFAULT NULL,
	`his_team_estimate_league_point` int DEFAULT NULL,
	`my_team_percent` varchar(45) DEFAULT NULL,
	`his_team_percent` varchar(45) DEFAULT NULL,
	`my_team_id` varchar(45) DEFAULT NULL,
	`his_team_id` varchar(45) DEFAULT NULL,
	`species` bigint NOT NULL,
	`gender` bigint NOT NULL,
	`fest_title` bigint DEFAULT NULL,
	`fest_exp` int DEFAULT NULL,
	`fest_title_after` bigint DEFAULT NULL,
	`fest_exp_after` int DEFAULT NULL,
	`fest_power` varchar(45) DEFAULT NULL,
	`my_team_estimate_fest_power` int DEFAULT NULL,
	`his_team_my_team_estimate_fest_power` int DEFAULT NULL,
	`my_team_fest_theme` varchar(45) DEFAULT NULL,
	`my_team_nickname` varchar(45) DEFAULT NULL,
	`his_team_nickname` varchar(45) DEFAULT NULL,
	`clout` int DEFAULT NULL,
	`total_clout` int DEFAULT NULL,
	`total_clout_after` int DEFAULT NULL,
	`my_team_win_streak` int DEFAULT NULL,
	`his_team_win_streak` int DEFAULT NULL,
	`synergy_bonus` float DEFAULT NULL,
	`special_battle` bigint DEFAULT NULL,
	`image_result` char(54) DEFAULT NULL,
	`image_gear` varchar(45) DEFAULT NULL,
	`gears` bigint NOT NULL,
	`period` int NOT NULL,
	`period_range` char(51) NOT NULL,
	`agent` bigint NOT NULL,
	`automated` tinyint(1) NOT NULL,
	`link_url` varchar(45) NOT NULL,
	`game_version` varchar(45) NOT NULL,
	`nawabari_bonus` int DEFAULT NULL,
	`start_at` bigint NOT NULL,
	`end_at` bigint NOT NULL,
	`register_at` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=379 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_agent` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`name` varchar(45) NOT NULL,
	`version` varchar(45) NOT NULL,
	`variables` bigint DEFAULT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=382 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_agent_variables` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`upload_mode` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_freshness` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`freshness` decimal(3,1) NOT NULL,
	`title` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=86 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_gears` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`headgear` bigint NOT NULL,
	`clothing` bigint NOT NULL,
	`shoes` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=383 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_gears_clothes` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`gear` bigint NOT NULL,
	`primary_ability` bigint NOT NULL,
	PRIMARY KEY (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=1149 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_gears_clothes_gear` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` bigint NOT NULL,
	`splatnet` int NOT NULL,
	`type` bigint NOT NULL,
	`brand` bigint NOT NULL,
	`primary_ability` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=1149 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_gears_clothes_sa_container` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`secondary_ability` bigint NOT NULL,
	PRIMARY KEY (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3445 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_map` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` bigint NOT NULL,
	`splatnet` int NOT NULL,
	`area` int NOT NULL,
	`release_at` bigint NOT NULL,
	`short_name` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=384 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_player` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`team` varchar(45) NOT NULL,
	`is_me` tinyint(1) NOT NULL,
	`weapon` bigint NOT NULL,
	`level` int NOT NULL,
	`rank` bigint DEFAULT NULL,
	`star_rank` int NOT NULL,
	`rank_in_team` int NOT NULL,
	`kill` int NOT NULL,
	`death` int NOT NULL,
	`kill_or_assist` int NOT NULL,
	`special` int NOT NULL,
	`point` int NOT NULL,
	`name` varchar(10) NOT NULL,
	`species` bigint NOT NULL,
	`gender` bigint NOT NULL,
	`fest_title` bigint DEFAULT NULL,
	`splatnet_id` char(16) NOT NULL,
	`top_500` tinyint(1) NOT NULL,
	`icon` char(71) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3007 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_rank` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` bigint NOT NULL,
	`zone` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3031 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_user` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` int NOT NULL,
	`name` varchar(45) NOT NULL,
	`screen_name` varchar(45) NOT NULL,
	`url` varchar(66) NOT NULL,
	`join_at` bigint NOT NULL,
	`profile` bigint NOT NULL,
	`stats` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=384 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_user_stats` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`v2` bigint DEFAULT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=385 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_user_stats_v2` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`updated_at` bigint NOT NULL,
	`entire` bigint NOT NULL,
	`nawabari` bigint NOT NULL,
	`gachi` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=385 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_user_stats_v2_entire` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`battles` int NOT NULL,
	`win_pct` float NOT NULL,
	`kill_ratio` float NOT NULL,
	`kill_total` int NOT NULL,
	`kill_avg` float NOT NULL,
	`kill_per_min` float NOT NULL,
	`death_total` int NOT NULL,
	`death_avg` float NOT NULL,
	`death_per_min` float NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=386 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_user_stats_v2_gachi` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`battles` int NOT NULL,
	`win_pct` float NOT NULL,
	`kill_ratio` float NOT NULL,
	`kill_total` int NOT NULL,
	`kill_avg` float NOT NULL,
	`kill_per_min` float NOT NULL,
	`death_total` int NOT NULL,
	`death_avg` float NOT NULL,
	`death_per_min` float NOT NULL,
	`rules` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=385 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_user_stats_v2_gachi_rules` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`area` bigint NOT NULL,
	`yagura` bigint NOT NULL,
	`hoko` bigint NOT NULL,
	`asari` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=385 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_user_stats_v2_gachi_rules_sub` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`rank_peak` varchar(45) NOT NULL,
	`rank_current` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=1537 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_user_stats_v2_nawabari` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`battles` int NOT NULL,
	`win_pct` float NOT NULL,
	`kill_ratio` float NOT NULL,
	`kill_total` int NOT NULL,
	`kill_avg` float NOT NULL,
	`kill_per_min` float NOT NULL,
	`death_total` int NOT NULL,
	`death_avg` float NOT NULL,
	`death_per_min` float NOT NULL,
	`total_inked` int NOT NULL,
	`max_inked` int NOT NULL,
	`avg_inked` float NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=386 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_weapon` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` bigint NOT NULL,
	`splatnet` int NOT NULL,
	`type` bigint NOT NULL,
	`reskin_of` varchar(45) DEFAULT NULL,
	`main_ref` varchar(45) NOT NULL,
	`sub` bigint NOT NULL,
	`special` bigint NOT NULL,
	`main_power_up` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3397 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_battles_battle_stat_ink_weapon_type` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` bigint NOT NULL,
	`category` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=3398 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`user_id` bigint NOT NULL,
	`player_splatnet_id` char(16) NOT NULL,
	`splatnet_number` bigint NOT NULL,
	`splatnet_json` bigint DEFAULT NULL,
	`stat_ink_json` bigint DEFAULT NULL,
	`start_time` bigint NOT NULL,
	`play_time` bigint NOT NULL,
	`end_time` bigint NOT NULL,
	`danger_rate` decimal(4,1) NOT NULL,
	`is_clear` tinyint NOT NULL,
	`job_failure_reason` varchar(10) DEFAULT NULL,
	`failure_wave` int DEFAULT NULL,
	`grade_point` int NOT NULL,
	`grade_point_delta` int NOT NULL,
	`job_score` int NOT NULL,
	`drizzler_count` int NOT NULL,
	`flyfish_count` int NOT NULL,
	`goldie_count` int NOT NULL,
	`griller_count` int NOT NULL,
	`maws_count` int NOT NULL,
	`scrapper_count` int NOT NULL,
	`steel_eel_count` int NOT NULL,
	`steelhead_count` int NOT NULL,
	`stinger_count` int NOT NULL,
	`stage` varchar(20) NOT NULL,
	`player_name` varchar(10) NOT NULL,
	`player_death_count` int NOT NULL,
	`player_revive_count` int NOT NULL,
	`player_golden_eggs` int NOT NULL,
	`player_power_eggs` int NOT NULL,
	`player_special` varchar(45) NOT NULL,
	`player_title` varchar(45) NOT NULL,
	`player_species` varchar(9) NOT NULL,
	`player_gender` varchar(4) NOT NULL,
	`player_w1_specials` int NOT NULL,
	`player_w2_specials` int DEFAULT NULL,
	`player_w3_specials` int DEFAULT NULL,
	`player_w1_weapon` varchar(19) NOT NULL,
	`player_w2_weapon` varchar(19) DEFAULT NULL,
	`player_w3_weapon` varchar(19) DEFAULT NULL,
	`player_drizzler_kills` int NOT NULL,
	`player_flyfish_kills` int NOT NULL,
	`player_goldie_kills` int NOT NULL,
	`player_griller_kills` int NOT NULL,
	`player_maws_kills` int NOT NULL,
	`player_scrapper_kills` int NOT NULL,
	`player_steel_eel_kills` int NOT NULL,
	`player_steelhead_kills` int NOT NULL,
	`player_stinger_kills` int NOT NULL,
	`teammate0_splatnet_id` char(16) DEFAULT NULL,
	`teammate0_name` varchar(10) DEFAULT NULL,
	`teammate0_death_count` int DEFAULT NULL,
	`teammate0_revive_count` int DEFAULT NULL,
	`teammate0_golden_eggs` int DEFAULT NULL,
	`teammate0_power_eggs` int DEFAULT NULL,
	`teammate0_special` varchar(45) DEFAULT NULL,
	`teammate0_species` varchar(9) DEFAULT NULL,
	`teammate0_gender` varchar(4) DEFAULT NULL,
	`teammate0_w1_specials` int DEFAULT NULL,
	`teammate0_w2_specials` int DEFAULT NULL,
	`teammate0_w3_specials` int DEFAULT NULL,
	`teammate0_w1_weapon` varchar(19) DEFAULT NULL,
	`teammate0_w2_weapon` varchar(19) DEFAULT NULL,
	`teammate0_w3_weapon` varchar(19) DEFAULT NULL,
	`teammate0_drizzler_kills` int DEFAULT NULL,
	`teammate0_flyfish_kills` int DEFAULT NULL,
	`teammate0_goldie_kills` int DEFAULT NULL,
	`teammate0_griller_kills` int DEFAULT NULL,
	`teammate0_maws_kills` int DEFAULT NULL,
	`teammate0_scrapper_kills` int DEFAULT NULL,
	`teammate0_steel_eel_kills` int DEFAULT NULL,
	`teammate0_steelhead_kills` int DEFAULT NULL,
	`teammate0_stinger_kills` int DEFAULT NULL,
	`teammate1_splatnet_id` char(16) DEFAULT NULL,
	`teammate1_name` varchar(10) DEFAULT NULL,
	`teammate1_death_count` int DEFAULT NULL,
	`teammate1_revive_count` int DEFAULT NULL,
	`teammate1_golden_eggs` int DEFAULT NULL,
	`teammate1_power_eggs` int DEFAULT NULL,
	`teammate1_special` varchar(45) DEFAULT NULL,
	`teammate1_species` varchar(9) DEFAULT NULL,
	`teammate1_gender` varchar(4) DEFAULT NULL,
	`teammate1_w1_specials` int DEFAULT NULL,
	`teammate1_w2_specials` int DEFAULT NULL,
	`teammate1_w3_specials` int DEFAULT NULL,
	`teammate1_w1_weapon` varchar(19) DEFAULT NULL,
	`teammate1_w2_weapon` varchar(19) DEFAULT NULL,
	`teammate1_w3_weapon` varchar(19) DEFAULT NULL,
	`teammate1_drizzler_kills` int DEFAULT NULL,
	`teammate1_flyfish_kills` int DEFAULT NULL,
	`teammate1_goldie_kills` int DEFAULT NULL,
	`teammate1_griller_kills` int DEFAULT NULL,
	`teammate1_maws_kills` int DEFAULT NULL,
	`teammate1_scrapper_kills` int DEFAULT NULL,
	`teammate1_steel_eel_kills` int DEFAULT NULL,
	`teammate1_steelhead_kills` int DEFAULT NULL,
	`teammate1_stinger_kills` int DEFAULT NULL,
	`teammate2_splatnet_id` char(16) DEFAULT NULL,
	`teammate2_name` varchar(10) DEFAULT NULL,
	`teammate2_death_count` int DEFAULT NULL,
	`teammate2_revive_count` int DEFAULT NULL,
	`teammate2_golden_eggs` int DEFAULT NULL,
	`teammate2_power_eggs` int DEFAULT NULL,
	`teammate2_special` varchar(45) DEFAULT NULL,
	`teammate2_species` varchar(9) DEFAULT NULL,
	`teammate2_gender` varchar(4) DEFAULT NULL,
	`teammate2_w1_specials` int DEFAULT NULL,
	`teammate2_w2_specials` int DEFAULT NULL,
	`teammate2_w3_specials` int DEFAULT NULL,
	`teammate2_w1_weapon` varchar(19) DEFAULT NULL,
	`teammate2_w2_weapon` varchar(19) DEFAULT NULL,
	`teammate2_w3_weapon` varchar(19) DEFAULT NULL,
	`teammate2_drizzler_kills` int DEFAULT NULL,
	`teammate2_flyfish_kills` int DEFAULT NULL,
	`teammate2_goldie_kills` int DEFAULT NULL,
	`teammate2_griller_kills` int DEFAULT NULL,
	`teammate2_maws_kills` int DEFAULT NULL,
	`teammate2_scrapper_kills` int DEFAULT NULL,
	`teammate2_steel_eel_kills` int DEFAULT NULL,
	`teammate2_steelhead_kills` int DEFAULT NULL,
	`teammate2_stinger_kills` int DEFAULT NULL,
	`schedule_end_time` bigint DEFAULT NULL,
	`schedule_start_time` bigint DEFAULT NULL,
	`schedule_weapon0` varchar(19) DEFAULT NULL,
	`schedule_weapon1` varchar(19) DEFAULT NULL,
	`schedule_weapon2` varchar(19) DEFAULT NULL,
	`schedule_weapon3` varchar(19) DEFAULT NULL,
	`wave1_water_level` varchar(45) NOT NULL,
	`wave1_event_type` varchar(45) NOT NULL,
	`wave1_golden_ikura_num` int NOT NULL,
	`wave1_golden_ikura_pop_num` int NOT NULL,
	`wave1_ikura_num` int NOT NULL,
	`wave1_quota_num` int NOT NULL,
	`wave2_water_level` varchar(45) DEFAULT NULL,
	`wave2_event_type` varchar(45) DEFAULT NULL,
	`wave2_golden_ikura_num` int DEFAULT NULL,
	`wave2_golden_ikura_pop_num` int DEFAULT NULL,
	`wave2_ikura_num` int DEFAULT NULL,
	`wave2_quota_num` int DEFAULT NULL,
	`wave3_water_level` varchar(45) DEFAULT NULL,
	`wave3_event_type` varchar(45) DEFAULT NULL,
	`wave3_golden_ikura_num` int DEFAULT NULL,
	`wave3_golden_ikura_pop_num` int DEFAULT NULL,
	`wave3_ikura_num` int DEFAULT NULL,
	`wave3_quota_num` int DEFAULT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`),
	UNIQUE KEY `user_id` (`user_id`,`splatnet_number`)
) ENGINE=InnoDB AUTO_INCREMENT=2111 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`job_id` int NOT NULL,
	`danger_rate` float NOT NULL,
	`job_result` bigint NOT NULL,
	`job_score` int NOT NULL,
	`job_rate` int NOT NULL,
	`grade_point` int NOT NULL,
	`grade_point_delta` int NOT NULL,
	`kuma_point` int NOT NULL,
	`start_time` bigint NOT NULL,
	`player_type` bigint NOT NULL,
	`play_time` bigint NOT NULL,
	`boss_counts` bigint NOT NULL,
	`end_time` bigint NOT NULL,
	`my_result` bigint NOT NULL,
	`grade` bigint NOT NULL,
	`schedule` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=2063 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_boss_counts` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`var3` bigint NOT NULL,
	`var6` bigint NOT NULL,
	`var9` bigint NOT NULL,
	`var12` bigint NOT NULL,
	`var13` bigint NOT NULL,
	`var14` bigint NOT NULL,
	`var15` bigint NOT NULL,
	`var16` bigint NOT NULL,
	`var21` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=10335 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_boss_counts_boss` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`boss` bigint NOT NULL,
	`count` int NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=93007 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_grade` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` varchar(45) NOT NULL,
	`short_name` varchar(45) NOT NULL,
	`long_name` varchar(45) NOT NULL,
	`name` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=2070 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_job_result` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`is_clear` tinyint NOT NULL,
	`failure_reason` varchar(45) DEFAULT NULL,
	`failure_wave` int DEFAULT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=2076 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_player` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`special` bigint NOT NULL,
	`pid` char(16) NOT NULL,
	`player_type` bigint NOT NULL,
	`name` varchar(10) NOT NULL,
	`dead_count` int NOT NULL,
	`golden_ikura_num` int NOT NULL,
	`boss_kill_counts` bigint NOT NULL,
	`ikura_num` int DEFAULT NULL,
	`help_count` int NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=8260 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_player_container` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`player` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=6187 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_player_weapon_list` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`id` varchar(45) NOT NULL,
	`weapon` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=20327 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_player_weapon_list_weapon` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` varchar(45) NOT NULL,
	`image` char(59) NOT NULL,
	`name` varchar(19) NOT NULL,
	`thumbnail` char(59) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=20327 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_schedule` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`start_time` bigint NOT NULL,
	`end_time` bigint NOT NULL,
	`stage` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=2069 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_schedule_stage` (
	`pk` int NOT NULL AUTO_INCREMENT,
	`image` char(63) NOT NULL,
	`name` varchar(20) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=2069 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_schedule_weapon` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`id` varchar(45) NOT NULL,
	`weapon` bigint DEFAULT NULL,
	`coop_special_weapon` bigint DEFAULT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=8252 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_schedule_weapon_special_weapon` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`image` char(65) NOT NULL,
	`name` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=1629 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_schedule_weapon_weapon` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` varchar(45) NOT NULL,
	`image` char(59) NOT NULL,
	`name` varchar(45) NOT NULL,
	`thumbnail` char(59) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=6625 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_splatnet_wave` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`water_level` bigint NOT NULL,
	`event_type` bigint NOT NULL,
	`golden_ikura_num` int NOT NULL,
	`golden_ikura_pop_num` int NOT NULL,
	`ikura_num` int NOT NULL,
	`quota_num` int NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=5104 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` int NOT NULL,
	`uuid` varchar(45) NOT NULL,
	`splatnet_number` int NOT NULL,
	`url` varchar(45) NOT NULL,
	`api_endpoint` varchar(45) NOT NULL,
	`user` bigint NOT NULL,
	`stage` bigint NOT NULL,
	`is_cleared` tinyint NOT NULL,
	`fail_reason` bigint DEFAULT NULL,
	`clear_waves` int NOT NULL,
	`danger_rate` varchar(45) NOT NULL,
	`title` bigint NOT NULL,
	`title_exp` int NOT NULL,
	`title_after` bigint NOT NULL,
	`title_exp_after` int NOT NULL,
	`my_data` bigint NOT NULL,
	`agent` bigint NOT NULL,
	`automated` tinyint NOT NULL,
	`shift_start_at` bigint NOT NULL,
	`start_at` bigint NOT NULL,
	`register_at` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_agent` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`name` varchar(45) NOT NULL,
	`version` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_boss_data` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`parent_table` varchar(64) NOT NULL,
	`boss` bigint NOT NULL,
	`count` int NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_boss_data_boss` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`splatnet` int NOT NULL,
	`splatnet_str` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_player` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`splatnet_id` char(16) NOT NULL,
	`name` varchar(10) NOT NULL,
	`special` bigint NOT NULL,
	`rescue` int NOT NULL,
	`death` int NOT NULL,
	`golden_egg_delivered` int NOT NULL,
	`power_egg_collected` int NOT NULL,
	`species` bigint NOT NULL,
	`gender` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_player_container` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`player` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_title` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`splatnet` int NOT NULL,
	`generic_name` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_triple` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`key` varchar(45) NOT NULL,
	`name` bigint NOT NULL,
	`splatnet` varchar(45) NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_triple_container` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`parent` bigint NOT NULL,
	`triple` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_user` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`id` int NOT NULL,
	`name` varchar(45) NOT NULL,
	`screen_name` varchar(45) NOT NULL,
	`url` varchar(45) NOT NULL,
	`salmon_url` varchar(45) NOT NULL,
	`battle_url` varchar(45) NOT NULL,
	`join_at` bigint NOT NULL,
	`profile` bigint NOT NULL,
	`stats` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `two_salmon_shift_stat_ink_user_stats` (
	`pk` bigint NOT NULL AUTO_INCREMENT,
	`work_count` int NOT NULL,
	`total_golden_eggs` int NOT NULL,
	`total_eggs` int NOT NULL,
	`total_rescued` int NOT NULL,
	`total_point` int NOT NULL,
	`as_of` bigint NOT NULL,
	`registered_at` bigint NOT NULL,
	PRIMARY KEY (`pk`),
	UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
