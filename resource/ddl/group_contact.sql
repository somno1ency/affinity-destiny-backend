CREATE TABLE `group_contact` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `group_id` BIGINT DEFAULT NULL COMMENT '群ID',
  `user_id` BIGINT DEFAULT NULL COMMENT '用户ID',
  `userNickname` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '群昵称',
  `remark` VARCHAR(120) NOT NULL DEFAULT '' COMMENT '群备注',
  `background` VARCHAR(120) NOT NULL DEFAULT '' COMMENT '背景',
  `is_disturb` BOOLEAN NOT NULL DEFAULT false COMMENT '是否免打扰',
  `is_top` BOOLEAN NOT NULL DEFAULT false COMMENT '是否置顶',
  `is_show_nickname` BOOLEAN NOT NULL DEFAULT false COMMENT '是否显示群昵称',
  `created_at` DATETIME DEFAULT NULL COMMENT "创建时间",
  `updated_at` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `pk_id` (`id`),
  KEY `idx_groupId` (`group_id`) USING BTREE,
  KEY `idx_userId` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;