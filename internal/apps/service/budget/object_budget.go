package budget

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

type ObjectBudgetService interface {
	AddToAccount(c *gin.Context)
	UpdateOnAccount(c *gin.Context)
	DeleteOnAccount(c *gin.Context)
}

type ObjectBudgetServiceImpl struct {
	ek repository.ExpendKegiatanRepo
	//k  repository.KegiatanRepo
	ea repository.ExpendAccountRepo
	eo repository.ExpendObjectRepo
	t  time.Duration
}

func NewObjectBudgetService(timeout time.Duration) ObjectBudgetService {
	expendKegiatanRepo := repository.GetRepo().ExpendKegiatanRepo
	//kegiatanRepo := repository.GetRepo().KegiatanRepo
	expendAccountRepo := repository.GetRepo().ExpendAccountRepo
	expendObjectRepo := repository.GetRepo().ExpendObjectRepo
	return &ObjectBudgetServiceImpl{ek: expendKegiatanRepo, ea: expendAccountRepo, eo: expendObjectRepo, t: timeout}
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

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, acc))
}

func (s *ObjectBudgetServiceImpl) UpdateOnAccount(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_objectID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_object creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	/* _, err = s.eo.GetByID(expend_objectID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	*/
	//oldTotal := data.Total
	var updated = &store.ExpendObject{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_object creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	updated.Total = float64(updated.Volume * updated.Satuan)
	updated.Total = updated.Total * float64(updated.Price)
	err = s.eo.Update(expend_objectID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	// value := updated.Total - oldTotal
	// err = s.ea.UpdatePagu(*updated, value)
	// if err != nil {
	// 	log.Err(err).Msg("Error when delete data. Error")
	// 	pkg.PanicException(constant.InvalidRequest)
	// }
	// err = s.ek.UpdateOnObject(*updated, value)
	// if err != nil {
	// 	log.Err(err).Msg("Error when delete data. Error")
	// 	pkg.PanicException(constant.InvalidRequest)
	// }
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, updated))

}

func (s *ObjectBudgetServiceImpl) DeleteOnAccount(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(s.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")

	expend_objectID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_object creation. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	data, err := s.eo.GetByID(expend_objectID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	oldTotal := data.Total
	value := 0 - oldTotal
	err = s.ea.UpdatePagu(*data, value)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.ek.UpdateOnObject(*data, value)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}
	err = s.eo.Delete(expend_objectID)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

}
