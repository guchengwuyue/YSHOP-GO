/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : yshop_go

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 07/03/2021 10:38:07
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `pid` bigint NOT NULL COMMENT '上级部门',
  `enabled` tinyint(1) NOT NULL COMMENT '状态',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime DEFAULT NULL,
  `is_del` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='部门';

-- ----------------------------
-- Records of dept
-- ----------------------------
BEGIN;
INSERT INTO `dept` VALUES (1, 'YSHOP', 0, 1, '2021-02-26 14:51:14', NULL, 0);
INSERT INTO `dept` VALUES (2, '研发部', 7, 1, '2019-03-25 09:15:32', NULL, 0);
INSERT INTO `dept` VALUES (5, '运维部', 7, 1, '2019-03-25 09:20:44', NULL, 0);
INSERT INTO `dept` VALUES (6, '测试部', 8, 0, '2021-02-25 16:22:53', NULL, 0);
INSERT INTO `dept` VALUES (7, '华南分部', 1, 1, '2019-03-25 11:04:50', NULL, 0);
INSERT INTO `dept` VALUES (8, '华北分部', 1, 1, '2019-03-25 11:04:53', NULL, 0);
INSERT INTO `dept` VALUES (11, '人事部', 8, 1, '2019-03-25 11:07:58', NULL, 0);
INSERT INTO `dept` VALUES (12, '7773888', 1, 1, '2021-02-26 15:04:43', NULL, 1);
INSERT INTO `dept` VALUES (13, '333', 1, 0, '2021-02-26 14:28:26', '2021-02-26 14:28:26', 1);
INSERT INTO `dept` VALUES (14, '9922', 13, 1, '2021-02-26 15:05:05', '2021-02-26 14:29:17', 1);
COMMIT;

-- ----------------------------
-- Table structure for dict
-- ----------------------------
DROP TABLE IF EXISTS `dict`;
CREATE TABLE `dict` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字典名称',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '描述',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime DEFAULT NULL,
  `is_del` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='数据字典';

-- ----------------------------
-- Records of dict
-- ----------------------------
BEGIN;
INSERT INTO `dict` VALUES (1, 'user_status', '用户状态', '2019-10-27 20:31:36', NULL, 0);
INSERT INTO `dict` VALUES (4, 'dept_status', '部门状态', '2019-10-27 20:31:36', NULL, 0);
INSERT INTO `dict` VALUES (5, 'job_status', '岗位状态', '2019-10-27 20:31:36', NULL, 0);
INSERT INTO `dict` VALUES (6, '33', '3', '2020-05-18 19:55:49', NULL, 1);
INSERT INTO `dict` VALUES (16, 'force_update', '强制升级', '2020-12-09 11:13:21', NULL, 0);
INSERT INTO `dict` VALUES (17, 'is_enable', '是否启用', '2020-12-10 12:02:57', NULL, 0);
INSERT INTO `dict` VALUES (18, 'sex2', '性别', '2021-02-23 15:20:40', '2021-02-23 14:12:04', 0);
INSERT INTO `dict` VALUES (19, 'sex4', '性别', '2021-02-23 15:20:20', '2021-02-23 14:14:29', 0);
COMMIT;

