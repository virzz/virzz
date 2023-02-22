package models

import (
	"fmt"

	"github.com/virzz/logger"
	"github.com/virzz/virzz/services/server/mariadb"
	"github.com/virzz/virzz/utils"
)

var (
	tokenLetters = []byte("abcdefghijklmnopqrstuvwxyz0123456789")
)

type Auth struct {
	ID       uint64 `json:"id" gorm:""`
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Created  int64  `json:"-" gorm:"autoCreateTime"`
	// Email    string `json:"email"`
}

func init() {
	mariadb.RegisterModel(&Auth{})
}

func NewAuth(username, password, email string) (auth Auth, err error) {
	if DB().First(&auth, &Auth{Username: username}).RowsAffected > 0 {
		return auth, fmt.Errorf("username is exists")
	}
	token := ""
	for {
		token = utils.RandomStringByLength(8, tokenLetters)
		if DB().First(&auth, &Auth{Token: token}).RowsAffected == 0 {
			logger.DebugF("token: %s", token)
			break
		}
	}
	auth = Auth{Username: username, Password: password, Token: token}
	if err = DB().Create(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil

}

func FindAuthByUsername(username string) (auth Auth, err error) {
	if err = DB().Where(&Auth{Username: username}).First(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil
}

func FindAuthByToken(token string) (auth Auth, err error) {
	if err = DB().Where(&Auth{Token: token}).First(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil
}
