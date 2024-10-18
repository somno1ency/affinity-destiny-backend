-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `Resource` (
  `Id` BIGINT NOT NULL AUTO_INCREMENT,
  `Src` VARCHAR(120) NOT NULL DEFAULT '' COMMENT '资源地址',
  `Type` TINYINT NOT NULL DEFAULT 0 COMMENT "资源类型",
  `NameZh` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '资源中文名称',
  `NameEn` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '资源英文名称',
  `CreatedAt` DATETIME DEFAULT NULL COMMENT "创建时间",
  `UpdatedAt` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `PK_Id` (`Id`),
  KEY `IDX_Type` (`Type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;