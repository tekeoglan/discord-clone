package controller

import (
	"github/tekeoglan/discord-clone/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterController struct {
	AccountService model.AccountService
}

func (rc *RegisterController) Register(c *gin.Context) {
	var request model.RegisterRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	isExist := rc.AccountService.IsEmailExist(c, request.Email)
	if isExist {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "User already exist with given email."})
		return
	}

	user := &model.User{
		UserName: request.UserName,
		Email:    request.Email,
		Password: request.Password,
		Image:    "",
	}

	err = rc.AccountService.Register(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User created.")
}
