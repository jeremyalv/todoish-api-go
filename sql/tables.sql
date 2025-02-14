CREATE DATABASE IF NOT EXISTS `todo_app`;

-- Use single line commands in MySQL if the multilines fail
CREATE TABLE IF NOT EXISTS todo_app.users (
    `id` BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)), -- UUIDv1
    `created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) UNIQUE NOT NULL,
    `hashed_password` VARCHAR NOT NULL,

    CONSTRAINT users_PK PRIMARY KEY (`id`)
);

CREATE TABLE IF NOT EXISTS todo_app.todos (
    `id` BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)), -- UUIDv1
    `owner_id` BINARY(16)
    `created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    `title` VARCHAR(255) NOT NULL,
    `description` VARCHAR(2000),
    `is_completed` BOOLEAN NOT NULL,
    `due_date` DATETIME, -- UTC TIME

    CONSTRAINT todos_PK PRIMARY KEY (`id`)
    CONSTRAINT todos_FK FOREIGN KEY (`owner_id`) REFERENCES todo_app.users(`id`)
);


INSERT INTO `todo_app`.`users` (`id`,`name`,`email`,`hashed_password`) VALUES (UUID_TO_BIN('a7e3dee2-ea9f-11ef-9cd2-0242ac120002', 1),'username','username@email.com','password');

-- Dummy Data
-- +------------------------------------+---------------------+---------------------+----------+--------------------+-----------------+
-- | id                                 | created             | updated             | name     | email              | hashed_password |
-- +------------------------------------+---------------------+---------------------+----------+--------------------+-----------------+
-- | 0x11EFEA9FA7E3DEE29CD20242AC120002 | 2025-02-14 08:11:21 | 2025-02-14 08:11:21 | username | username@email.com | password        |
-- +------------------------------------+---------------------+---------------------+----------+--------------------+-----------------+
