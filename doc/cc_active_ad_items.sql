/*
Navicat MySQL Data Transfer

Source Server         : apiadmin
Source Server Version : 50536
Source Host           : 192.168.134.128:3306
Source Database       : dbshop

Target Server Type    : MYSQL
Target Server Version : 50536
File Encoding         : 65001

Date: 2018-01-20 18:10:18
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `cc_active_ad_items`
-- ----------------------------
DROP TABLE IF EXISTS `cc_active_ad_items`;
CREATE TABLE `cc_active_ad_items` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `ad_id` char(32) NOT NULL DEFAULT '' COMMENT '推广ID',
  `ad_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '推广价格(以分为单位 程序需要/100)',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '推广状态：@0：正常  @1：首页固定展示 @2：下架 @3:推广过期 @4:库存不足 @5:测试商品',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modify_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_ad_id` (`ad_id`) USING BTREE,
  KEY `idx_ad_price` (`ad_price`) USING BTREE,
  KEY `idx_status_app_id` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=52268 DEFAULT CHARSET=utf8 COMMENT='推广期中的推广物料';

-- ----------------------------
-- Records of cc_active_ad_items
-- ----------------------------
INSERT INTO `cc_active_ad_items` VALUES ('35057', '170527114639001KI666666825551515', '17', '5', '1495856820', '1495856820');
INSERT INTO `cc_active_ad_items` VALUES ('52228', '1705271351200016R666666897985651', '20000', '0', '1495864294', '1495864294');
INSERT INTO `cc_active_ad_items` VALUES ('52229', '1705271353490016R666666865110248', '10001', '0', '1495864441', '1495864441');
INSERT INTO `cc_active_ad_items` VALUES ('52230', '1705271353270016R666666849555457', '1', '0', '1495864447', '1495864447');
INSERT INTO `cc_active_ad_items` VALUES ('52231', '1705271352590016R666666810057485', '312', '0', '1495864465', '1495864465');
INSERT INTO `cc_active_ad_items` VALUES ('52232', '170527114439001KI666666897499997', '2', '5', '1495865999', '1495865999');
INSERT INTO `cc_active_ad_items` VALUES ('52233', '170527142515001KI666666857545798', '16', '3', '1495866382', '1495987199');
INSERT INTO `cc_active_ad_items` VALUES ('52234', '170527142602001KI666666856554849', '15', '5', '1495866384', '1495866384');
INSERT INTO `cc_active_ad_items` VALUES ('52235', '1705311400000016R666666851525310', '10000', '0', '1496210425', '1496210425');
INSERT INTO `cc_active_ad_items` VALUES ('52236', '1706011152080016R666666805449489', '312', '0', '1496289353', '1496289353');
INSERT INTO `cc_active_ad_items` VALUES ('52248', '1706011428470016R666666895048559', '318', '0', '1496298692', '1496298692');
INSERT INTO `cc_active_ad_items` VALUES ('52249', '1706011429000016R666666810010054', '399', '0', '1496298706', '1496298706');
INSERT INTO `cc_active_ad_items` VALUES ('52250', '1706011429110016R666666805549521', '322', '0', '1496298716', '1496298716');
INSERT INTO `cc_active_ad_items` VALUES ('52251', '1706011429290016R666666850100539', '311', '0', '1496298724', '1496298724');
INSERT INTO `cc_active_ad_items` VALUES ('52252', '1706011429420016R666666848539753', '333', '0', '1496299051', '1496299051');
INSERT INTO `cc_active_ad_items` VALUES ('52253', '1706011430210016R666666869755100', '314', '0', '1496299064', '1496299064');
INSERT INTO `cc_active_ad_items` VALUES ('52254', '1706011430400016R666666852501015', '313', '0', '1496299081', '1496299081');
INSERT INTO `cc_active_ad_items` VALUES ('52255', '1706011513250016R666666810298505', '102', '0', '1496301219', '1496301219');
INSERT INTO `cc_active_ad_items` VALUES ('52256', '1706011512260016R666666824999485', '3', '0', '1496301225', '1496301225');
INSERT INTO `cc_active_ad_items` VALUES ('52257', '1706011512130016R666666869750564', '3', '0', '1496301230', '1496301230');
INSERT INTO `cc_active_ad_items` VALUES ('52258', '1706011512030016R666666852995457', '4', '0', '1496301236', '1496301236');
INSERT INTO `cc_active_ad_items` VALUES ('52259', '1706011511470016R666666849549848', '5', '0', '1496301241', '1496301241');
INSERT INTO `cc_active_ad_items` VALUES ('52260', '1706011511300016R666666850489710', '60', '0', '1496301247', '1496301247');
INSERT INTO `cc_active_ad_items` VALUES ('52261', '1706011511180016R666666850575210', '30', '0', '1496301253', '1496301253');
INSERT INTO `cc_active_ad_items` VALUES ('52262', '1706011510540016R666666895755101', '8', '0', '1496301266', '1496301266');
INSERT INTO `cc_active_ad_items` VALUES ('52263', '1706011535570016R666666889854102', '30', '0', '1496302605', '1496302605');
INSERT INTO `cc_active_ad_items` VALUES ('52264', '1706011535460016R666666810110053', '20', '0', '1496302611', '1496302611');
INSERT INTO `cc_active_ad_items` VALUES ('52265', '1706011555310016R666666809953985', '11', '0', '1496303743', '1496303743');
INSERT INTO `cc_active_ad_items` VALUES ('52266', '1706011555210016R666666851481004', '13', '0', '1496303748', '1496303748');
INSERT INTO `cc_active_ad_items` VALUES ('52267', '1706011714510016R666666856545710', '1009', '0', '1496308498', '1496308498');
