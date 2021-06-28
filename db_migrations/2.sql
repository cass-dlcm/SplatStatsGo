
CREATE TABLE `two_salmon_shift_stat_ink_triple_string` (
    `pk` bigint NOT NULL AUTO_INCREMENT,
    `key` varchar(45) NOT NULL,
    `name` bigint NOT NULL,
    `splatnet` varchar(63) NOT NULL,
    PRIMARY KEY (`pk`),
    UNIQUE KEY `pk_UNIQUE` (`pk`)
) ENGINE=InnoDB AUTO_INCREMENT=8346 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
