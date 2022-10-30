/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : couples_subtotal

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 30/10/2022 20:07:13
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for dynamic
-- ----------------------------
DROP TABLE IF EXISTS `dynamic`;
CREATE TABLE `dynamic`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '动态内容',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '动态发布者',
  `status` tinyint(3) NOT NULL DEFAULT 0 COMMENT '（0 双方可见；1 仅自己可见；）',
  `created_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dynamic
-- ----------------------------
INSERT INTO `dynamic` VALUES (1, '大头不听话', 1, 0, '2022-10-30 17:49:33', NULL, 0, NULL);
INSERT INTO `dynamic` VALUES (2, '脑子疼', 1, 0, '2022-10-30 17:50:22', NULL, 0, '2022-10-30 17:50:22');
INSERT INTO `dynamic` VALUES (3, '看医生', 1, 0, '2022-10-29 17:23:02', '2022-10-29 17:29:32', 1, '2022-10-29 17:28:58');
INSERT INTO `dynamic` VALUES (4, '嘎嘎嘎嘎嘎', 1, 1, '2022-10-30 17:41:13', NULL, 0, '2022-10-30 17:47:01');
INSERT INTO `dynamic` VALUES (5, '啦啦啦啦啦', 1, 0, '2022-10-30 17:46:02', NULL, 0, '2022-10-30 17:46:02');
INSERT INTO `dynamic` VALUES (6, '大头大头，下雨不愁', 1, 0, '2022-10-30 17:52:15', NULL, 0, '2022-10-30 17:52:20');
INSERT INTO `dynamic` VALUES (7, '大头凶死了', 1, 0, '2022-10-30 17:53:08', NULL, 0, '2022-10-30 17:53:12');

-- ----------------------------
-- Table structure for lovers_relationship
-- ----------------------------
DROP TABLE IF EXISTS `lovers_relationship`;
CREATE TABLE `lovers_relationship`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `couple_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '情侣Aid',
  `person_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '情侣Bid',
  `memorial_date` datetime NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of lovers_relationship
-- ----------------------------
INSERT INTO `lovers_relationship` VALUES (1, 1, 5, '2021-01-10 00:00:00', '2022-10-30 19:31:42', '2022-10-30 19:31:42', NULL, 0);
INSERT INTO `lovers_relationship` VALUES (28, 1, 5, '2021-01-10 00:00:00', '2022-10-30 19:31:57', '2022-10-30 19:31:57', NULL, 0);

-- ----------------------------
-- Table structure for photo
-- ----------------------------
DROP TABLE IF EXISTS `photo`;
CREATE TABLE `photo`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户id',
  `album_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0 COMMENT '相册id',
  `img_url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '照片路径',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of photo
-- ----------------------------

-- ----------------------------
-- Table structure for photo_album
-- ----------------------------
DROP TABLE IF EXISTS `photo_album`;
CREATE TABLE `photo_album`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '相册名称',
  `owner_id` bigint(20) UNSIGNED NULL DEFAULT 0 COMMENT '拥有者id',
  `type` tinyint(3) NOT NULL COMMENT '(0 情侣相册；1 个人相册)',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of photo_album
-- ----------------------------
INSERT INTO `photo_album` VALUES (1, '1111', 1, 0, '2022-10-27 16:41:00', '2022-10-27 16:41:00', '0000-00-00 00:00:00', 0);
INSERT INTO `photo_album` VALUES (2, '1111', 1, 0, '2022-10-27 16:55:18', '2022-10-27 16:55:18', '0000-00-00 00:00:00', 0);
INSERT INTO `photo_album` VALUES (4, '22222', 1, 1, '2022-10-27 17:19:13', '2022-10-27 17:19:13', NULL, 0);
INSERT INTO `photo_album` VALUES (5, '22222', 1, 1, '2022-10-27 17:21:43', '2022-10-27 17:21:43', NULL, 0);
INSERT INTO `photo_album` VALUES (6, '测试', 1, 1, '2022-10-29 17:45:15', '2022-10-29 17:45:15', NULL, 0);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户登录密码',
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` datetime NOT NULL COMMENT '出生日期',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户邮箱',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户手机号',
  `registration_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '注册码',
  `encrypted_registration` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '加密过后的注册码',
  `header_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `is_deleted` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '13183191657', '123456', '举个栗子', '2001-01-28 03:15:00', 'qiao_shu_rui@163.com', '13183191657', '123456', 'rS5gkx4a/M2KN0QMYkXcQQ==', 'uploads/file/head_A.jpg', NULL, NULL, 0);
INSERT INTO `users` VALUES (2, '15670533573', '123456', 'Duktig', '1998-04-08 00:00:00', 'ren_shi_wei@qq.com', '15670533573', '123456', 'rS5gkx4a/M2KN0QMYkXcQQ==', 'uploads/file/head_B.jpg', NULL, NULL, 0);
INSERT INTO `users` VALUES (3, '13623955429', '3132333435d41d8cd98f00b204e9800998ecf8427e', '', '0000-00-00 00:00:00', '1635513903@qq.com', '13623955429', '425425', 'x5WiK9kkGrVCR8WgpmT9wQ==', '', '2022-10-28 21:03:09', '2022-10-30 17:38:48', 0);
INSERT INTO `users` VALUES (4, '1234', '313233343536d41d8cd98f00b204e9800998ecf8427e', '', '0000-00-00 00:00:00', '1635513903@qq.com', '1234', '788738', 'YW1xiSizDJMNXMtoG2+d0A==', '', '2022-10-30 15:31:03', '2022-10-30 17:37:38', 0);
INSERT INTO `users` VALUES (5, '12345', '313233343536d41d8cd98f00b204e9800998ecf8427e', '测试', '2001-01-28 17:05:35', '1635513903@qq.com', '12345', '161571', 'aQNaXpy6Ukf2G4aXTejyiQ==', '', '2022-10-30 17:06:48', '2022-10-30 17:06:48', 0);

SET FOREIGN_KEY_CHECKS = 1;
