DROP TABLE IF EXISTS `tb_admin_user`;
CREATE TABLE IF NOT EXISTS `tb_admin_user` (
	`id` int(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`userid` varchar(100) NOT NULL DEFAULT '0' COMMENT '用户名',
	`password` varchar(100) NOT NULL DEFAULT '0' COMMENT '密码',
	`groupid` int(20) NOT NULL DEFAULT '0' COMMENT '权限组ID',
	`remarks` varchar(200) NOT NULL DEFAULT '0' COMMENT '备注',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_admin_group`;
CREATE TABLE IF NOT EXISTS `tb_admin_group` (
	`id` int(20) NOT NULL AUTO_INCREMENT COMMENT '角色组id',
	`name` varchar(20) NOT NULL DEFAULT '0' COMMENT '用户组名字',
	`actionlist` varchar(300) NOT NULL DEFAULT '[]' COMMENT '用户组权限',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_center_list`;
CREATE TABLE IF NOT EXISTS `tb_center_list` (
	`id` int(20) NOT NULL DEFAULT '0' COMMENT '平台Id',
	`name` varchar(100) NOT NULL DEFAULT '0' COMMENT '平台名字',
	`ip` varchar(50) NOT NULL DEFAULT '0' COMMENT 'IP地址',
	`port` varchar(20) NOT NULL DEFAULT '0' COMMENT '端口',
	`serverpath` varchar(200) NOT NULL DEFAULT '0' COMMENT '表存放路径',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_send_mail`;
CREATE TABLE IF NOT EXISTS `tb_send_mail` (
	`mailid` int(40) NOT NULL AUTO_INCREMENT COMMENT '',
	`title` varchar(40) NOT NULL DEFAULT '0' COMMENT '',
	`sendor` varchar(40) NOT NULL DEFAULT '0' COMMENT '',
	`recvname` varchar(400) NOT NULL DEFAULT '0' COMMENT '',
	`content` varchar(600) NOT NULL DEFAULT '0' COMMENT '',
	`itemlist` varchar(500) NOT NULL DEFAULT '0' COMMENT '',
	`isall` int(1) NOT NULL DEFAULT '0' COMMENT '',
	`createtime` varchar(50) NOT NULL DEFAULT '0' COMMENT '',
	`audittime` varchar(50) NOT NULL DEFAULT '0' COMMENT '',
	`status` int(10) NOT NULL DEFAULT '0' COMMENT '',
	`serverurl` varchar(200) NOT NULL DEFAULT '0' COMMENT '',
	`sendid` varchar(20) NOT NULL DEFAULT '0' COMMENT '',
	`auditid` varchar(20) NOT NULL DEFAULT '0' COMMENT '',
	`timetype` int(10) NOT NULL DEFAULT '0' COMMENT '',
	`starttime` varchar(50) NOT NULL DEFAULT '0' COMMENT '',
	`endtime` varchar(50) NOT NULL DEFAULT '0' COMMENT '',
	`lvstart` int(10) NOT NULL DEFAULT '0' COMMENT '',
	`lvend` int(10) NOT NULL DEFAULT '0' COMMENT '',
	`sex` int(10) NOT NULL DEFAULT '0' COMMENT '',
	`reason` varchar(50) NOT NULL DEFAULT '0' COMMENT '',
	PRIMARY KEY (`mailid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_publicnotice`;
CREATE TABLE IF NOT EXISTS `tb_publicnotice` (
	`id` int(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
	`platformid` int(20) NOT NULL DEFAULT '0' COMMENT '平台ID',
	`serverid` int(20) NOT NULL DEFAULT '0' COMMENT '服务器ID',
	`noticetype` int(20) NOT NULL DEFAULT '0' COMMENT '公告类型',
	`label` int(20) NOT NULL DEFAULT '0' COMMENT '标签',
	`priority` int(20) NOT NULL DEFAULT '0' COMMENT '优先级',
	`titleshort` varchar(50) NOT NULL DEFAULT '0' COMMENT '短标题',
	`title` varchar(50) NOT NULL DEFAULT '0' COMMENT '标题',
	`content` varchar(500) NOT NULL DEFAULT '0' COMMENT '内容',
	`starttime` bigint(50) NOT NULL DEFAULT '0' COMMENT '开始时间',
	`endtime` bigint(50) NOT NULL DEFAULT '0' COMMENT '结束时间',
	`createtime` bigint(50) NOT NULL DEFAULT '0' COMMENT '创建时间',
	`audittime` bigint(50) NOT NULL DEFAULT '0' COMMENT '审核时间',
	`sendid` varchar(20) NOT NULL DEFAULT '0' COMMENT '提交ID',
	`auditid` varchar(20) NOT NULL DEFAULT '0' COMMENT '审核ID',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

DROP TABLE IF EXISTS `tb_send_marquee`;
CREATE TABLE IF NOT EXISTS `tb_send_marquee` (
	`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '',
	`mtype` int(11) NOT NULL DEFAULT '0' COMMENT '',
	`num` int(11) NOT NULL DEFAULT '0' COMMENT '',
	`content` varchar(200) NOT NULL DEFAULT '0' COMMENT '',
	`createtime` varchar(50) NOT NULL DEFAULT '0' COMMENT '',
	`audittime` varchar(50) NOT NULL DEFAULT '0' COMMENT '',
	`endtime` varchar(50) NOT NULL DEFAULT '0' COMMENT '',
	`sendok` varchar(200) NOT NULL DEFAULT '0' COMMENT '',
	`serverurl` varchar(200) NOT NULL DEFAULT '0' COMMENT '',
	`sendid` varchar(50) NOT NULL DEFAULT '0' COMMENT '',
	`auditid` varchar(20) NOT NULL DEFAULT '0' COMMENT '',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;