/**
 * 用户群关系表
 * 
 * 展望: 以后还会出现群分组,当前用户与群的关系等定义
 */
CREATE TABLE `group_contact` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `group_id` bigint DEFAULT NULL COMMENT '群ID',
  `user_id` bigint DEFAULT NULL COMMENT '用户ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY `pk_id` (`id`),
  KEY `idx_groupId` (`group_id`) USING BTREE,
  -- 用于为用户添加群时提高性能
  KEY `idx_userId` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;