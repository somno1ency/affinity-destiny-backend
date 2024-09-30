/**
 * 用户关系表
 * 
 * 展望: 以后还会出现联系人分组,当前用户与联系人的关系等定义
 */
CREATE TABLE `user_contact` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `owner_id` bigint DEFAULT NULL COMMENT '所有者用户ID',
  `dst_id` bigint DEFAULT NULL COMMENT '目标用户ID',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY `pk_id` (`id`),
  KEY `idx_ownerId` (`owner_id`) USING BTREE,
  -- 用于查找好友时提高性能
  KEY `idx_dstId` (`dst_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;