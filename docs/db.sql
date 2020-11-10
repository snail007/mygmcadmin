-- Adminer 4.7.7 MySQL dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `gmc_user`;
CREATE TABLE `gmc_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `username` varchar(24) NOT NULL COMMENT '用户名',
  `nickname` varchar(24) NOT NULL COMMENT '昵称',
  `password` char(32) NOT NULL COMMENT '密码',
  `is_delete` tinyint(1) NOT NULL COMMENT '是否删除，1:是，0:否',
  `update_time` int(11) NOT NULL COMMENT '更新时间',
  `create_time` int(11) NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

INSERT INTO `gmc_user` (`user_id`, `username`, `nickname`, `password`, `is_delete`, `update_time`, `create_time`) VALUES
(1,	'root',	'root',	'2df594b9710111099edbdb7edaa43301',	0,	1605026224,	1605013042);

-- 2020-11-10 16:37:24