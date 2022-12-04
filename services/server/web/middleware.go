package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mozhu1024/virzz/logger"
	"github.com/mozhu1024/virzz/utils"
)

var middlewares = []gin.HandlerFunc{}

func RegisterMiddleware(ms ...gin.HandlerFunc) {
	middlewares = append(middlewares, ms...)
}

func corsMiddleware(c *gin.Context) {
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
		c.AbortWithStatusJSON(http.StatusBadRequest, Resp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	claims, err := utils.ParseToken(token)
	if err != nil {
		logger.Debug(token)
		logger.Error(err.Error())
		c.AbortWithStatusJSON(500, Resp{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	logger.Debug(claims)
	c.Set("token", claims.Token)
	c.Set("username", claims.Username)
	c.Set("jti", claims.Id)
	c.Next()
}

func init() {
	RegisterMiddleware(corsMiddleware, JWTAuthMiddleware)
}
