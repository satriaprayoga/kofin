package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type UnitRepo interface {
	Create(data *store.Unit) error
	GetByID(ID int) (*store.Unit, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	GetAllSubunit() (result *[]store.Unit, err error)
	GetFirstRootUnit() (result *store.Unit, err error)
}

type UnitRepoImpl struct {
	db *gorm.DB
}

func NewUnitRepo(db *gorm.DB) UnitRepo {
	return &UnitRepoImpl{db: db}
}

func (r *UnitRepoImpl) Create(data *store.Unit) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UnitRepoImpl) GetByID(ID int) (*store.Unit, error) {
	var result = &store.Unit{}
	query := r.db.Where("unit_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *UnitRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.Unit{}).Where("unit_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *UnitRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("unit_id=?", ID).Delete(&store.Unit{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UnitRepoImpl) GetAllSubunit() (result *[]store.Unit, err error) {
	query := r.db.Raw("SELECT * FROM unit WHERE root=FALSE ORDER BY kode ASC").Scan(&result)
	err = query.Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *UnitRepoImpl) GetFirstRootUnit() (result *store.Unit, err error) {
	query := r.db.Where("root=?", true).First(&result)
	err = query.Error
	if err != nil {
		return result, err
	}
	return result, nil
}
