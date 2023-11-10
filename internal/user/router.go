package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	repo IUserRepository
	us   IUserService
}

func NewUserRouter(repo IUserRepository, us IUserService) UserRouter {
	return UserRouter{
		repo: repo,
		us:   us,
	}
}

func (ur *UserRouter) InitUserHandler(path string, r *gin.RouterGroup) {
	r.POST(fmt.Sprintf("%s%s", path, ""), ur.createUser)
	r.POST(fmt.Sprintf("%s%s", path, "/login"), ur.loginUser)
}

/* type UserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
} */

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (ur *UserRouter) createUser(ctx *gin.Context) {
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

/*
type loginUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
} */

/* type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
} */

func (ur *UserRouter) loginUser(ctx *gin.Context) {
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

	/* 	accessToken, err := uh.tokenMaker.CreateToken(user.ID, uh.config.AccessTokenDuration)
	   	if err != nil {
	   		retErr := ownError.NewHttpError(
	   			"Could not create Token",
	   			"There are some internal errors",
	   			http.StatusInternalServerError,
	   		)
	   		ctx.Error(retErr)
	   		return
	   	}
	*/
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