-- ----------------------------
-- Table structure for dict_detail
-- ----------------------------
DROP TABLE IF EXISTS `dict_detail`;
CREATE TABLE `dict_detail` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `label` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字典标签',
  `value` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '字典值',
  `sort` int DEFAULT '0' COMMENT '排序',
  `dict_id` bigint DEFAULT NULL COMMENT '字典id',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime DEFAULT NULL,
  `is_del` tinyint(1) DEFAULT '0',
  `dict_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `FK5tpkputc6d9nboxojdbgnpmyb` (`dict_id`) USING BTREE,
  CONSTRAINT `FK5tpkputc6d9nboxojdbgnpmyb` FOREIGN KEY (`dict_id`) REFERENCES `dict` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='数据字典详情';

-- ----------------------------
-- Records of dict_detail
-- ----------------------------
BEGIN;
INSERT INTO `dict_detail` VALUES (1, '激活', '1', 1, 1, '2021-03-04 13:48:42', NULL, 0, 'user_status');
INSERT INTO `dict_detail` VALUES (2, '禁用', '0', 2, 1, '2021-03-04 13:48:49', NULL, 0, 'user_status');
INSERT INTO `dict_detail` VALUES (3, '启用', '1', 1, 4, '2021-02-25 16:12:46', NULL, 0, 'dept_status');
INSERT INTO `dict_detail` VALUES (4, '停用', '0', 2, 4, '2021-02-25 16:12:51', NULL, 0, 'dept_status');
INSERT INTO `dict_detail` VALUES (5, '启用', '1', 1, 5, '2021-02-26 16:23:39', NULL, 0, 'job_status');
INSERT INTO `dict_detail` VALUES (6, '停用', '0', 2, 5, '2021-02-26 16:23:46', NULL, 0, 'job_status');
INSERT INTO `dict_detail` VALUES (20, '是', '1', 999, 16, '2020-12-09 11:41:30', NULL, 0, 'force_update');
INSERT INTO `dict_detail` VALUES (21, '否', '0', 999, 16, '2020-12-09 11:41:36', NULL, 0, 'force_update');
INSERT INTO `dict_detail` VALUES (22, '是', '1', 999, 17, '2021-02-25 15:32:00', NULL, 0, 'is_enable');
INSERT INTO `dict_detail` VALUES (23, '否', '0', 999, 17, '2020-12-10 12:03:16', NULL, 0, 'is_enable');
COMMIT;

-- ----------------------------
-- Table structure for job
-- ----------------------------
DROP TABLE IF EXISTS `job`;
CREATE TABLE `job` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '岗位名称',
  `enabled` tinyint(1) NOT NULL COMMENT '岗位状态',
  `sort` bigint NOT NULL COMMENT '岗位排序',
  `dept_id` bigint DEFAULT NULL COMMENT '部门ID',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `update_time` datetime DEFAULT NULL,
  `is_del` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `FKmvhj0rogastlctflsxf1d6k3i` (`dept_id`) USING BTREE,
  CONSTRAINT `FKmvhj0rogastlctflsxf1d6k3i` FOREIGN KEY (`dept_id`) REFERENCES `dept` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='岗位';

-- ----------------------------
-- Records of job
-- ----------------------------
BEGIN;
INSERT INTO `job` VALUES (8, '人事专员', 1, 3, 11, '2021-03-07 10:15:49', NULL, 0);
INSERT INTO `job` VALUES (10, '产品经理', 1, 4, 2, '2019-03-29 14:55:51', NULL, 0);
INSERT INTO `job` VALUES (11, '全栈开发2', 1, 2, 2, '2019-03-31 13:39:30', NULL, 0);
INSERT INTO `job` VALUES (12, '软件测试', 1, 5, 2, '2019-03-31 13:39:43', NULL, 0);
INSERT INTO `job` VALUES (13, '6665', 1, 0, 1, '2021-03-02 15:45:16', '2021-03-02 15:42:55', 1);
INSERT INTO `job` VALUES (14, '99', 1, 0, 5, '2021-03-02 15:44:05', '2021-03-02 15:44:05', 1);
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `i_frame` tinyint(1) DEFAULT NULL COMMENT '是否外链',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '菜单名称',
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '组件',
  `pid` bigint NOT NULL COMMENT '上级菜单ID',
  `sort` int NOT NULL COMMENT '排序',
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '图标',
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '链接地址',
  `cache` tinyint(1) DEFAULT '0' COMMENT '缓存',
  `hidden` tinyint(1) DEFAULT '0' COMMENT '是否隐藏',
  `component_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '-' COMMENT '组件名称',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `permission` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '权限',
  `type` int DEFAULT NULL COMMENT '类型',
  `update_time` datetime DEFAULT NULL,
  `is_del` tinyint(1) DEFAULT '0',
  `router` varchar(255) DEFAULT NULL COMMENT '操作的路由',
  `router_method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT '' COMMENT '路由动作',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `FKqcf9gem97gqa5qjm4d3elcqt5` (`pid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=279 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, 0, '系统管理', NULL, 0, 99, 'system', 'system', 0, 0, NULL, '2018-12-18 15:11:29', NULL, 1, '2020-07-16 20:05:34', 0, NULL, NULL);
