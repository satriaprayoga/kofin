package service

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/satriaprayoga/kofin/internal/apps/repository"
	"github.com/satriaprayoga/kofin/internal/constant"
	budget "github.com/satriaprayoga/kofin/internal/dto/budget"
	"github.com/satriaprayoga/kofin/internal/pkg"
	"github.com/satriaprayoga/kofin/internal/store"
)

type ExpendKegiatanService interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	GetAvailable(c *gin.Context)
	GetAvailableRKA(c *gin.Context)
}

type ExpendKegiatanServiceImpl struct {
	r repository.ExpendKegiatanRepo
	d repository.ExpendProgramRepo
	t time.Duration
}

func NewExpendKegiatanService(timeout time.Duration) ExpendKegiatanService {
	accRepo := repository.GetRepo().ExpendKegiatanRepo
	ddRepo := repository.GetRepo().ExpendProgramRepo
	return &ExpendKegiatanServiceImpl{r: accRepo, d: ddRepo, t: timeout}
}

func (s *ExpendKegiatanServiceImpl) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var acc store.ExpendKegiatan
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	err := s.r.Create(&acc)
	if err != nil {
		log.Err(err).Msg("Error when saving new data to database. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *ExpendKegiatanServiceImpl) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_kegiatanID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Delete(expend_kegiatanID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))

}

func (s *ExpendKegiatanServiceImpl) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_kegiatanID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	var updated = store.ExpendKegiatan{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Update(expend_kegiatanID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *ExpendKegiatanServiceImpl) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_kegiatanID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.GetByID(expend_kegiatanID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s *ExpendKegiatanServiceImpl) GetAvailable(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("budgetId")

	year, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.GetAvailable(year)
	if err != nil {
		log.Err(err).Msg("Not Found")
		pkg.PanicException(constant.DataNotFound)
	}
	var responseK []KegiatanResponse
	for _, d := range *data {
		pr, err := s.d.GetByID(d.ExpendProgramID)
		if err != nil {
			log.Err(err).Msg("Not Found")
			pkg.PanicException(constant.DataNotFound)
		}
		responseK = append(responseK, KegiatanResponse{
			ExpendKegiatanID: d.ExpendKegiatanID,
			KegiatanKode:     d.KegiatanKode,
			KegiatanName:     d.KegiatanName,
			ProgramKode:      pr.ProgramKode,
			ProgramName:      pr.ProgramName,
		})
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, responseK))
}

func (s *ExpendKegiatanServiceImpl) GetAvailableRKA(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var acc budget.ExpendKegiatanRequest
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_kegiatan request. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	data, err := s.r.GetAvailableRKA(acc.BudgetID, acc.ExpendProgramID)
	if err != nil {
		log.Err(err).Msg("Not Found")
		pkg.PanicException(constant.DataNotFound)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

type KegiatanResponse struct {
	ExpendKegiatanID int    `json:"expend_kegiatan_id"`
	KegiatanKode     string `json:"kegiatan_kode"`
	KegiatanName     string `json:"kegiatan_name"`
	ExpendProgramID  int    `json:"expend_program_id"`
	ProgramKode      string `json:"program_kode"`
	ProgramName      string `json:"program_name"`
}
