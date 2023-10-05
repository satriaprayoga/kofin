package repository

import "github.com/satriaprayoga/kofin/internal/store"

type Repo struct {
	AccountRepo AccountRepo
}

var repo *Repo

func SetupRepo() {
	repo = &Repo{
		AccountRepo: NewAccountRepo(store.DB()),
	}
}

func GetRepo() *Repo {
	return repo
}
