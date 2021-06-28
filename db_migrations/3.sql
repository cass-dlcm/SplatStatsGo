CREATE TABLE `two_salmon_shift_stat_ink_wave` (
    `pk` bigint NOT NULL AUTO_INCREMENT,
    `parent` bigint NOT NULL,
    `known_occurrence` bigint DEFAULT NULL,
    `water_level` bigint DEFAULT NULL,
    `golden_egg_quota` int NOT NULL,
    `golden_egg_appearances` int NOT NULL,
    `golden_egg_delivered` int NOT NULL,
    `power_egg_collected` int NOT NULL,
    PRIMARY KEY (`pk`),
    UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=5063 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
