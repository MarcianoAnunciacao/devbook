CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS followers;

CREATE TABLE users(
    id int auto_increment PRIMARY KEY,
    name varchar(50) NOT NULL,
    nick_name varchar(50) NOT NULL unique,
    email varchar(50) NOT NULL unique,
    password longtext NOT NULL,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers(
    user_id int NOT NULL,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id INT NOT NULL,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    PRIMARY KEY(user_id, follower_id)
) ENGINE=INNODB;