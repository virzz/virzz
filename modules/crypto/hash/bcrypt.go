package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func bcryptGenerate(s string, cost int) (string, error) {
	dst, err := bcrypt.GenerateFromPassword([]byte(s), cost)
	if err != nil {
		return "", err
	}
	return string(dst), nil
}

func bcryptCompare(hashed, passwd string) error {
	_, err := bcrypt.Cost([]byte(hashed))
	if err != nil {
		_, err2 := bcrypt.Cost([]byte(passwd))
		if err2 != nil {
			return err
		} else {
			err2 := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(hashed))
			if err2 == nil {
				return nil
			}
			return err2
		}
		return err
	} else {
		err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwd))
		if err == nil {
			return nil
		}
	}
	return err
}
