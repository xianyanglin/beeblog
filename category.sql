/*
Navicat MariaDB Data Transfer

Source Server         : beeblog
Source Server Version : 100038
Source Host           : localhost:3307
Source Database       : beeblog

Target Server Type    : MariaDB
Target Server Version : 100038
File Encoding         : 65001

Date: 2019-04-27 18:33:38
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `Id` int(25) NOT NULL,
  `Title` varchar(255) NOT NULL,
  `Created` datetime(6) NOT NULL,
  `Views` int(25) NOT NULL,
  `TopicTime` datetime(6) NOT NULL,
  `TopicCount` int(25) NOT NULL,
  `TopicLastUserId` int(25) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------
