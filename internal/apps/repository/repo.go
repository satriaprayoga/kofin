package repository

import "github.com/satriaprayoga/kofin/internal/store"

type Repo struct {
	AccountRepo  AccountRepo
	UnitRepo     UnitRepo
	ProgramRepo  KProgramRepo
	KegiatanRepo KegiatanRepo
}

var repo *Repo

func SetupRepo() {
	repo = &Repo{
		AccountRepo:  NewAccountRepo(store.DB()),
		UnitRepo:     NewUnitRepo(store.DB()),
		ProgramRepo:  NewKProgramRepo(store.DB()),
		KegiatanRepo: NewKegiatanRepo(store.DB()),
	}
}

func GetRepo() *Repo {
	return repo
}
