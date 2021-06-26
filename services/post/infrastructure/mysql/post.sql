CREATE TABLE `posts` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `user_id` INT(11) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

INSERT INTO posts (text, user_id) VALUES ("this is text 1 by user 1", 1);
INSERT INTO posts (text, user_id) VALUES ("this is text 2 by user 1", 1);
INSERT INTO posts (text, user_id) VALUES ("this is text 3 by user 2", 2);