CREATE TABLE auth_user
(
    pk             bigserial NOT NULL,
    username       text      NOT NULL unique,
    password       text      NOT NULL,
    email          text      NOT NULL unique,
    email_verified boolean   NOT NULL,
    PRIMARY KEY (pk)
);

CREATE TABLE auth_user_session_token
(
    pk            bigserial NOT NULL,
    parent        bigint NOT NULL,
    session_token text NOT NULL unique,
    PRIMARY KEY (pk)
);
