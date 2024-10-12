-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `user_contact` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `owner_id` BIGINT NOT NULL DEFAULT 0 COMMENT '所有者用户ID',
  `dst_id` BIGINT NOT NULL DEFAULT 0 COMMENT '目标用户ID',
  `category_id` BIGINT NOT NULL DEFAULT 0 COMMENT "用户自定义分组ID",
  `background` VARCHAR(120) NOT NULL DEFAULT '' COMMENT '背景',
  `is_disturb` BOOLEAN NOT NULL DEFAULT false COMMENT '是否免打扰',
  `is_top` BOOLEAN NOT NULL DEFAULT false COMMENT '是否置顶',
  `isRemind` BOOLEAN NOT NULL DEFAULT false COMMENT '是否提醒',
  `created_at` DATETIME DEFAULT NULL COMMENT "创建时间",
  `updated_at` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `pk_id` (`id`),
  KEY `idx_ownerId` (`owner_id`) USING BTREE,
  KEY `idx_dstId` (`dst_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;