-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `UserContact` (
  `Id` BIGINT NOT NULL AUTO_INCREMENT,
  `OwnerId` BIGINT NOT NULL DEFAULT 0 COMMENT '所有者用户ID',
  `DstId` BIGINT NOT NULL DEFAULT 0 COMMENT '目标用户ID',
  `CategoryId` BIGINT NOT NULL DEFAULT 0 COMMENT "用户自定义分组ID",
  `Background` VARCHAR(120) NOT NULL DEFAULT '' COMMENT '背景',
  `IsDisturb` BOOLEAN NOT NULL DEFAULT false COMMENT '是否免打扰',
  `IsTop` BOOLEAN NOT NULL DEFAULT false COMMENT '是否置顶',
  `IsRemind` BOOLEAN NOT NULL DEFAULT false COMMENT '是否提醒',
  `ApprovalStatus` BOOLEAN NOT NULL DEFAULT false COMMENT "审批状态",
  `ApprovalAt` DATETIME DEFAULT NULL COMMENT "审批时间",
  `CreatedAt` DATETIME DEFAULT NULL COMMENT "创建时间",
  `UpdatedAt` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `PK_Id` (`Id`),
  KEY `IDX_OwnerId` (`OwnerId`) USING BTREE,
  KEY `IDX_DstId` (`DstId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;