INSERT INTO `menu` VALUES (2, 0, '用户管理', 'system/user/index', 1, 2, 'peoples', 'user', 0, 0, 'User', '2021-03-06 10:24:34', 'user:list', 1, NULL, 0, '/admin/user', 'get');
INSERT INTO `menu` VALUES (3, 0, '角色管理', 'system/role/index', 1, 3, 'role', 'role', 0, 0, 'Role', '2021-03-06 10:31:55', 'roles:list', 1, NULL, 0, '/admin/roles', 'get');
INSERT INTO `menu` VALUES (5, 0, '菜单管理', 'system/menu/index', 1, 5, 'menu', 'menu', 0, 0, 'Menu', '2021-03-06 10:33:17', 'menu:list', 1, NULL, 0, '/admin/menu', 'get');
INSERT INTO `menu` VALUES (35, 0, '部门管理', 'system/dept/index', 1, 6, 'dept', 'dept', 0, 0, 'Dept', '2021-03-06 10:34:35', 'dept:list', 1, NULL, 0, '/admin/dept', 'get');
INSERT INTO `menu` VALUES (36, 0, '系统工具', '', 0, 101, 'sys-tools', 'sys-tools', 0, 0, NULL, '2019-03-29 10:57:35', NULL, 1, NULL, 0, NULL, NULL);
INSERT INTO `menu` VALUES (37, 0, '岗位管理', 'system/job/index', 1, 7, 'Steve-Jobs', 'job', 0, 0, 'Job', '2021-03-06 10:35:42', 'user:list', 1, NULL, 0, '/admin/job', 'get');
INSERT INTO `menu` VALUES (39, 0, '字典管理', 'system/dict/index', 1, 8, 'dictionary', 'dict', 0, 0, 'Dict', '2021-03-06 12:54:29', 'dict:list', 1, NULL, 0, '/admin/dict', 'get');
INSERT INTO `menu` VALUES (117, 0, '图表库', 'components/Echarts', 10, 50, 'chart', 'echarts', 1, 0, 'Echarts', '2019-11-21 09:04:32', '', 1, NULL, 0, NULL, NULL);
INSERT INTO `menu` VALUES (120, 0, '商品删除', NULL, 45, 4, NULL, NULL, 0, 0, NULL, '2019-12-24 13:03:51', 'YXSTOREPRODUCT_DELETE', 2, '2020-07-10 16:22:51', 0, NULL, NULL);
INSERT INTO `menu` VALUES (123, 0, '后台接口文档', 'tools/swagger/index', 36, 31, 'swagger', 'swagger2', 0, 0, 'Swagger', '2020-01-07 18:05:52', NULL, 1, NULL, 0, NULL, NULL);
INSERT INTO `menu` VALUES (126, 0, '编辑', NULL, 54, 1, NULL, NULL, 0, 0, NULL, '2020-02-14 21:05:28', 'YXSTOREORDER_EDIT', 2, NULL, 0, NULL, NULL);
INSERT INTO `menu` VALUES (127, 0, '用户新增', '', 2, 2, '', '', 0, 0, '', '2021-03-06 10:30:11', 'user:add', 2, NULL, 0, '/admin/user', 'post');
INSERT INTO `menu` VALUES (128, 0, '用户编辑', '', 2, 3, '', '', 0, 0, '', '2021-03-06 10:30:38', 'user:edit', 2, NULL, 0, '/admin/user', 'edit');
INSERT INTO `menu` VALUES (129, 0, '用户删除', '', 2, 4, '', '', 0, 0, '', '2021-03-06 10:30:48', 'user:del', 2, NULL, 0, '/admin/user', 'delete');
INSERT INTO `menu` VALUES (130, 0, '角色创建', '', 3, 2, '', '', 0, 0, '', '2021-03-06 10:32:15', 'roles:add', 2, NULL, 0, '/admin/roles', 'post');
INSERT INTO `menu` VALUES (131, 0, '角色修改', '', 3, 3, '', '', 0, 0, '', '2021-03-06 10:32:27', 'roles:edit', 2, NULL, 0, '/admin/roles', 'put');
INSERT INTO `menu` VALUES (132, 0, '角色删除', '', 3, 999, '', '', 0, 0, '', '2021-03-06 10:32:40', 'roles:del', 2, NULL, 0, '/admin/roles', 'delete');
INSERT INTO `menu` VALUES (133, 0, '菜单新增', '', 5, 2, '', '', 0, 0, '', '2021-03-06 10:33:27', 'menu:add', 2, NULL, 0, '/admin/menu', 'post');
INSERT INTO `menu` VALUES (134, 0, '菜单编辑', '', 5, 3, '', '', 0, 0, '', '2021-03-06 10:33:37', 'menu:edit', 2, NULL, 0, '/admin/menu', 'put');
INSERT INTO `menu` VALUES (135, 0, '菜单删除', '', 5, 4, '', '', 0, 0, '', '2021-03-06 10:33:47', 'menu:del', 2, NULL, 0, '/admin/menu', 'delete');
INSERT INTO `menu` VALUES (136, 0, '部门新增', '', 35, 2, '', '', 0, 0, '', '2021-03-06 10:34:58', 'dept:add', 2, NULL, 0, '/admin/dept', 'post');
INSERT INTO `menu` VALUES (137, 0, '部门编辑', '', 35, 3, '', '', 0, 0, '', '2021-03-06 10:35:07', 'dept:edit', 2, NULL, 0, '/admin/dept', 'put');
INSERT INTO `menu` VALUES (138, 0, '部门删除', '', 35, 4, '', '', 0, 0, '', '2021-03-06 10:35:16', 'dept:del', 2, NULL, 0, '/admin/dept', 'delete');
INSERT INTO `menu` VALUES (139, 0, '岗位新增', '', 37, 2, '', '', 0, 0, '', '2021-03-06 10:35:53', 'job:add', 2, NULL, 0, '/admin/job', 'post');
INSERT INTO `menu` VALUES (140, 0, '岗位编辑', '', 37, 3, '', '', 0, 0, '', '2021-03-06 10:36:02', 'job:edit', 2, NULL, 0, '/admin/job', 'put');
INSERT INTO `menu` VALUES (141, 0, '岗位删除', '', 37, 4, '', '', 0, 0, '', '2021-03-06 10:36:10', 'job:del', 2, NULL, 0, '/admin/job', 'delete');
INSERT INTO `menu` VALUES (142, 0, '字典新增', '', 39, 2, '', '', 0, 0, '', '2021-03-06 10:36:51', 'dict:add', 2, NULL, 0, '/admin/dict', 'post');
INSERT INTO `menu` VALUES (143, 0, '字典编辑', '', 39, 3, '', '', 0, 0, '', '2021-03-06 10:36:59', 'dict:edit', 2, NULL, 0, '/admin/dict', 'put');
INSERT INTO `menu` VALUES (144, 0, '字典删除', '', 39, 4, '', '', 0, 0, '', '2021-03-06 10:37:10', 'dict:del', 2, NULL, 0, '/admin/dict', 'delete');
INSERT INTO `menu` VALUES (184, 0, '新增菜单', NULL, 49, 0, 'add', NULL, 0, 0, NULL, '2020-06-14 20:10:02', 'YxWechatMenu_CREATE', 2, NULL, 0, NULL, NULL);
INSERT INTO `menu` VALUES (265, 0, '顶级2', 'top/aa', 264, 999, 'alipay', 'top2', 0, 0, 'TOp', '2021-03-03 16:09:42', 'top2', 1, '2021-03-03 16:00:39', 0, '/top', 'get');
INSERT INTO `menu` VALUES (269, 0, '字典详情列表', '', 1, 999, 'configure', '/no', 0, 1, '', '2021-03-06 12:59:31', 'dict_detail:list', 1, '2021-03-06 10:40:50', 0, '/admin/dictDetail', 'get');
INSERT INTO `menu` VALUES (270, 0, '字典详情增加', '', 269, 999, '', '', 0, 0, '', '2021-03-06 12:57:17', 'dict_detail:add', 2, '2021-03-06 10:42:14', 0, '/admin/dictDetail', 'post');
INSERT INTO `menu` VALUES (271, 0, '字典详情编辑', '', 269, 999, '', '', 0, 0, '', '2021-03-06 12:57:25', 'dict_detail:edit', 2, '2021-03-06 10:43:33', 0, '/admin/dictDetail', 'put');
INSERT INTO `menu` VALUES (272, 0, '字典详情删除', '', 269, 999, '', '', 0, 0, '', '2021-03-06 12:57:33', 'dict_detail:delete', 2, '2021-03-06 10:44:09', 0, '/admin/dictDetail', 'delete');
INSERT INTO `menu` VALUES (273, 0, '角色菜单数', '', 3, 999, '', '', 0, 0, '', '2021-03-06 10:36:41', 'role_menu:tree', 2, '2021-03-06 10:46:21', 0, '/admin/menu/tree', 'get');
INSERT INTO `menu` VALUES (274, 0, '角色菜单保存', '', 3, 999, '', '', 0, 0, '', '2021-03-06 10:36:41', 'role_menu:save', 2, '2021-03-06 10:49:47', 0, '/admin/roles/menu', 'put');
INSERT INTO `menu` VALUES (275, 0, '单个角色', '', 3, 999, '', '', 0, 0, '', '2021-03-06 10:36:41', 'roles:one', 2, '2021-03-06 11:18:28', 0, '/admin/roles/*', 'get');
INSERT INTO `menu` VALUES (276, 0, '用户个人信息', '', 2, 999, '', '', 0, 0, '', '2021-03-06 10:36:41', 'user:info', 2, '2021-03-06 11:20:25', 0, '/admin/auth/info', 'get');
INSERT INTO `menu` VALUES (277, 0, '用户头像', '', 2, 999, '', '', 0, 0, '', '2021-03-06 10:36:41', 'user:avatar', 2, '2021-03-06 11:21:41', 0, '/admin/user/updateAvatar', 'post');
INSERT INTO `menu` VALUES (278, 0, '用户退出', '', 2, 999, '', '', 0, 0, '', '2021-03-06 11:25:00', 'user:logout', 2, '2021-03-06 11:25:00', 0, '/admin/auth/logout', 'delete');
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
  `data_scope` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '数据权限',
  `level` int DEFAULT NULL COMMENT '角色级别',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `permission` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '功能权限',
  `update_time` datetime DEFAULT NULL,
  `is_del` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='角色表';

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, '超级管理员', '-', '全部', 1, '2018-11-23 11:04:37', 'admin', NULL, 0);
INSERT INTO `role` VALUES (2, '普通用户', '-555', '本级', 2, '2021-03-07 10:09:08', 'common', NULL, 0);
INSERT INTO `role` VALUES (3, '管理员2', '222', '全部', 3, '2020-01-31 16:53:25', '22', NULL, 0);
INSERT INTO `role` VALUES (4, '667', '66', '全部', 3, '2021-03-02 16:16:43', '666', '2021-03-02 16:15:20', 1);
INSERT INTO `role` VALUES (5, '99', '9', '全部', 3, '2021-03-02 16:16:35', '9', '2021-03-02 16:16:35', 1);
COMMIT;

