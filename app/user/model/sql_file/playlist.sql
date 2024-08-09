CREATE TABLE `playlist` (
	`id` INT(1) NOT NULL AUTO_INCREMENT COMMENT '歌单ID,主键',
	`name` VARCHAR(50) NOT NULL COMMENT '歌单名称' COLLATE 'utf8mb4_general_ci',
	`last_play_time` TIMESTAMP NULL DEFAULT NULL COMMENT '最后一次播放时间',
	`play_count` INT(1) NULL DEFAULT NULL COMMENT '播放次数',
	`is_public` TINYINT(4) NULL DEFAULT NULL COMMENT '是否公开',
	`avatar` VARCHAR(50) NULL DEFAULT NULL COMMENT '歌曲头像，默认用最后一次收藏的歌曲图片' COLLATE 'utf8mb4_general_ci',
	`create_time` TIMESTAMP NULL DEFAULT NULL COMMENT '创建时间',
	`update_time` TIMESTAMP NULL DEFAULT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
ENGINE=InnoDB
;
