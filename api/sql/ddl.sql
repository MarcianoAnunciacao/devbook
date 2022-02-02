CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment PRIMARY KEY,
    name varchar(50) NOT NULL,
    nick_name varchar(50) NOT NULL unique,
    email varchar(50) NOT NULL unique,
    password longtext NOT NULL,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;