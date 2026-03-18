DROP TABLE IF EXISTS users;

CREATE TABLE users
(
    id         binary(16) default (uuid_to_bin(uuid())) NOT NULL,
    first_name varchar(64),
    last_name  varchar(64),
    nickname   varchar(32),
    password   varchar(255),
    email      varchar(64),
    country    varchar(64),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
    PRIMARY KEY (id)
);

INSERT INTO users(first_name,
                  last_name,
                  nickname,
                  password,
                  email,
                  country)
VALUES ("FirstName1",
        "LastName1",
        "Nickname1",
        "Password1",
        "email1@email.com",
        "Country1");

INSERT INTO users(first_name,
                  last_name,
                  nickname,
                  password,
                  email,
                  country)
VALUES ("FirstName2",
        "LastName2",
        "Nickname2",
        "Password2",
        "email2@email.com",
        "Country2");

INSERT INTO users(first_name,
                  last_name,
                  nickname,
                  password,
                  email,
                  country)
VALUES ("FirstName3",
        "LastName3",
        "Nickname3",
        "Password3",
        "email3@email.com",
        "Country3");

INSERT INTO users(first_name,
                  last_name,
                  nickname,
                  password,
                  email,
                  country)
VALUES ("FirstName4",
        "LastName4",
        "Nickname4",
        "Password4",
        "email4@email.com",
        "Country4");

INSERT INTO users(first_name,
                  last_name,
                  nickname,
                  password,
                  email,
                  country)
VALUES ("FirstName5",
        "LastName5",
        "Nickname5",
        "Password5",
        "email5@email.com",
        "Country5");

INSERT INTO users(first_name,
                  last_name,
                  nickname,
                  password,
                  email,
                  country)
VALUES ("FirstName6",
        "LastName6",
        "Nickname6",
        "Password6",
        "email6@email.com",
        "Country6");

INSERT INTO users(first_name,
                  last_name,
                  nickname,
                  password,
                  email,
                  country)
VALUES ("FirstName7",
        "LastName7",
        "Nickname7",
        "Password7",
        "email7@email.com",
        "Country7");