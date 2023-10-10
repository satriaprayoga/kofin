package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type ExpendKegiatanRepo interface {
	Create(data *store.ExpendKegiatan) error
	GetByID(ID int) (*store.ExpendKegiatan, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
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
