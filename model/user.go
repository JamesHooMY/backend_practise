package model

import (
	"gorm.io/gorm"

	"go_backend/util"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(50);not null;unique" json:"email"`
	Mobile   string `gorm:"type:varchar(11)" json:"mobile"`
	Name     string `gorm:"type:varchar(20);not null" json:"name"`
	Age      int    `gorm:"type:int(3);check:(age >= 0) AND (age <= 150)" json:"age"`
	Password string `gorm:"type:varchar(60);not null" json:"-"`
}

func (u *User) Encrypt() error {
	encryptedPassword, err := util.Encrypt(u.Password)
	if err != nil {
		return err
	}

	u.Password = encryptedPassword

	return nil
}
