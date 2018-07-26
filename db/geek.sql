/*
MySQL Backup
Source Server Version: 5.5.16
Source Database: gk
Date: 2018/7/26 23:06:55
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
--  Table structure for `Category`
-- ----------------------------
DROP TABLE IF EXISTS `Category`;
CREATE TABLE `Category` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Name` varchar(64) NOT NULL,
  `ParentId` bigint(20) DEFAULT NULL,
  `CreateAt` datetime DEFAULT NULL,
  `ModifyAt` datetime DEFAULT NULL,
  `ItemCount` bigint(20) DEFAULT '0',
  `SiteId` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `SiteId` (`SiteId`)
) ENGINE=MyISAM AUTO_INCREMENT=928 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
--  Table structure for `Fav`
-- ----------------------------
DROP TABLE IF EXISTS `Fav`;
CREATE TABLE `Fav` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` bigint(20) DEFAULT NULL,
  `EntityId` bigint(20) DEFAULT NULL,
  `CreateAt` datetime DEFAULT NULL,
  `SiteId` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `FavType` (`UserId`,`EntityId`,`SiteId`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
--  Table structure for `Follow`
-- ----------------------------
DROP TABLE IF EXISTS `Follow`;
CREATE TABLE `Follow` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` bigint(20) DEFAULT NULL,
  `FollowId` bigint(20) DEFAULT NULL,
  `CreateAt` datetime DEFAULT NULL,
  `SiteId` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UserId` (`UserId`,`FollowId`,`SiteId`)
) ENGINE=MyISAM AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `GkToken`
-- ----------------------------
DROP TABLE IF EXISTS `GkToken`;
CREATE TABLE `GkToken` (
  `UserId` bigint(20) NOT NULL,
  `Ua` varchar(8) NOT NULL,
  `Token` varchar(64) NOT NULL,
  `CreateAt` datetime NOT NULL,
  `SiteId` bigint(20) DEFAULT '0',
  PRIMARY KEY (`UserId`),
  UNIQUE KEY `Token` (`Token`) USING BTREE,
  UNIQUE KEY `UserId` (`UserId`,`Ua`,`SiteId`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- ----------------------------
--  Table structure for `GkUser`
-- ----------------------------
DROP TABLE IF EXISTS `GkUser`;
CREATE TABLE `GkUser` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Username` varchar(64) NOT NULL COMMENT '用户名',
  `Password` varchar(64) NOT NULL,
  `Email` varchar(64) NOT NULL COMMENT '邮箱',
  `Mobile` varchar(16) DEFAULT NULL COMMENT '手机号',
  `Weixin` varchar(160) DEFAULT NULL COMMENT '微信',
  `PasswordError` int(11) NOT NULL DEFAULT '0' COMMENT '密码错误次数',
  `CreateAt` datetime NOT NULL COMMENT '注册时间',
  `ModifyAt` datetime NOT NULL,
  `LastErrorTime` datetime DEFAULT NULL COMMENT '密码错误时间',
  `LoginDate` datetime DEFAULT NULL COMMENT '最后登录时间',
  `Score` bigint(20) NOT NULL DEFAULT '0' COMMENT '积分',
  `EditType` varchar(8) DEFAULT 'HTML' COMMENT ' 编辑器类型',
  `MsgCount` bigint(20) NOT NULL DEFAULT '0' COMMENT '消息总数',
  `FavCount` bigint(20) NOT NULL DEFAULT '0',
  `TopicCount` bigint(20) NOT NULL DEFAULT '0',
  `ReplyCount` bigint(20) NOT NULL DEFAULT '0',
  `FollowCount` int(11) NOT NULL DEFAULT '0' COMMENT '关注总人数',
  `LoginAward` smallint(6) DEFAULT '1' COMMENT '登录奖励',
  `Info` text,
  `Privacy` bigint(20) NOT NULL DEFAULT '0' COMMENT '隐私',
  `Dau` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '日刷量',
  `SiteId` int(11) NOT NULL DEFAULT '0',
  `DauCount` bigint(20) DEFAULT '0' COMMENT '总活跃度',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Email` (`Email`,`SiteId`) USING BTREE,
  UNIQUE KEY `Username` (`Username`,`SiteId`) USING BTREE,
  KEY `Dau` (`Dau`,`SiteId`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=11027 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
--  Table structure for `KeyVal`
-- ----------------------------
DROP TABLE IF EXISTS `KeyVal`;
CREATE TABLE `KeyVal` (
  `KeyName` varchar(160) NOT NULL,
  `Val` text,
  `SiteId` int(11) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`KeyName`,`SiteId`),
  UNIQUE KEY `KeyName` (`KeyName`,`SiteId`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `Msg`
-- ----------------------------
DROP TABLE IF EXISTS `Msg`;
CREATE TABLE `Msg` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` bigint(20) NOT NULL,
  `FromUserId` bigint(20) NOT NULL,
  `Title` text,
  `Body` text,
  `EntityId` varchar(64) DEFAULT NULL COMMENT '关联Id',
  `MsgType` varchar(64) DEFAULT NULL,
  `SiteId` int(11) DEFAULT '0',
  `Status` int(11) NOT NULL DEFAULT '0',
  `CreateAt` datetime DEFAULT NULL,
  `GroupId` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '消息组ID，for topic',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UserId_2` (`UserId`,`EntityId`,`SiteId`) USING BTREE,
  KEY `UserId` (`UserId`,`SiteId`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `NoteCategory`
-- ----------------------------
DROP TABLE IF EXISTS `NoteCategory`;
CREATE TABLE `NoteCategory` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `Name` varchar(64) NOT NULL,
  `ParentId` bigint(20) NOT NULL,
  `CreateAt` datetime DEFAULT NULL,
  `ModifyAt` datetime DEFAULT NULL,
  `ItemCount` bigint(20) NOT NULL DEFAULT '0',
  `UserId` bigint(20) NOT NULL,
  `SiteId` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM AUTO_INCREMENT=911 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
--  Table structure for `Notes`
-- ----------------------------
DROP TABLE IF EXISTS `Notes`;
CREATE TABLE `Notes` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` bigint(20) DEFAULT NULL,
  `CategoryId` bigint(20) DEFAULT NULL,
  `Title` varchar(240) DEFAULT NULL,
  `Body` text,
  `CreateAt` datetime DEFAULT NULL,
  `SourceType` varchar(8) NOT NULL DEFAULT '0' COMMENT '源码类型，1--md 2--html 3--text',
  `Source` text,
  `SiteId` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
--  Table structure for `ReplyThank`
-- ----------------------------
DROP TABLE IF EXISTS `ReplyThank`;
CREATE TABLE `ReplyThank` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` bigint(20) DEFAULT NULL,
  `SiteId` bigint(20) DEFAULT NULL,
  `ReplyId` bigint(20) DEFAULT NULL,
  `CreateAt` datetime DEFAULT NULL,
  `TopicId` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UserId` (`UserId`,`SiteId`,`ReplyId`),
  KEY `UserId_2` (`UserId`,`SiteId`,`TopicId`)
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `ScoreLog`
-- ----------------------------
DROP TABLE IF EXISTS `ScoreLog`;
CREATE TABLE `ScoreLog` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` bigint(20) NOT NULL,
  `Score` int(11) NOT NULL COMMENT '余额',
  `ScoreType` varchar(32) NOT NULL COMMENT '积分类型',
  `ScoreDesc` varchar(240) DEFAULT NULL,
  `EntityId` varchar(64) NOT NULL,
  `CreateAt` datetime DEFAULT NULL,
  `Fee` int(11) NOT NULL DEFAULT '0',
  `SiteId` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `UserId` (`UserId`,`EntityId`,`ScoreType`,`SiteId`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=188 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `Topic`
-- ----------------------------
DROP TABLE IF EXISTS `Topic`;
CREATE TABLE `Topic` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `UserId` bigint(20) NOT NULL,
  `Username` varchar(64) NOT NULL,
  `CategoryId` bigint(20) NOT NULL,
  `Title` varchar(240) NOT NULL,
  `Body` text,
  `CreateAt` datetime NOT NULL,
  `ModifyAt` datetime DEFAULT NULL,
  `ReplyCount` bigint(20) DEFAULT '0' COMMENT '回复总数',
  `ReplyUsername` varchar(64) DEFAULT '' COMMENT '最后回复',
  `ReplyUserId` bigint(20) DEFAULT '0',
  `ReplyTime` datetime DEFAULT NULL,
  `Source` text,
  `SourceType` varchar(8) NOT NULL DEFAULT 'Markdown' COMMENT '源码类型，1--md 2--html 3--text',
  `ShowTimes` bigint(20) NOT NULL DEFAULT '0' COMMENT '展现次数',
  `Today` int(11) DEFAULT '0' COMMENT '今日点击次数',
  `SiteId` int(11) DEFAULT '1',
  `Audit` int(11) DEFAULT '0' COMMENT '审核 0--未审核 1--已读 2--禁止',
  PRIMARY KEY (`Id`),
  KEY `CreateAt` (`CreateAt`) USING BTREE,
  KEY `ReplyTime` (`ReplyTime`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
--  Table structure for `TopicAppend`
-- ----------------------------
DROP TABLE IF EXISTS `TopicAppend`;
CREATE TABLE `TopicAppend` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `TopicId` bigint(20) NOT NULL,
  `AppendText` text NOT NULL,
  `CreateAt` datetime NOT NULL,
  `SiteId` bigint(20) NOT NULL,
  PRIMARY KEY (`Id`),
  KEY `TopicId` (`TopicId`,`SiteId`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `TopicReferer`
-- ----------------------------
DROP TABLE IF EXISTS `TopicReferer`;
CREATE TABLE `TopicReferer` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `TopicId` bigint(20) NOT NULL,
  `Referer` varchar(240) NOT NULL,
  `Click` bigint(20) NOT NULL,
  `CreateAt` datetime DEFAULT NULL,
  `ModifyAt` datetime DEFAULT NULL,
  `ClickAll` bigint(20) NOT NULL,
  `SiteId` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `TopicId` (`TopicId`,`Referer`)
) ENGINE=MyISAM AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
--  Table structure for `TopicReply`
-- ----------------------------
DROP TABLE IF EXISTS `TopicReply`;
CREATE TABLE `TopicReply` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `TopicId` bigint(20) NOT NULL,
  `UserId` bigint(20) NOT NULL,
  `Username` varchar(64) DEFAULT NULL,
  `Reply` text,
  `CreateAt` datetime DEFAULT NULL,
  `SiteId` int(11) DEFAULT '0',
  PRIMARY KEY (`Id`),
  KEY `UserId` (`UserId`) USING BTREE,
  KEY `Username` (`Username`) USING BTREE,
  KEY `CreateAt` (`CreateAt`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=80 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

-- ----------------------------
--  Table structure for `Zx`
-- ----------------------------
DROP TABLE IF EXISTS `Zx`;
CREATE TABLE `Zx` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `SiteId` bigint(11) DEFAULT NULL,
  `PageKey` varchar(32) DEFAULT NULL,
  `Ua` varchar(8) DEFAULT NULL,
  `ZxJson` text,
  `TestJson` text,
  `CreateAt` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`),
  UNIQUE KEY `SiteId` (`SiteId`,`PageKey`,`Ua`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Procedure definition for `fmt`
-- ----------------------------
DROP FUNCTION IF EXISTS `fmt`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` FUNCTION `fmt`(inD datetime) RETURNS varchar(20) CHARSET utf8mb4
BEGIN  
	DECLARE n BIGINT DEFAULT UNIX_TIMESTAMP(now())- UNIX_TIMESTAMP(inD);
	if n< 60 THEN
		RETURN '刚刚'; 
	end if;
 if n <   60 * 60 THEN
		RETURN CONCAT(ROUND(n /60),'分钟前') ;
 end if;

 if n <  60 * 60 * 24 THEN
		RETURN CONCAT(ROUND(n /(60*60)),'小时前') ;
	end if;

 if n <   60 * 60 * 24* 15 THEN
		RETURN CONCAT(ROUND(n /(60*60*24)),'天前') ;
	end if;

	if date_format(inD,'%Y')=date_format(now(),'%Y') THEN
		RETURN date_format(inD,'%m-%d %H:%i');
	end if; 

	RETURN date_format(inD,'%Y-%m-%d %H:%i');
END
;;
DELIMITER ;

-- ----------------------------
--  Event definition for `JobLoginAward`
-- ----------------------------
DROP EVENT IF EXISTS `JobLoginAward`;
CREATE DEFINER=`root`@`%` EVENT `JobLoginAward` ON SCHEDULE EVERY 1 DAY STARTS '2018-07-08 00:00:01' ON COMPLETION NOT PRESERVE ENABLE DO UPDATE GkUser set DauCount=DauCount+Dau,LoginAward=1,Dau=0;

-- ----------------------------
--  Event definition for `TopicTodayClick`
-- ----------------------------
DROP EVENT IF EXISTS `TopicTodayClick`;
CREATE DEFINER=`root`@`%` EVENT `TopicTodayClick` ON SCHEDULE EVERY 1 DAY STARTS '2018-07-08 00:00:01' ON COMPLETION NOT PRESERVE ENABLE DO UPDATE Topic set ClickCount=ClickCount+Today , Today=0;

-- ----------------------------
--  Records 
-- ----------------------------
INSERT INTO `Category` VALUES ('918','a','0','2018-07-19 08:42:54','2018-07-19 08:42:54','0','4'), ('919','bc','0','2018-07-19 08:42:56','2018-07-19 08:42:56','0','4'), ('920','a1','918','2018-07-19 08:43:00','2018-07-19 08:43:00','1','4'), ('921','a2','918','2018-07-19 08:43:03','2018-07-19 08:43:03','0','4'), ('922','b1','919','2018-07-19 08:43:05','2018-07-19 08:43:05','0','4'), ('923','b2','919','2018-07-19 08:43:08','2018-07-19 08:43:08','0','4'), ('924','a','0','2018-07-20 20:10:34','2018-07-20 20:10:34','0','6'), ('925','b','0','2018-07-20 20:10:38','2018-07-20 20:10:38','0','6'), ('926','a1','924','2018-07-20 20:10:40','2018-07-20 20:10:40','18','6'), ('927','b1','925','2018-07-20 20:10:42','2018-07-20 20:10:42','4','6');
INSERT INTO `Fav` VALUES ('5','11025','22','2018-07-22 10:17:39','6'), ('6','11023','41','2018-07-23 10:22:04','6');
INSERT INTO `Follow` VALUES ('11','11025','11023','2018-07-22 10:17:27','6');
INSERT INTO `GkToken` VALUES ('11021','web','955e1f0a29dbc4641358dfa67dda501b','2018-07-19 08:49:19','4'), ('11022','web','ede97375f98c39a4d7ed246dc18b931b','2018-07-19 10:29:13','4'), ('11023','web','c73df888ac8037a22c5679200853a271','2018-07-26 09:32:08','6'), ('11024','web','801ad9a173d7258235c4b1ff5c0d9e47','2018-07-20 20:31:41','6'), ('11025','web','7eed196b1460a9755d63ec57df78565b','2018-07-22 10:16:49','6'), ('11026','web','4b3755b8cc5c309678fc96479f013620','2018-07-24 17:36:17','6');
INSERT INTO `GkUser` VALUES ('11022','xypcn','6349cf24d81487251f23130bf6b996b9','xypcn@163.com','21212121211',NULL,'0','2018-07-19 08:57:37','2018-07-19 10:11:08','2018-07-19 08:57:37','2018-07-19 10:29:13','-80','Markdown','0','0','0','12','0','0',NULL,'0','179','4','0'), ('11023','xypcn','6349cf24d81487251f23130bf6b996b9','xypcn@163.com','18911417108',NULL,'0','2018-07-20 19:03:54','2018-07-26 22:41:04','2018-07-20 19:03:54','2018-07-26 09:32:08','1776','Markdown','0','1','6','9','0','0','{}','0','941','6','0'), ('11024','xieyp','c9b8211c41a7ebe91581cb03c4c9f59f','xieyp@66m.com','12345612341',NULL,'0','2018-07-20 20:31:41','2018-07-20 20:31:41','2018-07-20 20:31:41',NULL,'2000','Markdown','0','0','0','0','0','1',NULL,'0','21','6','0'), ('11025','xieyp2','d6de2a6a7b2483e14386851d6f0d9b85','xieyp@211.com','21212121211',NULL,'0','2018-07-22 10:16:49','2018-07-22 14:28:17','2018-07-22 10:16:49',NULL,'2020','Markdown','1','1','6','0','0','0',NULL,'0','98','6','0'), ('11026','xieyp1','962d90a6ba74ee1907211e6bd627e2ba','test@qq.com','21231321132',NULL,'0','2018-07-22 16:42:30','2018-07-23 00:28:30','2018-07-22 16:42:30','2018-07-24 17:36:17','1848','HTML','0','0','10','1','0','0','{}','0','203','6','0');
INSERT INTO `KeyVal` VALUES ('BaseInfo','{\"SiteName\":\"极客社区\",\"Miibeian\":\"冀ICP备16008655号-3\",\"Logo\":\"\"}','4'), ('UserCount','17','0'), ('ReplyCount','57','0'), ('TopicCount','0','4'), ('ReplyCount','12','4'), ('UserCount','1','4'), ('ScoreRule','{\"Thank\":10,\"LoginMax\":50,\"LoginMin\":5,\"Reply\":-5,\"GetReply\":5,\"Topic\":-20,\"Register\":2000}','4'), ('TopicCount','22','6'), ('ReplyCount','10','6'), ('UserCount','5','6'), ('BaseInfo','{\"SiteName\":\"test\",\"Logo\":\"\",\"Miibeian\":\"bah\",\"Sign\":\"网站签名\"}','6');
INSERT INTO `Msg` VALUES ('40','11025','11023','在  <a onclick=\"vgo(\'/p/topic/detail,23,11025\',this)\">test</a> 里回复了你','hehe','Reply:71','主题回复','6','0','2018-07-23 00:34:42','23'), ('42','11026','11023',' <a onclick=\"vgo(\'/p/member/xypcn\')\">xypcn</a>  感谢你在  <a onclick=\"vgo(\'/p/topic/detail,22,11023\',this)\">ta</a>  的回复','<nil>','thank:2','收到谢意','6','1','2018-07-24 17:27:15','22'), ('43','11026','11023',' <a onclick=\"vgo(\'/p/member/xypcn\')\">xypcn</a>  感谢你在  <a onclick=\"vgo(\'/p/topic/detail,22,11023\',this)\">ta</a>  的回复','<nil>','thank:3','收到谢意','6','1','2018-07-24 17:33:07','22');
INSERT INTO `ReplyThank` VALUES ('3','11023','6','70','2018-07-24 17:33:07','22');
INSERT INTO `ScoreLog` VALUES ('144','11025','-20','创建主题','创建主题 › <go onclick=\"vgo(\'/p/TopicDao/detail,26,11025\',this)\">test</go>','26','2018-07-22 12:36:20','-20','0'), ('145','11025','-20','创建主题','创建主题 › <go onclick=\"vgo(\'/p/TopicDao/detail,27,11025\',this)\">est</go>','27','2018-07-22 14:20:30','-20','0'), ('146','11025','2015','创建主题','创建主题 › <go onclick=\"vgo(\'/p/TopicDao/detail,28,11025\',this)\">test</go>','28','2018-07-22 14:28:17','-20','6'), ('147','11026','2000','初始资本','获得初始资本 2000','11026','2018-07-22 16:42:30','2000','6'), ('148','11026','2023','每日登录奖励','2018-07-22的每日登录奖励','2018-07-22','2018-07-22 16:43:26','23','6'), ('149','11026','2003','创建主题','创建主题 › <go onclick=\"vgo(\'/p/TopicDao/detail,29,11026\',this)\">test</go>','29','2018-07-22 16:43:49','-20','6'), ('150','11026','1983','创建主题','创建主题 › <go onclick=\"vgo(\'/p/TopicDao/detail,30,11026\',this)\">test</go>','30','2018-07-22 16:45:28','-20','6'), ('151','11026','1963','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,31,11026\',this)\">tet</go>','31','2018-07-22 16:53:56','-20','6'), ('152','11026','1943','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,32,11026\',this)\">test</go>','32','2018-07-22 17:27:24','-20','6'), ('153','11026','1923','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,33,11026\',this)\">test</go>','33','2018-07-22 18:41:20','-20','6'), ('154','11026','1903','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,34,11026\',this)\">b2</go>','34','2018-07-23 00:27:29','-20','6'), ('155','11026','1883','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,35,11026\',this)\">b4</go>','35','2018-07-23 00:27:41','-20','6'), ('156','11026','1878','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,22,11023\',this)\">ta</a>','70','2018-07-23 00:28:00','-5','6'), ('157','11023','2005','主题回复收益','收到  <a onclick=\"vgo(\'/p/member/xieyp1\')\">xieyp1</a> 的回复 ›  <a onclick=\"vgo(\'/p/topic/detail,22,11023\',this)\">ta</a>','70','2018-07-23 00:28:00','5','6'), ('158','11026','1858','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,36,11026\',this)\">aafd</go>','36','2018-07-23 00:28:12','-20','6'), ('159','11026','1838','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,37,11026\',this)\">fassfda</go>','37','2018-07-23 00:28:22','-20','6'), ('160','11026','1818','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,38,11026\',this)\">fsafsfasaf</go>','38','2018-07-23 00:28:30','-20','6'), ('161','11023','2051','每日登录奖励','2018-07-23的每日登录奖励','2018-07-23','2018-07-23 00:28:57','46','6'), ('162','11023','2031','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,39,11023\',this)\">fsaasfsafd</go>','39','2018-07-23 00:29:07','-20','6'), ('163','11023','2011','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,40,11023\',this)\">fsfsad</go>','40','2018-07-23 00:29:14','-20','6'), ('164','11023','1991','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,41,11023\',this)\">fsasfa</go>','41','2018-07-23 00:29:21','-20','6'), ('165','11023','1971','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,42,11023\',this)\">fssafd</go>','42','2018-07-23 00:29:32','-20','6'), ('166','11023','1966','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,23,11025\',this)\">test</a>','71','2018-07-23 00:34:42','-5','6'), ('167','11025','2020','主题回复收益','收到  <a onclick=\"vgo(\'/p/member/xypcn\')\">xypcn</a> 的回复 ›  <a onclick=\"vgo(\'/p/topic/detail,23,11025\',this)\">test</a>','71','2018-07-23 00:34:42','5','6'), ('168','11023','1961','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,41,11023\',this)\">fsasfa</a>','72','2018-07-23 10:18:38','-5','6'), ('169','11023','1956','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,41,11023\',this)\">fsasfa</a>','73','2018-07-23 10:19:33','-5','6'), ('170','11023','1951','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,41,11023\',this)\">fsasfa</a>','74','2018-07-23 10:20:37','-5','6'), ('171','11023','1946','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,41,11023\',this)\">fsasfa</a>','75','2018-07-23 10:20:44','-5','6'), ('172','11023','1941','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,41,11023\',this)\">fsasfa</a>','76','2018-07-23 10:22:00','-5','6'), ('173','11023','1931','发送谢意','感谢  <a onclick=\"vgo(\'/p/member/xieyp1\')\">xieyp1</a> 的回复 >  <a onclick=\"vgo(\'/p/topic/detail,22,11026\',this)\">ta</a>','thank:1','2018-07-24 17:19:23','-10','6'), ('174','11026','1828','收到谢意','感谢  <a onclick=\"vgo(\'/p/member/xypcn\')\">xypcn</a> 的回复','thank:1','2018-07-24 17:19:23','10','6'), ('175','11023','1921','发送谢意','感谢  <a onclick=\"vgo(\'/p/member/xieyp1\')\">xieyp1</a>  的回复 >  <a onclick=\"vgo(\'/p/topic/detail,22,11023\',this)\">ta</a> ','thank:2','2018-07-24 17:27:15','-10','6'), ('176','11026','1838','收到谢意','感谢  <a onclick=\"vgo(\'/p/member/xypcn\')\">xypcn</a>  的回复','thank:2','2018-07-24 17:27:15','10','6'), ('177','11023','1911','发送谢意','感谢  <a onclick=\"vgo(\'/p/member/xieyp1\')\">xieyp1</a>  的回复 >  <a onclick=\"vgo(\'/p/topic/detail,22,11023\',this)\">ta</a> ','thank:3','2018-07-24 17:33:07','-10','6'), ('178','11026','1848','收到谢意','感谢  <a onclick=\"vgo(\'/p/member/xypcn\')\">xypcn</a>  的回复','thank:3','2018-07-24 17:33:07','10','6'), ('179','11023','1906','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,42,11023\',this)\">fssafd</a>','77','2018-07-24 17:37:07','-5','6'), ('180','11023','1901','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,42,11023\',this)\">fssafd</a>','78','2018-07-24 17:37:11','-5','6'), ('181','11023','1881','创建主题','创建主题 › <go onclick=\"vgo(\'/p/topic/detail,43,11023\',this)\">test img</go>','43','2018-07-25 19:31:44','-20','6'), ('182','11023','1861','创建主题附言','创建了附言 >  <a onclick=\"vgo(\'/p/topic/detail,42,11023\',this)\">fssafd</a> ','append:4','2018-07-26 22:32:19','-20','6'), ('183','11023','1841','创建主题附言','创建了附言 >  <a onclick=\"vgo(\'/p/topic/detail,42,11023\',this)\">fssafd</a> ','append:5','2018-07-26 22:37:16','-20','6'), ('184','11023','1821','创建主题附言','创建了附言 >  <a onclick=\"vgo(\'/p/topic/detail,42,11023\',this)\">fssafd</a> ','append:6','2018-07-26 22:38:50','-20','6'), ('185','11023','1801','创建主题附言','创建了附言 >  <a onclick=\"vgo(\'/p/topic/detail,40,11023\',this)\">fsfsad</a> ','append:7','2018-07-26 22:40:54','-20','6'), ('186','11023','1781','创建主题附言','创建了附言 >  <a onclick=\"vgo(\'/p/topic/detail,40,11023\',this)\">fsfsad</a> ','append:8','2018-07-26 22:41:00','-20','6'), ('187','11023','1776','创建回复','创建复 ›   <a onclick=\"vgo(\'/p/topic/detail,40,11023\',this)\">fsfsad</a>','79','2018-07-26 22:41:04','-5','6');
INSERT INTO `Topic` VALUES ('21','11022','xypcn','920','test','<p>test</p>\n','2018-07-19 09:08:11',NULL,'12','xypcn','11022','2018-07-19 10:11:08','test','Markdown','17','17','4','0'), ('22','11023','xypcn','926','ta','<p>haha</p>\n','2018-07-20 20:10:54',NULL,'1','xieyp1','11026','2018-07-23 00:28:00','haha','Markdown','37','37','6','0'), ('23','11025','xieyp2','926','test','<p>test !!!</p>\n','2018-07-22 10:17:12',NULL,'1','xypcn','11023','2018-07-23 00:34:42','test !!!','Markdown','6','6','6','0'), ('24','11025','xieyp2','926','t','<p><img src=\"/upload/6/076/e3c/aed758a1c18c91a0e9cae3368f.jpg\" alt=\"Chrysanthemum.jpg\" /></p>\n','2018-07-22 12:28:04',NULL,'0','','0',NULL,'![Chrysanthemum.jpg](/upload/6/076/e3c/aed758a1c18c91a0e9cae3368f.jpg)','Markdown','2','2','6','0'), ('25','11025','xieyp2','926','test','<p>test</p>\n','2018-07-22 12:33:10',NULL,'0','','0',NULL,'test','Markdown','1','1','6','0'), ('26','11025','xieyp2','926','test','<p>testx</p>\n','2018-07-22 12:36:20',NULL,'0','','0',NULL,'testx','Markdown','0','0','6','0'), ('27','11025','xieyp2','926','est','<p>hhaa</p>\n','2018-07-22 14:20:30',NULL,'0','','0',NULL,'hhaa','Markdown','0','0','6','0'), ('28','11025','xieyp2','926','test','<p>a1</p>\n','2018-07-22 14:28:17',NULL,'0','','0',NULL,'a1','Markdown','1','1','6','0'), ('29','11026','xieyp1','926','test','<p>test</p>\n','2018-07-22 16:43:49',NULL,'0','','0',NULL,'test','Markdown','0','0','6','0'), ('30','11026','xieyp1','926','test','<p>fsfsdf</p>\n','2018-07-22 16:45:28',NULL,'0','','0',NULL,'fsfsdf','Markdown','0','0','6','0'), ('31','11026','xieyp1','926','tet','<p>aa</p>\n','2018-07-22 16:53:56',NULL,'0','','0',NULL,'aa','Markdown','2','2','6','0'), ('32','11026','xieyp1','926','test','\n<p><img src=\"../../upload/6/896/928/8f4245120e7c3870287cce0ff3.jpg\" alt=\"\" width=\"1024\" height=\"768\" /></p>\n','2018-07-22 17:27:24',NULL,'0','','0',NULL,'<nil>','Html','1','1','6','0'), ('33','11026','xieyp1','926','test','\n<p><img src=\"http://127.0.0.1/upload/6/076/e3c/aed758a1c18c91a0e9cae3368f_800.jpg\" width=\"800\" height=\"600\" /></p>\n','2018-07-22 18:41:20',NULL,'0','','0',NULL,'<nil>','Html','3','3','6','0'), ('34','11026','xieyp1','927','b2','\n<p>test</p>\n','2018-07-23 00:27:29',NULL,'0','','0',NULL,'<nil>','Html','0','0','6','0'), ('35','11026','xieyp1','927','b4','\n<p>test bbb</p>\n','2018-07-23 00:27:41',NULL,'0','','0',NULL,'<nil>','Html','0','0','6','0'), ('36','11026','xieyp1','927','aafd','\n<p>fdsa</p>\n','2018-07-23 00:28:12',NULL,'0','','0',NULL,'<nil>','Html','0','0','6','0'), ('37','11026','xieyp1','926','fassfda','\n<p>ffsasfad</p>\n','2018-07-23 00:28:22',NULL,'0','','0',NULL,'<nil>','Html','0','0','6','0'), ('38','11026','xieyp1','926','fsafsfasaf','\n<p>fssfsafsadf</p>\n','2018-07-23 00:28:30',NULL,'0','','0',NULL,'<nil>','Html','0','0','6','0'), ('39','11023','xypcn','926','fsaasfsafd','<p>fsfasafsaf</p>\n','2018-07-23 00:29:07',NULL,'0','','0',NULL,'fsfasafsaf','Markdown','2','2','6','0'), ('40','11023','xypcn','927','fsfsad','<p>fsfafsa</p>\n','2018-07-23 00:29:14',NULL,'1','xypcn','11023','2018-07-26 22:41:04','fsfafsa','Markdown','5','5','6','0'), ('41','11023','xypcn','926','fsasfa','<p>fsdsaffsadsa</p>\n','2018-07-23 00:29:21',NULL,'5','xypcn','11023','2018-07-23 10:22:00','fsdsaffsadsa','Markdown','17','17','6','0'), ('42','11023','xypcn','926','fssafd','<p>fsafsafsda</p>\n','2018-07-23 00:29:32',NULL,'2','xypcn','11023','2018-07-24 17:37:11','fsafsafsda','Markdown','19','19','6','0'), ('43','11023','xypcn','926','test img','\n<p><img src=\"http://127.0.0.1/upload/6/5b7/848/6b77e67cb2383bc6e8f304457c_800.jpg\" alt=\"\" /><img src=\"http://127.0.0.1:9000/upload/6/03a/471/5999dce4d5d1f2de75a279ed5d_800.jpg\" alt=\"\" width=\"727\" height=\"1292\" /></p>\n','2018-07-25 19:31:44',NULL,'0','','0',NULL,'<nil>','Html','39','39','6','0');
INSERT INTO `TopicAppend` VALUES ('1','43','\n<p>test gagag</p>\n','2018-07-26 00:00:00','6'), ('2','43','\n<p>test gagag!!</p>\n','2018-07-26 00:00:00','6'), ('3','43','\n<p>sfdfsfsd</p>\n','2018-07-26 00:00:00','6'), ('4','42','\n<p>gahaha</p>\n','2018-07-26 22:32:19','6'), ('5','42','\n<p>asffdsfaf</p>\n','2018-07-26 22:37:16','6'), ('6','42','\n<p>test !!</p>\n','2018-07-26 22:38:50','6'), ('7','40','\n<p>fafsa</p>\n','2018-07-26 22:40:54','6'), ('8','40','\n<p>fsddsfds</p>\n','2018-07-26 22:41:00','6');
INSERT INTO `TopicReferer` VALUES ('1','15','https://www.v2ex.com/t/469762','2','2018-07-24 21:05:27','2018-07-24 21:05:27','2','1'), ('2','15','https://www.v2ex.com/t/469763','2','2018-07-24 21:05:27','2018-07-24 21:05:27','2','1'), ('3','22','http://test.ecdiy.cn/p/topic/detail,22,11023','1','2018-07-24 22:28:53','2018-07-24 22:28:53','1','6'), ('4','42','http://test.ecdiy.cn/p/topic/detail,22,11023','1','2018-07-24 22:31:09','2018-07-24 22:31:09','1','6'), ('5','43','http://test.ecdiy.cn/p/topic/detail,43,11023','10','2018-07-26 09:40:32','2018-07-26 09:40:32','10','6'), ('6','42','http://test.ecdiy.cn/p/topic/append,43','11','2018-07-26 20:37:00','2018-07-26 20:37:00','11','6'), ('7','43','http://test.ecdiy.cn/p/topic/append,43','5','2018-07-26 21:49:43','2018-07-26 21:49:43','5','6'), ('8','39','http://test.ecdiy.cn/p/topic/append,43','2','2018-07-26 21:49:53','2018-07-26 21:49:53','2','6'), ('9','41','http://test.ecdiy.cn/p/topic/append,43','1','2018-07-26 21:50:24','2018-07-26 21:50:24','1','6'), ('10','40','http://test.ecdiy.cn/p/topic/append,43','4','2018-07-26 22:40:47','2018-07-26 22:40:47','4','6');
INSERT INTO `TopicReply` VALUES ('58','21','11022','xypcn','hehe','2018-07-19 09:19:39','4'), ('59','21','11022','xypcn','hxt','2018-07-19 09:19:44','4'), ('60','21','11022','xypcn','xttt','2018-07-19 09:24:44','4'), ('61','21','11022','xypcn','fsfsfsd','2018-07-19 09:29:14','4'), ('62','21','11022','xypcn','fdfsdsfd','2018-07-19 09:29:17','4'), ('63','21','11022','xypcn','xxxvvvv','2018-07-19 09:29:43','4'), ('64','21','11022','xypcn','xxxvvvv','2018-07-19 10:03:07','4'), ('65','21','11022','xypcn','fsffsd','2018-07-19 10:05:07','4'), ('66','21','11022','xypcn','xxxx231131','2018-07-19 10:05:20','4'), ('67','21','11022','xypcn','63443','2018-07-19 10:05:47','4'), ('68','21','11022','xypcn','fdsfsdxx','2018-07-19 10:06:54','4'), ('69','21','11022','xypcn','yyyyx','2018-07-19 10:11:08','4'), ('70','22','11026','xieyp1','hhaaaa','2018-07-23 00:28:00','6'), ('71','23','11023','xypcn','hehe','2018-07-23 00:34:42','6'), ('72','41','11023','xypcn','hehe','2018-07-23 10:18:38','6'), ('73','41','11023','xypcn','xhhh','2018-07-23 10:19:33','6'), ('74','41','11023','xypcn','ddx','2018-07-23 10:20:37','6'), ('75','41','11023','xypcn','yyudeg','2018-07-23 10:20:44','6'), ('76','41','11023','xypcn','xxxfdsfds','2018-07-23 10:22:00','6'), ('77','42','11023','xypcn','haha','2018-07-24 17:37:07','6'), ('78','42','11023','xypcn','agagagag','2018-07-24 17:37:11','6'), ('79','40','11023','xypcn','<nil>','2018-07-26 22:41:04','6');
INSERT INTO `Zx` VALUES ('1','6','Index','web','{\"main\":{\"list\":[{\"name\":\"自定义HTML\",\"key\":\"gk-html\",\"val\":{\"html\":\"\\n<p>test</p>\\n\"}}]},\"right\":{\"list\":[{\"name\":\"图片\",\"key\":\"gk-img\",\"val\":{\"list\":[{\"url\":\"http://127.0.0.1/upload//728/26e/0d5dbfdd5ee68dc1f909aa7cb3_800.png\"}]}}]}}','{\"main\":{\"list\":[]},\"right\":{\"list\":[{\"name\":\"图片\",\"key\":\"gk-img\",\"val\":{}}]}}','2018-07-21 09:22:35'), ('2','6','TopicDetail','web','{\"main\":{\"list\":[]},\"right\":{\"list\":[]}}','{\"main\":{\"list\":[]},\"right\":{\"list\":[]}}','2018-07-21 09:43:25');
