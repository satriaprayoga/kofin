package auth

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satriaprayoga/kofin/internal/apps/repository"
)

type AuthService interface {
	Login(c *gin.Context)
	GetByUsername(c *gin.Context)
}

type AuthServiceImpl struct {
	u repository.UserRepo
	r repository.RoleRepo
	t time.Duration
}

func NewAuthService(timeout time.Duration) AuthService {
	userRepo := repository.GetRepo().UserRepo
	roleRepo := repository.GetRepo().RoleRepo
	return &AuthServiceImpl{u: userRepo, r: roleRepo, t: timeout}
}

func (us *AuthServiceImpl) Login(c *gin.Context) {

}

func (us *AuthServiceImpl) GetByUsername(c *gin.Context) {

}
