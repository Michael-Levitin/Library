CREATE DATABASE IF NOT EXISTS `library` CHARACTER SET = 'utf8mb4' COLLATE = 'utf8mb4_general_ci';
USE `library`;

DROP TABLE if exists `books_authors`;
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

CREATE TABLE `books_authors`
(
    `book_id`   INT(11) NOT NULL,
    `author_id` INT(11) NOT NULL,
    PRIMARY KEY (`book_id`, `author_id`),
    FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON DELETE CASCADE,
    FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`) ON DELETE CASCADE
);
