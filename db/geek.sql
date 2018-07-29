/*
MySQL Backup
Source Server Version: 5.5.16
Source Database: gk
Date: 2018/7/29 15:37:41
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
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

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
  `LoginDay` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`Id`),
  UNIQUE KEY `Email` (`Email`,`SiteId`) USING BTREE,
  UNIQUE KEY `Username` (`Username`,`SiteId`) USING BTREE,
  KEY `Dau` (`Dau`,`SiteId`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=11033 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

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
) ENGINE=MyISAM AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4;

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
) ENGINE=MyISAM AUTO_INCREMENT=203 DEFAULT CHARSET=utf8mb4;

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
) ENGINE=MyISAM AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4;

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
  `Device` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `UserId` (`UserId`) USING BTREE,
  KEY `Username` (`Username`) USING BTREE,
  KEY `CreateAt` (`CreateAt`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

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
INSERT INTO `Fav` VALUES ('5','11025','22','2018-07-22 10:17:39','6'), ('6','11023','41','2018-07-23 10:22:04','6'), ('8','11032','40','2018-07-28 20:13:01','6');
INSERT INTO `Follow` VALUES ('11','11025','11023','2018-07-22 10:17:27','6');
INSERT INTO `GkToken` VALUES ('11021','web','955e1f0a29dbc4641358dfa67dda501b','2018-07-19 08:49:19','4'), ('11022','web','ede97375f98c39a4d7ed246dc18b931b','2018-07-19 10:29:13','4'), ('11023','web','b8b51b0e635eeb2ae3ed7484911e42cd','2018-07-29 08:13:07','6'), ('11024','web','801ad9a173d7258235c4b1ff5c0d9e47','2018-07-20 20:31:41','6'), ('11025','web','7eed196b1460a9755d63ec57df78565b','2018-07-22 10:16:49','6'), ('11026','web','4b3755b8cc5c309678fc96479f013620','2018-07-24 17:36:17','6'), ('11023','h5','17913c9ec44647c49d8e4bfce84d9929','2018-07-29 12:22:12','6'), ('11027','web','9192cc98ed9c3e0d38d4b179b65d3ef2','2018-07-28 16:12:18','6'), ('11028','web','7a813fd577558855232e90ab8d7900da','2018-07-28 16:19:24','6'), ('11029','web','563f20503eb69cd101f5e100f964e365','2018-07-28 16:22:47','6'), ('11030','web','b0e82721a80749ffe4acc1a2dd6d0f3a','2018-07-28 16:24:22','6'), ('11031','web','848f2a8ba52aeb63897ec47205954631','2018-07-28 16:28:29','6'), ('11032','h5','4310d53e2e2f988a4c3295a03217ffe0','2018-07-28 16:53:54','6');
INSERT INTO `GkUser` VALUES ('11022','xypcn','6349cf24d81487251f23130bf6b996b9','xypcn@163.com','21212121211',NULL,'0','2018-07-19 08:57:37','2018-07-19 10:11:08','2018-07-19 08:57:37','2018-07-19 10:29:13','-80','Markdown','0','0','0','12','0','1',NULL,'0','179','4','0','0'), ('11023','xypcn','6349cf24d81487251f23130bf6b996b9','xypcn@163.com','18911417108',NULL,'0','2018-07-20 19:03:54','2018-07-29 15:23:32','2018-07-20 19:03:54','2018-07-29 12:22:12','1766','Markdown','1','1','6','12','0','0','{}','0','2163','6','0','0'), ('11024','xieyp','c9b8211c41a7ebe91581cb03c4c9f59f','xieyp@66m.com','12345612341',NULL,'0','2018-07-20 20:31:41','2018-07-20 20:31:41','2018-07-20 20:31:41',NULL,'2000','Markdown','0','0','0','0','0','0',NULL,'0','21','6','0','0'), ('11025','xieyp2','d6de2a6a7b2483e14386851d6f0d9b85','xieyp@211.com','21212121211',NULL,'0','2018-07-22 10:16:49','2018-07-22 14:28:17','2018-07-22 10:16:49',NULL,'2020','Markdown','1','1','6','0','0','0',NULL,'0','98','6','0','0'), ('11026','xieyp1','962d90a6ba74ee1907211e6bd627e2ba','test@qq.com','21231321132',NULL,'0','2018-07-22 16:42:30','2018-07-23 00:28:30','2018-07-22 16:42:30','2018-07-24 17:36:17','1848','HTML','0','0','10','1','0','0','{}','0','203','6','0','0'), ('11027','xypcn2','b719000d698bbaee7def798b81619e50','xypcn@163.com2','21212121211',NULL,'0','2018-07-28 16:12:18','2018-07-28 16:12:18','2018-07-28 16:12:18',NULL,'2000','HTML','0','0','0','0','0','0',NULL,'0','0','6','0','0'), ('11028','xypcn3','e66aadbc58e4daf1c6a8644846b4cc5e','xypcn@13.com','21212121211',NULL,'0','2018-07-28 16:19:24','2018-07-28 16:19:24','2018-07-28 16:19:24',NULL,'2000','HTML','0','0','0','0','0','0',NULL,'0','9','6','0','0'), ('11029','xypcn4','bf3a370cb823acd8dd5f80f0a3ea52e4','ssaas@323.com','31232121323',NULL,'0','2018-07-28 16:22:47','2018-07-28 16:22:47','2018-07-28 16:22:47',NULL,'2000','HTML','0','0','0','0','0','0',NULL,'0','8','6','0','0'), ('11030','fasfdsa','6de964e7c281718d4c206ea824eab50a','dasad@1221','21212121211',NULL,'0','2018-07-28 16:24:22','2018-07-28 16:24:22','2018-07-28 16:24:22',NULL,'2000','HTML','0','0','0','0','0','0',NULL,'0','8','6','0','0'), ('11031','xypcn7','852c1ac098c451865b0728b12809d853','xypcn2121@fss.com','12341215123',NULL,'0','2018-07-28 16:28:29','2018-07-28 16:28:29','2018-07-28 16:28:29',NULL,'2043','HTML','0','0','0','0','0','0',NULL,'0','302','6','0','3'), ('11032','xypcn26','55ff3a8215f1629f9b89e9bc860f1aea','xypcn@163.com3','212 1212 1211',NULL,'0','2018-07-28 16:53:54','2018-07-28 20:30:15','2018-07-28 16:53:54',NULL,'2037','HTML','0','1','0','1','0','0',NULL,'0','281','6','0','1');
INSERT INTO `KeyVal` VALUES ('BaseInfo','{\"SiteName\":\"极客社区\",\"Miibeian\":\"冀ICP备16008655号-3\",\"Logo\":\"\"}','4'), ('UserCount','17','0'), ('ReplyCount','57','0'), ('TopicCount','0','4'), ('ReplyCount','12','4'), ('UserCount','1','4'), ('ScoreRule','{\"Thank\":10,\"LoginMax\":50,\"LoginMin\":5,\"Reply\":-5,\"GetReply\":5,\"Topic\":-20,\"Register\":2000}','4'), ('TopicCount','22','6'), ('ReplyCount','14','6'), ('UserCount','10','6'), ('BaseInfo','{\"SiteName\":\"test\",\"Logo\":\"\",\"Miibeian\":\"bah\",\"Sign\":\"网站签名\"}','6');
INSERT INTO `Msg` VALUES ('40','11025','11023','在  <a onclick=\"vgo(\'/p/topic/detail,23,11025\',this)\">test</a> 里回复了你','hehe','Reply:71','主题回复','6','0','2018-07-23 00:34:42','23'), ('44','11023','11032','在  <a onclick=\"vgo(\'/p/topic/detail,40,11023\',this)\">fsfsad</a> 里回复了你','hehe android','Reply:83','主题回复','6','0','2018-07-28 20:30:15','40'), ('42','11026','11023',' <a onclick=\"vgo(\'/p/member/xypcn\')\">xypcn</a>  感谢你在  <a onclick=\"vgo(\'/p/topic/detail,22,11023\',this)\">ta</a>  的回复','<nil>','thank:2','收到谢意','6','1','2018-07-24 17:27:15','22'), ('43','11026','11023',' <a onclick=\"vgo(\'/p/member/xypcn\')\">xypcn</a>  感谢你在  <a onclick=\"vgo(\'/p/topic/detail,22,11023\',this)\">ta</a>  的回复','<nil>','thank:3','收到谢意','6','1','2018-07-24 17:33:07','22');
INSERT INTO `ReplyThank` VALUES ('3','11023','6','70','2018-07-24 17:33:07','22');
INSERT INTO `ScoreLog` VALUES ('202','11031','2043','每日登录奖励','2018-07-29的每日登录奖励','2018-07-29','2018-07-29 07:22:44','16','6');
INSERT INTO `Topic` VALUES ('21','11022','xypcn','920','test','<p>test</p>\n','2018-07-19 09:08:11',NULL,'12','xypcn','11022','2018-07-19 10:11:08','test','Markdown','17','17','4','0'), ('22','11023','xypcn','926','ta','<p>haha</p>\n','2018-07-20 20:10:54',NULL,'1','xieyp1','11026','2018-07-23 00:28:00','haha','Markdown','37','37','6','0'), ('23','11025','xieyp2','926','test','<p>test !!!</p>\n','2018-07-22 10:17:12',NULL,'1','xypcn','11023','2018-07-23 00:34:42','test !!!','Markdown','7','7','6','0'), ('24','11025','xieyp2','926','t','<p><img src=\"/upload/6/076/e3c/aed758a1c18c91a0e9cae3368f.jpg\" alt=\"Chrysanthemum.jpg\" /></p>\n','2018-07-22 12:28:04',NULL,'0','','0',NULL,'![Chrysanthemum.jpg](/upload/6/076/e3c/aed758a1c18c91a0e9cae3368f.jpg)','Markdown','2','2','6','0'), ('25','11025','xieyp2','926','test','<p>test</p>\n','2018-07-22 12:33:10',NULL,'0','','0',NULL,'test','Markdown','1','1','6','0'), ('26','11025','xieyp2','926','test','<p>testx</p>\n','2018-07-22 12:36:20',NULL,'0','','0',NULL,'testx','Markdown','0','0','6','0'), ('27','11025','xieyp2','926','est','<p>hhaa</p>\n','2018-07-22 14:20:30',NULL,'0','','0',NULL,'hhaa','Markdown','0','0','6','0'), ('28','11025','xieyp2','926','test','<p>a1</p>\n','2018-07-22 14:28:17',NULL,'0','','0',NULL,'a1','Markdown','1','1','6','0'), ('29','11026','xieyp1','926','test','<p>test</p>\n','2018-07-22 16:43:49',NULL,'0','','0',NULL,'test','Markdown','0','0','6','0'), ('30','11026','xieyp1','926','test','<p>fsfsdf</p>\n','2018-07-22 16:45:28',NULL,'0','','0',NULL,'fsfsdf','Markdown','0','0','6','0'), ('31','11026','xieyp1','926','tet','<p>aa</p>\n','2018-07-22 16:53:56',NULL,'0','','0',NULL,'aa','Markdown','2','2','6','0'), ('32','11026','xieyp1','926','test','\n<p><img src=\"../../upload/6/896/928/8f4245120e7c3870287cce0ff3.jpg\" alt=\"\" width=\"1024\" height=\"768\" /></p>\n','2018-07-22 17:27:24',NULL,'0','','0',NULL,'<nil>','Html','1','1','6','0'), ('33','11026','xieyp1','926','test','\n<p><img src=\"http://127.0.0.1/upload/6/076/e3c/aed758a1c18c91a0e9cae3368f_800.jpg\" width=\"800\" height=\"600\" /></p>\n','2018-07-22 18:41:20',NULL,'0','','0',NULL,'<nil>','Html','3','3','6','0'), ('34','11026','xieyp1','927','b2','\n<p>test</p>\n','2018-07-23 00:27:29',NULL,'0','','0',NULL,'<nil>','Html','0','0','6','0'), ('35','11026','xieyp1','927','b4','\n<p>test bbb</p>\n','2018-07-23 00:27:41',NULL,'0','','0',NULL,'<nil>','Html','0','0','6','0'), ('36','11026','xieyp1','927','aafd','\n<p>fdsa</p>\n','2018-07-23 00:28:12',NULL,'0','','0',NULL,'<nil>','Html','1','1','6','0'), ('37','11026','xieyp1','926','fassfda','\n<p>ffsasfad</p>\n','2018-07-23 00:28:22',NULL,'0','','0',NULL,'<nil>','Html','0','0','6','0'), ('38','11026','xieyp1','926','fsafsfasaf','\n<p>fssfsafsadf</p>\n','2018-07-23 00:28:30',NULL,'0','','0',NULL,'<nil>','Html','1','1','6','0'), ('39','11023','xypcn','926','fsaasfsafd','<p>fsfasafsaf</p>\n','2018-07-23 00:29:07',NULL,'0','','0',NULL,'fsfasafsaf','Markdown','2','2','6','0'), ('40','11023','xypcn','927','fsfsad','<p>fsfafsa</p>\n','2018-07-23 00:29:14',NULL,'5','xypcn26','11032','2018-07-28 20:30:15','fsfafsa','Markdown','33','33','6','0'), ('41','11023','xypcn','926','fsasfa','<p>fsdsaffsadsa</p>\n','2018-07-23 00:29:21',NULL,'5','xypcn','11023','2018-07-23 10:22:00','fsdsaffsadsa','Markdown','18','18','6','0'), ('42','11023','xypcn','926','fssafd','<p>fsafsafsda</p>\n','2018-07-23 00:29:32',NULL,'2','xypcn','11023','2018-07-24 17:37:11','fsafsafsda','Markdown','19','19','6','0'), ('43','11023','xypcn','926','test img','\n<p><img src=\"http://127.0.0.1/upload/6/5b7/848/6b77e67cb2383bc6e8f304457c_800.jpg\" alt=\"\" /><img src=\"http://127.0.0.1:9000/upload/6/03a/471/5999dce4d5d1f2de75a279ed5d_800.jpg\" alt=\"\" width=\"727\" height=\"1292\" /></p>\n','2018-07-25 19:31:44',NULL,'0','','0',NULL,'<nil>','Html','43','43','6','0');
INSERT INTO `TopicAppend` VALUES ('1','43','\n<p>test gagag</p>\n','2018-07-26 00:00:00','6'), ('2','43','\n<p>test gagag!!</p>\n','2018-07-26 00:00:00','6'), ('3','43','\n<p>sfdfsfsd</p>\n','2018-07-26 00:00:00','6'), ('4','42','\n<p>gahaha</p>\n','2018-07-26 22:32:19','6'), ('5','42','\n<p>asffdsfaf</p>\n','2018-07-26 22:37:16','6'), ('6','42','\n<p>test !!</p>\n','2018-07-26 22:38:50','6'), ('7','40','\n<p>fafsa</p>\n','2018-07-26 22:40:54','6'), ('8','40','\n<p>fsddsfds</p>\n','2018-07-26 22:41:00','6');
INSERT INTO `TopicReferer` VALUES ('1','15','https://www.v2ex.com/t/469762','2','2018-07-24 21:05:27','2018-07-24 21:05:27','2','1'), ('2','15','https://www.v2ex.com/t/469763','2','2018-07-24 21:05:27','2018-07-24 21:05:27','2','1'), ('3','22','http://test.ecdiy.cn/p/topic/detail,22,11023','1','2018-07-24 22:28:53','2018-07-24 22:28:53','1','6'), ('4','42','http://test.ecdiy.cn/p/topic/detail,22,11023','1','2018-07-24 22:31:09','2018-07-24 22:31:09','1','6'), ('5','43','http://test.ecdiy.cn/p/topic/detail,43,11023','10','2018-07-26 09:40:32','2018-07-26 09:40:32','10','6'), ('6','42','http://test.ecdiy.cn/p/topic/append,43','11','2018-07-26 20:37:00','2018-07-26 20:37:00','11','6'), ('7','43','http://test.ecdiy.cn/p/topic/append,43','5','2018-07-26 21:49:43','2018-07-26 21:49:43','5','6'), ('8','39','http://test.ecdiy.cn/p/topic/append,43','2','2018-07-26 21:49:53','2018-07-26 21:49:53','2','6'), ('9','41','http://test.ecdiy.cn/p/topic/append,43','1','2018-07-26 21:50:24','2018-07-26 21:50:24','1','6'), ('10','40','http://test.ecdiy.cn/p/topic/append,43','5','2018-07-26 22:40:47','2018-07-26 22:40:47','5','6'), ('11','40','http://test.ecdiy.cn/p/topic/detail,40,11023','1','2018-07-27 02:05:28','2018-07-27 02:05:28','1','6'), ('12','40','http://test.ecdiy.cn/p/topic/list,0,0,1','1','2018-07-28 02:17:59','2018-07-28 02:17:59','1','6');
INSERT INTO `TopicReply` VALUES ('58','21','11022','xypcn','hehe','2018-07-19 09:19:39','4',NULL), ('59','21','11022','xypcn','hxt','2018-07-19 09:19:44','4',NULL), ('60','21','11022','xypcn','xttt','2018-07-19 09:24:44','4',NULL), ('61','21','11022','xypcn','fsfsfsd','2018-07-19 09:29:14','4',NULL), ('62','21','11022','xypcn','fdfsdsfd','2018-07-19 09:29:17','4',NULL), ('63','21','11022','xypcn','xxxvvvv','2018-07-19 09:29:43','4',NULL), ('64','21','11022','xypcn','xxxvvvv','2018-07-19 10:03:07','4',NULL), ('65','21','11022','xypcn','fsffsd','2018-07-19 10:05:07','4',NULL), ('66','21','11022','xypcn','xxxx231131','2018-07-19 10:05:20','4',NULL), ('67','21','11022','xypcn','63443','2018-07-19 10:05:47','4',NULL), ('68','21','11022','xypcn','fdsfsdxx','2018-07-19 10:06:54','4',NULL), ('69','21','11022','xypcn','yyyyx','2018-07-19 10:11:08','4',NULL), ('70','22','11026','xieyp1','hhaaaa','2018-07-23 00:28:00','6',NULL), ('71','23','11023','xypcn','hehe','2018-07-23 00:34:42','6',NULL), ('72','41','11023','xypcn','hehe','2018-07-23 10:18:38','6',NULL), ('73','41','11023','xypcn','xhhh','2018-07-23 10:19:33','6',NULL), ('74','41','11023','xypcn','ddx','2018-07-23 10:20:37','6',NULL), ('75','41','11023','xypcn','yyudeg','2018-07-23 10:20:44','6',NULL), ('76','41','11023','xypcn','xxxfdsfds','2018-07-23 10:22:00','6',NULL), ('77','42','11023','xypcn','haha','2018-07-24 17:37:07','6',NULL), ('78','42','11023','xypcn','agagagag','2018-07-24 17:37:11','6',NULL), ('79','40','11023','xypcn','<nil>','2018-07-26 22:41:04','6',NULL), ('80','40','11023','xypcn','haha','2018-07-27 18:48:56','6',NULL), ('81','40','11023','xypcn','afsfs','2018-07-27 18:49:19','6',NULL), ('82','40','11023','','fdsfs','2018-07-27 18:51:04','6',NULL), ('83','40','11032','xypcn26','hehe android','2018-07-28 20:30:15','6',NULL);
INSERT INTO `Zx` VALUES ('1','6','Index','web','{\"main\":{\"list\":[{\"name\":\"自定义HTML\",\"key\":\"gk-html\",\"val\":{\"html\":\"\\n<p>test</p>\\n\"}}]},\"right\":{\"list\":[{\"name\":\"图片\",\"key\":\"gk-img\",\"val\":{\"list\":[{\"url\":\"http://127.0.0.1/upload//728/26e/0d5dbfdd5ee68dc1f909aa7cb3_800.png\"}]}}]}}','{\"main\":{\"list\":[]},\"right\":{\"list\":[{\"name\":\"图片\",\"key\":\"gk-img\",\"val\":{}}]}}','2018-07-21 09:22:35'), ('2','6','TopicDetail','web','{\"main\":{\"list\":[]},\"right\":{\"list\":[]}}','{\"main\":{\"list\":[]},\"right\":{\"list\":[]}}','2018-07-21 09:43:25');
