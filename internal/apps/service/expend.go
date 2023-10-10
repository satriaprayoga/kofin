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

type ExpendService interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
}

type ExpendServiceImpl struct {
	r repository.ExpendRepo
	t time.Duration
}

func NewExpendService(timeout time.Duration) ExpendService {
	accRepo := repository.GetRepo().ExpendRepo
	return &ExpendServiceImpl{r: accRepo, t: timeout}
}

func (s *ExpendServiceImpl) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var acc store.Expend
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Err(err).Msg("Error when mapping request for expend creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	err := s.r.Create(&acc)
	if err != nil {
		log.Err(err).Msg("Error when saving new data to database. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *ExpendServiceImpl) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expendID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Delete(expendID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))

}

func (s *ExpendServiceImpl) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expendID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	var updated = store.Expend{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for expend creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Update(expendID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *ExpendServiceImpl) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expendID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.GetByID(expendID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}
