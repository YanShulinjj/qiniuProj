CREATE DATABASE IF NOT EXISTS qiniu;
USE qiniu;
CREATE TABLE `user` (
        `user_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
        `user_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '用户ip',
        `next_page` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '已经存在的Page数量',
        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        PRIMARY KEY (`user_id`),
        UNIQUE KEY `uniq_userip` (`user_ip`),
        KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';