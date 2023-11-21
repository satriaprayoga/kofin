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

type KegiatanBudgetService interface {
	Import(c *gin.Context)
}

type KegiatanBudgetServiceImpl struct {
	ek repository.ExpendKegiatanRepo
	k  repository.KegiatanRepo
	t  time.Duration
}

func NewKegiatanBudgetService(timeout time.Duration) KegiatanBudgetService {
	expendKegiatanRepo := repository.GetRepo().ExpendKegiatanRepo
	kegiatanRepo := repository.GetRepo().KegiatanRepo
	return &KegiatanBudgetServiceImpl{ek: expendKegiatanRepo, k: kegiatanRepo, t: timeout}
}

func (pb *KegiatanBudgetServiceImpl) Import(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(pb.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	id := c.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		log.Err(errors.New("id is invalid or empty")).Msg("Error when mapping request for expend_program creation. Error")
		pkg.PanicException(constant.InvalidRequest)
		pkg.PanicHandler(c)
	}

	//var data store.ExpendKegiatan{}
	data, err := pb.ek.GetByID(ID)
	if err != nil {
		log.Err(errors.New("data not found")).Msg("Error when check expend program. Data not found Error")
		pkg.PanicException(constant.DataNotFound)
		pkg.PanicHandler(c)
	}
	if data.ExpendKegiatanID < 1 {
		log.Err(errors.New("data not found")).Msg("Error when check expend program. Data not found Error")
		pkg.PanicException(constant.DataNotFound)
		pkg.PanicHandler(c)
	}
	var updated = store.ExpendKegiatan{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
		pkg.PanicHandler(c)
	}

	updated.Included = true
	updated.BudgetYear = data.BudgetYear

	err = pb.ek.Update(ID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
		//pkg.PanicHandler(c)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}
