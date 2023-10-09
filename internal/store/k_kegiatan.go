package store

import (
	"time"

	"gorm.io/gorm"
)

type Kegiatan struct {
	KegiatanID   int    `json:"kegiatan_id" gorm:"primary_key;auto_increment:true"`
	KegiatanKode string `json:"kegiatan_kode" binding:"required" gorm:"type:varchar(10)"`
	KegiatanName string `json:"kegiatan_name" binding:"required" gorm:"type:text"`
	Slug         string `json:"slug" gorm:"type:text"`
	ProgramID    int    `json:"program_id" gorm:"type:int"`
	ProgramKode  string `json:"program_kode" binding:"required" gorm:"type:varchar(10)"`
	ProgramName  string `json:"program_name" binding:"required" gorm:"type:text"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

func (a *Kegiatan) BeforeSave(tx *gorm.DB) (err error) {
	a.Slug = a.ProgramName + " " + a.KegiatanName
	return
}
