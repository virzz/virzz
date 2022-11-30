package http

import (
	"fmt"

	"github.com/mozhu1024/virzz/logger"
	"github.com/mozhu1024/virzz/services/server/mariadb"
	"github.com/mozhu1024/virzz/utils"
)

var (
	tokenLetters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	db           = mariadb.DB
)

type Auth struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func init() {
	mariadb.RegisterModel(&Auth{})
}

func newUser(username, password, email string) (auth Auth, err error) {
	if db.First(&auth, &Auth{Username: username}).RowsAffected > 0 {
		return auth, fmt.Errorf("username is exists")
	}
	token := ""
	for {
		token = utils.RandomStringByLength(8, tokenLetters)
		if db.First(&auth, &Auth{Token: token}).RowsAffected == 0 {
			logger.DebugF("token: %s", token)
			break
		}
	}
	auth = Auth{Username: username, Password: password, Token: token}
	if err = db.Create(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil

}

func findAuthByUsername(username string) (auth Auth, err error) {
	if err = db.Where(&Auth{Username: username}).First(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil
}

func FindAuthByToken(token string) (auth Auth, err error) {
	if err = db.Where(&Auth{Token: token}).First(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil
}
