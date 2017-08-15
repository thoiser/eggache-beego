/*
Navicat MySQL Data Transfer

Source Server         : 本地的mysql
Source Server Version : 50547
Source Host           : localhost:3306
Source Database       : eggache

Target Server Type    : MYSQL
Target Server Version : 50547
File Encoding         : 65001

Date: 2017-08-15 17:07:57
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `a_article`
-- ----------------------------
DROP TABLE IF EXISTS `a_article`;
CREATE TABLE `a_article` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `content` varchar(2000) NOT NULL,
  `del` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of a_article
-- ----------------------------
INSERT INTO `a_article` VALUES ('1', '第一篇日记', '第一篇日记的内容。啊1', '0', '2017-06-15 08:53:18');
INSERT INTO `a_article` VALUES ('2', '记！！！', '第2篇日记的内容。啊2', '0', '2017-06-15 08:54:09');
INSERT INTO `a_article` VALUES ('3', '第2篇日记的内容。啊', '第2篇日记的内容。啊', '0', '2017-06-15 08:54:19');
INSERT INTO `a_article` VALUES ('4', '第2篇日记的内容。啊', '第2篇日记的内容。啊', '0', '2017-06-15 08:54:34');
INSERT INTO `a_article` VALUES ('5', '记！！！', '第5篇日记的内容', '0', '2017-06-15 08:54:51');
INSERT INTO `a_article` VALUES ('10', '123222', '1111111111', '0', '2017-08-07 16:02:35');
INSERT INTO `a_article` VALUES ('11', '123', '123', '0', '2017-08-08 16:10:39');
INSERT INTO `a_article` VALUES ('12', '123', '123', '0', '2017-08-08 09:08:42');
INSERT INTO `a_article` VALUES ('13', 'test', '1234567\r\n', '0', '2017-08-08 09:34:36');
INSERT INTO `a_article` VALUES ('14', 'test', '123456\r\n7', '0', '2017-08-08 09:38:29');
INSERT INTO `a_article` VALUES ('15', 'test', '1234561\r\n', '0', '2017-08-08 09:41:39');
INSERT INTO `a_article` VALUES ('16', 'test', '12345612\r\n啊啊啊啊', '0', '2017-08-08 16:24:25');
INSERT INTO `a_article` VALUES ('17', '1111', '2222222', '0', '2017-08-09 13:35:46');

-- ----------------------------
-- Table structure for `a_photo`
-- ----------------------------
DROP TABLE IF EXISTS `a_photo`;
CREATE TABLE `a_photo` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(600) DEFAULT '',
  `title` varchar(255) NOT NULL DEFAULT '未命名',
  `del` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=65 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of a_photo
-- ----------------------------
INSERT INTO `a_photo` VALUES ('63', '15022582033.jpg', '未命名', '0', '2017-08-09 15:12:09');
INSERT INTO `a_photo` VALUES ('64', '15022597022.jpg', 'test', '0', '2017-08-09 15:12:08');

-- ----------------------------
-- Table structure for `a_user`
-- ----------------------------
DROP TABLE IF EXISTS `a_user`;
CREATE TABLE `a_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(60) NOT NULL COMMENT '登录用户名',
  `password` varchar(60) NOT NULL DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of a_user
-- ----------------------------
INSERT INTO `a_user` VALUES ('1', 'thoise', '49ba59abbe56e057');

-- ----------------------------
-- Table structure for `a_video`
-- ----------------------------
DROP TABLE IF EXISTS `a_video`;
CREATE TABLE `a_video` (
  `id` int(20) unsigned NOT NULL AUTO_INCREMENT,
  `url` varchar(600) DEFAULT NULL,
  `title` varchar(255) NOT NULL DEFAULT '',
  `ctime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  `del` tinyint(1) unsigned zerofill NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of a_video
-- ----------------------------
INSERT INTO `a_video` VALUES ('1', 'thoise作品072801.mp4', '嘿嘿', '2017-08-09 16:01:45', '0');
