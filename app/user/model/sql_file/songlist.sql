CREATE TABLE `songlist` (
	`id` INT(1) NOT NULL AUTO_INCREMENT COMMENT '歌单ID,主键',
	`own_user_id` INT(1) NOT NULL DEFAULT 0 COMMENT `歌单拥有者ID`,
	`name` VARCHAR(50) NOT NULL DEFAULT '啸啸啸歌单' COMMENT '歌单名称' COLLATE 'utf8mb4_general_ci',
	`last_play_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次播放时间',
	`play_count` INT(1) NULL DEFAULT 0 COMMENT '播放次数',
	`is_public` TINYINT(1) NOT NULL DEFAULT 0 COMMENT '是否公开',
	`avatar` VARCHAR(50) NULL DEFAULT '' COMMENT '歌曲头像，默认用最后一次收藏的歌曲图片' COLLATE 'utf8mb4_general_ci',
    `describe` VARCHAR(128) NULL DEFAULT '独到的内涵，值得细细品味' COMMENT `歌单描述`,
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
COMMENT `歌单表`
ENGINE=InnoDB
;
