package mw

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/services/server/web/resp"
	"github.com/virzz/virzz/utils"
)

var Middlewares = []gin.HandlerFunc{
	CorsMiddleware,
	JWTAuthMiddleware,
}

func RegisterMiddleware(ms ...gin.HandlerFunc) {
	Middlewares = append(Middlewares, ms...)
}

func CorsMiddleware(c *gin.Context) {
	c.Header("Server", "Go-Gin-1.8.1")

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
	c.Header("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Accept")
	c.Header("Access-Control-Max-Age", "1800")
	if strings.ToUpper(c.Request.Method) == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}
	c.Next()
}

func JWTAuthMiddleware(c *gin.Context) {
	token, err := utils.GetHeaderToken(c.GetHeader("Authorization"))
	if err != nil {
		logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, resp.E(err.Error()))
		return
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		logger.Debug(token)
		logger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, resp.E(err.Error()))
		return
	}
	logger.Debug(claims)
	c.Set("token", claims.Token)
	c.Set("username", claims.Username)
	c.Set("jti", claims.ID)
	c.Next()
}
