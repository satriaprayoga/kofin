package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(data *store.User) error
	GetByID(ID int) (*store.User, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	GetByUsername(username string) (*store.User, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{db: db}
}

func (r *UserRepoImpl) Create(data *store.User) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepoImpl) GetByID(ID int) (*store.User, error) {
	var result = &store.User{}
	query := r.db.Where("user_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *UserRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.User{}).Where("user_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *UserRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("user_id=?", ID).Delete(&store.User{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepoImpl) GetByUsername(username string) (*store.User, error) {
	var result = &store.User{}
	query := r.db.Where("user_name=?", username).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}
