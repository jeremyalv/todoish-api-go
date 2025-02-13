CREATE DATABASE IF NOT EXISTS `todo_app`;

CREATE TABLE IF NOT EXISTS todo_app.users (
    `id` BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY, -- UUIDv1
    `created_at` DATETIME NOT NULL, -- UTC TIME

    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) UNIQUE NOT NULL,
    `hashed_password` VARCHAR NOT NULL,
);

CREATE TABLE IF NOT EXISTS todo_app.todos (
    `id` BINARY(16) DEFAULT (UUID_TO_BIN(UUID(), 1)) PRIMARY KEY, -- UUIDv1
    `created_at` DATETIME NOT NULL, -- UTC TIME

    `title` VARCHAR(255) NOT NULL,
    `description` VARCHAR(2000),
    `is_completed` BOOLEAN NOT NULL,
    `due_date` DATETIME, -- UTC TIME
);