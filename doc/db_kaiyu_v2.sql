/*
Navicat MySQL Data Transfer

Source Server         : localhost1
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : db_kaiyu_v2

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2022-05-19 11:40:16
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for hz_ad
-- ----------------------------
DROP TABLE IF EXISTS `hz_ad`;
CREATE TABLE `hz_ad` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `link` varchar(255) NOT NULL DEFAULT '' COMMENT '链接',
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类',
  `img` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_ad
-- ----------------------------
INSERT INTO `hz_ad` VALUES ('1', '2022-05-16 05:39:55', '2022-05-16 05:39:55', '0', '0', '0', '百度', '', '0', '1', '4', '');
INSERT INTO `hz_ad` VALUES ('2', '2022-05-15 21:41:02', '2022-05-16 05:41:02', '0', '0', '0', '百度', 'http://www.baidu.com', '0', '1', '5', '');

-- ----------------------------
-- Table structure for hz_article
-- ----------------------------
DROP TABLE IF EXISTS `hz_article`;
CREATE TABLE `hz_article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `info` varchar(1000) NOT NULL DEFAULT '' COMMENT '简介',
  `content` longtext NOT NULL COMMENT '内容',
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态1 正常 0 禁用',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '类别',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `img` varchar(255) NOT NULL DEFAULT '' COMMENT '图片',
  `memo` varchar(1000) NOT NULL DEFAULT '' COMMENT '扩展',
  `user` varchar(255) NOT NULL DEFAULT '' COMMENT '作者',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_article
-- ----------------------------
INSERT INTO `hz_article` VALUES ('1', '2022-05-16 15:00:34', '2022-05-16 15:00:34', '0', '0', '0', '发布华能泰安光电科技有限公司2021年度社会责任报告》的通知', '热热我', '<p>而我热无热无热无</p>', '0', '12', '0', 'upload/7deac2b1-3cec-9e81-8d50-44c0e04131c2.jpg', '', '');
INSERT INTO `hz_article` VALUES ('2', '2022-05-17 02:31:48', '2022-05-17 02:31:48', '0', '0', '0', '开始了', '34343', '<p>发生的辅导书发多少</p>', '0', '16', '0', 'upload/baf8b95f-a064-ca33-e00e-f8c539005cca.jpg', '', '');
INSERT INTO `hz_article` VALUES ('3', '2022-05-16 18:32:13', '2022-05-17 02:32:13', '0', '0', '0', '提示话', '', '<p>提示化语言5555</p>', '0', '9', '0', 'upload/38fb490c-b8f6-3d8e-ffd2-0a6cc359f270.jpg', '', '');
INSERT INTO `hz_article` VALUES ('4', '2022-05-16 18:39:59', '2022-05-17 02:39:59', '0', '0', '0', '软件资质', '', '<p>343434343454545</p>', '0', '8', '0', 'upload/1ca64b27-ba7a-2077-fa0c-304c2e36b524.jpg', '', '');
INSERT INTO `hz_article` VALUES ('5', '2022-05-18 14:11:12', '2022-05-18 14:11:12', '0', '0', '0', '集成案例', '2323', '<p>23232</p>', '0', '16', '2', '/upload/a30bf1fd-45a9-d09c-8f48-8d86ddb016d1.jpg', '', '');

-- ----------------------------
-- Table structure for hz_category
-- ----------------------------
DROP TABLE IF EXISTS `hz_category`;
CREATE TABLE `hz_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级编号',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '类型',
  `sort` int(11) NOT NULL DEFAULT '0',
  `banner` varchar(800) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_category
-- ----------------------------
INSERT INTO `hz_category` VALUES ('1', '2022-05-15 10:36:10', '2022-05-16 10:36:12', '0', '0', '0', '首页', '0', '0', '0', '[{\"url\":\"upload/70461e0d-e134-727f-5465-9a011866f445.jpg\"},{\"url\":\"upload/32208fa2-9bc3-c431-fb0e-9010c0b87775.jpg\"}]', '/');
INSERT INTO `hz_category` VALUES ('2', '2022-05-16 10:36:26', '2022-05-16 10:36:28', '0', '0', '0', '主营业务', '0', '2', '0', '', '');
INSERT INTO `hz_category` VALUES ('4', '2022-05-16 10:37:41', '2022-05-16 10:37:43', '0', '0', '0', '所有企业链接', '0', '1', '0', '', '');
INSERT INTO `hz_category` VALUES ('5', '2022-05-16 10:39:08', '2022-05-16 10:39:10', '0', '0', '0', '相关行业链接', '0', '1', '0', '', '');
INSERT INTO `hz_category` VALUES ('6', '2022-05-16 10:39:44', '2022-05-16 10:39:46', '0', '0', '0', '新闻媒体链接', '0', '1', '0', '', '');
INSERT INTO `hz_category` VALUES ('7', '2022-05-14 18:40:10', '2022-05-16 10:40:13', '0', '0', '0', '公司资质', '0', '0', '1', '[{\"url\":\"upload/b8131ec5-9b9d-2a89-a29b-fdacc9a4d73a.jpg\"},{\"url\":\"upload/93dbe807-96e8-da21-6b35-b4f5649526de.jpg\"}]', '/quality');
INSERT INTO `hz_category` VALUES ('8', '2022-05-15 18:41:00', '2022-05-16 10:41:12', '0', '0', '0', '商务资质', '7', '0', '0', '', '/quality');
INSERT INTO `hz_category` VALUES ('9', '2022-05-15 18:41:26', '2022-05-16 10:41:29', '0', '0', '0', '技术资质', '7', '0', '0', '', '/quality');
INSERT INTO `hz_category` VALUES ('10', '2022-05-15 18:42:01', '2022-05-16 10:42:04', '0', '0', '0', '自主软件', '7', '0', '0', '', '/quality');
INSERT INTO `hz_category` VALUES ('11', '2022-05-15 10:42:41', '2022-05-16 10:42:43', '0', '0', '0', '客户和案例', '0', '0', '5', '[{\"url\":\"upload/7f529651-24fc-3882-75a9-1515c4e8cc56.jpg\"}]', '/case');
INSERT INTO `hz_category` VALUES ('12', '2022-05-16 02:42:56', '2022-05-16 10:42:59', '0', '0', '0', '主要案例', '11', '0', '0', '', '/case');
INSERT INTO `hz_category` VALUES ('13', '2022-05-15 10:37:07', '2022-05-16 10:37:10', '0', '0', '0', '行业客户', '11', '0', '0', '', '/case?id=13#1');
INSERT INTO `hz_category` VALUES ('14', '2022-05-14 19:53:43', '2022-05-16 03:53:43', '0', '0', '0', '关于开赟', '0', '0', '0', '[{\"url\":\"upload/8ed2687f-e80d-8df1-39fa-b46c902e0348.jpg\"},{\"url\":\"upload/ee610ff5-8924-abb5-38e5-2761197006ae.jpg\"}]', '/about');
INSERT INTO `hz_category` VALUES ('15', '2022-05-16 03:55:16', '2022-05-16 03:55:16', '0', '0', '0', '公司概况', '14', '0', '1', '1', '/about');
INSERT INTO `hz_category` VALUES ('16', '2022-05-15 19:57:02', '2022-05-16 03:57:02', '0', '0', '0', '主营业务', '14', '0', '2', '2', '/about#2');
INSERT INTO `hz_category` VALUES ('17', '2022-05-14 20:02:18', '2022-05-16 04:02:18', '0', '0', '0', '技术服务', '0', '0', '3', '[{\"url\":\"/upload/d10bd77d-8c5e-23c2-dc06-a5a7ae28f356.jpg\"}]', '/service');
INSERT INTO `hz_category` VALUES ('18', '2022-05-16 04:02:52', '2022-05-16 04:02:52', '0', '0', '0', '我们的服务', '17', '0', '0', '0', '/service');
INSERT INTO `hz_category` VALUES ('19', '2022-05-16 04:03:31', '2022-05-16 04:03:31', '0', '0', '0', '项目规划和实施', '17', '0', '1', '', '/service#1');
INSERT INTO `hz_category` VALUES ('20', '2022-05-15 20:04:56', '2022-05-16 04:04:56', '0', '0', '0', '自主产品', '0', '0', '4', '[{\"url\":\"upload/0ea79dca-f515-6f10-226c-1c22c2871b69.jpg\"}]', '/product');
INSERT INTO `hz_category` VALUES ('21', '2022-05-15 20:09:01', '2022-05-16 04:09:01', '0', '0', '0', 'ccLab智能算力平台', '20', '0', '1', '0', '/product');
INSERT INTO `hz_category` VALUES ('22', '2022-05-16 04:10:02', '2022-05-16 04:10:02', '0', '0', '0', 'ccLab Work Flow工作流程管理系统', '20', '0', '2', '', '/product');
INSERT INTO `hz_category` VALUES ('24', '2022-05-15 20:11:53', '2022-05-16 04:11:53', '0', '0', '0', '合作伙伴', '0', '0', '5', '[{\"url\":\"upload/63ed6590-4699-dcdc-5471-c94b23230126.jpg\"}]', '/partner');
INSERT INTO `hz_category` VALUES ('25', '2022-05-15 12:12:14', '2022-05-16 04:12:14', '0', '0', '0', '联系我们', '0', '0', '7', '', '/contact');
INSERT INTO `hz_category` VALUES ('26', '2022-05-16 04:14:07', '2022-05-16 04:14:07', '0', '0', '0', 'ccLab MES系统', '20', '0', '3', '', '/product');

-- ----------------------------
-- Table structure for hz_config
-- ----------------------------
DROP TABLE IF EXISTS `hz_config`;
CREATE TABLE `hz_config` (
  `key` varchar(255) NOT NULL DEFAULT '' COMMENT '键',
  `value` varchar(255) NOT NULL DEFAULT '' COMMENT '值',
  `group` varchar(255) NOT NULL DEFAULT '' COMMENT '分组',
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `type` int(11) NOT NULL DEFAULT '0',
  `description` varchar(255) NOT NULL DEFAULT '',
  `name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_config
-- ----------------------------
INSERT INTO `hz_config` VALUES ('title', '上海莫某公司', 'site', '1', '2022-05-15 10:40:48', '2022-05-14 18:40:51', '0', '0', '0', '0', '', '网站名称');
INSERT INTO `hz_config` VALUES ('tel', '18817710533', 'site', '2', '2022-05-15 10:52:37', '2022-05-14 10:52:39', '0', '0', '0', '0', '', '电话');
INSERT INTO `hz_config` VALUES ('address', '上海浦东新区龙阳路', 'site', '3', '2022-05-15 10:53:48', '2022-05-15 10:53:51', '0', '0', '0', '0', '', '地址');
INSERT INTO `hz_config` VALUES ('dets', '小', 'site', '4', '2022-05-15 15:24:51', '2022-05-15 07:24:54', '0', '0', '0', '1', '大,中,小', '选项');
INSERT INTO `hz_config` VALUES ('orcode_wx', 'upload/82310507-eb33-e932-b136-3ca308dfcdee.jpg', 'site', '5', '2022-05-15 15:32:54', '2022-05-14 15:32:56', '0', '0', '0', '4', '', '公众号');
INSERT INTO `hz_config` VALUES ('orcode_wb', 'upload/2fa684cc-980a-7fe6-af68-6b43c48c479b.jpg', 'site', '6', '2022-05-15 15:33:16', '2022-05-14 15:33:18', '0', '0', '0', '4', '', '微博');
INSERT INTO `hz_config` VALUES ('coordinate', '121.458601,31.27192', 'site', '7', '2022-05-15 07:35:44', '2022-05-15 07:35:46', '0', '0', '0', '0', '', '地图坐标');
INSERT INTO `hz_config` VALUES ('email', '313690636@qq.com', 'site', '8', '2022-05-15 15:36:50', '2022-05-15 07:36:52', '0', '0', '0', '0', '', '邮箱');
INSERT INTO `hz_config` VALUES ('right', '© *** 光电科技有限公司版权所有', 'site', '9', '2022-05-17 06:47:21', '2022-05-17 14:47:24', '0', '0', '0', '0', '', '版权');
INSERT INTO `hz_config` VALUES ('icp', '京ICP备案号***号', 'site', '10', '2022-05-17 06:47:42', '2022-05-17 14:47:44', '0', '0', '0', '0', '', '备案');
INSERT INTO `hz_config` VALUES ('logo', 'upload/7424ef43-e28d-27c6-c050-73c9c2f3ee37.png', 'site', '11', '2022-05-17 06:53:10', '2022-05-17 14:53:13', '0', '0', '0', '4', '', 'logo');

-- ----------------------------
-- Table structure for hz_content
-- ----------------------------
DROP TABLE IF EXISTS `hz_content`;
CREATE TABLE `hz_content` (
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `img` varchar(255) NOT NULL DEFAULT '' COMMENT 'banner图片',
  `top_img` varchar(255) NOT NULL DEFAULT '' COMMENT '导航图片',
  `content` longtext NOT NULL COMMENT '内容',
  `bs` longtext NOT NULL COMMENT '中部业务',
  `video` varchar(255) NOT NULL DEFAULT '' COMMENT '视频地址',
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `code` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_content
-- ----------------------------
INSERT INTO `hz_content` VALUES ('公司概况', 'upload/d452cf78-ae7e-0b94-4d50-56191edccd1c.jpg', 'upload/e3d97894-83f4-6026-70d8-d874a41abcc8.jpg', '上海开赟软件服务有限公司上海开赟软件服务有限公司上海开赟软件服务有限公司<br><br>上海开赟软件服务有限公司上海开赟软件服务有限公司上海开赟软件服务有限公司上海开赟软件服务有限公司上海开赟软件服务有限公司上海开赟软件服务有限公司上海开赟软件服务有限公司上海开赟软件服务有限公司<br><br>上海开赟软件服务有限公司上海开赟软件服务有限公司上海开赟软件服务有限公司上海开赟软件服务有限公司\r\n     ', '[{\"ico\":\"\",\"text\":\"公司注册资本<span>5000</span>万元\",\"id\":0,\"icon\":\"/static/web/images/a1.jpg\"},{\"ico\":\"\",\"text\":\"公司成员<span>60</span>多人\",\"id\":1,\"icon\":\"/static/web/images/a2.jpg\"},{\"ico\":\"\",\"text\":\"营业规模<span>3亿以上</span>（包括下属分子公司）\",\"id\":2,\"icon\":\"/static/web/images/a3.jpg\"},{\"icon\":\"1212\",\"text\":\"1212\",\"id\":4}]', '/static/web/video/mp4', '1', '2022-05-13 14:54:05', '2022-05-15 15:27:33', '0', '0', '0', 'about');

-- ----------------------------
-- Table structure for hz_permission
-- ----------------------------
DROP TABLE IF EXISTS `hz_permission`;
CREATE TABLE `hz_permission` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '权限名',
  `description` varchar(6000) NOT NULL DEFAULT '' COMMENT '描述',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '访问路径',
  `perms` varchar(255) NOT NULL DEFAULT '' COMMENT '权限码',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '图标',
  `parent_id` bigint(20) NOT NULL DEFAULT '0',
  `show` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_permission
-- ----------------------------
INSERT INTO `hz_permission` VALUES ('1', '2022-05-15 12:21:53', '2022-05-15 12:21:56', '0', '0', '0', '系统设置', '系统设置', '', '', '0', '10', 'el-icon-s-tools', '0', '1');
INSERT INTO `hz_permission` VALUES ('2', '2022-05-15 12:26:26', '2022-05-15 12:26:29', '0', '0', '0', '部门管理', '部门管理', '', 'role-list', '1', '0', '', '1', '1');
INSERT INTO `hz_permission` VALUES ('3', '2022-05-15 14:12:08', '2022-05-15 14:12:11', '0', '0', '0', '配置管理', '配置管理', '', 'system-cfg', '1', '2', '', '1', '1');
INSERT INTO `hz_permission` VALUES ('4', '2022-05-15 18:29:01', '2022-05-15 18:29:03', '0', '0', '0', '用户管理', '用户管理', '', 'user-list', '1', '3', '', '1', '1');
INSERT INTO `hz_permission` VALUES ('5', '2022-05-15 20:05:28', '2022-05-15 20:05:31', '0', '0', '0', ' 关于开赟', ' 关于开赟', '', '', '0', '0', 'el-icon-set-up', '0', '1');
INSERT INTO `hz_permission` VALUES ('6', '2022-05-15 20:11:47', '2022-05-15 20:11:49', '0', '0', '0', '公司概况', '公司概况', '', 'content', '1', '1', '', '5', '1');
INSERT INTO `hz_permission` VALUES ('7', '2022-05-15 20:12:36', '2022-05-15 20:12:38', '0', '0', '0', '主营业务', '16', '', 'article-list', '1', '2', '', '5', '1');
INSERT INTO `hz_permission` VALUES ('8', '2022-05-16 09:47:28', '2022-05-16 09:47:30', '0', '0', '0', '客户和案例', '客户和案例', '', '', '0', '0', 'el-icon-collection', '0', '1');
INSERT INTO `hz_permission` VALUES ('9', '2022-05-16 09:50:15', '2022-05-16 09:50:17', '0', '0', '0', '行业客户', '13', '', 'partner-list', '0', '0', '', '8', '1');
INSERT INTO `hz_permission` VALUES ('10', '2022-05-16 09:51:44', '2022-05-16 09:51:47', '0', '0', '0', '主要案例', '12', '', 'article-list', '0', '0', '', '8', '1');
INSERT INTO `hz_permission` VALUES ('11', '2022-05-16 09:52:39', '2022-05-16 09:52:42', '0', '0', '0', '站点导航', '站点导航', '', 'category-list', '0', '0', '', '12', '1');
INSERT INTO `hz_permission` VALUES ('12', '2022-05-16 09:54:26', '2022-05-16 09:54:29', '0', '0', '0', '站点管理', '站点管理', '', 'el-icon-m', '0', '0', 'el-icon-map-location', '0', '1');
INSERT INTO `hz_permission` VALUES ('13', '2022-05-16 09:57:37', '2022-05-16 09:57:39', '0', '0', '0', '友情链接', '1', '', 'ad-list', '0', '0', '', '12', '1');
INSERT INTO `hz_permission` VALUES ('14', '2022-05-16 10:00:47', '2022-05-16 10:00:49', '0', '0', '0', '自主产品', '自主产品', '', '', '0', '0', 'el-icon-sell', '0', '1');
INSERT INTO `hz_permission` VALUES ('15', '2022-05-16 10:02:09', '2022-05-16 10:02:12', '0', '0', '0', '产品管理', '1', '', 'product-list', '0', '0', '', '14', '1');
INSERT INTO `hz_permission` VALUES ('16', '2022-05-16 10:05:44', '2022-05-16 10:05:47', '0', '0', '0', '合作伙伴', '2', '', '', '0', '0', 'el-icon-user', '0', '1');
INSERT INTO `hz_permission` VALUES ('17', '2022-05-16 10:06:29', '2022-05-16 10:06:31', '0', '0', '0', '合作伙伴管理', '24', '', 'partner-list', '0', '0', '', '16', '1');
INSERT INTO `hz_permission` VALUES ('18', '2022-05-16 10:09:06', '2022-05-16 10:09:09', '0', '0', '0', '公司资质', '', '', '', '0', '0', 'el-icon-trophy', '0', '1');
INSERT INTO `hz_permission` VALUES ('19', '2022-05-16 10:09:53', '2022-05-16 10:09:55', '0', '0', '0', '资质管理', '7', '', 'article-list', '0', '0', '', '18', '1');
INSERT INTO `hz_permission` VALUES ('20', '2022-05-18 11:56:47', '2022-05-18 11:56:50', '0', '0', '0', '技术服务', '', '', '', '0', '3', 'el-icon-data-line', '0', '1');
INSERT INTO `hz_permission` VALUES ('21', '2022-05-18 11:58:32', '2022-05-18 11:58:35', '0', '0', '0', '服务管理', '', '', 'service', '0', '0', '', '20', '1');

-- ----------------------------
-- Table structure for hz_product
-- ----------------------------
DROP TABLE IF EXISTS `hz_product`;
CREATE TABLE `hz_product` (
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `content` longtext NOT NULL COMMENT '内容',
  `category_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类 1产品 2 合作伙伴 3 行业客户',
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `info` varchar(255) NOT NULL DEFAULT '',
  `img` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_product
-- ----------------------------
INSERT INTO `hz_product` VALUES ('ccLab智能算力平台', '<div>\r\n                        <span>许可资源模块</span>\r\n                        <em>EDA软件License状态监控与管理，可查看各License的状态、利用率和国企时间等，可导出相应报表，并有过期预警提醒。</em>\r\n                    </div>\r\n                    <div>\r\n                        <span>硬件资源模块</span>\r\n                        <em>集群服务器硬件资源监控与管理，可实时查看集群系统性能，包括集群各节点的CPU、内存、硬盘利用率、I/O负载、网络流量情况等，可导出相应报表。</em>\r\n                    </div>\r\n                    <div>\r\n                        <span>Job监控模块</span>\r\n                        <em>集群人物状态监控与管理，可查看任务调度的统计信息和任务实时状态信息等，可导出相应的报表，并配有异常预警提醒。</em>\r\n                    </div>\r\n                    <div>\r\n                        <span>硬件资源模块</span>\r\n                        <em>用户管理与邮件管理，可对集群用户、系统用户、数据源用户进行管理，并可设置异常预警时的邮件帐户。<br />远程桌面登陆工具，可通过SSH和VNC方式登陆到远程主机。</em>\r\n                    </div>', '21', '1', '2022-05-15 12:09:01', '2022-05-15 20:09:01', '0', '0', '0', '一个基于先进的集群管理、调度和监控技术的分布式系统平台，为有高算力需求的企业提供智能化解决方案，为企业发展赋能。', '/upload/a8d8f6d4-6a5e-eef1-6f09-a346c053a8c1.jpg');
INSERT INTO `hz_product` VALUES ('ccLab Work Flow工作流程管理系统', '<div>\r\n                        <span>自定义工艺流程模块</span>\r\n                        <em>用户可自定义各项目的工艺流程图，并配置相应参数。</em>\r\n                    </div>\r\n                    <div>\r\n                        <span>工艺流程监控管理模块</span>\r\n                        <em>实时监控项目的各个工艺流程状态，设置项目执行的优先级，并可中断、停止和重启工作流。</em>\r\n                    </div>\r\n                    <div>\r\n                        <span>人物监控模块</span>\r\n                        <em>用户可实时查看各项目中任务的状态，并监控是否有异常状况出现。</em>\r\n                    </div>\r\n                    <div>\r\n                        <span>智能数据搜集分析模块</span>\r\n                        <em>帮助用户搜集和分析数据，达到更智能化的业务目标。</em>\r\n                    </div>', '22', '2', '2022-05-15 20:10:02', '2022-05-15 20:10:02', '0', '0', '0', '一个用户可对工艺流程自定义和管理，并可按预设工艺流程执行的软件系统，为制造业企业实施生产过程中任务协调与分解的自动化提供了有效的技术支持', '/upload/b66ab021-d58e-1a8a-b5f3-900a19162d6d.jpg');
INSERT INTO `hz_product` VALUES ('ccLab MES系统', ' <div>\r\n                        <span>生产管理模块</span>\r\n                        <em>根据半导体行业各客户的需求自定义配置工艺流程，<br />并可查询生产计划和进行任务管理等。</em>\r\n                    </div>\r\n                    <div>\r\n                        <span>设备管理模块</span>\r\n                        <em>实现了设备状态实时监控，包括设备的能耗管理、<br />运维点检、异常监控等。</em>\r\n                    </div>\r\n                    <div>\r\n                        <span>质量管理模块</span>\r\n                        <em>用户可自定义质检要求，并进行工序质检、过程质检、不合格管理和质量分析等。</em>\r\n                    </div>\r\n                    <div>\r\n                        <span>生产监控模块</span>\r\n                        <em>实现了生产进度监控和实时库存监控等。</em>\r\n                    </div>', '26', '3', '2022-05-15 20:14:07', '2022-05-15 20:14:07', '0', '0', '0', 'MES系统是一个服务于半导体行业的生产信息化管理系统，有效保证产品的稳定 、有序的生产，提高企业运营效率。', '/upload/466bff25-5c76-727e-7b31-6abc79e96bd2.jpg');
INSERT INTO `hz_product` VALUES ('主机 Server', '[{\"url\":\"upload/ca0fb14a-c363-a588-5c94-7f084655e1c3.png\",\"params\":0},{\"url\":\"upload/890a7566-16a4-807b-e942-89fa01cbbd3d.png\",\"params\":\"0\"}]', '24', '4', '2022-05-12 17:37:06', '2022-05-16 17:37:08', '0', '0', '0', '', '');
INSERT INTO `hz_product` VALUES ('存储 Storage', '[{\"url\":\"upload/c80a5382-aca4-7217-db5c-993fc50489ea.jpg\",\"params\":1}]', '24', '5', '2022-05-12 17:38:53', '2022-05-16 17:38:56', '0', '0', '0', '', '');
INSERT INTO `hz_product` VALUES ('系统网络 Network', '[{\"url\":\"upload/c43947b9-9de7-1e5b-6d99-e8aced5d7768.png\",\"params\":\"2\"}]', '24', '6', '2022-05-13 21:11:19', '2022-05-16 13:11:19', '0', '0', '0', '', '');
INSERT INTO `hz_product` VALUES ('信息安全  Security', '[{\"url\":\"upload/525cc5e3-32f3-93ae-05a6-9d05a797f6f0.png\",\"params\":\"3\"}]', '24', '7', '2022-05-15 05:23:49', '2022-05-16 13:23:49', '0', '0', '0', '', '');
INSERT INTO `hz_product` VALUES ('算力调度 Job Scheduling', '', '24', '12', '2022-05-16 13:48:51', '2022-05-16 13:48:51', '0', '0', '0', '', '');
INSERT INTO `hz_product` VALUES ('云与AI Cloud', '', '24', '13', '2022-05-16 13:48:51', '2022-05-16 13:48:51', '0', '0', '0', '', '');
INSERT INTO `hz_product` VALUES ('软件系统 System', '', '24', '14', '2022-05-16 13:48:51', '2022-05-16 13:48:51', '0', '0', '0', '', '');
INSERT INTO `hz_product` VALUES ('自主软件 Software', '', '24', '15', '2022-05-16 13:48:51', '2022-05-16 13:48:51', '0', '0', '0', '', '');
INSERT INTO `hz_product` VALUES ('行业客户', '[{\"url\":\"upload/133e0d37-47e3-527b-466f-bdaee6eceeed.png\",\"params\":\"0\"}]', '13', '16', '2022-05-16 13:57:23', '2022-05-16 13:57:23', '0', '0', '0', '', '');
INSERT INTO `hz_product` VALUES ('技术服务', '{\"text\":\"是否获得免费的广告覆盖发酒疯了你呢领导风格是否获得免费的广告覆盖发酒疯了你呢领导风格是否获得免费的广告覆盖发酒疯了你呢领导风格\",\"foot_content\":\"北方的方面的辅导辅导水电费那地方<br>北方的方面的辅导辅导水电费那地方<br>北方的方面的辅导辅导水电费那地方<br>北方的方面的辅导辅导水电费那地方<br>北方的方面的辅导辅导水电费那地方<br>北方的方面的辅导辅导水电费那地方<br>北方的方面的辅导辅导水电费那地方<br>北方的方面的辅导辅导水电费那地方\",\"bg\":[{\"url\":\"/upload/459de7ea-67d2-a763-2b69-b10fd5b2d663.jpg\"}],\"foot\":[{\"url\":\"/upload/9e491ef9-c9c2-d137-643f-a34f4b384fdd.jpg\"},{\"url\":\"/upload/8e0d3ea5-80d8-f634-7bcc-574961ccd3b4.jpg\"}],\"content\":\" <ul class=\\\"service_tab\\\">\\n            <li>数据中心</li>\\n            <li>系统集成</li>\\n            <li>运维服务</li>\\n        </ul>\\n        <ul class=\\\"service_list\\\">\\n            <li>\\n                <div><img src=\\\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyNpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTQ4IDc5LjE2NDAzNiwgMjAxOS8wOC8xMy0wMTowNjo1NyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIDIxLjAgKFdpbmRvd3MpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjcwQ0Y0RDVGRDAyMjExRUM4NDdCQUIwNjkxRUI5MTk0IiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjcwQ0Y0RDYwRDAyMjExRUM4NDdCQUIwNjkxRUI5MTk0Ij4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6NzBDRjRENUREMDIyMTFFQzg0N0JBQjA2OTFFQjkxOTQiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6NzBDRjRENUVEMDIyMTFFQzg0N0JBQjA2OTFFQjkxOTQiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz4KWrX1AAAAb1BMVEWtuOj9/v/8/P78/f5edNM+Wco1Ucj7+/5Sas/y9PyotOc3UsisuOjP1vK3wez09vz7/P48VsnR2PMuS8U7VslletX6+v1KY81XbtGAktz3+P1PZ8/Hz/A/WcqsuOmSoeEtSsU+Wcvo6/ksScX////Xw7pGAAAAJXRSTlP///////////////////////////////////////////////8AP89CTwAAAIxJREFUeNpiUCEAGICYXxkX4IUoUMYNBpkCLjkeBmTAgaZAGt3/zIPRF6QrAIsjFCgwwfwnA5NHVSCLCAGwPCO6AjTAroJfgZQKfgVKKvgVKIriVyCuwsQCkufGpYBPRYVRQkVFBDWgBBGeBFOsbGghKQyVF4AoYMYX1GKckoMwPQjhlJeHKMALAAIMANiRiLqGGq0BAAAAAElFTkSuQmCC\\\"><br></div>\\n                <p>需求调研</p>\\n            </li>\\n            <li><span>调研</span></li>\\n            <li>\\n                <div><img src=\\\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyNpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTQ4IDc5LjE2NDAzNiwgMjAxOS8wOC8xMy0wMTowNjo1NyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIDIxLjAgKFdpbmRvd3MpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjRCODcyNDA5RDAyMjExRUM4RkYwQjg3NDM2N0FBRDM1IiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjRCODcyNDBBRDAyMjExRUM4RkYwQjg3NDM2N0FBRDM1Ij4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6NEI4NzI0MDdEMDIyMTFFQzhGRjBCODc0MzY3QUFEMzUiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6NEI4NzI0MDhEMDIyMTFFQzhGRjBCODc0MzY3QUFEMzUiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz4v3feJAAAAP1BMVEXx8/u0v+vt8Pp4i9qAktyXpeM7VsnQ1/I8V8qHmN6ZqOOXpeLIz/Dr7fphd9QyTsf2+P26xOx3itosScX////W13n9AAAAfElEQVR42uST3QqAIAxGTWf2/2N7/2ctQZvBVArvOlf75MCmMoEFRB0BrEdJRKlcZWIBDgLxLkiwD2Hq2oshLTAzvBQWCoYVVgoj30IEtq9DNoE9CIp7KEfvBUlHGlFTmqv9ZrKFLA2p2GtGwf5LgJQAmcVxQMXdzHEKMAC9DUp03LTG0AAAAABJRU5ErkJggg==\\\" style=\\\"max-width:100%;\\\"><br></div>\\n                <p>功能需求分析</p>\\n            </li>\\n            <li><span>访谈</span></li>\\n            <li>\\n                <div><img src=\\\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyNpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTQ4IDc5LjE2NDAzNiwgMjAxOS8wOC8xMy0wMTowNjo1NyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIDIxLjAgKFdpbmRvd3MpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjU5RTBEMEQwRDAyMjExRUNCQTdCQTVFNEI1Qjc3NTk3IiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjU5RTBEMEQxRDAyMjExRUNCQTdCQTVFNEI1Qjc3NTk3Ij4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6NTlFMEQwQ0VEMDIyMTFFQ0JBN0JBNUU0QjVCNzc1OTciIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6NTlFMEQwQ0ZEMDIyMTFFQ0JBN0JBNUU0QjVCNzc1OTciLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz7w8F0sAAAArlBMVEWCk93P1vLy9Pz19vzu8frs7/rh5fc7VsmMnOCJmd5zh9lnfNX8/f7v8fs1UcdXbtFUa9BfddN7jdt7jts3Usja3/U+WcpyhtnR1/M0UMfW2/RacdJtgdcyTsczT8cyT8dofdbT2fMwTMaBk93Y3vVYb9FbctKir+Zpftbn6vlwhNiEld3U2vPo6/k3U8hDXcvS2fPk5/c4U8jk6Pj9/f/S2PPO1fI6VcksScX///+bocrVAAAAOnRSTlP///////////////////////////////////////////////////////////////////////////8AN8D/CgAAANhJREFUeNrkk9cOgkAQRbEgKqDYe++96+X/f8yVYdiwRiW8el42uXOS3dzJau4PtChCCR4Fme4p6ZBQgqkJKgDPgfIrWaLoCTAprkmh6XYNwRgk+O9ISUE/0B2qEKAjppDwUYUKCbZ8wyUkFFBLCWxDCmd/3qOiWL+zsOXECqoGkuGGM1zbHwlcy42nU06unlDFKC3YPWTV1dUraaNPyyq7RV6NL5zo3PC6F1AEnc4sC3koxBY+XTF8E450NkiYo64KmOQEazhU1IzGAyvosUWJE/XzfuUpwAAe4swojVhY9AAAAABJRU5ErkJggg==\\\" style=\\\"max-width:100%;\\\"><br></div>\\n                <p>方案造型</p>\\n            </li>\\n            <li><span>讨论</span></li>\\n            <li>\\n                <div><img src=\\\"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAMAAABEpIrGAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyNpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTQ4IDc5LjE2NDAzNiwgMjAxOS8wOC8xMy0wMTowNjo1NyAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIDIxLjAgKFdpbmRvd3MpIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjY2MEIzNkZBRDAyMjExRUNBQzI0OEI5RDRFQjE5OTlEIiB4bXBNTTpEb2N1bWVudElEPSJ4bXAuZGlkOjY2MEIzNkZCRDAyMjExRUNBQzI0OEI5RDRFQjE5OTlEIj4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6NjYwQjM2RjhEMDIyMTFFQ0FDMjQ4QjlENEVCMTk5OUQiIHN0UmVmOmRvY3VtZW50SUQ9InhtcC5kaWQ6NjYwQjM2RjlEMDIyMTFFQ0FDMjQ4QjlENEVCMTk5OUQiLz4gPC9yZGY6RGVzY3JpcHRpb24+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+IDw/eHBhY2tldCBlbmQ9InIiPz6fTk5qAAAAVFBMVEVVbNHt7/pUa9Dc4fYtSsa4wuu2wOvs7vqzvuq3wez09vy8xe3d4vbq7fnV2/SptegzT8d8jtvu8fvHz/CGl961v+uqtui0vupWbdHg5fcsScX///8gtoscAAAAHHRSTlP///////////////////////////////////8AF7Li1wAAAIVJREFUeNrMk8kOgCAMRCvivu9a//8/BQkHkRaPzqFN4KVNJh04A4JPAKBXYAEkZQBBA3ADiNyIXwHR05/sBbgOeoGq5AFVW9XqOKaBXTWJOJFAwq4YsB/NV+cFpG4HPWG1T7rOHqBJinTZctYH1ijH6uiX98AAEDr7YHDI6Ijv4eV0CTAAu79nGOGZWG8AAAAASUVORK5CYII=\\\" style=\\\"max-width:100%;\\\"><br></div>\\n                <p>概要设计</p>\\n            </li>\\n        </ul>\",\"id\":19}', '17', '19', '2022-05-18 19:37:23', '2022-05-18 05:03:07', '0', '0', '0', '', '');

-- ----------------------------
-- Table structure for hz_role
-- ----------------------------
DROP TABLE IF EXISTS `hz_role`;
CREATE TABLE `hz_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '角色名称',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `code` varchar(255) NOT NULL DEFAULT '',
  `parent_id` int(11) NOT NULL DEFAULT '0',
  `sort` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_role
-- ----------------------------
INSERT INTO `hz_role` VALUES ('1', '2022-05-15 12:13:43', '2022-05-15 10:18:46', '0', '0', '0', '管理员', '管理员', '002232', '0', '44');
INSERT INTO `hz_role` VALUES ('4', '2022-05-15 10:24:58', '2022-05-15 10:25:33', '0', '0', '0', '财务', '财务', '002', '0', '1');

-- ----------------------------
-- Table structure for hz_role_permission
-- ----------------------------
DROP TABLE IF EXISTS `hz_role_permission`;
CREATE TABLE `hz_role_permission` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `permission_id` int(11) NOT NULL DEFAULT '0' COMMENT '权限id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_role_permission
-- ----------------------------
INSERT INTO `hz_role_permission` VALUES ('20', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '18');
INSERT INTO `hz_role_permission` VALUES ('21', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '19');
INSERT INTO `hz_role_permission` VALUES ('22', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '16');
INSERT INTO `hz_role_permission` VALUES ('23', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '17');
INSERT INTO `hz_role_permission` VALUES ('24', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '14');
INSERT INTO `hz_role_permission` VALUES ('25', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '15');
INSERT INTO `hz_role_permission` VALUES ('26', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '12');
INSERT INTO `hz_role_permission` VALUES ('27', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '13');
INSERT INTO `hz_role_permission` VALUES ('28', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '11');
INSERT INTO `hz_role_permission` VALUES ('29', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '8');
INSERT INTO `hz_role_permission` VALUES ('30', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '10');
INSERT INTO `hz_role_permission` VALUES ('31', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '9');
INSERT INTO `hz_role_permission` VALUES ('32', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '5');
INSERT INTO `hz_role_permission` VALUES ('33', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '7');
INSERT INTO `hz_role_permission` VALUES ('34', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '6');
INSERT INTO `hz_role_permission` VALUES ('35', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '1');
INSERT INTO `hz_role_permission` VALUES ('36', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '4');
INSERT INTO `hz_role_permission` VALUES ('37', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '3');
INSERT INTO `hz_role_permission` VALUES ('38', '2022-05-17 05:56:11', '2022-05-17 05:56:11', '0', '0', '0', '1', '2');
INSERT INTO `hz_role_permission` VALUES ('39', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '18');
INSERT INTO `hz_role_permission` VALUES ('40', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '19');
INSERT INTO `hz_role_permission` VALUES ('41', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '16');
INSERT INTO `hz_role_permission` VALUES ('42', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '17');
INSERT INTO `hz_role_permission` VALUES ('43', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '8');
INSERT INTO `hz_role_permission` VALUES ('44', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '10');
INSERT INTO `hz_role_permission` VALUES ('45', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '9');
INSERT INTO `hz_role_permission` VALUES ('46', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '5');
INSERT INTO `hz_role_permission` VALUES ('47', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '7');
INSERT INTO `hz_role_permission` VALUES ('48', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '6');
INSERT INTO `hz_role_permission` VALUES ('49', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '1');
INSERT INTO `hz_role_permission` VALUES ('50', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '4');
INSERT INTO `hz_role_permission` VALUES ('51', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '3');
INSERT INTO `hz_role_permission` VALUES ('52', '2022-05-17 05:56:17', '2022-05-17 05:56:17', '0', '0', '0', '4', '2');

-- ----------------------------
-- Table structure for hz_user
-- ----------------------------
DROP TABLE IF EXISTS `hz_user`;
CREATE TABLE `hz_user` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '账号',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `sex` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
  `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0,1',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `reg_ip` varchar(255) NOT NULL DEFAULT '' COMMENT '注册ip',
  `login_ip` varchar(255) NOT NULL DEFAULT '' COMMENT '登陆ip',
  `role_id` int(11) NOT NULL DEFAULT '0',
  `role_name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_user
-- ----------------------------
INSERT INTO `hz_user` VALUES ('1', '2022-05-10 20:09:25', '2022-05-15 12:00:28', '0', '0', '0', 'admin', 'c33367701511b4f6020ec61ded352059', '0', '管理员', '0', '313690636@qq.com', '', '127.0.0.1', '1', '管理员');
INSERT INTO `hz_user` VALUES ('2', '2022-05-15 11:56:51', '2022-05-15 12:00:20', '0', '0', '0', 'user', '9bb396986c0d74cf0333fdf685f6fb34', '1', '张三', '0', '2@qq.com', '', '', '4', '财务');

-- ----------------------------
-- Table structure for hz_user_log
-- ----------------------------
DROP TABLE IF EXISTS `hz_user_log`;
CREATE TABLE `hz_user_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `user_id` varchar(255) NOT NULL DEFAULT '',
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `agent` varchar(255) NOT NULL DEFAULT '' COMMENT '客户端信息',
  `ip` varchar(255) NOT NULL DEFAULT '' COMMENT '客户端ip',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '请求url',
  `type` int(11) NOT NULL DEFAULT '0' COMMENT '业务类型',
  `param` longtext NOT NULL COMMENT '业务信息',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '操作说明',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_user_log
-- ----------------------------

-- ----------------------------
-- Table structure for hz_user_role
-- ----------------------------
DROP TABLE IF EXISTS `hz_user_role`;
CREATE TABLE `hz_user_role` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  `creator_id` int(11) NOT NULL DEFAULT '0' COMMENT '添加',
  `updater_id` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `delete_at` tinyint(4) NOT NULL DEFAULT '0' COMMENT '软删除状态',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of hz_user_role
-- ----------------------------
