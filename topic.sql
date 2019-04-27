/*
Navicat MariaDB Data Transfer

Source Server         : beeblog
Source Server Version : 100038
Source Host           : localhost:3307
Source Database       : beeblog

Target Server Type    : MariaDB
Target Server Version : 100038
File Encoding         : 65001

Date: 2019-04-27 18:33:28
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic` (
  `Id` int(25) NOT NULL,
  `uid` int(25) NOT NULL,
  `title` varchar(255) NOT NULL,
  `content` varchar(5000) NOT NULL,
  `attachment` varchar(255) NOT NULL,
  `created` datetime(6) NOT NULL,
  `updated` datetime(6) NOT NULL,
  `views` int(25) NOT NULL,
  `author` varchar(255) NOT NULL,
  `reply_time` datetime(6) NOT NULL,
  `reply_count` int(25) NOT NULL,
  `repley_last_user_id` int(25) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of topic
-- ----------------------------
