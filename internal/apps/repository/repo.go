package repository

import "github.com/satriaprayoga/kofin/internal/store"

type Repo struct {
	AccountRepo        AccountRepo
	UnitRepo           UnitRepo
	ProgramRepo        KProgramRepo
	KegiatanRepo       KegiatanRepo
	BudgetRepo         BudgetRepo
	ExpendRepo         ExpendRepo
	ExpendProgramRepo  ExpendProgramRepo
	ExpendKegiatanRepo ExpendKegiatanRepo
	ExpendAccountRepo  ExpendAccountRepo
}

var repo *Repo

func SetupRepo() {
	repo = &Repo{
		AccountRepo:        NewAccountRepo(store.DB()),
		UnitRepo:           NewUnitRepo(store.DB()),
		ProgramRepo:        NewKProgramRepo(store.DB()),
		KegiatanRepo:       NewKegiatanRepo(store.DB()),
		BudgetRepo:         NewBudgetRepo(store.DB()),
		ExpendRepo:         NewExpendRepo(store.DB()),
		ExpendProgramRepo:  NewExpendProgramRepo(store.DB()),
		ExpendKegiatanRepo: NewExpendKegiatanRepo(store.DB()),
		ExpendAccountRepo:  NewExpendAccountRepo(store.DB()),
	}
}

func GetRepo() *Repo {
	return repo
}
