package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type ExpendProgramRepo interface {
	Create(data *store.ExpendProgram) error
	GetByID(ID int) (*store.ExpendProgram, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
}

type ExpendProgramRepoImpl struct {
	db *gorm.DB
}

func NewExpendProgramRepo(db *gorm.DB) ExpendProgramRepo {
	return &ExpendProgramRepoImpl{db: db}
}

func (r *ExpendProgramRepoImpl) Create(data *store.ExpendProgram) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ExpendProgramRepoImpl) GetByID(ID int) (*store.ExpendProgram, error) {
	var result = &store.ExpendProgram{}
	query := r.db.Where("expend_program_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *ExpendProgramRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.ExpendProgram{}).Where("expend_program_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *ExpendProgramRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("expend_program_id=?", ID).Delete(&store.ExpendProgram{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}
