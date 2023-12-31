package budget

import (
	"context"
	"errors"
	"fmt"
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

type ProgramBudgetService interface {
	Import(c *gin.Context)
	Hapus(c *gin.Context)
}

type ProgramBudgetServiceImpl struct {
	ep repository.ExpendProgramRepo
	ek repository.ExpendKegiatanRepo
	k  repository.KegiatanRepo
	t  time.Duration
}

func NewProgramBudgetService(timeout time.Duration) ProgramBudgetService {
	expendProgramRepo := repository.GetRepo().ExpendProgramRepo
	extendKegiatanRepo := repository.GetRepo().ExpendKegiatanRepo
	kegiatanRepo := repository.GetRepo().KegiatanRepo
	return &ProgramBudgetServiceImpl{ep: expendProgramRepo, ek: extendKegiatanRepo, k: kegiatanRepo, t: timeout}
}

func (pb *ProgramBudgetServiceImpl) Import(c *gin.Context) {
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

	//var data store.ExpendProgram{}
	data, err := pb.ep.GetByID(ID)
	if err != nil {
		log.Err(errors.New("data not found")).Msg("Error when check expend program. Data not found Error")
		pkg.PanicException(constant.DataNotFound)
		pkg.PanicHandler(c)
	}
	if data.ExpendProgramID < 1 {
		log.Err(errors.New("data not found")).Msg("Error when check expend program. Data not found Error")
		pkg.PanicException(constant.DataNotFound)
		pkg.PanicHandler(c)
	}
	var updated = store.ExpendProgram{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
		pkg.PanicHandler(c)
	}

	updated.Included = true
	updated.BudgetYear = data.BudgetYear

	err = pb.ep.Update(ID, updated)
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
		//pkg.PanicHandler(c)
	}

	go pb.initializeImport(data.ExpendProgramID, updated)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (pb *ProgramBudgetServiceImpl) Hapus(c *gin.Context) {
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

	//var data store.ExpendProgram{}
	data, err := pb.ep.GetByID(ID)
	if err != nil {
		log.Err(errors.New("data not found")).Msg("Error when check expend program. Data not found Error")
		pkg.PanicException(constant.DataNotFound)
		pkg.PanicHandler(c)
	}
	if data.ExpendProgramID < 1 {
		log.Err(errors.New("data not found")).Msg("Error when check expend program. Data not found Error")
		pkg.PanicException(constant.DataNotFound)
		pkg.PanicHandler(c)
	}
	var updated = store.ExpendProgram{}
	if err := c.ShouldBindJSON(&updated); err != nil {
		log.Err(err).Msg("Error when mapping request for expend_kegiatan creation. Error")
		pkg.PanicException(constant.InvalidRequest)
		pkg.PanicHandler(c)
	}

	log.Info().Msg(fmt.Sprintf("updated: %v", updated))
	err = pb.ep.Update(data.ExpendProgramID, map[string]interface{}{
		"included": false,
	})
	if err != nil {
		log.Err(err).Msg("Error when delete data. Error")
		pkg.PanicException(constant.InvalidRequest)
		//pkg.PanicHandler(c)
	}
	log.Info().Msg(fmt.Sprintf("updated: %v", updated))

	go pb.initializeHapus(updated)

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, "OK"))
}

func (pb *ProgramBudgetServiceImpl) initializeImport(exPrID int, updated store.ExpendProgram) {
	var (
		err       error
		kegiatans *[]store.Kegiatan
	)

	kegiatans, err = pb.k.GetByProgramID(updated.ProgramID)
	log.Info().Msg(fmt.Sprintf("updated: %v", kegiatans))
	if err != nil {
		log.Err(errors.New("data not found")).Msg("Error when check expend program. Data not found Error")
		pkg.PanicException(constant.DataNotFound)

	}
	for _, kk := range *kegiatans {
		ex, _ := pb.ek.GetByKegiatanID(kk.KegiatanID)
		log.Info().Msg(fmt.Sprintf("updated: %v", ex))
		if ex.ExpendKegiatanID == 0 {
			newKgtn := &store.ExpendKegiatan{
				KegiatanID:      kk.KegiatanID,
				KegiatanName:    kk.KegiatanName,
				KegiatanKode:    kk.KegiatanKode,
				ExpendProgramID: exPrID,
				KegiatanPagu:    0.0,
				Included:        true,
				BudgetID:        updated.BudgetID,
				BudgetYear:      updated.BudgetYear,
			}
			log.Info().Msg(fmt.Sprintf("updated: %v", newKgtn))
			err = pb.ek.Create(newKgtn)
			if err != nil {
				log.Err(err).Msg("Error when check expend program. Data not found Error")
				pkg.PanicException(constant.UnknownError)
			}
		} else {

			err = pb.ek.Update(ex.ExpendKegiatanID, map[string]interface{}{
				"included": true,
			})
			if err != nil {
				log.Err(err).Msg("Error when delete data. Error")
				pkg.PanicException(constant.InvalidRequest)
				//pkg.PanicHandler(c)
			}
		}

	}
}

func (pb *ProgramBudgetServiceImpl) initializeHapus(updated store.ExpendProgram) {
	var (
		err       error
		kegiatans *[]store.Kegiatan
	)

	kegiatans, err = pb.k.GetByProgramID(updated.ProgramID)
	log.Info().Msg(fmt.Sprintf("updated: %v", kegiatans))
	if err != nil {
		log.Err(errors.New("data not found")).Msg("Error when check expend program. Data not found Error")
		pkg.PanicException(constant.DataNotFound)

	}
	for _, kk := range *kegiatans {
		ex, _ := pb.ek.GetByKegiatanID(kk.KegiatanID)
		log.Info().Msg(fmt.Sprintf("updated: %v", ex))
		err = pb.ek.Update(ex.ExpendKegiatanID, map[string]interface{}{
			"included": false,
		})
		if err != nil {
			log.Err(err).Msg("Error when delete data. Error")
			pkg.PanicException(constant.InvalidRequest)
			//pkg.PanicHandler(c)
		}

	}
}
