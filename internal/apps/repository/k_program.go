package repository

import (
	"fmt"

	budget "github.com/satriaprayoga/kofin/internal/dto/budget"
	pagination "github.com/satriaprayoga/kofin/internal/dto/pagination"
	"github.com/satriaprayoga/kofin/internal/store"
	"github.com/satriaprayoga/kofin/internal/utils"
	"gorm.io/gorm"
)

type KProgramRepo interface {
	Create(data *store.KProgram) error
	GetByID(ID int) (*store.KProgram, error)
	Update(ID int, data interface{}) error
	Delete(ID int) error
	GetAll() (result *[]store.KProgram, err error)
	TextSearch(term string) (result *[]budget.ProgramResult, err error)
	PaginateSearch(params pagination.ParamList) (result *[]budget.ProgramResult, err error)
	Count(params pagination.ParamList) (int64, error)
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

func (r *KProgramRepoImpl) TextSearch(term string) (result *[]budget.ProgramResult, err error) {
	searchTerm := utils.FormatSearch(term)
	queryString := fmt.Sprintf(`
	k.program_id,k.program_name, k.program_kode, k.unit_name, k.unit_kode, 
	ts_rank(text_search, to_tsquery('indonesian','%s')) as rank`, searchTerm)
	// query := r.db.Raw(queryString).Scan(&result)
	query := r.db.Table("k_program k").Select(queryString).Where(`text_search @@ to_tsquery('indonesian',?)`, searchTerm).Find(&result)
	err = query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *KProgramRepoImpl) PaginateSearch(params pagination.ParamList) (result *[]budget.ProgramResult, err error) {
	var (
		pageNum     = 0
		pageSize    = 5
		orderBy     = "rank desc"
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
		queryString := fmt.Sprintf(`k.program_id,k.program_name, k.program_kode, k.unit_name, k.unit_kode,ts_rank(text_search, to_tsquery('indonesian','%s')) as rank`, searchTerm)
		whereString = `text_search @@ to_tsquery('indonesian',?)`
		query = r.db.Table("k_program k").Select(queryString).Where(whereString, searchTerm).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)

	} else {
		queryString := fmt.Sprintf(`k.program_id,k.program_name, k.program_kode, k.unit_name, k.unit_kode,ts_rank(text_search, to_tsquery('indonesian','')) as rank`)
		query = r.db.Table("k_program k").Select(queryString).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	}
	err = query.Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *KProgramRepoImpl) Count(params pagination.ParamList) (int64, error) {
	var (
		result      int64
		searchTerm  string
		whereString string
		query       *gorm.DB
		err         error
	)

	if params.Search != "" {
		searchTerm = utils.FormatSearch(params.Search)
		queryString := fmt.Sprintf(`k.program_id,k.program_name, k.program_kode, k.unit_name, k.unit_kode,ts_rank(text_search, to_tsquery('indonesian','%s')) as rank`, searchTerm)
		whereString = `text_search @@ to_tsquery('indonesian',?)`
		query = r.db.Table("k_program k").Select(queryString).Where(whereString, searchTerm).Count(&result)

	} else {
		queryString := fmt.Sprintf(`k.program_id,k.program_name, k.program_kode, k.unit_name, k.unit_kode,ts_rank(text_search, to_tsquery('indonesian','')) as rank`)
		query = r.db.Table("k_program k").Select(queryString).Count(&result)
	}
	err = query.Error
	if err != nil {
		return 0, err
	}
	return result, nil

}
