/*
 Navicat Premium Data Transfer

 Source Server         : @localhost
 Source Server Type    : MySQL
 Source Server Version : 80021
 Source Host           : localhost:3306
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 80021
 File Encoding         : 65001

 Date: 11/06/2022 23:01:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论id',
  `video_id` int(0) UNSIGNED NOT NULL COMMENT '评论视频id',
  `user_id` int(0) UNSIGNED NOT NULL COMMENT '评论用户id',
  `content` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
  `create_date` date NOT NULL COMMENT '评论时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `comments`(`user_id`) USING BTREE,
  INDEX `videoIdIncom`(`video_id`) USING BTREE,
  CONSTRAINT `userIdIncom` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `videoIdIncom` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '点赞记录id',
  `video_id` int(0) UNSIGNED NOT NULL COMMENT '被点赞视频id',
  `user_id` int(0) UNSIGNED NOT NULL COMMENT '点赞用户id',
  `created_at` timestamp(3) NOT NULL COMMENT '创建时间',
  `updated_at` timestamp(3) NOT NULL COMMENT '更新时间',
  `favorite_type` int(0) UNSIGNED NOT NULL COMMENT '操作类型:1为点赞,2为取消',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `video_id`(`video_id`, `user_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `userIdInfav` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `videoIdInfav` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for relations
-- ----------------------------
DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户关系id',
  `follow_id` int(0) UNSIGNED NOT NULL COMMENT '被关注人id',
  `follower_id` int(0) UNSIGNED NOT NULL COMMENT '关注人id',
  `created_at` timestamp(3) NOT NULL COMMENT '创建时间',
  `updated_at` timestamp(3) NOT NULL COMMENT '更新时间',
  `type` int(0) UNSIGNED NOT NULL COMMENT '类型1为关注2为取消关注',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `follow_user`(`follow_id`) USING BTREE,
  INDEX `follower_user`(`follower_id`) USING BTREE,
  CONSTRAINT `follow_user` FOREIGN KEY (`follow_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `follower_user` FOREIGN KEY (`follower_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for userinfos
-- ----------------------------
DROP TABLE IF EXISTS `userinfos`;
CREATE TABLE `userinfos`  (
  `user_id` int(0) UNSIGNED NOT NULL COMMENT '用户唯一id',
  `avatar` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户头像',
  `background` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户背景',
  `signature` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户签名',
  PRIMARY KEY (`user_id`) USING BTREE,
  CONSTRAINT `userId` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户密码',
  `created_at` datetime(0) NOT NULL COMMENT '用户注册时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `nameunique`(`name`) USING BTREE COMMENT '保证用户名唯一索引'
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '视频id',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频title',
  `user_id` int(0) UNSIGNED NOT NULL COMMENT '视频作者id',
  `playurl` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '视频地址',
  `coverurl` varchar(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '视频封面',
  `published_at` datetime(0) NULL DEFAULT NULL COMMENT '视频发布时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `userIdInvideo`(`user_id`) USING BTREE,
  CONSTRAINT `userIdInvideo` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
