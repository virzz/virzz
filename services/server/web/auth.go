package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virzz/logger"
	"github.com/virzz/virzz/services/server/models"
	"github.com/virzz/virzz/utils"
)

// LoginRegister - Login or Register Data Struct
type LoginRegister struct {
	Username string `form:"username" json:"username" binding:"required" example:"mozhu1024"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" example:"mozhu233@outlook.com"`
}

func AuthLoginHandle(c *gin.Context) {
	var req LoginRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, Resp{Code: -1, Msg: err.Error()})
		return
	}
	user, err := models.FindAuthByUsername(req.Username)
	if err != nil {
		logger.Debug(err)
		c.JSON(500, Resp{Code: -1, Msg: "Username is not exists!"})
		return
	}
	if utils.VerifyPassword(user.Password, req.Password) {
		token, err := utils.GenerateToken(user.Token, user.Username)
		if err != nil || token == "" {
			logger.Debug(err)
			c.JSON(500, Resp{Code: -1, Msg: "Could not generate token"})
			return
		}
		c.JSON(200, Resp{Code: 0, Msg: "Success", Data: token})
		return
	}
	c.JSON(500, Resp{Code: -1, Msg: "Password error!"})
}

func AuthRegisterHandle(c *gin.Context) {
	var req LoginRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, Resp{Code: -1, Msg: err.Error()})
		return
	}
	password := utils.GeneratePassword(req.Password)
	_, err := models.NewAuth(req.Username, password, req.Email)
	if err != nil {
		logger.Debug(err)
		c.JSON(500, Resp{Code: -1, Msg: "Register error"})
		return
	}
	c.JSON(200, Resp{Code: 0, Msg: "Register success"})
}

func AuthRefreshHandle(c *gin.Context) {
	token, err := utils.GetHeaderToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(400, Resp{Code: -1, Msg: err.Error()})
		return
	}
	reToken, err := utils.RefreshToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Resp{Code: -1, Msg: err.Error()})
		return
	}
	c.Header("Token", reToken)
	c.JSON(200, Resp{Code: 0, Msg: "success", Data: reToken})
}

func init() {
	RegisterRoute(func(g *gin.Engine) {
		authGroup := g.Group("/auth")
		{
			authGroup.POST("/login", AuthLoginHandle)
			authGroup.POST("/register", AuthRegisterHandle)
			authGroup.GET("/refresh", AuthRefreshHandle, JWTAuthMiddleware)
		}
	})
}
