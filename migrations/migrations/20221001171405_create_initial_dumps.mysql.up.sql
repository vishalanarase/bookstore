
DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
  `id` VARCHAR(40) NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `authorname` VARCHAR(30) NOT NULL,
  `rating` INT NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  CONSTRAINT rating_1_to_10 CHECK (rating IS NULL OR
    (`rating` >= 1 and `rating` <= 10)),
  PRIMARY KEY (`id`), 
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `name` (`name`)
);