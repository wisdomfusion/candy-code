DROP DATABASE IF EXISTS candydb;

CREATE DATABASE candydb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE candydb;

DROP TABLE IF EXISTS candies;

CREATE TABLE candies
(
    id         mediumint    NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title      varchar(200) NOT NULL,
    candy      text         NOT NULL,
    created_at datetime     NOT NULL,
    updated_at datetime     NOT NULL,
    expired_at datetime     NULL,
    deleted_at datetime     NULL
);

INSERT INTO candies (title, candy, created_at, updated_at, expired_at)
VALUES ('golang interface', '// golang interface code sample', UTC_TIMESTAMP(), UTC_TIMESTAMP(), NULL),
       ('code2', '// code2\n// code here', UTC_TIMESTAMP(), UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP, INTERVAL 365 DAY)),
       ('code3', '// code3\n// code here', UTC_TIMESTAMP(), UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP, INTERVAL 7 DAY)),
       ('代码4', '// 代码4\n// code here', UTC_TIMESTAMP(), UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP, INTERVAL 7 DAY));
