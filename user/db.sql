
CREATE DATABASE user;
USE user;

CREATE TABLE `user` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
    `password` varchar(50) NOT NULL DEFAULT '' COMMENT '用户密码，MD5加密',
    `phone` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
