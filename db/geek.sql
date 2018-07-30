/*
MySQL Backup
Source Server Version: 5.5.16
Source Database: gk
Date: 2018/7/30 18:59:31
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
) ENGINE=MyISAM AUTO_INCREMENT=937 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

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
) ENGINE=MyISAM AUTO_INCREMENT=11035 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

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
) ENGINE=MyISAM AUTO_INCREMENT=208 DEFAULT CHARSET=utf8mb4;

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
) ENGINE=MyISAM AUTO_INCREMENT=45 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

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
) ENGINE=MyISAM AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

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

 