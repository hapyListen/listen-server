CREATE TABLE `songlist_collectlist` (
	`id` INT(1) NOT NULL AUTO_INCREMENT COMMENT '收藏Id,主键',
	`user_id` INT(1) NOT NULL COMMENT `用户Id`
	`songlist_id` INT(1) NOT NULL COMMENT '歌单Id' COLLATE 'utf8mb4_general_ci',
	`play_count` INT(1) NULL DEFAULT NULL COMMENT '播放次数',
	 PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
COMMENT `音乐收藏表`
ENGINE=InnoDB
;
