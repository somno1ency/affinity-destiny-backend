/**
 * 用户信息表,这张表要改的地方很多
 */
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `mobile` varchar(20) DEFAULT NULL COMMENT "手机",
  `password` varchar(40) DEFAULT NULL COMMENT "密码",
  `avatar` varchar(150) DEFAULT NULL COMMENT "头像",
  `sex` tinyint DEFAULT NULL COMMENT "性别(0:未设置 1:男 2:女)",
  `nickname` varchar(20) DEFAULT NULL COMMENT "昵称",
  `salt` varchar(10) DEFAULT NULL COMMENT "盐值",
  `online` tinyint DEFAULT NULL COMMENT "在线(0:否 1:是)",
  `token` varchar(40) DEFAULT NULL COMMENT "token",
  `memo` varchar(140) DEFAULT NULL COMMENT "备注",
  `created_at` datetime DEFAULT NULL COMMENT "创建时间",
  `updated_at` datetime DEFAULT NULL COMMENT "更新时间",
  `last_login_at` datetime DEFAULT NULL COMMENT "最后登录时间",
  PRIMARY KEY `pk_id` (`id`),
  UNIQUE KEY `uk_mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4;