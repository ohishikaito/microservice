CREATE TABLE `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `last_name` VARCHAR(255) NOT NULL,
  `first_name` VARCHAR(255) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `telephone_number` varchar(255) NOT NULL,
  `gender` int NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
)ENGINE=INNODB DEFAULT CHARSET=utf8mb4;

INSERT INTO users (last_name, first_name, email, telephone_number, gender) VALUES ("ohishi", "kaito", "test1@gmail.com", "09000007777", 0);
INSERT INTO users (last_name, first_name, email, telephone_number, gender) VALUES ("ohishi", "kaito", "test2@gmail.com", "09000007777", 0);
INSERT INTO users (last_name, first_name, email, telephone_number, gender) VALUES ("ohishi", "kaito", "test3@gmail.com", "09000007777", 1);
INSERT INTO users (last_name, first_name, email, telephone_number, gender) VALUES ("ohishi", "kaito", "test4@gmail.com", "09000007777", 1);
INSERT INTO users (last_name, first_name, email, telephone_number, gender) VALUES ("ohishi", "kaito", "test5@gmail.com", "09000007777", 1);