-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `resource` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `src` VARCHAR(120) NOT NULL DEFAULT '' COMMENT '资源地址',
  `type` TINYINT NOT NULL DEFAULT 0 COMMENT "资源类型",
  `name_zh` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '资源中文名称',
  `name_en` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '资源英文名称',
  `created_at` DATETIME DEFAULT NULL COMMENT "创建时间",
  `updated_at` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `pk_id` (`id`),
  KEY `idx_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;