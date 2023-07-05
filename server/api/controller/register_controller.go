package controller

import (
	"github/tekeoglan/discord-clone/bootstrap"
	"github/tekeoglan/discord-clone/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {
	RegisterService model.RegisterService
	Env             *bootstrap.Env
}

func (rc *RegisterController) Register(c *gin.Context) {
	var request model.RegisterRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = rc.RegisterService.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, model.ErrorResponse{Message: "User already exist with given email."})
		return
	}

	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptPassword)

	user := model.User{
		BaseModel: model.BaseModel{
			ID:        primitive.NewObjectID(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		UserName: request.UserName,
		Email:    request.Email,
		Password: request.Password,
		Image:    "",
	}

	err = rc.RegisterService.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User created.")
}
