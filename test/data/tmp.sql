/*
 Navicat Premium Data Transfer

 Source Server         : mac-leeprince
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : localhost:3306
 Source Schema         : tmp

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 03/06/2024 00:03:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `expire_time` datetime DEFAULT NULL COMMENT '授权过期时间',
  `access_time` int(11) NOT NULL DEFAULT '0' COMMENT '访问时间',
  `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
  `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user_auth
-- ----------------------------
BEGIN;
INSERT INTO `user_auth` VALUES (1, 1, NULL, 0, NULL, NULL, NULL);
INSERT INTO `user_auth` VALUES (2, 2, '2023-08-31 09:55:34', 0, NULL, NULL, NULL);
INSERT INTO `user_auth` VALUES (3, 3, NULL, 0, '2023-08-30 23:35:08.000', NULL, NULL);
INSERT INTO `user_auth` VALUES (4, 4, NULL, 0, NULL, '2023-08-30 09:56:22.000', NULL);
INSERT INTO `user_auth` VALUES (5, 5, NULL, 0, NULL, NULL, '2023-08-30 09:56:34.000');
INSERT INTO `user_auth` VALUES (6, 6, NULL, 1693387020, NULL, NULL, NULL);
INSERT INTO `user_auth` VALUES (7, 7, NULL, 0, '2023-08-30 23:35:08.001', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for user_base
-- ----------------------------
DROP TABLE IF EXISTS `user_base`;
CREATE TABLE `user_base` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '姓名 \\r 姓名',
  `age` int(11) NOT NULL COMMENT '年龄 \\n 年龄',
  `schoool` varchar(255) NOT NULL COMMENT '学校 \\r\\n 学校',
  `name_1` varchar(255) NOT NULL COMMENT '姓名\n姓名换行',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `age` (`age`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8 COMMENT='用户基础表\n用户基础表换行';

-- ----------------------------
-- Records of user_base
-- ----------------------------
BEGIN;
INSERT INTO `user_base` VALUES (1, 'n001', 1, '', '');
INSERT INTO `user_base` VALUES (2, 'p002', 2, '', '');
INSERT INTO `user_base` VALUES (3, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (4, 'tt-02', 10, '', '');
INSERT INTO `user_base` VALUES (5, 'tt-03', 10, '', '');
INSERT INTO `user_base` VALUES (6, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (7, 'tt-02', 10, '', '');
INSERT INTO `user_base` VALUES (8, 'tt-03', 10, '', '');
INSERT INTO `user_base` VALUES (9, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (10, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (11, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (12, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (13, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (14, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (15, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (16, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (17, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (18, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (19, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (20, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (21, 'tt-0101', 10, '', '');
INSERT INTO `user_base` VALUES (22, 'tt-0101', 10, '', '');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `age` int(11) NOT NULL DEFAULT '18' COMMENT '年龄',
  `card_no` varchar(18) NOT NULL DEFAULT '' COMMENT '身份证',
  `head_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'xxx' COMMENT '头像',
  `school` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `null_no_default` varchar(255) DEFAULT NULL,
  `null_default_nil` varchar(255) DEFAULT NULL,
  `null_default_no_nil` varchar(255) DEFAULT 'x',
  `no_null_no_default` varchar(255) NOT NULL,
  `no_null_default_no_nil` varchar(255) NOT NULL DEFAULT 'y',
  `no_null_default_nil` varchar(255) NOT NULL,
  `created_at` int(11) NOT NULL COMMENT '创建时间',
  `updated_at` int(11) NOT NULL COMMENT '更新时间',
  `deleted_at` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `card_no` (`card_no`),
  UNIQUE KEY `unq_name_card` (`name`,`card_no`),
  KEY `age` (`age`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'insert-prince01', 3, '', 'http://pp.com', '', NULL, NULL, 'x', '', 'y', '', 1639411296, 1701537032, 0);
INSERT INTO `users` VALUES (2, 'insert-prince01', 0, '2', '', '', NULL, NULL, 'x', '', 'y', '', 1639411296, 1701537032, 1642337297);
INSERT INTO `users` VALUES (3, 'insert-prince01', 0, '3', '', '', NULL, NULL, 'x', '', 'y', '', 1639411296, 1701537032, 1642337297);
INSERT INTO `users` VALUES (4, 'name01', 20, '4', '', '', NULL, NULL, 'x', '', 'y', '', 1639411296, 1639411296, 0);
INSERT INTO `users` VALUES (6, 'name03', 18, '5', '', '', NULL, NULL, 'x', '', 'y', '', 1639411296, 1639411296, 0);
INSERT INTO `users` VALUES (7, 'insert-prince01', 18, '46000', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1642342051, 1656520036, 0);
INSERT INTO `users` VALUES (8, 'insert-prince02', 18, '460001', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1642342268, 1642342268, 0);
INSERT INTO `users` VALUES (9, 'insert-prince003', 12, '46100001', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1652112852, 1652112852, 1);
INSERT INTO `users` VALUES (14, 'insert-prince003', 0, '4610000101', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1652113064, 1);
INSERT INTO `users` VALUES (15, 'insert-prince004', 0, '4610000102', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1652113064, 1643399938, 1);
INSERT INTO `users` VALUES (16, 'insert-prince005', 0, '4610000103', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1652113064, 1652113064, 1);
INSERT INTO `users` VALUES (20, 'insert-prince003', 0, '46100001001', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1652113511, 1);
INSERT INTO `users` VALUES (21, 'insert-prince004', 0, '46100001002', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1652113511, 1643399938, 1);
INSERT INTO `users` VALUES (22, 'insert-prince005', 0, '46100001003', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1652113511, 1652113511, 1);
INSERT INTO `users` VALUES (23, 'insert-prince005', 0, '46100001004', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1652113511, 1652113511, 0);
INSERT INTO `users` VALUES (24, 'insert-prince003', 0, '461000010001', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1652113835, 1);
INSERT INTO `users` VALUES (25, 'insert-prince004', 0, '461000010002', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1652113835, 1643399938, 1);
INSERT INTO `users` VALUES (26, 'insert-prince005', 0, '461000010003', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1652113835, 1652113835, 1);
INSERT INTO `users` VALUES (27, 'insert-prince005', 0, '461000010004', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1652113835, 1652113835, 0);
INSERT INTO `users` VALUES (32, 'insert-prince003', 0, '4610000100101', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1654848152, 1);
INSERT INTO `users` VALUES (33, 'insert-prince004', 0, '4610000100102', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1654848152, 1643399938, 1);
INSERT INTO `users` VALUES (34, 'insert-prince005', 0, '4610000100103', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1654848152, 1654848152, 1);
INSERT INTO `users` VALUES (35, 'insert-prince005', 0, '4610000100104', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1654848152, 1654848152, 0);
INSERT INTO `users` VALUES (40, 'insert-prince1', 0, '46100101', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1654927732, 1);
INSERT INTO `users` VALUES (41, 'insert-prince1', 0, '46100102', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1654927935, 0);
INSERT INTO `users` VALUES (42, 'insert-prince1', 0, '46100103', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1654928048, 0);
INSERT INTO `users` VALUES (45, 'insert-prince1', 18, '46100104', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1654928145, 0);
INSERT INTO `users` VALUES (50, 'insert-prince2', 18, '32131da40', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 164339993, 1710782761, 0);
INSERT INTO `users` VALUES (53, 'insert-prince2', 18, '46100106', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1654929571, 0);
INSERT INTO `users` VALUES (56, 'insert-prince2', 18, '46100206', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1655652245, 0);
INSERT INTO `users` VALUES (57, 'insert-prince2', 18, '46100207', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1655652561, 0);
INSERT INTO `users` VALUES (58, 'insert-prince2', 18, '46100208', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1655652575, 0);
INSERT INTO `users` VALUES (59, 'insert-prince2', 18, '46100209', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1655652767, 0);
INSERT INTO `users` VALUES (60, 'insert-prince2', 18, '46100210', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1655652822, 0);
INSERT INTO `users` VALUES (63, 'insert-prince2', 18, '46100211', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1655652868, 0);
INSERT INTO `users` VALUES (64, 'insert-prince2', 18, '46100212', 'https://dd.xx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1655652924, 0);
INSERT INTO `users` VALUES (65, 'insert-prince2', 18, '22134d433', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1710211002, 0);
INSERT INTO `users` VALUES (70, 'insert-prince2', 18, '32134d433', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1710211444, 0);
INSERT INTO `users` VALUES (73, 'insert-prince2', 19, '32134d434', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1710211475, 0);
INSERT INTO `users` VALUES (77, 'insert-prince2', 19, '32134d435', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1710211519, 0);
INSERT INTO `users` VALUES (81, 'insert-prince2', 19, '32134d436', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 1643399938, 1710211612, 0);
INSERT INTO `users` VALUES (88, 'insert-prince2', 18, '32134d437', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 164339993, 1710211690, 0);
INSERT INTO `users` VALUES (89, 'insert-prince2', 18, '32134d447', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 164339993, 1710211790, 0);
INSERT INTO `users` VALUES (93, '', 18, '32134d448', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 164339993, 1710211981, 0);
INSERT INTO `users` VALUES (97, '', 19, '32134d449', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 164339993, 1710212063, 0);
INSERT INTO `users` VALUES (101, '', 19, '32134d440', 'xxx', '', NULL, NULL, 'x', '', 'y', '', 164339993, 1710212173, 0);
INSERT INTO `users` VALUES (102, '', 19, '32134da40', '', '', NULL, NULL, 'x', '', 'y', '', 164339993, 1710212711, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
