
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
