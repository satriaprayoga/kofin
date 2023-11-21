package service

import (
	"context"
	"errors"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/satriaprayoga/kofin/internal/apps/repository"
	"github.com/satriaprayoga/kofin/internal/constant"
	"github.com/satriaprayoga/kofin/internal/pkg"
	"github.com/satriaprayoga/kofin/internal/store"

	pagination "github.com/satriaprayoga/kofin/internal/dto/pagination"
)

type KProgramService interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	TextSearch(c *gin.Context)
	PaginateSearch(c *gin.Context)
}

type KProgramServiceImpl struct {
	r repository.KProgramRepo
	t time.Duration
}

func NewKProgramService(timeout time.Duration) KProgramService {
	accRepo := repository.GetRepo().ProgramRepo
	return &KProgramServiceImpl{r: accRepo, t: timeout}
}

func (s *KProgramServiceImpl) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var acc store.KProgram
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Err(err).Msg("Error when mapping request for program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	err := s.r.Create(&acc)
	if err != nil {
		log.Err(err).Msg("Error when saving new data to database. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *KProgramServiceImpl) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	programID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Delete(programID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))

}

func (s *KProgramServiceImpl) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	programID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	var updated = store.KProgram{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	updated.Slug = updated.ProgramName + " " + updated.UnitName
	err = s.r.Update(programID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *KProgramServiceImpl) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	programID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.GetByID(programID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s *KProgramServiceImpl) TextSearch(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	searchTerm := c.Query("searchTerm")

	data, err := s.r.TextSearch(searchTerm)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))

}

func (s *KProgramServiceImpl) PaginateSearch(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var params pagination.ParamList
	var total int64
	var lastPage int
	var page int
	var err error
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Err(err).Msg("Error when mapping request for program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.PaginateSearch(params)
	if err != nil {
		log.Err(err).Msg("Error when searching data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	//response.Data = data
	total, err = s.r.Count(params)
	if err != nil {
		log.Err(err).Msg("Error when searching data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	lastPage = int(math.Ceil(float64(total) / float64(params.PerPage)))
	page = params.Page

	c.JSON(http.StatusOK, pkg.BuildResponseList(constant.Success, data, page, total, lastPage))

}
