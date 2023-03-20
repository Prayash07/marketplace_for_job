
-- +migrate Up
CREATE TABLE IF NOT EXISTS `Education`(
  `ID` INT NOT NULL UNIQUE AUTO_INCREMENT,
  `DegreeName` VARCHAR(255) NOT NULL,
  `Years` INT NOT NULL,
  `UserId` INT NOT NULL,
  PRIMARY KEY (`ID`),
  FOREIGN KEY (`UserId`) REFERENCES `User`(`ID`)
);
-- +migrate Down
drop table `Education`;
