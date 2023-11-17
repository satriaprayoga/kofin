package store

import (
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	UserID    int       `json:"user_id" gorm:"primary_key;auto_increment:true"`
	Username  string    `json:"user_name" binding:"required" gorm:"type:varchar(30)"`
	Fullname  string    `json:"full_name" binding:"required" gorm:"type:varchar(250)"`
	Password  string    `json:"password" binding:"required" gorm:"type:varchar(255)"`
	RoleID    uint      `gorm:"not null;DEFAULT:3" json:"role_id"`
	Role      Role      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}
