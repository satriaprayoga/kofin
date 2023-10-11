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
	"github.com/satriaprayoga/kofin/internal/pkg"
	"github.com/satriaprayoga/kofin/internal/store"
)

type ExpendProgramService interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	GetAvailable(c *gin.Context)
}

type ExpendProgramServiceImpl struct {
	r repository.ExpendProgramRepo
	t time.Duration
}

func NewExpendProgramService(timeout time.Duration) ExpendProgramService {
	accRepo := repository.GetRepo().ExpendProgramRepo
	return &ExpendProgramServiceImpl{r: accRepo, t: timeout}
}

func (s *ExpendProgramServiceImpl) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var acc store.ExpendProgram
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	err := s.r.Create(&acc)
	if err != nil {
		log.Err(err).Msg("Error when saving new data to database. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *ExpendProgramServiceImpl) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_programID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Delete(expend_programID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))

}

func (s *ExpendProgramServiceImpl) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_programID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	var updated = store.ExpendProgram{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Update(expend_programID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *ExpendProgramServiceImpl) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_programID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.GetByID(expend_programID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s *ExpendProgramServiceImpl) GetAvailable(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("year")

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

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))

}
