package dto

import "time"

type BudgetSetup struct {
	Year int       `json:"year"`
	Date time.Time `json:"date"`
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
	Year            int `json:"year"`
}
