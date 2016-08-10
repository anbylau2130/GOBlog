CREATE  PROCEDURE `usp_addoperator`(
Corp bigint,
LoginName varchar(100),
RealName varchar(100),
Password varchar(100),
IdCard varchar(100),
Email varchar(100),
Mobile varchar(100),
Creator bigint,
Role varchar(100)
)
label_top:
BEGIN
	 declare OperatorID bigint ;
   DECLARE `_rollback` VARCHAR(255);
   DECLARE ProcMsg VARCHAR(250);
   DECLARE IsSuccess VARCHAR(250);
	 DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET `_rollback` = 'ERROR';
	 START TRANSACTION;

		if exists(select 1 from SysOperator where LoginName=LoginName and Corp=Corp)
		THEN
				set IsSuccess='false';
				set ProcMsg='已存在相同的用户名';
				SELECT IsSuccess as IsSuccess,ProcMsg as ProcMsg ;
				LEAVE label_top;
		ELSE 

	  insert into SysOperator(
			Corp,
			LoginName,
			RealName,
			Password,
			IdCard,
			Email,
			Mobile,
			Creator,		
			Province,
			Area,
			County,
			Community,
			CreateTime) 
	    values(
			 Corp, 
			 LoginName, 
			 RealName,
			 Password,
			 IdCard,
			 Email,
       Mobile,
			 Creator,
			 0,
			 0,
			 0,
			 0,
	     NOW());

set @a=LAST_INSERT_ID()	;
set @b=Creator;
set @c=Role;
	    PREPARE addRoleOpeartor FROM ' insert into SysRoleOperator(Role,Operator,Creator,CreateTime) select ID,?,?,NOW() from SysRole where ID in (?) ';
		 
		  EXECUTE addRoleOpeartor USING  @a,@b,@c;

			IF `_rollback`='ERROR' THEN
			  set IsSuccess='false';
			  set ProcMsg='系统错误';
        ROLLBACK;
			ELSE
				COMMIT;
			END IF;
	   SELECT IsSuccess as IsSuccess,ProcMsg as ProcMsg ;
END IF;

END label_top



go


CREATE PROCEDURE `usp_addcorp`(CorpName varchar(100),
   CorpType bigint,
   Operator bigint,
   ParentCorp bigint,
   LoginName varchar(50),
   pwd  varchar(50))
label_top:
BEGIN
   DECLARE `_rollback` VARCHAR(255);
   DECLARE ProcMsg VARCHAR(250);
   DECLARE IsSuccess VARCHAR(250);
	 DECLARE CONTINUE HANDLER FOR SQLEXCEPTION SET `_rollback` = 'ERROR';
	 START TRANSACTION;
	 set @SysCorpID:=0;
   set @SysOperatorID:=0;
   set @SysRoleID:=0;
   set @datetime :=NOW();
   set IsSuccess='true';
   set ProcMsg=''; 
   #判断类型表是否存在相应的数据
   if  (exists(select 1 from SysCorp where Name=@CorpName))
   then
		  set IsSuccess='false';
		  set ProcMsg='已存在相同的公司名称';
		  SELECT IsSuccess as IsSuccess,ProcMsg as ProcMsg ;
		  LEAVE label_top;
   end if ;

   
  #增加一条操作员数据
  if  (exists(select 1 from SysOperator where SysOperator.LoginName=@LoginName))
  then
		 set IsSuccess='false';
		 set ProcMsg='已存在相同的管理员账号';
		 SELECT IsSuccess as IsSuccess,ProcMsg as ProcMsg ;
		 LEAVE label_top;
  end if;
  
			insert into SysCorp
			(
				Parent, Priority, Name,
				BriefName, Certificate, CertificateNumber,
				Ceo, Postcode, Faxcode,
				Linkman, Mobile, Email,
				Qq, Wechat, Weibo,
				VirtualIntegral, RealIntegral, FeeType,
				PrepayValve, Balance, FrozenBalance,
				IncomingBalance, Commission, Discount,
				Province, Area, County,
				Community, Address, Status,
				Type, Grade, Vocation,
				Reserve, Remark, Creator,
				CreateTime, Auditor, AuditTime,
				Canceler, CancelTime
			)
			values
			(
				ParentCorp, 0, CorpName,
				CorpName, '', '',
				'', '', '',
				'', '', '',
				'', '', '',
				0, 0, 0,
				0, 0, 0,
				0, 0, 1,
				'110000', '110000', '110000',
				'110000', '', 0,
				CorpType, 0, 0,
				'', '', @Operator,
				@datetime, null, null,
				null, null
			 );

			set @SysCorpID=LAST_INSERT_ID();

			insert into SysOperator
				(
					Corp, LoginName, RealName,
					Password, Mobile, IdCard,
					Email, WechatOpenid, AlipayOpenid,
					Weibo,AvailableIP, WeatherCode,
					VirtualIntegral,RealIntegral, Balance,
					FrozenBalance, IncomingBalance, Commission,
					Discount, Province, Area,
					County, Community,Address,
					Status, Skin, Grade,
					Star,Session, LoginTime,
					LoginIP, LoginAgent,LoginCount,
					LoginErrorCount, FrozenStartTime,FrozenEndTime,
					Reserve, Remark, Creator,
					CreateTime,Auditor, AuditTime,
					Canceler, CancelTime
				)
				values (
						@SysCorpID, LoginName, CorpName,
						PWD,@datetime, @datetime,
						'', '', '',
						'', '', '110000',
						0, 0, 0,
						0, 0, 0,
						1, '110000', '110000',
						'110000', '110000', '',
						0, 0, 0,
						0, '', null,
						null, null, 0,
						0, null, null,
						'', '', @Operator,
						@datetime, null, null,
						null, null
					);
			set @SysOperatorID=LAST_INSERT_ID();
			 #增加一条角色信息

			insert into SysRole (
					Corp, Name, Type,
					Reserve, Remark, Creator,
					CreateTime, Auditor, AuditTime,
					Canceler, CancelTime
			 )
			values (
					@SysCorpID, 'administrators', 1,
					null, null, Operator,
					@datetime, null, null,
					null, null
			 );

			set @SysRoleID=LAST_INSERT_ID();

			#临时表插入到正式表
			insert into SysRoleMenu
			(
				Role,Menu,Reserve,
				Remark,Creator,CreateTime,
				Auditor,AuditTime,Canceler,
				CancelTime
			)  
			select
			@SysRoleID,ID,null,
			null,@Operator,@datetime,
			null,null,null,
			null 
		  from SysMenu ;

			IF `_rollback`='ERROR' THEN
			  set IsSuccess='false';
			  set ProcMsg='系统错误';
        ROLLBACK;
			ELSE
				COMMIT;
			END IF;
	   SELECT IsSuccess as IsSuccess,ProcMsg as ProcMsg ;
END label_top
go