-- ----------------------------
-- Table structure for roles_depts
-- ----------------------------
DROP TABLE IF EXISTS `roles_depts`;
CREATE TABLE `roles_depts` (
  `role_id` bigint NOT NULL,
  `dept_id` bigint NOT NULL,
  `id` bigint NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `FK7qg6itn5ajdoa9h9o78v9ksur` (`dept_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='角色部门关联';

-- ----------------------------
-- Records of roles_depts
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for roles_menus
-- ----------------------------
DROP TABLE IF EXISTS `roles_menus`;
CREATE TABLE `roles_menus` (
  `menu_id` bigint NOT NULL COMMENT '菜单ID',
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `id` bigint NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`),
  KEY `FKcngg2qadojhi3a651a5adkvbq` (`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=472 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='角色菜单关联';

-- ----------------------------
-- Records of roles_menus
-- ----------------------------
BEGIN;
INSERT INTO `roles_menus` VALUES (1, 3, 199);
INSERT INTO `roles_menus` VALUES (2, 3, 200);
INSERT INTO `roles_menus` VALUES (127, 3, 201);
INSERT INTO `roles_menus` VALUES (128, 3, 202);
INSERT INTO `roles_menus` VALUES (129, 3, 203);
INSERT INTO `roles_menus` VALUES (3, 3, 204);
INSERT INTO `roles_menus` VALUES (130, 3, 205);
INSERT INTO `roles_menus` VALUES (131, 3, 206);
INSERT INTO `roles_menus` VALUES (132, 3, 207);
INSERT INTO `roles_menus` VALUES (5, 3, 208);
INSERT INTO `roles_menus` VALUES (133, 3, 209);
INSERT INTO `roles_menus` VALUES (134, 3, 210);
INSERT INTO `roles_menus` VALUES (135, 3, 211);
INSERT INTO `roles_menus` VALUES (35, 3, 212);
INSERT INTO `roles_menus` VALUES (136, 3, 213);
INSERT INTO `roles_menus` VALUES (137, 3, 214);
INSERT INTO `roles_menus` VALUES (138, 3, 215);
INSERT INTO `roles_menus` VALUES (37, 3, 216);
INSERT INTO `roles_menus` VALUES (139, 3, 217);
INSERT INTO `roles_menus` VALUES (140, 3, 218);
INSERT INTO `roles_menus` VALUES (141, 3, 219);
INSERT INTO `roles_menus` VALUES (39, 3, 220);
INSERT INTO `roles_menus` VALUES (142, 3, 221);
INSERT INTO `roles_menus` VALUES (143, 3, 222);
INSERT INTO `roles_menus` VALUES (144, 3, 223);
INSERT INTO `roles_menus` VALUES (1, 1, 411);
INSERT INTO `roles_menus` VALUES (2, 1, 412);
INSERT INTO `roles_menus` VALUES (127, 1, 413);
INSERT INTO `roles_menus` VALUES (128, 1, 414);
INSERT INTO `roles_menus` VALUES (129, 1, 415);
INSERT INTO `roles_menus` VALUES (276, 1, 416);
INSERT INTO `roles_menus` VALUES (277, 1, 417);
INSERT INTO `roles_menus` VALUES (278, 1, 418);
INSERT INTO `roles_menus` VALUES (3, 1, 419);
INSERT INTO `roles_menus` VALUES (130, 1, 420);
INSERT INTO `roles_menus` VALUES (131, 1, 421);
INSERT INTO `roles_menus` VALUES (132, 1, 422);
INSERT INTO `roles_menus` VALUES (273, 1, 423);
INSERT INTO `roles_menus` VALUES (274, 1, 424);
INSERT INTO `roles_menus` VALUES (275, 1, 425);
INSERT INTO `roles_menus` VALUES (5, 1, 426);
INSERT INTO `roles_menus` VALUES (133, 1, 427);
INSERT INTO `roles_menus` VALUES (134, 1, 428);
INSERT INTO `roles_menus` VALUES (135, 1, 429);
INSERT INTO `roles_menus` VALUES (35, 1, 430);
INSERT INTO `roles_menus` VALUES (136, 1, 431);
INSERT INTO `roles_menus` VALUES (137, 1, 432);
INSERT INTO `roles_menus` VALUES (138, 1, 433);
INSERT INTO `roles_menus` VALUES (37, 1, 434);
INSERT INTO `roles_menus` VALUES (139, 1, 435);
INSERT INTO `roles_menus` VALUES (140, 1, 436);
INSERT INTO `roles_menus` VALUES (141, 1, 437);
INSERT INTO `roles_menus` VALUES (39, 1, 438);
INSERT INTO `roles_menus` VALUES (142, 1, 439);
INSERT INTO `roles_menus` VALUES (143, 1, 440);
INSERT INTO `roles_menus` VALUES (144, 1, 441);
INSERT INTO `roles_menus` VALUES (269, 1, 442);
INSERT INTO `roles_menus` VALUES (270, 1, 443);
INSERT INTO `roles_menus` VALUES (271, 1, 444);
INSERT INTO `roles_menus` VALUES (272, 1, 445);
INSERT INTO `roles_menus` VALUES (36, 1, 446);
INSERT INTO `roles_menus` VALUES (123, 1, 447);
INSERT INTO `roles_menus` VALUES (1, 2, 461);
INSERT INTO `roles_menus` VALUES (35, 2, 462);
INSERT INTO `roles_menus` VALUES (37, 2, 463);
INSERT INTO `roles_menus` VALUES (136, 2, 464);
INSERT INTO `roles_menus` VALUES (137, 2, 465);
INSERT INTO `roles_menus` VALUES (139, 2, 466);
INSERT INTO `roles_menus` VALUES (140, 2, 467);
INSERT INTO `roles_menus` VALUES (269, 2, 468);
INSERT INTO `roles_menus` VALUES (270, 2, 469);
INSERT INTO `roles_menus` VALUES (271, 2, 470);
INSERT INTO `roles_menus` VALUES (272, 2, 471);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `avatar` varchar(200) DEFAULT NULL COMMENT '头像',
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '邮箱',
  `enabled` tinyint(1) DEFAULT NULL COMMENT '状态：1启用、0禁用',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '密码',
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '用户名',
  `dept_id` bigint DEFAULT NULL COMMENT '部门名称',
  `phone` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '手机号码',
  `job_id` bigint DEFAULT NULL COMMENT '岗位名称',
  `create_time` datetime DEFAULT NULL COMMENT '创建日期',
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `sex` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `is_del` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE,
  KEY `FK5rwmryny6jthaaxkogownknqp` (`dept_id`) USING BTREE,
  KEY `FKfftoc2abhot8f2wu6cl9a5iky` (`job_id`) USING BTREE,
  KEY `FKpq2dhypk2qgt68nauh2by22jb` (`avatar`) USING BTREE,
  CONSTRAINT `FK5rwmryny6jthaaxkogownknqp` FOREIGN KEY (`dept_id`) REFERENCES `dept` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `FKfftoc2abhot8f2wu6cl9a5iky` FOREIGN KEY (`job_id`) REFERENCES `job` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB AUTO_INCREMENT=9998 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='系统用户';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'https://goapi.yixiang.co/static/upload/1.jpg', 'yshop@qq.com', 1, '$2a$10$fP.426qKaTmix50Oln8L.uav55gELhAd0Eg66Av4oG86u8km7D/Ky', 'admin', 2, '18888888888', 11, '2021-03-07 10:08:30', '管理员', '男', '2020-06-27 12:05:56', 0);
INSERT INTO `user` VALUES (3, '', 'test@yshopnet', 1, '$2a$04$tw0sZ2EOqt7y.tkrnWHdgeyTG6ku7iv0vAACBTBwym5lNr2oQkG.y', 'test', 2, '17777777777', 12, '2021-03-06 12:46:08', '测试2', '男', NULL, 0);
INSERT INTO `user` VALUES (4, NULL, 'test2@qq.com', 1, '$2a$10$IjehtV8MiXb8ni.Qz0wBteE7FjVn49cEcsSj2.ZBUqqHjnC3umSh.', 'test2', 2, '15136175247', 11, '2020-02-15 20:39:16', 'test2', '男', NULL, 0);
INSERT INTO `user` VALUES (5, '', '444@qq.com', 1, '$2a$10$f/VH35NBOBszycV9KEA1HenQ0qVjazDm8LacQU9PO.A4UizFxLMuq', 'qqqqq', 11, '15136175249', 8, '2021-03-04 16:22:56', 'eeeeee', '男', NULL, 0);
INSERT INTO `user` VALUES (6, '', '666@qq.com', 0, '$2a$10$3Vlo24eOmHHW7.3vAjqPSusfIINNo4JiujzxgqsaoWLx/d5de/jEm', '7777', 8, '15136175246', 8, '2021-03-04 15:31:40', '777', '男', NULL, 1);
INSERT INTO `user` VALUES (9992, '', '66@qq.com', 1, '$2a$04$zI7VSH/WspYF9Kea1lCMkO/0pMh.d0U72EO6T0YinOxs.d97dTImi', 'hupeng33333', 7, '18888888888', 8, '2021-03-05 10:42:42', 'hu', '男', '2021-03-04 15:07:36', 1);
INSERT INTO `user` VALUES (9993, '', '44@qq.com', 1, '$2a$04$IBgD8Un9Xgbi0lzbklZlC.uK1srn5w1Y5ntN1qfpzDgcipLLRp1aG', 'hupeng2', 11, '', 10, '2021-03-04 15:28:40', 'hu2', '男', '2021-03-04 15:09:11', 1);
INSERT INTO `user` VALUES (9996, '', '777@qq.com', 0, '$2a$04$Ep/UAK2kmJKPto6efG9qyOYlwQ5NGL0T/PkLVrUcXJxTcq6WG5Zca', 'zhang', 5, '18888888888', 11, '2021-03-04 16:23:49', 'zhang', '男', '2021-03-04 16:14:59', 1);
COMMIT;

-- ----------------------------
-- Table structure for users_roles
-- ----------------------------
DROP TABLE IF EXISTS `users_roles`;
CREATE TABLE `users_roles` (
  `user_id` bigint NOT NULL COMMENT '用户ID',
  `role_id` bigint NOT NULL COMMENT '角色ID',
  `id` bigint NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户角色关联';

-- ----------------------------
-- Records of users_roles
-- ----------------------------
BEGIN;
INSERT INTO `users_roles` VALUES (1, 1, 1);
INSERT INTO `users_roles` VALUES (3, 2, 2);
INSERT INTO `users_roles` VALUES (4, 2, 3);
INSERT INTO `users_roles` VALUES (5, 2, 4);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
