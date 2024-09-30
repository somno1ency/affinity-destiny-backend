/**
 * 群信息表
 */
CREATE TABLE `group` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `group_id` varchar(10) DEFAULT NULL COMMENT "群标识",
  `name` varchar(30) DEFAULT NULL COMMENT "群名称",
  `owner_id` bigint DEFAULT NULL COMMENT "群创建者ID",
  `icon` varchar(250) DEFAULT NULL COMMENT "群头像",
  `category` tinyint DEFAULT NULL COMMENT "群类型",
  `memo` varchar(120) DEFAULT NULL COMMENT "群备注",
  `created_at` datetime DEFAULT NULL COMMENT "创建时间",
  `updated_at` datetime DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `pk_id` (`id`),
  UNIQUE KEY `uk_groupId` (`group_id`) USING BTREE,
  UNIQUE KEY `uk_name` (`name`) USING BTREE,
  KEY `idx_ownerId` (`owner_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;