-- Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
-- Use of this source code is governed by a MIT style
-- license that can be found in the LICENSE file.

CREATE TABLE `GroupContact` (
  `Id` BIGINT NOT NULL AUTO_INCREMENT,
  `GroupId` BIGINT DEFAULT NULL COMMENT '群ID',
  `UserId` BIGINT DEFAULT NULL COMMENT '用户ID',
  `CategoryId` BIGINT NOT NULL DEFAULT 0 COMMENT "用户自定义分组ID",
  `UserNickname` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '群昵称',
  `Remark` VARCHAR(120) NOT NULL DEFAULT '' COMMENT '群备注',
  `Background` VARCHAR(120) NOT NULL DEFAULT '' COMMENT '背景',
  `IsDisturb` INT NOT NULL DEFAULT 0 COMMENT '是否免打扰',
  `IsTop` INT NOT NULL DEFAULT 0 COMMENT '是否置顶',
  `IsShowNickname` INT NOT NULL DEFAULT 0 COMMENT '是否显示群昵称',
  `ApprovalStatus` INT NOT NULL DEFAULT 0 COMMENT "审批状态",
  `ApprovalAt` DATETIME DEFAULT NULL COMMENT "审批时间",
  `CreatedAt` DATETIME DEFAULT NULL COMMENT "创建时间",
  `UpdatedAt` DATETIME DEFAULT NULL COMMENT "更新时间",
  PRIMARY KEY `PK_Id` (`Id`),
  KEY `IDX_GroupId` (`GroupId`) USING BTREE,
  KEY `IDX_UserId` (`UserId`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;