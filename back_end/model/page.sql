CREATE DATABASE IF NOT EXISTS qiniu;
USE qiniu;
CREATE TABLE `page` (
        `page_id`  bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '页面ID',
        `user_id`  bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
        `page_idx` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户的第idx个页面',
        `page_name` varchar(64) NOT NULL DEFAULT '' COMMENT '页面别名',
        `svg_path` varchar(108) NOT NULL DEFAULT '' COMMENT '页面svg文件路径',
        `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
        PRIMARY KEY (`page_id`),
        UNIQUE KEY `uniq_user_page` (`user_id`, `page_idx`),
        KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';