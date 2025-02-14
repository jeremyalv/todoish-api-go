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