-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `category` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `owner_id` BIGINT NOT NULL DEFAULT 0 COMMENT "创建者ID",
  `name_zh` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '分类中文名称',
  `name_en` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '分类英文名称',
  `star` TINYINT NOT NULL DEFAULT 0 COMMENT "分类星级",
  `created_at` DATETIME DEFAULT NULL COMMENT "创建时间",
  `updated_at` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `pk_id` (`id`),
  KEY `idx_ownerId` (`owner_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;