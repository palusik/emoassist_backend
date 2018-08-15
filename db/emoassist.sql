-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema emodb
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Table `actions`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `actions` ;

CREATE TABLE IF NOT EXISTS `actions` (
  `idactions` INT(11) NOT NULL AUTO_INCREMENT,
  `userid` VARCHAR(45) NULL DEFAULT NULL,
  `haction` VARCHAR(45) NULL DEFAULT NULL,
  `naction` VARCHAR(45) NULL DEFAULT NULL,
  `aaction` VARCHAR(45) NULL DEFAULT NULL,
  `saction` VARCHAR(45) NULL DEFAULT NULL,
  `faction` VARCHAR(45) NULL DEFAULT NULL,
  `activity` VARCHAR(45) NULL DEFAULT NULL,
  `book` VARCHAR(45) NULL DEFAULT NULL,
  `music` VARCHAR(45) NULL DEFAULT NULL,
  `call` VARCHAR(45) NULL DEFAULT NULL,
  `picture` VARCHAR(45) NULL DEFAULT NULL,
  `updated` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`idactions`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = latin1;

CREATE UNIQUE INDEX `userid_UNIQUE` ON `actions` (`userid` ASC);


-- -----------------------------------------------------
-- Table `alerts`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `alerts` ;

CREATE TABLE IF NOT EXISTS `alerts` (
  `idalerts` INT(11) NOT NULL AUTO_INCREMENT,
  `created` VARCHAR(45) NULL DEFAULT NULL,
  `userid` VARCHAR(45) NULL DEFAULT NULL,
  `type` VARCHAR(45) NULL DEFAULT NULL,
  `probability` FLOAT NULL DEFAULT NULL,
  `hr` INT(11) NULL DEFAULT NULL,
  `alert_type` VARCHAR(45) NULL DEFAULT NULL,
  `target` VARCHAR(45) NULL DEFAULT NULL,
  `status` VARCHAR(45) NULL DEFAULT 'NEW',
  `data` VARCHAR(500) NULL DEFAULT NULL,
  PRIMARY KEY (`idalerts`))
ENGINE = InnoDB
AUTO_INCREMENT = 133
DEFAULT CHARACTER SET = latin1;


-- -----------------------------------------------------
-- Table `preferences`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `preferences` ;

CREATE TABLE IF NOT EXISTS `preferences` (
  `idpreferences` INT(11) NOT NULL AUTO_INCREMENT,
  `userid` VARCHAR(45) NULL DEFAULT NULL,
  `hprobability` INT(11) NULL DEFAULT NULL,
  `nprobability` INT(11) NULL DEFAULT NULL,
  `aprobability` INT(11) NULL DEFAULT NULL,
  `sprobability` INT(11) NULL DEFAULT NULL,
  `fprobability` INT(11) NULL DEFAULT NULL,
  `hhrmin` INT(11) NULL DEFAULT NULL,
  `nhrmin` INT(11) NULL DEFAULT NULL,
  `ahrmin` INT(11) NULL DEFAULT NULL,
  `shrmin` INT(11) NULL DEFAULT NULL,
  `fhrmin` INT(11) NULL DEFAULT NULL,
  `hhrmax` INT(11) NULL DEFAULT NULL,
  `nhrmax` INT(11) NULL DEFAULT NULL,
  `ahrmax` INT(11) NULL DEFAULT NULL,
  `shrmax` INT(11) NULL DEFAULT NULL,
  `fhrmax` INT(11) NULL DEFAULT NULL,
  `remote` VARCHAR(1) NULL DEFAULT NULL,
  `age` INT(11) NULL DEFAULT NULL,
  `updated` DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (`idpreferences`))
ENGINE = InnoDB
DEFAULT CHARACTER SET = latin1;

CREATE UNIQUE INDEX `userid_UNIQUE` ON `preferences` (`userid` ASC);


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
