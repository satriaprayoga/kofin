package repository

import (
	"fmt"

	budget "github.com/satriaprayoga/kofin/internal/dto/budget"
	pagination "github.com/satriaprayoga/kofin/internal/dto/pagination"
	"github.com/satriaprayoga/kofin/internal/store"
	"github.com/satriaprayoga/kofin/internal/utils"
	"gorm.io/gorm"
)

type KegiatanRepo interface {
	Create(data *store.Kegiatan) error
	GetByID(ID int) (*store.Kegiatan, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	GetByProgramID(prID int) (result *[]store.Kegiatan, err error)
	PaginateSearch(programID int, params pagination.ParamList) (result *[]budget.KegiatanResult, err error)
	Count(programID int, params pagination.ParamList) (int64, error)
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

func (r *KegiatanRepoImpl) PaginateSearch(programID int, params pagination.ParamList) (result *[]budget.KegiatanResult, err error) {
	var (
		pageNum     = 0
		pageSize    = 5
		orderBy     = "rank asc"
		searchTerm  string
		whereString string
		query       *gorm.DB
	)

	if params.Page > 0 {
		pageNum = (params.Page - 1) * params.PerPage
	}

	if params.PerPage > 0 {
		pageSize = params.PerPage
	}

	if params.SortField.Key != "" {
		orderBy = params.SortField.Key + " " + params.SortField.Order
	}

	if params.Search != "" {
		searchTerm = utils.FormatSearch(params.Search)
		queryString := fmt.Sprintf(`k.kegiatan_id, k.kegiatan_kode, k.kegiatan_name, k.program_id,k.program_name, k.program_kode, ts_rank(text_search, to_tsquery('indonesian','%s')) as rank`, searchTerm)
		whereString = `text_search @@ to_tsquery('indonesian',?) AND k.program_id=?`
		query = r.db.Table("kegiatan k").Select(queryString).Where(whereString, searchTerm, programID).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)

	} else {
		queryString := fmt.Sprintf(`k.kegiatan_id, k.kegiatan_kode, k.kegiatan_name, k.program_id,k.program_name, k.program_kode, ts_rank(text_search, to_tsquery('indonesian','')) as rank`)
		query = r.db.Table("kegiatan k").Select(queryString).Where("k.program_id=?", programID).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	}
	err = query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (r *KegiatanRepoImpl) Count(programID int, params pagination.ParamList) (int64, error) {
	var (
		result      int64
		searchTerm  string
		whereString string
		query       *gorm.DB
		err         error
	)

	if params.Search != "" {
		searchTerm = utils.FormatSearch(params.Search)
		queryString := fmt.Sprintf(`k.kegiatan_id, k.kegiatan_kode, k.kegiatan_name, k.program_id,k.program_name, k.program_kode, ts_rank(text_search, to_tsquery('indonesian','%s')) as rank`, searchTerm)
		whereString = `text_search @@ to_tsquery('indonesian',?) AND k.program_id=?`
		query = r.db.Table("k_program k").Select(queryString).Where(whereString, searchTerm, programID).Count(&result)

	} else {
		queryString := fmt.Sprintf(`k.kegiatan_id, k.kegiatan_kode, k.kegiatan_name, k.program_id,k.program_name, k.program_kode, ts_rank(text_search, to_tsquery('indonesian','')) as rank`)
		query = r.db.Table("k_program k").Select(queryString).Where("k.program_id=?", programID).Count(&result)
	}
	err = query.Error
	if err != nil {
		return 0, err
	}
	return result, nil

}
