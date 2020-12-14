DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '账号',
  `password` VARCHAR(50) NOT NULL DEFAULT '' COMMENT '密码',
  `gender` TINYINT NOT NULL DEFAULT 3 COMMENT '性别(1:男,2:女,3:未知)',
  `create_user_id` BIGINT NOT NULL COMMENT '创建用户id',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_user_id` BIGINT NOT NULL COMMENT '更新用户id',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户';

INSERT INTO `user`(username, PASSWORD,create_user_id,create_time,update_user_id,update_time) 
VALUES ('admin', '21232f297a57a5a743894a0e4a801fc3', 1,NOW(), 1,NOW());


DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名字',
  `path` varchar(50) NOT NULL DEFAULT '' COMMENT '访问路径',
  `method` varchar(50) NOT NULL DEFAULT '' COMMENT '资源请求方式',
  `create_user_id` BIGINT NOT NULL COMMENT '创建用户id',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_user_id` BIGINT NOT NULL COMMENT '更新用户id',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='菜单';

INSERT INTO `menu` VALUES ('1', '查询所有菜单', '/api/menu', 'GET', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('2', '查询单个菜单', '/api/menu/:id', 'GET', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('3', '创建单个菜单', '/api/menu', 'POST', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('4', '更新单个菜单', '/api/menu/:id', 'PUT', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('5', '删除单个菜单', '/api/menu/:id', 'DELETE', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('6', '查询所有用户', '/api/user', 'GET', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('7', '查询单个用户', '/api/user/:id', 'GET', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('8', '创建单个用户', '/api/user', 'POST', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('9', '更新单个用户', '/api/user/:id', 'PUT', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('10', '删除单个用户', '/api/user/:id', 'DELETE', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('11', '查询所有角色', '/api/role', 'GET', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('12', '查询单个角色', '/api/role/:id', 'GET', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('13', '创建单个角色', '/api/role', 'POST', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('14', '更新单个角色', '/api/role/:id', 'PUT', 1,NOW(), 1,NOW());
INSERT INTO `menu` VALUES ('15', '删除单个角色', '/api/role/:id', 'DELETE', 1,NOW(), 1,NOW());


DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名字',
  `create_user_id` BIGINT NOT NULL COMMENT '创建用户id',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_user_id` BIGINT NOT NULL COMMENT '更新用户id',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='角色';

INSERT INTO `role` VALUES ('1', '开发部', 1,NOW(), 1,NOW());
INSERT INTO `role` VALUES ('2', '运维部', 1,NOW(), 1,NOW());
INSERT INTO `role` VALUES ('3', '测试部', 1,NOW(), 1,NOW());


DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role_id` bigint unsigned NOT NULL COMMENT '角色ID',
  `menu_id` bigint unsigned NOT NULL COMMENT '菜单ID',
  `create_user_id` BIGINT NOT NULL COMMENT '创建用户id',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_user_id` BIGINT NOT NULL COMMENT '更新用户id',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='角色菜单关联';

INSERT INTO `role_menu` VALUES ('1', '2', '1', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('2', '2', '2', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('3', '2', '3', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('4', '2', '4', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('5', '2', '5', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('6', '2', '6', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('7', '2', '7', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('8', '2', '8', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('9', '2', '9', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('10', '2', '10', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('11', '2', '11', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('12', '2', '12', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('13', '2', '13', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('14', '2', '14', 1,NOW(), 1,NOW());
INSERT INTO `role_menu` VALUES ('15', '2', '15', 1,NOW(), 1,NOW());


DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) unsigned NOT NULL COMMENT '用户ID',
  `role_id` int(11) unsigned NOT NULL COMMENT '角色ID',
  `create_user_id` BIGINT NOT NULL COMMENT '创建用户id',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_user_id` BIGINT NOT NULL COMMENT '更新用户id',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='用户角色关联';

INSERT INTO `user_role` VALUES ('1', '2', '2', 1,NOW(), 1,NOW());
