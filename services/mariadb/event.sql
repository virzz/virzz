-- 开启定时器

SET GLOBAL event_scheduler = 1;

-- 删除存储过程

DROP PROCEDURE IF EXISTS auto_clear_record;

-- 创建无参存储过程

delimiter //

CREATE PROCEDURE AUTO_CLEAR_RECORD() BEGIN 
	DECLARE v_name VARCHAR(128);
	-- 退出标志
	declare done BOOLEAN default 0;
	-- 结果游标
	declare name_list cursor for
	SELECT TABLE_NAME
	FROM
	    INFORMATION_SCHEMA.TABLES
	WHERE
	    TABLE_NAME like '%_record%';
	declare continue handler for not found set done=1;
	-- 打开游标
	open name_list;
	read_loop: LOOP FETCH name_list into v_name;
	DELETE FROM v_name WHERE unix_timestamp() - created > 604800;
	IF done THEN LEAVE read_loop;
	END IF;
	END LOOP;
	-- 关闭游标
	close name_list ;
	END// 


delimiter ; 

-- 创建任务

CREATE EVENT IF 
	not exists auto_clear_record_event on schedule every 1 day on completion PRESERVE
	do
	call auto_clear_record();
	-- 开启事件
	ALTER EVENT eventJob ON COMPLETION PRESERVE ENABLE;
	-- 关闭事件
	ALTER EVENT eventJob ON COMPLETION PRESERVE DISABLE;
