package models

import "github.com/virzz/virzz/services/server/mariadb"

type Record struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Created int64  `json:"created" gorm:"autoCreateTime"`
	Token   string `json:"-"`
	Name    string `json:"name"`
	Record  string `json:"record" gorm:"varchar(255)"`
	Type    string `json:"type" `
}

func init() {
	mariadb.RegisterModel(&Record{})
}

// Type	Name	Record
// A    test	127.0.0.1

func FindRecordByToken(token, typ string) (r Record, err error) {
	err = DB().Where(&Record{Token: token, Type: typ}).First(&r).Error
	return
}

// NewRecord -
func NewRecord(token, name, typ, record string) (r Record, err error) {
	err = DB().Where(Record{Token: token, Name: name, Type: typ}).
		Attrs(Record{Record: record}).FirstOrCreate(&r).Error
	return
}

// UpdateRecord -
func UpdateRecord(token, name, typ, record string) (err error) {
	err = DB().Where(&Record{Token: token, Type: typ, Name: name}).
		Updates(Record{Record: record}).Error
	return
}

// DeleteRecordByName -
func DeleteRecordByName(token, name string) (err error) {
	var out Record
	if err = DB().First(&out, Record{Token: token, Name: name}).Error; err != nil {
		return err
	}
	err = DB().Delete(out).Error
	return
}

// DeleteRecordByID -
func DeleteRecordByID(token string, id uint) (err error) {
	var out Record
	if err = DB().First(&out, Record{Token: token, ID: id}).Error; err != nil {
		return err
	}
	err = DB().Delete(out).Error
	return
}
