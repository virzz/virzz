package mariadb

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/common"
	"github.com/virzz/virzz/utils"
)

var (
	DB   *gorm.DB
	once utils.OncePlus
	conf = &common.Conf.MySQL

	models = []interface{}{}
)

/*
SQL
	CREATE user virzz@localhost identified by 'virzz9999';
	CREATE DATABASE `virzz_platform` CHARACTER SET = 'utf8mb4' COLLATE = 'utf8mb4_general_ci';
	grant all on `virzz_platform`.* to virzz@localhost;
	flush privileges;
*/

func RegisterModel(t ...interface{}) {
	models = append(models, t...)
}

func Connect() error {

	logger.Debug("Init Mariadb...")

	return once.Do(func() (err error) {
		dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=30s",
			conf.User, conf.Pass, conf.Host, conf.Name, conf.Charset,
		)
		DB, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			logger.Error(err.Error())
			return
		}

		err = DB.AutoMigrate(models...)

		if err != nil {
			logger.Error(err)
			return err
		}

		return nil
	})

}

func GetDB() *gorm.DB {
	return DB
}
