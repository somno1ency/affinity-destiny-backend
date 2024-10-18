-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `User` (
  `Id` BIGINT NOT NULL AUTO_INCREMENT,
  `CustomId` VARCHAR(20) NOT NULL DEFAULT '' COMMENT "自定义ID",
  `Mobile` VARCHAR(20) NOT NULL DEFAULT '' COMMENT "手机",
  `Password` VARCHAR(40) NOT NULL DEFAULT '' COMMENT "密码",
  `Nickname` VARCHAR(20) NOT NULL DEFAULT '' COMMENT "昵称",
  `Avatar` VARCHAR(120) NOT NULL DEFAULT '' COMMENT "头像",
  `Sex` INT NOT NULL DEFAULT 0 COMMENT "性别(0:未设置 1:男 2:女)",
  `Memo` VARCHAR(120) NOT NULL DEFAULT '' COMMENT "备注",
  `Salt` VARCHAR(10) NOT NULL DEFAULT '' COMMENT "盐值",
  `CreatedAt` DATETIME DEFAULT NULL COMMENT "创建时间",
  `UpdatedAt` DATETIME DEFAULT NULL COMMENT "更新时间",
  `LastLoginAt` DATETIME DEFAULT NULL COMMENT "最后登录时间",
  PRIMARY KEY `PK_Id` (`Id`),
  KEY `IDX_CustomId` (`CustomId`) USING BTREE,
  UNIQUE KEY `UK_Mobile` (`Mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;