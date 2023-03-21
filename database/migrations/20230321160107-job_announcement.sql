-- +migrate Up
CREATE TABLE IF NOT EXISTS `JobAnnouncement`(
    `id` INT NOT NULL UNIQUE AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `description` VARCHAR(255) NOT NULL,
    `url` VARCHAR(255) NOT NULL,
    `company_id` INT NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`company_id`) REFERENCES `Company`(`id`)
);
-- +migrate Down
drop table `JobAnnouncement`;
