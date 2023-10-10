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

type ExpendAccountService interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
}

type ExpendAccountServiceImpl struct {
	r repository.ExpendAccountRepo
	t time.Duration
}

func NewExpendAccountService(timeout time.Duration) ExpendAccountService {
	accRepo := repository.GetRepo().ExpendAccountRepo
	return &ExpendAccountServiceImpl{r: accRepo, t: timeout}
}

func (s *ExpendAccountServiceImpl) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var acc store.ExpendAccount
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_account creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	err := s.r.Create(&acc)
	if err != nil {
		log.Err(err).Msg("Error when saving new data to database. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *ExpendAccountServiceImpl) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_accountID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_account creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Delete(expend_accountID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))

}

func (s *ExpendAccountServiceImpl) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_accountID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_account creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	var updated = store.ExpendAccount{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_account creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Update(expend_accountID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *ExpendAccountServiceImpl) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_accountID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_account creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.GetByID(expend_accountID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}
