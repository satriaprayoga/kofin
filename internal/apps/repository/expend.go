package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type ExpendRepo interface {
	Create(data *store.Expend) error
	GetByID(ID int) (*store.Expend, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
}

type ExpendRepoImpl struct {
	db *gorm.DB
}

func NewExpendRepo(db *gorm.DB) ExpendRepo {
	return &ExpendRepoImpl{db: db}
}

func (r *ExpendRepoImpl) Create(data *store.Expend) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ExpendRepoImpl) GetByID(ID int) (*store.Expend, error) {
	var result = &store.Expend{}
	query := r.db.Where("expend_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *ExpendRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.Expend{}).Where("expend_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *ExpendRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("expend_id=?", ID).Delete(&store.Expend{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}
