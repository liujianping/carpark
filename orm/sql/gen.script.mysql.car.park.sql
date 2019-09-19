
USE `db_carpark`;
CREATE TABLE `carparks` (
	`car_park_no` VARCHAR(16) NOT NULL DEFAULT '',
	`address` VARCHAR(256) NOT NULL DEFAULT '',
	`latitude` FLOAT(10, 6) NOT NULL DEFAULT '0',
	`longitude` FLOAT(10, 6) NOT NULL DEFAULT '0',
	`car_park_type` VARCHAR(32) NOT NULL DEFAULT '',
	`type_of_parking_system` VARCHAR(32) NOT NULL DEFAULT '',
	`short_term_parking` VARCHAR(32) NOT NULL DEFAULT '',
	`short_term_parking_from` BIGINT(20) NOT NULL DEFAULT '0' COMMENT 'short term parking start second',
	`short_term_parking_to` BIGINT(20) NOT NULL DEFAULT '0' COMMENT 'short term parking end second',
	`free_parking` VARCHAR(32) NOT NULL DEFAULT '',
	`night_parking` TINYINT(1) NOT NULL DEFAULT '0',
	`car_park_decks` INT(11) NOT NULL DEFAULT '0',
	`gantry_height` FLOAT NOT NULL DEFAULT '0',
	`car_park_basement` TINYINT(1) NOT NULL DEFAULT '0',
	PRIMARY KEY(`car_park_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'carparks';
CREATE INDEX `short_term_parking_from_short_term_parking_to_of_car_park_idx` ON `carparks`(`short_term_parking_from`,`short_term_parking_to`);

