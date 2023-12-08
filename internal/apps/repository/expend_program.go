package repository

import (
	"errors"
	"fmt"

	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type ExpendProgramRepo interface {
	Create(data *store.ExpendProgram) error
	GetByID(ID int) (store.ExpendProgram, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	GetAvailable(budgetId int) (result *[]store.ExpendProgram, err error)
	BatchImport(IDs []int) error
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

func (r *ExpendProgramRepoImpl) GetByID(ID int) (result store.ExpendProgram, err error) {
	query := r.db.Where("expend_program_id=?", ID).Find(&result)
	err = query.Error
	if err != nil {
		return result, errors.New("data not found")
	}

	return result, nil
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

func (r *ExpendProgramRepoImpl) GetAvailable(budgetId int) (result *[]store.ExpendProgram, err error) {

	queryString := fmt.Sprintf(`SELECT * FROM expend_program WHERE budget_id=%d AND included=FALSE`, budgetId)

	query := r.db.Raw(queryString).Scan(&result)
	err = query.Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *ExpendProgramRepoImpl) BatchImport(IDs []int) error {
	query := r.db.Table("expend_program").Where("expend_program_id IN ?", IDs).Updates(map[string]bool{"included": true})
	err := query.Error
	if err != nil {
		return err
	}
	return nil

}
