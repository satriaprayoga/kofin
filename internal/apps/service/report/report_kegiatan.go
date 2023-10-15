package report

import (
	"time"

	"github.com/satriaprayoga/kofin/internal/apps/repository"
)

type ExpKegiatanReportService interface {
	GenerateExpKegiatan()
}

type ExpKegiatanReportServiceImpl struct {
	ep repository.ExpendKegiatanRepo
	t  time.Duration
}

func NewExpendKegiatanService(timeout time.Duration) ExpKegiatanReportService {
	ep := repository.GetRepo().ExpendKegiatanRepo
	return &ExpKegiatanReportServiceImpl{ep: ep, t: timeout}
}

func (s *ExpKegiatanReportServiceImpl) GenerateExpKegiatan() {

}
