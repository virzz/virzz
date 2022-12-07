package models

import "github.com/virzz/virzz/services/server/mariadb"

// Log - Database module
type Log struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Created int64  `json:"created" gorm:"autoCreateTime"`
	Token   string `json:"-"`
	Data    string `json:"data" gorm:"varchar(255)"`
	IP      string `json:"ip"` // remote ip
}

func init() {
	mariadb.RegisterModel(&Log{})
}

func NewLog(token, data, ip string) (m Log, err error) {
	m = Log{Token: token, Data: data, IP: ip}
	err = db.Create(&m).Error
	return
}

// FindLogByToken -
func FindLogByToken(token string) (ls []Log, err error) {
	err = db.Find(&ls, &Log{Token: token}).Limit(50).Error
	return
}

// DeleteLogByToken Batch
func DeleteLogByToken(token string) (err error) {
	var out Log
	if err = db.First(&out, Log{Token: token}).Error; err != nil {
		return err
	}
	err = db.Delete(out).Error
	return
}

// DeleteLogByID Single
func DeleteLogByID(token string, id uint) (err error) {
	var out Log
	if err = db.First(&out, Log{Token: token, ID: id}).Error; err != nil {
		return err
	}
	err = db.Delete(out).Error
	return
}
