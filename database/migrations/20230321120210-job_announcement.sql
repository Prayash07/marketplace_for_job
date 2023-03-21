
-- +migrate Up
CREATE TABLE IF NOT EXISTS `JobAnnouncement`(
    `id` INT NOT NULL UNIQUE AUTO_INCREMENT,
    `degree_name` VARCHAR(255) NOT NULL,
    `years` INT NOT NULL,
    `user_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `User`(`ID`)
);
-- +migrate Down
drop table `JobAnnouncement`;