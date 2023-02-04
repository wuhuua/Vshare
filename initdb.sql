CREATE DATABASE IF NOT EXISTS `vshare`;
USE `vshare`;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
	`id`			bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
	`name`          varchar(128)        NOT NULL DEFAULT '' COMMENT '用户昵称',
	`password`		varchar(32)			NOT NULL DEFAULT '' COMMENT '用户密码',
	`followcount`   bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '关注人数',
	`followercount` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '粉丝人数',
	`headicon`		varchar(128)		NOT NULL DEFAULT 0 COMMENT '头像地址',
	PRIMARY KEY(`id`)
) ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

 INSERT INTO `user`
 VALUES (20230001, 'Iori','12345678',0,0,"https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/d94b31bfa5a74db9b7d448e6ab3bc3e8~tplv-k3u1fbpfcp-watermark.image?"),
		(20230002, 'Misoda','12345678',0,0,"https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/05ea8220091d4a2a82c85f68c900368d~tplv-k3u1fbpfcp-watermark.image?");

DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
	`id`			bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
	`userid`		bigint(20) unsigned NOT NULL COMMENT '作者id',
	`playurl`		varchar(256)        NOT NULL DEFAULT '' COMMENT '播放链接',
	`coverurl`		varchar(256)        NOT NULL DEFAULT '' COMMENT '封面链接',
	`favoritecount`	bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '点赞数目',
	`commentcount`	bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '评论数目',
	`videoname`		varchar(128)		NOT NULL DEFAULT '' COMMENT '视频标题',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`userid`) REFERENCES user(`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='视频表';

 INSERT INTO `video`
 VALUES (20230001,
 		 20230001,
 		"https://www.w3schools.com/html/movie.mp4",
		"https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		0,
		1,
		'title test'
		);

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
	`id`			bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
	`userid`		bigint(20) unsigned NOT NULL COMMENT '作者id',
	`videoid`		bigint(20) unsigned NOT NULL COMMENT '视频id',
	`content`		varchar(256) 		NOT NULL DEFAULT '' COMMENT '评论内容',
	`createdate`	varchar(128)		NOT NULL DEFAULT '' COMMENT '创建日期',
	PRIMARY KEY (`id`),
	FOREIGN KEY (`userid`) REFERENCES user(`id`),
	FOREIGN KEY (`videoid`) REFERENCES video(`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='评论区';

  INSERT INTO `comment`
  VALUES (10000001,20230001,20230001,"test message","01-21")
