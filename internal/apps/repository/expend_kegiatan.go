package repository

import (
	"fmt"

	dto "github.com/satriaprayoga/kofin/internal/dto/budget"
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type ExpendKegiatanRepo interface {
	Create(data *store.ExpendKegiatan) error
	GetByID(ID int) (*store.ExpendKegiatan, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	GetAvailable(setup dto.ExpendKegiatanSetup) (result *[]store.ExpendKegiatan, err error)
	UpdateOnAccount(A store.ExpendObject) error
	UpdateOnObject(A store.ExpendObject, value float64) error
}

type ExpendKegiatanRepoImpl struct {
	db *gorm.DB
}

func NewExpendKegiatanRepo(db *gorm.DB) ExpendKegiatanRepo {
	return &ExpendKegiatanRepoImpl{db: db}
}

func (r *ExpendKegiatanRepoImpl) Create(data *store.ExpendKegiatan) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ExpendKegiatanRepoImpl) GetByID(ID int) (*store.ExpendKegiatan, error) {
	var result = &store.ExpendKegiatan{}
	query := r.db.Where("expend_kegiatan_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *ExpendKegiatanRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.ExpendKegiatan{}).Where("expend_kegiatan_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *ExpendKegiatanRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("expend_kegiatan_id=?", ID).Delete(&store.ExpendKegiatan{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ExpendKegiatanRepoImpl) GetAvailable(setup dto.ExpendKegiatanSetup) (result *[]store.ExpendKegiatan, err error) {
	queryString := fmt.Sprintf(`SELECT * FROM expend_kegiatan WHERE budget_year=%d AND expend_program_id=%d AND included=FALSE`, setup.Year, setup.ExpendProgramID)
	query := r.db.Raw(queryString).Scan(&result)
	err = query.Error
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *ExpendKegiatanRepoImpl) UpdateOnAccount(A store.ExpendObject) error {
	var result *store.ExpendKegiatan
	q := r.db.Where("expend_kegiatan_id=?", A.ExpendKegiatanID).First(&result)
	if q.RowsAffected == 1 {
		result.KegiatanPagu = result.KegiatanPagu + A.Total
		r.Update(result.ExpendKegiatanID, result)
	}
	err := q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ExpendKegiatanRepoImpl) UpdateOnObject(A store.ExpendObject, value float64) error {
	var result *store.ExpendKegiatan
	q := r.db.Where("expend_kegiatan_id=?", A.ExpendKegiatanID).First(&result)
	err := q.Error
	if err != nil {
		return err
	}
	result.KegiatanPagu += value
	err = r.Update(result.ExpendKegiatanID, result)
	if err != nil {
		return err
	}
	return nil
}
