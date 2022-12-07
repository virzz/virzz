package models

import (
	"github.com/virzz/virzz/services/server/mariadb"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMariadb() {
	db = mariadb.GetDB()
}
