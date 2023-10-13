package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type ExpendObjectRepo interface {
	Create(data *store.ExpendObject) error
	GetByID(ID int) (*store.ExpendObject, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
}

type ExpendObjectRepoImpl struct {
	db *gorm.DB
}

func NewExpendObjectRepo(db *gorm.DB) ExpendObjectRepo {
	return &ExpendObjectRepoImpl{db: db}
}

func (r *ExpendObjectRepoImpl) Create(data *store.ExpendObject) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ExpendObjectRepoImpl) GetByID(ID int) (*store.ExpendObject, error) {
	var result = &store.ExpendObject{}
	query := r.db.Where("expend_object_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *ExpendObjectRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.ExpendObject{}).Where("expend_object_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *ExpendObjectRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("expend_object_id=?", ID).Delete(&store.ExpendObject{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}
