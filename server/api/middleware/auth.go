package middleware

import (
	"github/tekeoglan/discord-clone/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(ss model.SessionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, _ := c.Cookie(model.COOKIE_PREFIX_SESSION)
		if sessionId == "" {
			c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "Session not provided."})
			c.Abort()
			return
		}

		val, _ := ss.RetriveSession(c, sessionId)
		if val == "" {
			c.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "Session invalid"})
			c.Abort()
			return
		}

		c.Set(model.CONTEXT_USER_KEY, val)

		c.Next()
	}
}
