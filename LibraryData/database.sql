CREATE DATABASE IF NOT EXISTS `library` CHARACTER SET = 'utf8mb4' COLLATE = 'utf8mb4_general_ci';
USE `library`;

DROP TABLE if exists `books`;
DROP TABLE if exists `authors`;

CREATE TABLE `authors`
(
    `id`   INT(11)      NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL
);

CREATE TABLE `books`
(
    `id`        INT(11)      NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `author_id` INT(11)      NOT NULL,
    `title`     VARCHAR(255) NOT NULL,
    FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`)
);