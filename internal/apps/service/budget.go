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

type BudgetService interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
}

type BudgetServiceImpl struct {
	r repository.BudgetRepo
	t time.Duration
}

func NewBudgetService(timeout time.Duration) BudgetService {
	accRepo := repository.GetRepo().BudgetRepo
	return &BudgetServiceImpl{r: accRepo, t: timeout}
}

func (s *BudgetServiceImpl) Create(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var acc store.Budget
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Err(err).Msg("Error when mapping request for budget creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	err := s.r.Create(&acc)
	if err != nil {
		log.Err(err).Msg("Error when saving new data to database. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *BudgetServiceImpl) Delete(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	budgetID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for budget creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Delete(budgetID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))

}

func (s *BudgetServiceImpl) Update(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	budgetID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for budget creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	var updated = store.Budget{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for budget creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.r.Update(budgetID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (s *BudgetServiceImpl) Get(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	budgetID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for budget creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	data, err := s.r.GetByID(budgetID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}
