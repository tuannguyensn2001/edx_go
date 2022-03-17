/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 50733
 Source Host           : localhost:3306
 Source Schema         : edx

 Target Server Type    : MySQL
 Target Server Version : 50733
 File Encoding         : 65001

 Date: 17/03/2022 08:11:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for chapters
-- ----------------------------
DROP TABLE IF EXISTS `chapters`;
CREATE TABLE `chapters`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `course_id` int(11) NULL DEFAULT NULL,
  `order` int(11) NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of chapters
-- ----------------------------
INSERT INTO `chapters` VALUES (1, 'hello123', 1, 1, '2022-03-15 00:05:40', '2022-03-15 22:31:11', '2022-03-15 22:53:42');
INSERT INTO `chapters` VALUES (2, 'Tìm hiểu JS', 1, 2, '2022-03-15 00:05:41', '2022-03-15 00:05:41', '2022-03-15 23:14:01');
INSERT INTO `chapters` VALUES (3, 'Tìm hiểu JS', 1, 3, '2022-03-15 00:05:42', '2022-03-15 00:05:42', '2022-03-15 23:12:23');
INSERT INTO `chapters` VALUES (4, 'Tìm hiểu JS', 1, 4, '2022-03-15 00:05:43', '2022-03-15 00:05:43', '2022-03-15 23:03:09');
INSERT INTO `chapters` VALUES (5, 'Tìm hiểu JS', 1, 5, '2022-03-15 00:05:44', '2022-03-15 00:05:44', '2022-03-15 23:14:04');
INSERT INTO `chapters` VALUES (6, 'Tìm hiểu JS123', 1, 6, '2022-03-15 00:05:44', '2022-03-15 22:38:47', '2022-03-15 23:03:05');
INSERT INTO `chapters` VALUES (7, 'abc', 1, 7, '2022-03-15 22:13:50', '2022-03-15 22:13:50', '2022-03-15 23:12:16');
INSERT INTO `chapters` VALUES (8, 'abc123', 1, 8, '2022-03-15 22:30:11', '2022-03-15 22:30:11', '2022-03-15 23:03:02');
INSERT INTO `chapters` VALUES (9, 'mk', 1, 1, '2022-03-15 23:14:10', '2022-03-15 23:14:10', NULL);
INSERT INTO `chapters` VALUES (10, 'hello', 1, 2, '2022-03-16 19:28:38', '2022-03-16 19:28:38', NULL);

-- ----------------------------
-- Table structure for courses
-- ----------------------------
DROP TABLE IF EXISTS `courses`;
CREATE TABLE `courses`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `imageUrl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `price` decimal(10, 2) NULL DEFAULT NULL,
  `status` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_id` int(11) NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of courses
-- ----------------------------
INSERT INTO `courses` VALUES (1, 'Học javascript', 'mô tả có 1 tí', 'https://wiki.tino.org/wp-content/uploads/2021/09/pasted-image-0.png', 10001.00, 'ACTIVE', 1, '2022-03-14 23:09:29', '2022-03-14 23:09:29', NULL);
INSERT INTO `courses` VALUES (2, 'Học php', 'mô tả có 1 tí', 'https://wiki.tino.org/wp-content/uploads/2021/09/pasted-image-0.png', 10001.00, 'ACTIVE', 1, '2022-03-14 23:15:18', '2022-03-14 23:15:18', NULL);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'tuannguyensn2001', 'devpro2001', 'java2001', '2022-03-13 23:47:51', '2022-03-13 23:47:51');
INSERT INTO `users` VALUES (2, 'tuannguyensn2001', 'devpro2001', 'java2001', '2022-03-13 23:48:35', '2022-03-13 23:48:35');
INSERT INTO `users` VALUES (3, 'tuannguyensn2001', 'devpro2001', 'java2001', '2022-03-13 23:53:36', '2022-03-13 23:53:36');
INSERT INTO `users` VALUES (4, 'tuannguyensn2001', 'devpro2001', 'java2001', '2022-03-13 23:53:37', '2022-03-13 23:53:37');
INSERT INTO `users` VALUES (5, 'tuannguyensn2001', 'devpro2001', 'java2001', '2022-03-13 23:53:38', '2022-03-13 23:53:38');
INSERT INTO `users` VALUES (6, 'tuannguyensn2001', 'devpro2001', 'java2001', '2022-03-13 23:53:39', '2022-03-13 23:53:39');
INSERT INTO `users` VALUES (7, 'tuannguyensn2001', 'devpro2001@gmail.com', 'java2001', '2022-03-13 23:59:27', '2022-03-13 23:59:27');

SET FOREIGN_KEY_CHECKS = 1;
