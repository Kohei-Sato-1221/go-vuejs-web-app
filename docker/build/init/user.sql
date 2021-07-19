-- noinspection SqlDialectInspectionForFile

-- noinspection SqlNoDataSourceInspectionForFile

UPDATE mysql.user SET plugin = 'mysql_native_password';

CREATE TABLE `users` (
     `id` int(50) unsigned NOT NULL AUTO_INCREMENT,
     `name` varchar(100) DEFAULT NULL,
     `email` varchar(100) DEFAULT NULL,
     `gender` varchar(30) DEFAULT NULL,
     `createAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
     `updateAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

INSERT INTO `users` (`name`, `email`, `gender`) VALUES ('Sugar', 'sugar@hogehoge.com', 'Male');
INSERT INTO `users` (`name`, `email`, `gender`) VALUES ('Salt', 'salt@hogehoge.com', 'Female');