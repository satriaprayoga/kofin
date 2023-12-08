package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type BudgetRepo interface {
	Create(data *store.Budget) (*store.Budget, error)
	GetByID(ID int) (*store.Budget, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	FindAll() (result *[]store.Budget, err error)
	GetIsActive() (result *store.Budget, err error)
}

type BudgetRepoImpl struct {
	db *gorm.DB
}

func NewBudgetRepo(db *gorm.DB) BudgetRepo {
	return &BudgetRepoImpl{db: db}
}

func (r *BudgetRepoImpl) Create(data *store.Budget) (*store.Budget, error) {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *BudgetRepoImpl) GetByID(ID int) (*store.Budget, error) {
	var result = &store.Budget{}
	query := r.db.Where("budget_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *BudgetRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.Budget{}).Where("budget_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *BudgetRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("budget_id=?", ID).Delete(&store.Budget{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *BudgetRepoImpl) FindAll() (result *[]store.Budget, err error) {
	q := r.db.Find(&result)
	err = q.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BudgetRepoImpl) GetIsActive() (result *store.Budget, err error) {
	q := r.db.Where("budget_is_active", true).First(&result)
	err = q.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
