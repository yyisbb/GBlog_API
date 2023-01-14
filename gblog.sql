/*
 Navicat Premium Data Transfer

 Source Server         : gblog
 Source Server Type    : MySQL
 Source Server Version : 50650 (5.6.50-log)
 Source Host           : 106.52.76.86:3306
 Source Schema         : gblog

 Target Server Type    : MySQL
 Target Server Version : 50650 (5.6.50-log)
 File Encoding         : 65001

 Date: 14/01/2023 20:02:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for articles
-- ----------------------------
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(256) DEFAULT NULL,
  `back_ground` varchar(256) DEFAULT NULL,
  `content` longtext,
  `category_id` bigint(20) DEFAULT NULL,
  `author_name` varchar(256) DEFAULT NULL,
  `avatar` varchar(256) DEFAULT NULL,
  `share_num` bigint(20) DEFAULT NULL,
  `watch_num` bigint(20) DEFAULT NULL,
  `comment_num` bigint(20) DEFAULT NULL,
  `category_name` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_articles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=30 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(256) DEFAULT NULL,
  `description` varchar(256) DEFAULT NULL,
  `banner` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for settings
-- ----------------------------
DROP TABLE IF EXISTS `settings`;
CREATE TABLE `settings` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(256) DEFAULT NULL,
  `logo` varchar(256) DEFAULT NULL,
  `email` varchar(256) DEFAULT NULL,
  `avatar` varchar(256) DEFAULT NULL,
  `author_name` varchar(256) DEFAULT NULL,
  `home_text` varchar(256) DEFAULT NULL,
  `about_content` longtext,
  `github_url` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_settings_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(256) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
