CREATE DEFINER=`root`@`localhost` PROCEDURE `usp_addcorp`(PCorpName varchar(100),
   PCorpType bigint,
   POperator bigint,
   ParentCorp bigint,
   PLoginName varchar(50),
   pwd  varchar(50))
label_top:
BEGIN
	 DECLARE `_rollback` INT DEFAULT 0;
	 DECLARE CONTINUE HANDLER FOR SQLEXCEPTION  set  `_rollback`= 1;  
   set @datetime :=NOW();

   select count(1) into @corpcnt from SysCorp where `Name`=PCorpName LIMIT 1;
   if  (@corpcnt>0)
   then
		  set @IsSuccess='false';
		  set @ProcMsg='已存在相同的公司名';
		  SELECT @IsSuccess as IsSuccess,@ProcMsg as ProcMsg ;
		  LEAVE label_top;
   end if ;

   select count(1) INTO @cnt from SysOperator where SysOperator.LoginName=PLoginName LIMIT 1;
 
  if  (@cnt>0)
  then
		 set @IsSuccess='false';
		 set @ProcMsg='已存在相同的用户';
		 SELECT @IsSuccess as IsSuccess,@ProcMsg as ProcMsg ;
		 LEAVE label_top;
  end if;
      START TRANSACTION;
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
				ParentCorp, 0, PCorpName,
				PCorpName, '', '',
				'', '', '',
				'', '', '',
				'', '', '',
				0, 0, 0,
				0, 0, 0,
				0, 0, 1,
				'110000', '110000', '110000',
				'110000', '', 0,
				PCorpType, 0, 0,
				'', '', POperator,
				@datetime, null, null,
				null, null
			 );

			set @SysCorpID=LAST_INSERT_ID();

			insert into SysOperator
				(
					Corp, LoginName, RealName,
					`Password`, Mobile, IdCard,
					Email, WechatOpenid, AlipayOpenid,
					Weibo,AvailableIP, WeatherCode,
					VirtualIntegral,RealIntegral, Balance,
					FrozenBalance, IncomingBalance, Commission,
					Discount, Province, Area,
					County, Community,Address,
					`Status`, Skin, Grade,
					Star,`Session`, LoginTime,
					LoginIP, LoginAgent,LoginCount,
					LoginErrorCount, FrozenStartTime,FrozenEndTime,
					Reserve, Remark, Creator,
					CreateTime,Auditor, AuditTime,
					Canceler, CancelTime
				)
				values (
						@SysCorpID, PLoginName, PCorpName,
						pwd,@datetime, @datetime,
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
						'', '', POperator,
						@datetime, null, null,
						null, null
					);
			set @SysOperatorID=LAST_INSERT_ID();

		     
                      
                          
          insert into SysRole (              
              Corp, Name, Type,              
              Reserve, Remark, Creator,              
              CreateTime, Auditor, AuditTime,              
              Canceler, CancelTime              
           )              
          values (
              @SysCorpID, CONCAT(PCorpName,'admin') , 1,              
              NULL, PCorpName, POperator,              
              @datetime, 0, @datetime,              
              null, null              
           );              
            
          set @SysRoleID=LAST_INSERT_ID();  
          

          insert into SysRoleOperator(Role,Operator,Creator)
          values(@SysRoleID,@SysOperatorID,POperator);
                     
               
          insert into SysRoleMenu              
          (              
            Role,Menu,Reserve,              
            Remark,Creator,CreateTime,              
            Auditor,AuditTime,Canceler,              
            CancelTime              
          )  select              
          @SysRoleID,Menu,null,  
          null,POperator,@datetime,  
          0,@datetime,null,  
          null from SysMenuTemplate where `CorpType`=PCorpType;
          
          insert into SysRolePrivilege
          (Role,Privilege,Creator)
          select             
          @SysRoleID,Privilege,Operator from SysPrivilegeTemplate where `CorpType`=PCorpType;
                   

		IF `_rollback`=1 THEN
			set @IsSuccess='false';
			set @ProcMsg= CONCAT(RETURNED_SQLSTATE,':',MESSAGE_TEXT);
			ROLLBACK;
		ELSE
			set @IsSuccess='true';
			set @ProcMsg='';
			COMMIT;
		END IF;

	 SELECT @IsSuccess as IsSuccess ,@ProcMsg  as ProcMsg ;
	 
END label_top





