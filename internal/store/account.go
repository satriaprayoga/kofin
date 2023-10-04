package store

import "gorm.io/gorm"

type Account struct {
	AccountID   int    `json:"account_id" gorm:"primary_key;auto_increment:true"`
	AccKode     string `json:"acc_kode" gorm:"type:varchar(10)"`
	AccName     string `json:"acc_name" gorm:"type:text"`
	Root        bool   `json:"root" gorm:"type:bool"`
	Report      string `json:"report" gorm:"type:varchar(5)"`
	Slug        string `json:"slug" gorm:"type:text"`
	AccType     string `json:"acc_type" gorm:"type:varchar(20)"`
	AccGroup    string `json:"acc_group" gorm:"type:varchar(20)"`
	ParentAccID int    `json:"acc_parent_id" gorm:"type:int"`
}

func (a *Account) BeforeCreate(tx *gorm.DB) {
	a.Slug = a.AccName + " " + a.AccType + " " + a.AccGroup

}

func (a *Account) BeforeUpdate(tx *gorm.DB) (err error) {
	a.Slug = a.AccName + " " + a.AccType + " " + a.AccGroup
	return
}
