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

 Date: 29/01/2022 23:14:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '名称',
  `age` int(11) NOT NULL DEFAULT '18' COMMENT '年龄',
  `card_no` varchar(18) NOT NULL DEFAULT '' COMMENT '身份证',
  `head_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `created_at` int(11) NOT NULL COMMENT '创建时间',
  `updated_at` int(11) NOT NULL COMMENT '更新时间',
  `deleted_at` int(11) NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `card_no` (`card_no`),
  UNIQUE KEY `unq_name_card` (`name`,`card_no`),
  KEY `age` (`age`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, 'prince-update', 10, '1', '', 1639411296, 1642339987, 1642337297);
INSERT INTO `users` VALUES (2, 'name02', 0, '2', '', 1639411296, 1639411296, 1639411297);
INSERT INTO `users` VALUES (3, 'name01', 12, '3', '', 1639411296, 1639411296, 0);
INSERT INTO `users` VALUES (4, 'name01', 20, '4', '', 1639411296, 1639411296, 0);
INSERT INTO `users` VALUES (6, 'name03', 18, '5', '', 1639411296, 1639411296, 0);
INSERT INTO `users` VALUES (7, 'insert-prince01', 18, '46000', 'https://dd.xx', 1642342051, 1642342051, 0);
INSERT INTO `users` VALUES (8, 'insert-prince02', 18, '460001', 'https://dd.xx', 1642342268, 1642342268, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
