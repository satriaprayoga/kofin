package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type AccountRepo interface {
	Create(data *store.Account) error
	GetByID(ID int) (*store.Account, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
}

type AccountRepoImpl struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) AccountRepo {
	return &AccountRepoImpl{db: db}
}

func (r *AccountRepoImpl) Create(data *store.Account) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *AccountRepoImpl) GetByID(ID int) (*store.Account, error) {
	var result = &store.Account{}
	query := r.db.Where("account_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *AccountRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.Account{}).Where("account_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *AccountRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("account_id=?", ID).Delete(&store.Account{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}
