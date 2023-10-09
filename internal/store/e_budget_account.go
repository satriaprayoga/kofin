package store

import "time"

type BudgetAccount struct {
	BudgetAccountID int     `json:"budget_acc_id" gorm:"primary_key;auto_increment:true"`
	AccountID       int     `json:"account_id" binding:"required" gorm:"type:int"`
	AccKode         string  `json:"acc_kode" binding:"required" gorm:"type:varchar(10)"`
	AccName         string  `json:"acc_name" binding:"required" gorm:"type:text"`
	AccountPagu     float64 `json:"acc_pagu" gorm:"type:decimal(12,2)"`
	BudgetYear      int     `json:"budget_year" gorm:"type:int"`
	Root            bool    `json:"root" gorm:"type:bool"`
	ParentAccID     int     `json:"acc_parent_id" gorm:"type:int"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}
