CREATE TABLE `user` (
	`id` INT(1) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
	`user_id` VARCHAR(40) NOT NULl DEFAULT '' COMMENT '用户Id',
	`password` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '密码',
	`name` VARCHAR(40) NOT NULL DEFAULT '' COMMENT '用户名' COLLATE 'utf8mb4_general_ci',
	`phone_number` VARCHAR(40) NULL DEFAULT '' COMMENT '电话号码' COLLATE 'utf8mb4_general_ci',
	`email` VARCHAR(40) NULL DEFAULT '' COMMENT '邮箱' COLLATE 'utf8mb4_general_ci',
	`status_flag` TINYINT(1) NULL DEFAULT '' COMMENT '1:听歌中~; 2:活跃者; 3:默默无闻的点歌达人',
	`avatar` VARCHAR(64) NULL DEFAULT '' COMMENT '用户头像，存储头像图片的链接' COLLATE 'utf8mb4_general_ci',
	`person_signature` VARCHAR(256) NULL DEFAULT '' COMMENT '个人简介' COLLATE 'utf8mb4_general_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8mb4_general_ci' ENGINE=InnoDB;