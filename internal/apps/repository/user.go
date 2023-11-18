package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(data *store.KUser) error
	GetByID(ID int) (*store.KUser, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	GetByUsername(username string) (*store.KUser, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{db: db}
}

func (r *UserRepoImpl) Create(data *store.KUser) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepoImpl) GetByID(ID int) (*store.KUser, error) {
	var result = &store.KUser{}
	query := r.db.Where("user_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *UserRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.KUser{}).Where("user_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *UserRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("user_id=?", ID).Delete(&store.KUser{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepoImpl) GetByUsername(username string) (*store.KUser, error) {
	var result = &store.KUser{}
	query := r.db.Where("username=?", username).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}
