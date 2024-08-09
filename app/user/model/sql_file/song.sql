CREATE TABLE `song` (
	`id` INT(1) NOT NULL AUTO_INCREMENT COMMENT'歌曲ID,主键',
	`name` VARCHAR(50) NOT NULL DEFAULT "啸啸啸" COMMENT '歌曲名称' COLLATE 'utf8mb4_general_ci',
	`songlist_id` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '歌单ID' COLLATE 'utf8mb4_general_ci',
	`address` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '歌曲资源地址' COLLATE 'utf8mb4_general_ci',
	`picture_address` VARCHAR(50) NULL DEFAULT '' COMMENT '歌曲封面地址' COLLATE 'utf8mb4_general_ci',
	`platform` VARCHAR(50) NULL DEFAULT ''  COMMENT '所属平台(bilibili、QQ、网易、酷我、酷狗...)' COLLATE 'utf8mb4_general_ci',
	`singer_id` VARCHAR(50) NULL DEFAULT ''  COMMENT '歌手ID(平台下的歌手ID)' COLLATE 'utf8mb4_general_ci',
	`album` VARCHAR(50) NULL DEFAULT ''  COMMENT '歌曲所属专辑（如果有）' COLLATE 'utf8mb4_general_ci',
	`album_id` VARCHAR(50) NULL DEFAULT ''  COMMENT '歌曲所属专辑ID(如果有,记录平台下的专辑ID)' COLLATE 'utf8mb4_general_ci',
	`collection_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加收藏时间',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci'
COMMENT `音乐表`
ENGINE=InnoDB
;