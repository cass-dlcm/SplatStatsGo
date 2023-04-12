create table three_salmon_shift
(
    userId                   bigint        not null,
    id                       text          not null,
    typename                 text          not null,
    afterGrade               text          not null,
    playerByname             text          not null,
    playerBackground         text          not null,
    playerName               text          not null,
    playerNameId             text          not null,
    playerUniform            text          not null,
    playerId                 text          not null,
    playerSpecies            text          not null,
    playerSpecialWeapon      int           not null,
    playerDefeatEnemyCount   int           not null,
    playerDeliverCount       int           not null,
    playerGoldenAssistCount  int           not null,
    playerGoldenDeliverCount int           not null,
    playerRescueCount        int           not null,
    playerRescuedCount       int           not null,
    resultWave               int           not null,
    playedTime               timestamp     not null,
    rule                     text          not null,
    stage                    text          not null,
    dangerRate               decimal(4, 1) not null,
    scenarioCode             text,
    smellMeter               int           not null,
    afterGradePoint          int           not null,
    scaleBronze              int,
    scaleSilver              int,
    scaleGold                int,
    jobPoint                 int           not null,
    jobScore                 int           not null,
    jobRate                  decimal(3, 2) not null,
    jobBonus                 int           not null,
    hasDefeatBoss            boolean,
    boss                     text,
    weapon0                  text not null,
    weapon1                  text not null,
    weapon2                  text not null,
    weapon3                  text not null,
    primary key (userId, id),
    constraint three_salmon_shift_user foreign key (userId) references auth_user (pk) on delete cascade
);

create table three_salmon_user_weapon
(
    userId  bigint not null,
    shiftId text   not null,
    wave    int    not null,
    weapon  text   not null,
    primary key (userId, shiftId, wave),
    constraint three_salmon_user_weapon_shift foreign key (userId, shiftId) references three_salmon_shift (userId, id)
);

create table three_salmon_user_badge
(
    userId    bigint not null,
    shiftId   text   not null,
    badgeSlot int    not null,
    badge     text   not null,
    primary key (userId, shiftId, badgeSlot),
    constraint three_salmon_user_badge_shift foreign key (userId, shiftId) references three_salmon_shift (userId, id)
);

create table three_salmon_wave
(
    userId           bigint not null,
    shiftId          text   not null,
    waveNumber       int    not null,
    waterLevel       int    not null,
    eventWave        text,
    deliverNorm      int    not null,
    goldenPopCount   int    not null,
    teamDeliverCount int    not null,
    primary key (userId, shiftId, waveNumber),
    constraint three_salmon_wave_shift foreign key (userId, shiftId) references three_salmon_shift (userId, id)
);

create table three_salmon_wave_special
(
    pk         bigserial not null,
    userId     bigint    not null,
    shiftId    text      not null,
    waveNumber int       not null,
    special    text      not null,
    primary key (pk),
    constraint three_salmon_wave_special_wave foreign key (userId, shiftId, waveNumber) references three_salmon_wave (userId, shiftId, waveNumber)
);

create table three_salmon_enemy_result
(
    userId          bigint not null,
    shiftId         text   not null,
    defeatCount     int    not null,
    teamDefeatCount int    not null,
    popCount        int    not null,
    enemy           text   not null,
    primary key (userId, shiftId, enemy),
    constraint three_salmon_enemy_shift foreign key (userId, shiftId) references three_salmon_shift (userId, id)
);

create table three_salmon_player
(
    userId             bigint not null,
    shiftId            text   not null,
    byname             text   not null,
    name               text   not null,
    nameId             text   not null,
    background         text   not null,
    uniform            text   not null,
    id                 text   not null,
    species            text   not null,
    special            int    not null,
    defeatEnemyCount   int    not null,
    deliverCount       int    not null,
    goldenAssistCount  int    not null,
    goldenDeliverCount int    not null,
    rescueCount        int    not null,
    rescuedCount       int    not null,
    primary key (userId, shiftId, id),
    constraint three_salmon_player_shift foreign key (userId, shiftId) references three_salmon_shift (userId, id)
);

create table three_salmon_player_badge
(
    userId    bigint not null,
    shiftId   text   not null,
    playerId  text   not null,
    badgeSlot int    not null,
    badge     text   not null,
    primary key (userId, shiftId, playerId, badgeSlot),
    constraint three_salmon_player_badge_player foreign key (userId, shiftId, playerId) references three_salmon_player (userId, shiftID, id)
);

create table three_salmon_player_weapon
(
    userId   bigint not null,
    shiftId  text   not null,
    playerId text   not null,
    wave     int    not null,
    weapon   text   not null,
    primary key (userId, shiftId, playerId, wave),
    constraint three_salmon_player_weapon_player foreign key (userId, shiftId, playerId) references three_salmon_player (userId, shiftID, id)
)