CREATE DEFINER=`root`@`localhost` PROCEDURE `usp_addoperator`(PCorp bigint,
PLoginName varchar(100),
PRealName varchar(100),
Pwd varchar(100),
PIdCard varchar(100),
PEmail varchar(100),
PMobile varchar(100),
PCreator bigint,PRole varchar(100))
label_top:
BEGIN

	 DECLARE `_rollback` INT DEFAULT 0;
	 DECLARE CONTINUE HANDLER FOR SQLEXCEPTION  set  `_rollback`= 1;  
		select count(1) into @cnt from SysOperator where `LoginName`=PLoginName  and `Corp`=PCorp  LIMIT 1;
		if (@cnt>0)
		THEN
        
				set @IsSuccess='false';
				set @ProcMsg='已存在相同的用户名';
				SELECT @IsSuccess as IsSuccess ,@ProcMsg  as ProcMsg ;
				LEAVE label_top;
		ELSE 
    START TRANSACTION;
	  insert into SysOperator(
			`Corp`,
			`LoginName`,
			`RealName`,
			`Password`,
			`IdCard`,
			`Email`,
			`Mobile`,
			`Creator`,		
			`Province`,
			`Area`,
			`County`,
			`Community`,
			`CreateTime`) 
	    values(
			     PCorp, 
					 PLoginName, 
					 PRealName,
					 Pwd,
					 DATE_FORMAT(NOW(),'%Y-%m-%d-%T'),
					 PEmail,
					 DATE_FORMAT(NOW(),'%Y-%m-%d-%T'),
					 PCreator,
					 0,
					 0,
					 0,
					 0,
					DATE_FORMAT(NOW(),'%Y-%m-%d-%T'));
	
			set @sqlRoleOperator=CONCAT('insert into SysRoleOperator(Role,Operator,Creator,CreateTime) ',
												'select ID,',LAST_INSERT_ID(),',',PCreator,',NOW() from SysRole',
												' where ID in (',PRole,')');
			
	    PREPARE addRoleOperator FROM @sqlRoleOperator;
    
		  EXECUTE addRoleOperator ;
      DEALLOCATE PREPARE addRoleOperator;

		IF `_rollback`=1 THEN
			set @IsSuccess='false';
			set @ProcMsg= CONCAT(RETURNED_SQLSTATE,':',MESSAGE_TEXT);
			ROLLBACK;
		ELSE
			set @IsSuccess='true';
			set @ProcMsg='';
			COMMIT;
		END IF;

	 SELECT @IsSuccess as IsSuccess ,@ProcMsg  as ProcMsg ;

   
    
END IF;

END label_top











CREATE DEFINER=`root`@`localhost` PROCEDURE `usp_addrole`(RoleCorp bigint,
RoleName varchar(100),
Type tinyint,
Remark varchar(250),
Creator bigint,
Menus varchar(250),
PrivilegesRole varchar(250))
label_top:
BEGIN
	 DECLARE `_rollback` INT DEFAULT 0;
	 DECLARE CONTINUE HANDLER FOR SQLEXCEPTION  set  `_rollback`= 1;  
	select count(*) INTO @cnt from SysRole where `Name`=RoleName  and `Corp`=RoleCorp ;
  IF (@cnt<>0)THEN	
			set @IsSuccess='false';
			set @ProcMsg='´æÔÚÏàÍ¬µÄ½ÇÉ«Ãû';
			SELECT @IsSuccess as IsSuccess ,@ProcMsg  as ProcMsg ;
			LEAVE label_top;
	ELSE


  START TRANSACTION;
	insert into SysRole(Corp,`Name`,`Type`,Remark,Creator,Auditor) values(RoleCorp,RoleName,Type,Remark,Creator,Creator);
	SET @RoleID:=LAST_INSERT_ID();
	

	
	set @sqlmenu:=CONCAT('insert into SysRoleMenu(Role,Menu,Creator,Auditor) select ',
											@RoleID,',ID,',Creator ,',',Creator,
											' from SysMenu where ID in (?)');
   PREPARE addmenu FROM @sqlmenu;
	 set @menu:=Menus;
	 EXECUTE addmenu USING @menu;


	 set @sqlprivilege:=CONCAT('insert into SysRolePrivilege(Role,Privilege,Creator,Auditor) select ',
												@RoleID ,',ID,',Creator ,',',Creator,
												' from SysPrivilege where ID in (?)');
	 PREPARE addprivilege FROM @sqlprivilege;
	 set @privilege:=PrivilegesRole ;
	 EXECUTE addprivilege USING @privilege;
	
		IF `_rollback`=1 THEN
			set @IsSuccess='false';
			set @ProcMsg= CONCAT(RETURNED_SQLSTATE,':',MESSAGE_TEXT);
			ROLLBACK;
		ELSE
			set @IsSuccess='true';
			set @ProcMsg='';
			COMMIT;
		END IF;

	 SELECT @IsSuccess as IsSuccess ,@ProcMsg  as ProcMsg ;
    	

END IF;

END label_top
