
-- +migrate Up
CREATE TABLE
    IF NOT EXISTS `User`(
       `ID` INT NOT NULL UNIQUE AUTO_INCREMENT,
       `Name` VARCHAR (127) NOT NULL UNIQUE,
       `Address` VARCHAR (127) NOT NULL,
       PRIMARY KEY (`ID`)
    );
-- +migrate Down
drop table `User`;
