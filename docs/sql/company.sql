CREATE DATABASE IF NOT EXISTS lankaiyun
    DEFAULT CHARACTER SET utf8mb4
    DEFAULT COLLATE utf8mb4_general_ci;

use lankaiyun;

CREATE TABLE `subscription` (
                                `id` int AUTO_INCREMENT,
                                `email` varchar(100) UNIQUE,
                                `time` date,
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

drop table subscription;