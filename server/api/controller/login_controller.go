package controller

import (
	"github/tekeoglan/discord-clone/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	AccountService model.AccountService
	SessionService model.SessionService
}

func (lc *LoginController) Login(c *gin.Context) {
	var request model.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}

	var user model.User
	user, err = lc.AccountService.GetByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrorResponse{Message: "User not found with given email."})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "Invalid Credantials"})
		return
	}

	var sessionId string
	sessionId, err = lc.SessionService.CreateSession(c, user.ID.String())

	c.SetCookie("session_id", sessionId, 24*60*60, "/", "localhost", false, true)
}
