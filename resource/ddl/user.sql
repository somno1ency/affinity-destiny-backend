-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `user` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `custom_id` VARCHAR(20) NOT NULL DEFAULT '' COMMENT "自定义ID",
  `mobile` VARCHAR(20) NOT NULL DEFAULT '' COMMENT "手机",
  `password` VARCHAR(40) NOT NULL DEFAULT '' COMMENT "密码",
  `nickname` VARCHAR(20) NOT NULL DEFAULT '' COMMENT "昵称",
  `avatar` VARCHAR(120) NOT NULL DEFAULT '' COMMENT "头像",
  `sex` TINYINT NOT NULL DEFAULT 0 COMMENT "性别(0:未设置 1:男 2:女)",
  `memo` VARCHAR(120) NOT NULL DEFAULT '' COMMENT "备注",
  `salt` VARCHAR(10) NOT NULL DEFAULT '' COMMENT "盐值",
  `created_at` DATETIME DEFAULT NULL COMMENT "创建时间",
  `updated_at` DATETIME DEFAULT NULL COMMENT "更新时间",
  `last_login_at` DATETIME DEFAULT NULL COMMENT "最后登录时间",
  PRIMARY KEY `pk_id` (`id`),
  KEY `idx_customId` (`custom_id`) USING BTREE,
  UNIQUE KEY `uk_mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;