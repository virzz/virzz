package models

import (
	"github.com/virzz/virzz/services/server/mariadb"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	return mariadb.GetDB()
}
