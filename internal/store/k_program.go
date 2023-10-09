package store

import (
	"time"

	"gorm.io/gorm"
)

type KProgram struct {
	ProgramID   int     `json:"program_id" gorm:"primary_key;auto_increment:true"`
	ProgramKode string  `json:"program_kode" binding:"required" gorm:"type:varchar(10)"`
	ProgramName string  `json:"program_name" binding:"required" gorm:"type:text"`
	Slug        string  `json:"slug" gorm:"type:text"`
	UnitID      int     `json:"unit_id" gorm:"type:int"`
	UnitKode    string  `json:"unit_kode" binding:"required" gorm:"type:varchar(10)"`
	UnitName    string  `json:"unit_name" binding:"required" gorm:"type:text"`
	ProgramPagu float64 `json:"program_pagu" gorm:"type:decimal(12,2)"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

func (a *KProgram) BeforeSave(tx *gorm.DB) (err error) {
	a.Slug = a.ProgramName + " " + a.UnitName
	return
}
