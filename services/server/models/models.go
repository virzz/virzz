package models

import (
	"github.com/mozhu1024/virzz/services/server/mariadb"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMariadb() {
	db = mariadb.GetDB()
}
