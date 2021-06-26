CREATE TABLE `posts` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `user_id` INT(11) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `post_details` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `post_id` BIGINT UNSIGNED NOT NULL,
  `description` TEXT NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `index_post_details_on_post_id` (`post_id`),
  CONSTRAINT `index_post_details_on_post_id` FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

INSERT INTO posts (text, user_id) VALUES ("this is text 1 by user 1", 1);
INSERT INTO posts (text, user_id) VALUES ("this is text 2 by user 1", 1);
INSERT INTO posts (text, user_id) VALUES ("this is text 3 by user 2", 2);

INSERT INTO post_details (description, post_id) VALUES ("description!!", 1);
INSERT INTO post_details (description, post_id) VALUES ("description!!", 1);