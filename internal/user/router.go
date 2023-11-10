package user

import (
	"net/http"

	"github.com/Ridju/ticktr/config"
	"github.com/Ridju/ticktr/internal/token"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userRouter struct {
	repo   IUserRepository
	us     IUserService
	tm     token.Maker
	config config.Config
}

func InitUserRouter(r *gin.RouterGroup, db *gorm.DB, config config.Config, tokenMaker token.Maker) {
	userRepo := NewGORMRepository(db)
	userService := NewUserService(userRepo)

	userRouter := userRouter{
		repo:   userRepo,
		us:     userService,
		config: config,
		tm:     tokenMaker,
	}

	r.POST("", userRouter.createUser)
	r.POST("/login", userRouter.loginUser)

}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (ur *userRouter) createUser(ctx *gin.Context) {
	req := struct {
		Username string `json:"username" binding:"required,alphanum"`
		Password string `json:"password" binding:"required,min=6"`
		Email    string `json:"email" binding:"required,email"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	user, err := ur.us.CreateUser(req.Username, req.Password, req.Email)
	if err != nil {
		ctx.Error(err)
		return
	}

	resp := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	ctx.JSON(http.StatusOK, resp)
}

func (ur *userRouter) loginUser(ctx *gin.Context) {
	req := struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	user, err := ur.us.LoginUser(req.Email, req.Password)
	if err != nil {
		ctx.Error(err)
		return
	}

	accessToken, err := ur.tm.CreateToken(user.ID, ur.config.AccessTokenDuration)
	if err != nil {
		ctx.Error(err)
		return
	}

	rsp := struct {
		AccessToken string       `json:"access_token"`
		User        UserResponse `json:"user"`
	}{
		AccessToken: accessToken,
		User: UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}

	ctx.JSON(http.StatusOK, rsp)
}
