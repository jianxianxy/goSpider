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
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8;

/*Table structure for table `data_price` */

DROP TABLE IF EXISTS `data_price`;

CREATE TABLE `data_price` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(64) DEFAULT NULL COMMENT '小区',
  `signkey` varchar(64) DEFAULT NULL COMMENT '标识',
  `day` date DEFAULT NULL COMMENT '时间',
  `area` varchar(32) DEFAULT NULL COMMENT '面积',
  `price` float DEFAULT NULL COMMENT '价格',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=93 DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB AUTO_INCREMENT=784 DEFAULT CHARSET=utf8;

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
) ENGINE=InnoDB AUTO_INCREMENT=51633 DEFAULT CHARSET=utf8;
