DROP DATABASE IF EXISTS `user`;
CREATE DATABASE IF NOT EXISTS `user`;

USE `user`;

CREATE TABLE `user` (
    `user` BIGINT(11) auto_increment,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(255) NOT NULL,
    `active` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`user`)
);

INSERT INTO `user`(`name`, `email`, `active`) VALUES("Igor Silva", "igor.desenvolvedor.full@gmail.com", 1);

-- CREATE TABLE `user_strategy` (
--     `user_strategy` BIGINT(11) auto_increment,
--     `user` BIGINT(11) NOT NULL,
--     `trade_config` BIGINT(11) NOT NULL,
--     PRIMARY KEY(`user_strategy`),
--     FOREIGN KEY(`user`) REFERENCES `user`(`user`)
-- );