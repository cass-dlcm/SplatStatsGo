create table three_battle
(
    userId                 bigint           not null,
    id                     text             not null,
    typename               text             not null,
    vsRule                 text             not null,
    vsMode                 text             not null,
    playerByname           text             not null,
    playerName             text             not null,
    playerNameId           text             not null,
    playerBackground       text             not null,
    playerId               text             not null,
    playerHeadgearName     text             not null,
    playerHeadgearPrimary  text             not null,
    playerClothingName     text             not null,
    playerClothingPrimary  text             not null,
    playerShoesName        text             not null,
    playerShoesPrimary     text             not null,
    playerPaint            int              not null,
    playerkill             int              not null,
    playerdeath            int              not null,
    playerassist           int              not null,
    playerspecial          int              not null,
    playercrown            boolean          not null,
    judgement              text             not null,
    myTeamColorA           double precision not null,
    myTeamColorB           double precision not null,
    myTeamColorG           double precision not null,
    myTeamColorR           double precision not null,
    myTeamResultPaintRatio double precision,
    myTeamScore            int,
    myTeamJudgement        text             not null,
    myTeamOrder            int              not null,
    knockout               text             not null,
    duration               int              not null,
    playedTime             timestamp        not null,
    primary key (userId, id)
);

create table three_battle_user_badge
(
    userId    bigint not null,
    battleId  text   not null,
    badgeSlot int    not null,
    badge     text   not null,
    primary key (userId, battleId, badgeSlot),
    constraint three_battle_user_badge_battle foreign key (userId, battleId) references three_battle (userId, id)
);

create table three_battle_my_player
(
    userId              bigint  not null,
    battleId            text    not null,
    id                  text    not null,
    byname              text    not null,
    playerName          text    not null,
    weapon              text    not null,
    species             text    not null,
    isPlayer            text    not null,
    nameId              text    not null,
    nameplateBackground text    not null,
    headgearName        text    not null,
    headgearMain        text    not null,
    clothingName        text    not null,
    clothingMain        text    not null,
    shoesName           text    not null,
    shoesMain           text    not null,
    paint               int     not null,
    kill                int     not null,
    death               int     not null,
    assist              int     not null,
    special             int     not null,
    crown               boolean not null,
    festDragonCert      text    not null,
    primary key (userId, battleId, id),
    constraint three_battle_my_player_battle foreign key (userId, battleId) references three_battle (userId, id)
);

create table three_battle_my_player_badge
(
    userId    bigint not null,
    battleId  text   not null,
    playerId  text   not null,
    badgeSlot int    not null,
    badge     text   not null,
    primary key (userId, battleId, playerId, badgeSlot),
    constraint three_battle_my_badge_player foreign key (userId, battleId, playerId) references three_battle_my_player (userId, battleId, id)
);

create table three_battle_other_team
(
    userId     bigint           not null,
    battleId   text             not null,
    teamIndex  int              not null,
    colorA     double precision not null,
    colorB     double precision not null,
    colorG     double precision not null,
    colorR     double precision not null,
    paintRatio double precision,
    score      int,
    judgement  text             not null,
    teamOrder  int              not null,
    primary key (userId, battleId, teamIndex),
    constraint three_battle_other_team_battle foreign key (userId, battleId) references three_battle (userId, id)
);

create table three_battle_other_team_player
(
    userId              bigint  not null,
    battleId            text    not null,
    teamIndex           int     not null,
    id                  text    not null,
    byname              text    not null,
    playerName          text    not null,
    weapon              text    not null,
    species             text    not null,
    isPlayer            text    not null,
    nameId              text    not null,
    nameplateBackground text    not null,
    headgearName        text    not null,
    headgearMain        text    not null,
    clothingName        text    not null,
    clothingMain        text    not null,
    shoesName           text    not null,
    shoesGear           text    not null,
    paint               int     not null,
    kill                int     not null,
    death               int     not null,
    assist              int     not null,
    special             int     not null,
    crown               boolean not null,
    festDragonCert      text    not null,
    primary key (userId, battleId, id),
    constraint three_battle_other_player_team foreign key (userId, battleId, teamIndex) references three_battle_other_team (userId, battleId, teamIndex)
);

create table three_battle_other_player_badge
(
    userId    bigint not null,
    battleId  text   not null,
    playerId  text   not null,
    badgeSlot int    not null,
    badge     text   not null,
    primary key (userId, battleId, playerId, badgeSlot),
    constraint three_battle_other_badge_player foreign key (userId, battleId, playerId) references three_battle_other_team_player (userId, battleId, id)
);

create table three_battle_award (
    userId bigint not null,
    battleId text not null,
    awardName text not null,
    primary key (userId, battleId, awardName),
    constraint three_battle_award_battle foreign key (userId, battleId) references three_battle (userId, id)
);