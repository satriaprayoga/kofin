package store

import "time"

type BudgetAccount struct {
	BudgetAccountID int     `json:"budget_acc_id" gorm:"primary_key;auto_increment:true"`
	AccountID       int     `json:"account_id" gorm:"primary_key;auto_increment:true"`
	AccKode         string  `json:"acc_kode" binding:"required" gorm:"type:varchar(20)"`
	AccName         string  `json:"acc_name" binding:"required" gorm:"type:text"`
	Root            bool    `json:"root" gorm:"type:bool"`
	Report          string  `json:"report" binding:"required" gorm:"type:varchar(20)"`
	Slug            string  `json:"slug" gorm:"type:text"`
	AccType         string  `json:"acc_type" binding:"required" gorm:"type:varchar(20)"`
	AccGroup        string  `json:"acc_group" binding:"required" gorm:"type:varchar(20)"`
	ParentAccID     int     `json:"acc_parent_id" gorm:"type:int"`
	AccPagu         float64 `json:"acc_pagu" gorm:"type:decimal(12,2)"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}
