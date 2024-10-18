-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `Group` (
  `Id` BIGINT NOT NULL AUTO_INCREMENT,
  `CustomId` VARCHAR(10) NOT NULL DEFAULT '' COMMENT "群标识",
  `Name` VARCHAR(20) NOT NULL DEFAULT '' COMMENT "群名称",
  `OwnerId` BIGINT NOT NULL DEFAULT 0 COMMENT "群创建者ID",
  `Avatar` VARCHAR(120) NOT NULL DEFAULT '' COMMENT "群头像",
  `Memo` VARCHAR(120) NOT NULL DEFAULT '' COMMENT "群备注",
  `IsApproval` BOOLEAN NOT NULL DEFAULT false COMMENT "是否需要审批",
  `CreatedAt` DATETIME DEFAULT NULL COMMENT "创建时间",
  `UpdatedAt` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `PK_Id` (`Id`),
  KEY `IDX_CustomId` (`CustomId`) USING BTREE,
  KEY `IDX_OwnerId` (`OwnerId`) USING BTREE,
  UNIQUE KEY `UK_Name` (`Name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;