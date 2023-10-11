package repository

import (
	"github.com/satriaprayoga/kofin/internal/store"
	"gorm.io/gorm"
)

type KegiatanRepo interface {
	Create(data *store.Kegiatan) error
	GetByID(ID int) (*store.Kegiatan, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	GetByProgramID(prID int) (result *[]store.Kegiatan, err error)
}

type KegiatanRepoImpl struct {
	db *gorm.DB
}

func NewKegiatanRepo(db *gorm.DB) KegiatanRepo {
	return &KegiatanRepoImpl{db: db}
}

func (r *KegiatanRepoImpl) Create(data *store.Kegiatan) error {
	query := r.db.Create(data)
	err := query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *KegiatanRepoImpl) GetByID(ID int) (*store.Kegiatan, error) {
	var result = &store.Kegiatan{}
	query := r.db.Where("kegiatan_id=?", ID).Find(result)
	err := query.Error
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *KegiatanRepoImpl) Update(ID int, data interface{}) error {
	var err error

	q := r.db.Model(&store.Kegiatan{}).Where("kegiatan_id=?", ID).Updates(data)
	err = q.Error
	if err != nil {
		return err
	}
	return nil

}

func (r *KegiatanRepoImpl) Delete(ID int) error {
	var err error
	q := r.db.Where("kegiatan_id=?", ID).Delete(&store.Kegiatan{})
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *KegiatanRepoImpl) GetByProgramID(prID int) (result *[]store.Kegiatan, err error) {
	//queryString := fmt.Sprintf("select * from kegiatan where program_id=%d", prID)
	//q := r.db.Raw(queryString).Scan(&result)
	q := r.db.Where("program_id=?", prID).Find(&result)
	err = q.Error
	if err != nil {
		return nil, err
	}
	return result, err
}
