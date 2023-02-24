package models

import (
	"github.com/virzz/virzz/services/mariadb"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	return mariadb.GetDB()
}
