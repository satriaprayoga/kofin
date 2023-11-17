package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type RoleRepo interface {
	Create(data *store.Role) error
	GetByID(ID int) (*store.Role, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
}

type RoleRepoImpl struct {
	db *gorm.DB
}

func NewRoleRepo(db *gorm.DB) RoleRepo {
	return &RoleRepoImpl{db: db}
}

func (r *RoleRepoImpl) Create(data *store.Role) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleRepoImpl) GetByID(ID int) (*store.Role, error) {
	var result = &store.Role{}
	query := r.db.Where("role_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *RoleRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.Role{}).Where("role_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *RoleRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("role_id=?", ID).Delete(&store.Role{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}
