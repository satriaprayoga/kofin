package store

import "time"

type Budget struct {
	BudgetID     int       `json:"budget_id" gorm:"primary_key;auto_increment:true"`
	BudgetYear   int       `json:"budget_year" gorm:"type:int"`
	BudgetStatus int       `json:"budget_status" gorm:"type:int"`
	BudgetDate   time.Time `json:"budget_date" gorm:"type:timestamp(0)"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type Expend struct {
	ExpendID       int     `json:"expend_id" gorm:"primary_key;auto_increment:true"`
	ExpendStatus   int     `json:"expend_status" gorm:"type:int"`
	ExpendYear     int     `json:"expend_year" gorm:"type:int"`
	ExpendPagu     float64 `json:"expend_pagu" gorm:"type:decimal(12,2)"`
	UnitID         int     `json:"unit_id" gorm:"type:int"`
	UnitKode       string  `json:"unit_kode" binding:"required" gorm:"type:varchar(10)"`
	UnitName       string  `json:"unit_name" binding:"required" gorm:"type:text"`
	BudgetID       int     `json:"budget_id" gorm:"type:int"`
	Root           bool    `json:"root" gorm:"type:bool"`
	ParentExpendID int     `json:"expend_parent_id" gorm:"type:int"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type ExpendProgram struct {
	ExpendProgramID int     `json:"expend_program_id" gorm:"primary_key;auto_increment:true"`
	ProgramID       int     `json:"program_id" gorm:"type:int"`
	ProgramKode     string  `json:"program_kode" binding:"required" gorm:"type:varchar(10)"`
	ProgramName     string  `json:"program_name" binding:"required" gorm:"type:text"`
	Slug            string  `json:"slug" gorm:"type:text"`
	ProgramPagu     float64 `json:"program_pagu" gorm:"type:decimal(12,2)"`
	BudgetYear      int     `json:"budget_year" gorm:"type:int"`
	Included        bool    `json:"included" gorm:"type:bool"`

	UnitID         int    `json:"unit_id" gorm:"type:int"`
	UnitKode       string `json:"unit_kode" binding:"required" gorm:"type:varchar(10)"`
	UnitName       string `json:"unit_name" binding:"required" gorm:"type:text"`
	ParentExpendID int    `json:"expend_parent_id" gorm:"type:int"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type ExpendKegiatan struct {
	ExpendKegiatanID int       `json:"expend_kegiatan_id" gorm:"primary_key;auto_increment:true"`
	KegiatanID       int       `json:"kegiatan_id" gorm:"type:int"`
	KegiatanKode     string    `json:"kegiatan_kode" binding:"required" gorm:"type:varchar(10)"`
	KegiatanName     string    `json:"kegiatan_name" binding:"required" gorm:"type:text"`
	Slug             string    `json:"slug" gorm:"type:text"`
	KegiatanPagu     float64   `json:"kegiatan_pagu" gorm:"type:decimal(12,2)"`
	ExpendProgramID  int       `json:"expend_program_id" gorm:"type:int"`
	BudgetYear       int       `json:"budget_year" gorm:"type:int"`
	Included         bool      `json:"included" gorm:"type:bool"`
	UserInput        string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit         string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput        time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit         time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type ExpendAccount struct {
	ExpendAccountID  int     `json:"expend_account_id" gorm:"primary_key;auto_increment:true"`
	AccountID        int     `json:"account_id"`
	AccKode          string  `json:"acc_kode" binding:"required" gorm:"type:varchar(20)"`
	AccName          string  `json:"acc_name" binding:"required" gorm:"type:text"`
	Root             bool    `json:"root" gorm:"type:bool"`
	Report           string  `json:"report" binding:"required" gorm:"type:varchar(20)"`
	AccType          string  `json:"acc_type" binding:"required" gorm:"type:varchar(20)"`
	AccGroup         string  `json:"acc_group" binding:"required" gorm:"type:varchar(20)"`
	AccountPagu      float64 `json:"acc_pagu" gorm:"type:decimal(12,2)"`
	UnitID           int     `json:"unit_id" gorm:"type:int"`
	ExpendKegiatanID int     `json:"expend_kegiatan_id" gorm:"type:int"`
	BudgetYear       int     `json:"budget_year" gorm:"type:int"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}

type ExpendObject struct {
	ExpendObjekID    int     `json:"expend_objek_id" gorm:"primary_key;auto_increment:true"`
	ObjekName        string  `json:"objek_name" binding:"required" gorm:"type:text"`
	Volume           int     `json:"volume" binding:"required" gorm:"type:int"`
	VolumnName       string  `json:"volume_name" binding:"required" gorm:"type:varchar(20)"`
	Satuan           int     `json:"satuan" binding:"required" gorm:"type:int"`
	SatuanName       string  `json:"satuan_name" binding:"required" gorm:"type:varchar(20)"`
	Price            int     `json:"price" binding:"required" gorm:"type:int"`
	Total            float64 `json:"total" gorm:"type:decimal(12,2)"`
	ExpendKegiatanID int     `json:"expend_kegiatan_id" gorm:"type:int"`
	AccountID        int     `json:"account_id"`
	AccKode          string  `json:"acc_kode" binding:"required" gorm:"type:varchar(20)"`
	AccName          string  `json:"acc_name" binding:"required" gorm:"type:text"`

	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
}
