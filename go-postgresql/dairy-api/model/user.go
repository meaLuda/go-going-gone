package model

import (
	"dairyapi/database"
	"golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "html"
    "strings"
)

type User struct{
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Entries []Entry
}



// functions
func (user *User) save(*User, error){
	err := database.Database.Create(&user).Error

	if err != nil{
		// return empty nil item and the error.
		return &User{},err 
	}
	return user, err
}

func (user *User) BeforeSave(*gorm.DB) error{
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(passwordHash)
    user.Username = html.EscapeString(strings.TrimSpace(user.Username))
    return nil
}