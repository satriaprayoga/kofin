package store

import "time"

type Unit struct {
	UnitID      int    `json:"unit_id" gorm:"primary_key;auto_increment:true"`
	Kode        string `json:"unit_kode" binding:"required" gorm:"type:varchar(20)"`
	Name        string `json:"unit_name" binding:"required" gorm:"type:text"`
	Loc         string `json:"unit_loc" binding:"required" gorm:"type:text"`
	Abbr        string `json:"unit_abbr" binding:"required" gorm:"type:varchar(20)"`
	Head        string `json:"unit_head" binding:"required" gorm:"type:varchar(150)"`
	HeadID      string `json:"unit_head_id" binding:"required" gorm:"type:varchar(20)"`
	IndukID     int    `json:"unit_induk_id" gorm:"type:int"`
	IndukHead   string `json:"unit_induk_head" binding:"required" gorm:"type:varchar(150)"`
	IndukHeadID string `json:"unit_induk_head_id" binding:"required" gorm:"type:varchar(20)"`
	Root        bool   `json:"root" gorm:"type:bool"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}
