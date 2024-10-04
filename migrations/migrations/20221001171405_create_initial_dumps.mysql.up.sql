
DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
  `id` CHAR(36) NOT NULL,
  `title` VARCHAR(100) NOT NULL,
  `author` VARCHAR(50) NOT NULL,
  `publisher` VARCHAR(50) NOT NULL,
  `isbn` VARCHAR(13) NOT NULL,
  `year` INT NOT NULL,
  `edition` INT,
  `rating` TINYINT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` DATETIME DEFAULT NULL,
  CONSTRAINT rating_1_to_10 CHECK (rating >= 1 AND rating <= 10),
  PRIMARY KEY (`id`),
  UNIQUE KEY `isbn` (`isbn`), 
  UNIQUE KEY `title` (`title`)
);

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` CHAR(36) NOT NULL,
  `username` VARCHAR(50) NOT NULL,
  `email` VARCHAR(100) NOT NULL,
  `password` VARCHAR(255) NOT NULL,  -- Hash the password!
  `role` ENUM('admin', 'user') NOT NULL DEFAULT 'user',  -- Define roles
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
);

DROP TABLE IF EXISTS `ratings`;

CREATE TABLE `ratings` (
  `id` CHAR(36) NOT NULL,
  `book_id` CHAR(36) NOT NULL,
  `user_id` CHAR(36) NOT NULL,
  `rating` TINYINT NOT NULL CHECK (rating >= 1 AND rating <= 10),
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`book_id`) REFERENCES `books`(`id`) ON DELETE CASCADE,
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
);
