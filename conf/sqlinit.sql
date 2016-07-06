

		insert into SysCorpVocation (Parent, Name, Creator) values (1,'root', 1);
		insert into SysCorpStatus (Name, Creator) values ('正常', 1);
		insert into SysCorpStatus (Name, Creator) values ('商户被冻结', 1);
		insert into SysCorpStatus (Name, Creator) values ('商户结算异常', 1);
		insert into SysCorpType (Name, Creator) values ('区域商户', 1);
		insert into SysCorpGrade (Name, Creator) values ('0', 1);
		insert into SysCorpFeeType (Name, Creator) values ('预付费', 1);
		insert into SysCorpFeeType (Name, Creator) values ('后付费', 1);
		insert into SysCorpFeeType (Name, Creator) values ('混合模式', 1);
		insert into SysCorp (Parent, Priority, Name, BriefName, Certificate, CertificateNumber, Ceo, Postcode, Faxcode, Linkman, Mobile, Email, Qq, Wechat, Weibo, VirtualIntegral, RealIntegral, FeeType, PrepayValve, Balance, FrozenBalance, IncomingBalance, Commission, Discount, Province, Area, County, Community, Address, Status, Type, Grade, Vocation, Reserve, Remark, Creator, CreateTime, Auditor, AuditTime, Canceler, CancelTime)
		values (1, 0, 'System', 'System', '', '', '', '', '', '', '', '', '', '', '', 0, 0, 0, 0, 0, 0, 0, 0, 1, '110000', '110000', '110000', '110000', '', 0, 0, 0, 0, '', '', 1, NOW(), null, null, null, null);
		insert into SysSkin(Name, Path, DemoPicture, Creator,CreateTime) values('默认','default.css','default.png', 1,NOW());
		insert into SysOperatorStatus (Name, Creator) values ('正常', 1);
		insert into SysOperatorStatus (Name, Creator) values ('冻结', 1);
		insert into SysOperatorStatus (Name, Creator) values ('注销', 1);
		insert into SysOperatorStatus (Name, Creator) values ('待审核', 1);
		insert into SysOperatorStatus (Name, Creator) values ('待验证', 1);
		insert into SysOperatorGrade (Name, Creator) values ('初级', 1);
		insert into SysOperatorStar (Name, Creator) values ('0星', 1);
		insert into SysOperator (Uimage,Corp, LoginName, RealName, Password, Mobile, IdCard, Email, WechatOpenid, AlipayOpenid, Weibo, AvailableIP, WeatherCode, VirtualIntegral, RealIntegral, Balance, FrozenBalance, IncomingBalance, Commission, Discount, Province, Area, County, Community, Address, Status, Skin, Grade, Star, Session, LoginTime, LoginIP, LoginAgent, LoginCount, LoginErrorCount, FrozenStartTime, FrozenEndTime, Reserve, Remark, Creator, CreateTime, Auditor, AuditTime, Canceler, CancelTime)
		values ('/static/images/uimage.jpg',1, 'admin', '系统管理员', '0d26b6d2be93a768bec4134f33a2c53d', 'admin', 'admin', '', '', '', '', '', '110000', 0, 0, 0, 0, 0, 0, 1, '110000', '110000', '110000', '110000', '', 1, 1, 1, 1, '', null, null, null, 0, 0, null, null, '', '', 1, NOW(), null, null, null, null);
		insert into SysMenu (Direction, Parent, Name, Icon, Clazz, Area, Controller, Method, Parameter, Url, Reserve, Remark, Creator, CreateTime, Auditor, AuditTime, Canceler, CancelTime)
		values (0, 1, 'root', '', '', '', '', '', '', '', null, null, 1, NOW(), 1, NOW(), null, null);
		insert into SysPrivilege (Parent, Menu, Name, Clazz, Area, Controller, Method, Parameter, Url, Reserve, Remark, Creator, CreateTime, Auditor, AuditTime, Canceler, CancelTime)
		values (0, 0, 'root', '', '', '', '', '', '', null, null, 1, NOW(), 1, NOW(), null, null);
		insert into SysRole (Corp, Name, Type, Reserve, Remark, Creator, CreateTime, Auditor, AuditTime, Canceler, CancelTime)
		values (1, 'administrators', 1, null, null, 1, NOW(), 1, NOW(), null, null);
		insert into SysRoleOperator (Role, Operator, Reserve, Remark, Creator, CreateTime, Auditor, AuditTime, Canceler, CancelTime)
		values (1, 1, null, null, 1, NOW(), 1, NOW(), null, null);
		insert into SysRoleMenu (Role, Menu, Reserve, Remark, Creator, CreateTime, Auditor, AuditTime, Canceler, CancelTime)
		values (1, 1, null, null, 1, NOW(), 1, NOW(), null, null);
		insert into SysRolePrivilege (Role, Privilege, Reserve, Remark, Creator, CreateTime, Auditor, AuditTime, Canceler, CancelTime)
		values (1, 1, null, null, 1, NOW(), 1, NOW(), null, null);