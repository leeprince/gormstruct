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

 Date: 08/05/2022 16:33:41
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_base
-- ----------------------------
DROP TABLE IF EXISTS `user_base`;
CREATE TABLE `user_base` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `age` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `age` (`age`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_base
-- ----------------------------
BEGIN;
INSERT INTO `user_base` VALUES (1, 'n001', 1);
INSERT INTO `user_base` VALUES (2, 'p002', 2);
INSERT INTO `user_base` VALUES (3, 'tt-01', 10);
INSERT INTO `user_base` VALUES (4, 'tt-02', 10);
INSERT INTO `user_base` VALUES (5, 'tt-03', 10);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
