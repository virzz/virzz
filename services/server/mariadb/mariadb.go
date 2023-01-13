package mariadb

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/virzz/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/virzz/virzz/utils"
)

var (
	DB   *gorm.DB
	once utils.OncePlus

	models = []interface{}{}
)

func RegisterModel(t ...interface{}) {
	models = append(models, t...)
}

func Migrate() error {
	logger.Debug(models)
	return DB.AutoMigrate(models...)
}

func Procedure() string {
	logger.InfoF(`需要管理权限运行以下脚本 
need the SUPER privilege(s) for this
	%s mariadb -p > event.sql && mysql -uroot -p < event.sql`, os.Args[0])

	return fmt.Sprintf(`
-- 开启定时器
SET GLOBAL event_scheduler = 1;

-- 选择数据库
USE %s;

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
	    TABLE_NAME like '%%_record%%';
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
create event
    if not exists auto_clear_record_event on schedule every 1 day on completion PRESERVE
do
call auto_clear_record();

-- 开启事件
ALTER EVENT auto_clear_record_event ON COMPLETION PRESERVE ENABLE;

-- 关闭事件
ALTER EVENT auto_clear_record_event ON COMPLETION PRESERVE DISABLE;`, viper.GetString("mariadb.name"))
}

func ExecSQL(sql string) (string, error) {
	// Use Raw SQL
	row := DB.Raw(sql).Row()
	var result string
	err := row.Scan(&result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func Connect(debug ...bool) error {
	if !viper.IsSet("mariadb") {
		logger.Fatal("Not set mariadb config")
	}
	return once.Do(func() (err error) {
		logger.Info("Database Mariadb Connecting ...")

		var config = viper.Get("mariadb").(map[string]interface{})

		DB, err = gorm.Open(mysql.Open(fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=30s",
			config["user"], config["pass"], config["host"],
			config["name"], config["charset"],
		)))
		if err != nil {
			logger.Fatal(err)
			return
		}

		if len(debug) > 0 && debug[0] {
			GetDebugDB()
		}
		return nil
	})
}

func GetDebugDB() {
	DB = DB.Debug()
}

func GetDB() *gorm.DB {
	return DB
}
