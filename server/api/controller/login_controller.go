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
	err := c.Bind(&request)
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
	sessionId, err = lc.SessionService.CreateSession(c, user.ID.Hex())
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.SetCookie(model.COOKIE_PREFIX_SESSION, sessionId, lc.SessionService.GetCokiExpr(),
		lc.SessionService.GetCokiPath(), lc.SessionService.GetCokiDomain(),
		lc.SessionService.IsCokiSecure(), lc.SessionService.IsCokiHttpOnly())

	c.JSON(http.StatusOK, user)
}

func (lc *LoginController) LogOut(c *gin.Context) {
	cookie, err := c.Cookie(model.COOKIE_PREFIX_SESSION)

	err = lc.SessionService.RemoveSession(c, cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "user loged out")
}

func (lc *LoginController) FetchUser(c *gin.Context) {
	cookie, err := c.Cookie(model.COOKIE_PREFIX_SESSION)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid session"})
		return
	}

	var userId string
	userId, err = lc.SessionService.RetriveSession(c, cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	var user model.User
	user, err = lc.AccountService.FetchUser(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
