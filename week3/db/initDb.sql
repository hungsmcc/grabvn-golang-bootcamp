DROP DATABASE IF EXISTS `feedback`;
CREATE DATABASE `feedback` CHARACTER SET utf8 COLLATE utf8_bin;
CREATE USER 'user'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON feedback.* TO 'user'@'%';
FLUSH PRIVILEGES;