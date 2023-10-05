package service

import (
	"context"

	"github.com/satriaprayoga/kofin/apps/repository"
	"github.com/satriaprayoga/kofin/internal/store"
)

type AccountService interface {
	Create(ctx context.Context, data *store.Account) error
}

type AccountServiceImpl struct {
	r repository.AccountRepo
}

func (a *AccountServiceImpl) NewAccountService() AccountService {
	accRepo := repository.GetRepo().AccountRepo
	return &AccountServiceImpl{r: accRepo}
}

func (s *AccountServiceImpl) Create(ctx context.Context, data *store.Account) error {
	return s.r.Create(data)
}
