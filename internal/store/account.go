package store

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	AccountID   int    `json:"account_id" gorm:"primary_key;auto_increment:true"`
	AccKode     string `json:"acc_kode" binding:"required" gorm:"type:varchar(10)"`
	AccName     string `json:"acc_name" binding:"required" gorm:"type:text"`
	Root        bool   `json:"root" gorm:"type:bool"`
	Report      string `json:"report" binding:"required" gorm:"type:varchar(5)"`
	Slug        string `json:"slug" gorm:"type:text"`
	AccType     string `json:"acc_type" binding:"required" gorm:"type:varchar(20)"`
	AccGroup    string `json:"acc_group" binding:"required" gorm:"type:varchar(20)"`
	ParentAccID int    `json:"acc_parent_id" gorm:"type:int"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.Slug = a.AccName + " " + a.AccType + " " + a.AccGroup
	return

}

func (a *Account) BeforeSave(tx *gorm.DB) (err error) {
	a.Slug = a.AccName + " " + a.AccType + " " + a.AccGroup
	return
}
