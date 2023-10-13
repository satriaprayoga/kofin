package budget

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/satriaprayoga/kofin/internal/apps/repository"
	"github.com/satriaprayoga/kofin/internal/constant"
	"github.com/satriaprayoga/kofin/internal/pkg"
	"github.com/satriaprayoga/kofin/internal/store"
)

type ObjectBudgetService interface {
	AddToAccount(c *gin.Context)
}

type ObjectBudgetServiceImpl struct {
	ek repository.ExpendKegiatanRepo
	k  repository.KegiatanRepo
	eo repository.ExpendObjectRepo
	t  time.Duration
}

func NewObjectBudgetService(timeout time.Duration) ObjectBudgetService {
	expendKegiatanRepo := repository.GetRepo().ExpendKegiatanRepo
	kegiatanRepo := repository.GetRepo().KegiatanRepo
	expendObjectRepo := repository.GetRepo().ExpendObjectRepo
	return &ObjectBudgetServiceImpl{ek: expendKegiatanRepo, k: kegiatanRepo, eo: expendObjectRepo, t: timeout}
}

func (s *ObjectBudgetServiceImpl) AddToAccount(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var acc store.ExpendObject
	if err := c.ShouldBindJSON(&acc); err != nil {
		log.Err(err).Msg("Error when mapping request for expend object creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	err := s.eo.Create(&acc)
	if err != nil {
		log.Err(err).Msg("Error when saving new data to database. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}
