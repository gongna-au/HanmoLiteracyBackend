drop database if exists `HanmoLiteracy`;

create database `HanmoLiteracy`;

use  `HanmoLiteracy`;

-- 用户信息表(user)
create table `tbl_user`(
   `id` int not null AUTO_INCREMENT comment "用户id" ,       
   `phone` varchar(20) UNIQUE comment "账户",
   `password` varchar(30) not null comment "密码",
   `name` varchar(20) not null comment "昵称",
   `gender` varchar(20) not null  comment "性别",
-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

-- 汉字信息表(character)
create table `tbl_character`(
   `id` int not null AUTO_INCREMENT comment "汉字id" ,       
   `name` varchar(20) UNIQUE comment "名称",
-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 视频表(character)
create table `tbl_video`(
   `id` int not null AUTO_INCREMENT comment "视频id" ,  
   `path` varchar(100) null comment "视频文件路径以及名称",
-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 视频对应的汉字表(character)
create table `tbl_video_records`(
   `id`  int not null AUTO_INCREMENT comment "学习记录id" ,
   `vid` int  null comment "视频id",
   `cid` int  null comment "汉字id",
-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;


-- 学习信息表(character)
create table `tbl_study_records`(
   `id` int not null AUTO_INCREMENT comment "学习记录id" ,
   `uid` int  null comment "用户id",
   `cid` int  null comment "汉字id",
   `cname` varchar(20) null comment "汉字名称",
   `time` varchar(20) null comment "时间",
-- 添加约束
primary key (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;










