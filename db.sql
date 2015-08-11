create table user(
id int auto_increment primary key,
name char(12) not null unique,
passwd char(64) not null
)default charset=utf8;

create table topic(
id int auto_increment primary key,
userid int not null,
title char(100) not null,
content text
)default charset=utf8;

create table remark(
id int auto_increment primary key,
userid int not null,
topicid int not null,
content text
)default charset=utf8;