# 建表脚本
drop database if exists partner_matching;
create database partner_matching default character set = 'utf8';;
use partner_matching;

DROP TABLE IF EXISTS tag;
create table if not exists tag
(
    id         bigint auto_increment comment 'id' primary key,
    tagName    varchar(256)                       null comment '标签名称',
    userId     bigint                             null comment '用户id',
    parentId   bigint                             null comment '父标签 id',
    isParent   tinyint                            null comment '0-不是父标签 1-父标签',
    createTime datetime default CURRENT_TIMESTAMP null comment '创建时间',
    updateTime datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    isDelete   tinyint  default 0                 not null comment '是否删除',
    index idx_userId (userId),
    unique tag_tagName_uindex (tagName)
)
    comment '标签';

DROP TABLE IF EXISTS user;
create table if not exists user
(
    id           bigint auto_increment comment 'id' primary key,
    username     varchar(256)                       null comment '用户昵称',
    userAccount  varchar(256)                       null comment '账号',
    avatarUrl    varchar(1024)                      null comment '用户头像',
    gender       tinyint                            null comment '性别',
    userPassword varchar(512)                       not null comment '密码',
    phone        varchar(128)                       null comment '电话',
    email        varchar(512)                       null comment '邮箱',
    userStatus   int      default 0                 not null comment '用户状态 0-正常',
    createTime   datetime default CURRENT_TIMESTAMP null comment '创建时间',
    updateTime   datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    isDelete     tinyint  default 0                 not null comment '是否删除',
    role         int      default 0                 not null comment '用户角色 0-普通用户 1-管理员',
    tags         varchar(1024)                      null comment '标签列表, json格式',
    profile      varchar(512)                       null comment '个人简介'
)
    comment '用户';

insert into user value (null, '绍幻', 'Carey',
                        'http://cdn.u2.huluxia.com/g3/M00/36/56/wKgBOVwPmcmAB2cnAACcXKrjLlw989.jpg',
                        0, '25d55ad283aa400af464c76d713c07ad', null, null, 0, '2023-05-01 13:56:40', null, 0, 1,
                        '{"Java":1,"C++":1,"Python":1}',
                        '未来一定会很好，即使现在有诸多的不幸。相信糟糕的日子熬过去了，剩下的就是好运气。');

insert into user value (null, '顾平', 'Eckard',
                        'http://cdn.u2.huluxia.com/g3/M02/36/6D/wKgBOVwPoDmADKeoAADArH7B8qc650.jpg',
                        0, '25d55ad283aa400af464c76d713c07ad', null, null, 0, '2023-05-01 13:56:40', null, 0, 1,
                        '{"Golang":1,"Docker":1,"Rust":1}',
                        '所有的美好，都不负归期，选一种姿态让自己活的无可代替，没有所谓的运气，只有绝对的努力。');

insert into user value (null, '戴玉', 'Wally',
                        'http://cdn.u2.huluxia.com/g3/M01/36/70/wKgBOVwPoOyAZZQ4AAIt-Z3iUwQ316.jpg',
                        0, '25d55ad283aa400af464c76d713c07ad', null, null, 0, '2023-05-01 13:56:40', null, 0, 1,
                        '{"大一":1,"C++":1,"emo":1}', '人生就像沙漏，今天它要补满所有的缺口；明天它要修補所有的窟窿。');

insert into user value (null, '惠俊哲', 'Velarde',
                        'http://cdn.u2.huluxia.com/g3/M03/29/B8/wKgBOVwKZPyAeZyOAAGBKo7j4sY097.jpg',
                        0, '25d55ad283aa400af464c76d713c07ad', null, null, 0, '2023-05-01 13:56:40', null, 0, 1,
                        '{"随和":1,"程序员":1,"Python":1}',
                        '当你不再寻找爱情，只是去爱；当你不再渴望成功，只是去做；你不再追求成长，只是去修，一切才真正开始。');

insert into user value (null, '弥芷珍', 'Jesse',
                        'http://cdn.u2.huluxia.com/g3/M01/36/6B/wKgBOVwPn8SAfHRcAAC046awE3Q885.jpg',
                        0, '25d55ad283aa400af464c76d713c07ad', null, null, 0, '2023-05-01 13:56:40', null, 0, 1,
                        '{"伤心":1,"萌妹子":1,"Web3":1}', '我是一个废物，但并不代表我没有尊严。');

insert into user value (null, '玉晗玥', 'Ridley',
                        'http://cdn.u2.huluxia.com/g3/M03/26/B6/wKgBOVwJE7KAA5CyAAOIpyg9x28469.jpg',
                        0, '25d55ad283aa400af464c76d713c07ad', null, null, 0, '2023-05-01 13:56:40', null, 0, 1,
                        '{"开心":1,"萌妹子":1,"Docker":1}', '我不是一个完美的人，但是我接受所有的不完美。');

insert into user value (null, '局曼雁', 'Lowell',
                        'http://cdn.u2.huluxia.com/g3/M01/27/D3/wKgBOVwJonyAKm54AAH-N0Lk_io211.jpg',
                        0, '25d55ad283aa400af464c76d713c07ad', null, null, 0, '2023-05-01 13:56:40', null, 0, 1,
                        '{"emo":1,"乐子人":1,"C++":1}', '你是我的小确幸，可就算你不喜欢我，我也不会再去追寻别人。');

DROP TABLE IF EXISTS team;
create table if not exists team
(
    id          bigint auto_increment comment 'id' primary key,
    name    varchar(256)                       not null comment '队伍名称',
    description varchar(1024)                      null comment '描述',
    maxNum      int      default 1                 not null comment '最大人数',
    expireTime  datetime                           null comment '过期时间',
    userId      bigint comment '用户id',
    status      int      default 0                 not null comment '状态 0-公开 1-私有 2-加密',
    password    varchar(512)                       null comment '密码',
    createTime  datetime default CURRENT_TIMESTAMP null comment '创建时间',
    updateTime  datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    isDelete    tinyint  default 0                 not null comment '是否删除'
)
    comment '队伍';

DROP TABLE IF EXISTS user_team;
create table if not exists user_team
(
    id         bigint auto_increment comment 'id' primary key,
    userId     bigint comment '用户id',
    teamId     bigint comment '队伍id',
    joinTime   datetime                           null comment '加入时间',
    createTime datetime default CURRENT_TIMESTAMP null comment '创建时间',
    updateTime datetime default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP comment '更新时间',
    isDelete   tinyint  default 0                 not null comment '是否删除'
)
    comment '用户-队伍';





