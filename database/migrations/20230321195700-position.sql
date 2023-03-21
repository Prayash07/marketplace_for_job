-- +migrate Up
CREATE TABLE IF NOT EXISTS `Position`(
    `id` INT NOT NULL UNIQUE AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `experience` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);
-- +migrate Down
drop table `Position`;