package controller

import (
	"github/tekeoglan/discord-clone/bootstrap"
	"github/tekeoglan/discord-clone/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginService model.LoginService
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request model.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.LoginService.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "User not found with given email."})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "Incalid Credantials"})
		return
	}

}
