package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type KProgramRepo interface {
	Create(data *store.KProgram) error
	GetByID(ID int) (*store.KProgram, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	GetAll() (result *[]store.KProgram, err error)
}

type KProgramRepoImpl struct {
	db *gorm.DB
}

func NewKProgramRepo(db *gorm.DB) KProgramRepo {
	return &KProgramRepoImpl{db: db}
}

func (r *KProgramRepoImpl) Create(data *store.KProgram) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *KProgramRepoImpl) GetByID(ID int) (*store.KProgram, error) {
	var result = &store.KProgram{}
	query := r.db.Where("program_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *KProgramRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.KProgram{}).Where("program_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *KProgramRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("program_id=?", ID).Delete(&store.KProgram{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *KProgramRepoImpl) GetAll() (result *[]store.KProgram, err error) {
	q := r.db.Find(&result)
	err = q.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
