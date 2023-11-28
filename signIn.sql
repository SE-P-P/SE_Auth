use activityDiary;
create table signIn(
    `id` INT(8) not null auto_increment,
    `email` VARCHAR(30) NOT NULL COMMENT '邮箱',
    `password` VARCHAR(20) NOT NULL COMMENT '密码',
    PRIMARY KEY(`id`)
)