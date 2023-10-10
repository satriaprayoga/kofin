package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type ExpendAccountRepo interface {
	Create(data *store.ExpendAccount) error
	GetByID(ID int) (*store.ExpendAccount, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
}

type ExpendAccountRepoImpl struct {
	db *gorm.DB
}

func NewExpendAccountRepo(db *gorm.DB) ExpendAccountRepo {
	return &ExpendAccountRepoImpl{db: db}
}

func (r *ExpendAccountRepoImpl) Create(data *store.ExpendAccount) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ExpendAccountRepoImpl) GetByID(ID int) (*store.ExpendAccount, error) {
	var result = &store.ExpendAccount{}
	query := r.db.Where("expend_account_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *ExpendAccountRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.ExpendAccount{}).Where("expend_account_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *ExpendAccountRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("expend_account_id=?", ID).Delete(&store.ExpendAccount{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}
