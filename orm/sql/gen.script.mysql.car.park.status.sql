
USE `db_carpark`;
CREATE TABLE `carpark_status` (
	`car_park_no` VARCHAR(16) NOT NULL DEFAULT '',
	`total_lots` INT(11) NOT NULL DEFAULT '0',
	`available_lots` INT(11) NOT NULL DEFAULT '0',
	`report_at` BIGINT(20) NOT NULL DEFAULT '0',
	`created_at` BIGINT(20) NOT NULL DEFAULT '0',
	PRIMARY KEY(`car_park_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT 'carpark_status';
CREATE INDEX `available_lots_of_car_park_status_idx` ON `carpark_status`(`available_lots`);

