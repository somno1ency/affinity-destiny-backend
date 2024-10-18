-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `Category` (
  `Id` BIGINT NOT NULL AUTO_INCREMENT,
  `OwnerId` BIGINT NOT NULL DEFAULT 0 COMMENT "创建者ID",
  `NameZh` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '分类中文名称',
  `NameEn` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '分类英文名称',
  `Star` INT NOT NULL DEFAULT 0 COMMENT "分类星级",
  `CreatedAt` DATETIME DEFAULT NULL COMMENT "创建时间",
  `UpdatedAt` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `PK_Id` (`Id`),
  KEY `IDX_OwnerId` (`OwnerId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;