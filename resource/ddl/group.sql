CREATE TABLE `group` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `custom_id` VARCHAR(10) NOT NULL DEFAULT '' COMMENT "群标识",
  `category_id` BIGINT NOT NULL DEFAULT 0 COMMENT "用户自定义分组ID",
  `name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT "群名称",
  `owner_id` BIGINT NOT NULL DEFAULT 0 COMMENT "群创建者ID",
  `avatar` VARCHAR(120) NOT NULL DEFAULT '' COMMENT "群头像",
  `memo` VARCHAR(120) NOT NULL DEFAULT '' COMMENT "群备注",
  `created_at` DATETIME DEFAULT NULL COMMENT "创建时间",
  `updated_at` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `pk_id` (`id`),
  KEY `idx_customId` (`custom_id`) USING BTREE,
  KEY `idx_ownerId` (`owner_id`) USING BTREE,
  UNIQUE KEY `uk_name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;