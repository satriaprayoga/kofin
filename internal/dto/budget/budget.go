package dto

import "time"

type BudgetSetup struct {
	Year   int       `json:"year"`
	Date   time.Time `json:"date"`
	Level  int       `json:"level"`
	Desc   string    `json:"desc"`
	LastID int       `json:"lastId"`
}

type ExpendBudgetSetup struct {
	ProgramID   int    `json:"program_id"`
	ProgramKode string `json:"program_kode"`
	ProgramName string `json:"program_name"`
	UnitID      int    `json:"unit_id"`
	UnitKode    string `json:"unit_kode"`
	UnitName    string `json:"unit_name"`
}
type ExpendKegiatanSetup struct {
	ExpendProgramID int `json:"expend_program_id"`
	BudgetID        int `json:"budget_id"`
}
