-- +migrate Up
CREATE TABLE IF NOT EXISTS `Company`(
    `id` INT NOT NULL UNIQUE AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `no_of_employees` INT NOT NULL,
    PRIMARY KEY (`id`)
);
-- +migrate Down
drop table `Company`;