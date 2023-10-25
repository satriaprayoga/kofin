package store

import (
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

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

func (e *ExpendObject) BeforeCreate(tx *gorm.DB) (err error) {
	e.Total = float64(e.Volume * e.Satuan)
	e.Total = e.Total * float64(e.Price)
	return
}

func (e *ExpendObject) AfterCreate(tx *gorm.DB) (err error) {
	var result *ExpendAccount
	q := tx.Where("account_id=?", e.AccountID).First(&result)

	if q.RowsAffected == 0 {
		result = &ExpendAccount{
			AccountID:        e.AccountID,
			AccKode:          e.AccKode,
			AccName:          e.AccName,
			AccType:          "BELANJA",
			AccGroup:         "LRA",
			Report:           "rincian",
			Root:             false,
			AccountPagu:      e.Total,
			ExpendKegiatanID: e.ExpendKegiatanID,
			UnitID:           1,
			BudgetYear:       2024,
		}
		q = tx.Create(result)
		err := q.Error
		if err != nil {
			return err
		}
	} else {
		result.AccountPagu = result.AccountPagu + e.Total
		tx.Model(&ExpendAccount{}).Where("account_id=?", e.AccountID).Updates(result)

	}
	e.updateBudgetKegiatan(tx)
	return
}

func (e *ExpendObject) AfterUpdate(tx *gorm.DB) (err error) {
	log.Info().Msg("AfterUpdate called")
	var allObject []ExpendObject
	acc_ID := e.AccountID
	tx.Where("account_id=?", acc_ID).Find(&allObject)
	var result *ExpendAccount
	tx.Where("account_id=?", acc_ID).First(&result)
	var total float64
	for _, j := range allObject {
		total = total + j.Total
	}
	result.AccountPagu = total
	tx.Model(&ExpendAccount{}).Where("expend_account_id=?", result.ExpendAccountID).Updates(result)
	e.updateBudgetKegiatan(tx)

	return
}

/* func (e *ExpendObject) BeforeDelete(tx *gorm.DB) (err error) {
	log.Info().Msg("AfterDelete called")
	var result *ExpendAccount
	tx.Where("account_id=?", e.AccountID).First(&result)
	result.AccountPagu = result.AccountPagu - e.Total
	tx.Model(&ExpendAccount{}).Where("expend_account_id=?", result.ExpendAccountID).Updates(result)
	//e.updateBudgetKegiatan(tx)
	var resultKgtn *ExpendKegiatan
	tx.Where("expend_kegiatan_id=?", e.ExpendKegiatanID).First(&resultKgtn)
	resultKgtn.KegiatanPagu = resultKgtn.KegiatanPagu - e.Total
	tx.Model(&ExpendKegiatan{}).Where("expend_kegiatan_id=?", e.ExpendKegiatanID).Updates(resultKgtn)

	return
} */

func (e *ExpendObject) updateBudgetKegiatan(tx *gorm.DB) (err error) {
	var result *ExpendKegiatan
	q := tx.Where("expend_kegiatan_id=?", e.ExpendKegiatanID).First(&result)
	if q.RowsAffected == 1 {
		var allObject []ExpendObject
		tx.Where("expend_kegiatan_id=?", e.ExpendKegiatanID).Find(&allObject)
		var total float64
		for _, j := range allObject {
			total = total + j.Total
		}
		result.KegiatanPagu = total
		tx.Model(&ExpendKegiatan{}).Where("expend_kegiatan_id=?", e.ExpendKegiatanID).Updates(result)
	}
	return
}

func (e *ExpendKegiatan) AfterUpdate(tx *gorm.DB) (err error) {
	var result *ExpendProgram
	q := tx.Where("expend_program_id=?", e.ExpendProgramID).First(&result)
	if q.RowsAffected == 1 {
		var allObject []ExpendKegiatan
		tx.Where("expend_program_id=?", e.ExpendKegiatanID).Find(&allObject)
		var total float64
		for _, j := range allObject {
			total = total + j.KegiatanPagu
		}
		result.ProgramPagu = total
		tx.Model(&ExpendProgram{}).Where("expend_program_id=?", e.ExpendProgramID).Updates(result)
	}
	return

}
