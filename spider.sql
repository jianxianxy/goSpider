/*
SQLyog Enterprise v12.09 (64 bit)
MySQL - 5.5.48-log : Database - spider_db
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`spider_db` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `spider_db`;

/*Table structure for table `data_chart` */

DROP TABLE IF EXISTS `data_chart`;

CREATE TABLE `data_chart` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `anday` date NOT NULL COMMENT '分析日期',
  `add` int(11) NOT NULL COMMENT '新增条数',
  `reduce` int(11) NOT NULL COMMENT '减少条数',
  `bare` int(11) NOT NULL COMMENT '净差',
  `info` text NOT NULL COMMENT '详细json',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=31 DEFAULT CHARSET=utf8;

/*Table structure for table `data_sale` */

DROP TABLE IF EXISTS `data_sale`;

CREATE TABLE `data_sale` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `anday` date NOT NULL COMMENT '时间',
  `rea_id` int(11) NOT NULL COMMENT '记录ID',
  `onday` date NOT NULL COMMENT '上架时间',
  `offday` date NOT NULL COMMENT '下架时间',
  `showday` int(11) NOT NULL COMMENT '展示时间',
  `name` varchar(128) NOT NULL COMMENT '小区',
  `area` float NOT NULL COMMENT '面积',
  `price` int(11) NOT NULL COMMENT '价格',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8;

/*Table structure for table `realty` */

DROP TABLE IF EXISTS `realty`;

CREATE TABLE `realty` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(64) NOT NULL DEFAULT '' COMMENT '标题',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '小区',
  `pos_1` varchar(64) NOT NULL DEFAULT '' COMMENT '地区',
  `pos_2` varchar(64) NOT NULL DEFAULT '' COMMENT '街道',
  `style` varchar(32) NOT NULL DEFAULT '' COMMENT '格局',
  `area` varchar(32) NOT NULL DEFAULT '' COMMENT '面积',
  `layer` varchar(32) NOT NULL DEFAULT '' COMMENT '楼层',
  `extra` varchar(64) NOT NULL DEFAULT '' COMMENT '其他',
  `price` int(11) NOT NULL DEFAULT '0' COMMENT '价格',
  `price_m2` int(11) NOT NULL DEFAULT '0' COMMENT '每平价格',
  `from` tinyint(4) NOT NULL DEFAULT '1' COMMENT '来源',
  `signkey` varchar(64) NOT NULL DEFAULT '' COMMENT '唯一标识',
  `create` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  `href` varchar(255) DEFAULT '' COMMENT '地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=34124 DEFAULT CHARSET=utf8;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
