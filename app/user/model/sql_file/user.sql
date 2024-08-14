CREATE TABLE `user` (
	`id` INT(1) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
	`user_id` INT(1) NOT NULL DEFAULT '0' COMMENT '用户ID',
	`password` VARCHAR(128) NOT NULL DEFAULT '0' COMMENT '密码' COLLATE 'latin1_swedish_ci',
	`name` VARCHAR(40) NULL DEFAULT '0' COMMENT '用户名' COLLATE 'utf8mb4_general_ci',
	`phone_number` VARCHAR(40) NULL DEFAULT '0' COMMENT '电话号码' COLLATE 'utf8mb4_general_ci',
	`email` VARCHAR(40) NULL DEFAULT '0' COMMENT '邮箱' COLLATE 'utf8mb4_general_ci',
	`status_flag` TINYINT(1) NULL DEFAULT '0' COMMENT '1:听歌中~; 2:活跃者; 3:默默无闻的点歌达人',
	`avatar` VARCHAR(64) NULL DEFAULT '0' COMMENT '用户头像，存储头像图片的链接' COLLATE 'utf8mb4_general_ci',
	`person_signature` VARCHAR(256) NULL DEFAULT '0' COMMENT '个人简介' COLLATE 'utf8mb4_general_ci',
	`token` VARCHAR(128) NULL DEFAULT '' COMMENT '用户Token' COLLATE 'utf8mb4_general_ci',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `user_id` (`user_id`) USING BTREE
)
COLLATE='latin1_swedish_ci'
ENGINE=InnoDB
AUTO_INCREMENT=3
;
