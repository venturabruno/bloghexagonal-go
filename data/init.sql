create database if not exists blog;
use blog;
create table if not exists post (id varchar(50) NOT NULL, title varchar(255) NOT NULL, subtitle varchar(255) NOT NULL, status varchar(50) NOT NULL, content text, created_at datetime NOT NULL, published_at datetime, primary key (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;