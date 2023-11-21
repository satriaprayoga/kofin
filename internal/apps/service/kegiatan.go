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
	pagination "github.com/satriaprayoga/kofin/internal/dto/pagination"
	"github.com/satriaprayoga/kofin/internal/pkg"
	"github.com/satriaprayoga/kofin/internal/store"
)

type KegiatanService interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	GetByProgramID(c *gin.Context)
	PaginateSearch(c *gin.Context)
}

type KegiatanServiceImpl struct {
	r repository.KegiatanRepo
	t time.Duration
}

func NewKegiatanService(timeout time.Duration) KegiatanService {
	accRepo := repository.GetRepo().KegiatanRepo
	return &KegiatanServiceImpl{r: accRepo, t: timeout}
}

func (s *KegiatanServiceImpl) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var acc store.Kegiatan
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Err(err).Msg("Error when mapping request for kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	err := s.r.Create(&acc)
	if err != nil {
		log.Err(err).Msg("Error when saving new data to database. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *KegiatanServiceImpl) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	kegiatanID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Delete(kegiatanID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))

}

func (s *KegiatanServiceImpl) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	kegiatanID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	var updated = store.Kegiatan{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Update(kegiatanID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *KegiatanServiceImpl) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	kegiatanID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.GetByID(kegiatanID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s *KegiatanServiceImpl) GetByProgramID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	prID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.GetByProgramID(prID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s *KegiatanServiceImpl) PaginateSearch(c *gin.Context) {

	var params pagination.ParamList
	var total int64
	var lastPage int
	var page int
	var err error
	var prID int
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	prID, err = strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	if err := c.ShouldBindJSON(&params); err != nil {
		log.Err(err).Msg("Error when mapping request for program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.PaginateSearch(prID, params)
	if err != nil {
		log.Err(err).Msg("Error when searching data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	//response.Data = data
	total, err = s.r.Count(prID, params)
	if err != nil {
		log.Err(err).Msg("Error when searching data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	lastPage = int(math.Ceil(float64(total) / float64(params.PerPage)))
	page = params.Page

	c.JSON(http.StatusOK, pkg.BuildResponseList(constant.Success, data, page, total, lastPage))
}
