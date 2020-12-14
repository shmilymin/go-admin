DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '角色名称',
  `create_by` BIGINT NOT NULL COMMENT '创建用户',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_by` BIGINT NOT NULL COMMENT '更新用户',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT='角色';
INSERT INTO `role` VALUES ('1', '开发部', 1,NOW(), 1,NOW());
INSERT INTO `role` VALUES ('2', '运维部', 1,NOW(), 1,NOW());
INSERT INTO `role` VALUES ('3', '测试部', 1,NOW(), 1,NOW());

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` VARCHAR(255) NOT NULL COMMENT '用户名',
  `password` VARCHAR(255) NOT NULL COMMENT '密码',
  `avatar` VARCHAR(255) COMMENT '头像',
  `phone` VARCHAR(255) COMMENT '手机号',
  `email` VARCHAR(255) COMMENT '邮箱',
  `enabled` INT DEFAULT 1 COMMENT '状态：1启用、0禁用',
  `create_by` BIGINT NOT NULL COMMENT '创建用户',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_by` BIGINT NOT NULL COMMENT '更新用户',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT='用户';


DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` BIGINT NOT NULL COMMENT '用户ID',
  `role_id` BIGINT NOT NULL COMMENT '角色ID',
  `create_by` BIGINT NOT NULL COMMENT '创建用户',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_by` BIGINT NOT NULL COMMENT '更新用户',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT='用户角色关联';

DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` VARCHAR(255) COMMENT '图标',
  `pid` BIGINT NOT NULL COMMENT '上级菜单ID',
  `sort` INT NOT NULL COMMENT '排序',
  `i_frame` BIT(1) DEFAULT b'0' COMMENT '是否外链',
  `component_name` VARCHAR(255) NOT NULL COMMENT '组件名称',
  `component` VARCHAR(255) NOT NULL DEFAULT '-' COMMENT '组件路径',
  `cache` BIT(1) DEFAULT b'0' COMMENT '菜单缓存',
  `hidden` BIT(1) DEFAULT b'0' COMMENT '菜单可见',
  `path` VARCHAR(255) DEFAULT NULL COMMENT '链接地址',
  `create_by` BIGINT NOT NULL COMMENT '创建用户',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_by` BIGINT NOT NULL COMMENT '更新用户',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT='菜单';

DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` BIGINT NOT NULL COMMENT '角色ID',
  `menu_id` BIGINT NOT NULL COMMENT '菜单ID',
  `create_by` BIGINT NOT NULL COMMENT '创建用户',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_by` BIGINT NOT NULL COMMENT '更新用户',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT='角色菜单关联';

DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `alias` VARCHAR(255) NOT NULL COMMENT '别名',
  `name` VARCHAR(255) NOT NULL COMMENT '名称',
  `pid` INT(11) NOT NULL COMMENT '上级权限',
  `create_by` BIGINT NOT NULL COMMENT '创建用户',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_by` BIGINT NOT NULL COMMENT '更新用户',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT='权限';

DROP TABLE IF EXISTS `role_permission`;
CREATE TABLE `role_permission`  (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `role_id` BIGINT NOT NULL COMMENT '角色ID',
  `permission_id` BIGINT NOT NULL COMMENT '权限ID',
  `create_by` BIGINT NOT NULL COMMENT '创建用户',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  `update_by` BIGINT NOT NULL COMMENT '更新用户',
  `update_time` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) COMMENT='角色权限关联';

