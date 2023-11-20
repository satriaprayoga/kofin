package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	token "github.com/satriaprayoga/kofin/internal/apps/middlewares/jwt"
	"github.com/satriaprayoga/kofin/internal/apps/repository"
	"github.com/satriaprayoga/kofin/internal/constant"
	authDto "github.com/satriaprayoga/kofin/internal/dto/auth"
	"github.com/satriaprayoga/kofin/internal/pkg"
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
	var (
		userInfo authDto.UserInfo
		loginRsp authDto.LoginResponse
	)
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Duration(us.t)*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)

	var login authDto.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		log.Err(err).Msg("Error when mapping request for login. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	user, err := us.u.GetByUsername(login.Username)
	if err != nil {
		log.Err(err).Msg("Username not found. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	err = user.ValidateUserPassword(login.Password)
	if err != nil {
		log.Err(err).Msg("Password error. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	jwt, errJwt := token.GenerateToken(*user)
	if errJwt != nil {
		log.Err(err).Msg("Token creation failed. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	role, errRole := us.r.GetByID(user.RoleID)
	if errRole != nil {
		log.Err(err).Msg("Token creation failed. Error")
		pkg.PanicException(constant.InvalidRequest)
	}

	userInfo = authDto.UserInfo{
		Username: user.Username,
		Role:     []string{role.Name},
	}

	loginRsp = authDto.LoginResponse{
		UserInfo: userInfo,
		Token:    jwt,
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, loginRsp))

}

func (us *AuthServiceImpl) GetByUsername(c *gin.Context) {

}
