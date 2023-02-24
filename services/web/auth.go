package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virzz/logger"

	"github.com/virzz/virzz/services/models"
	"github.com/virzz/virzz/services/web/mw"
	"github.com/virzz/virzz/services/web/resp"
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
		c.JSON(400, resp.E(err.Error()))
		return
	}
	user, err := models.FindAuthByUsername(req.Username)
	if err != nil {
		logger.Debug(err)
		c.JSON(500, resp.E("Username is not exists!"))
		return
	}
	if utils.VerifyPassword(user.Password, req.Password) {
		token, err := utils.GenerateToken(user.Token, user.Username)
		if err != nil || token == "" {
			logger.Debug(err)
			c.JSON(500, resp.E("Could not generate token"))
			return
		}
		c.JSON(200, resp.S("Success", token))
		return
	}
	c.JSON(500, resp.E("Password error!"))
}

func AuthRegisterHandle(c *gin.Context) {
	var req LoginRegister
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, resp.E(err.Error()))
		return
	}
	password := utils.GeneratePassword(req.Password)
	_, err := models.NewAuth(req.Username, password, req.Email)
	if err != nil {
		logger.Debug(err)
		c.JSON(500, resp.E("Register error"))
		return
	}
	c.JSON(200, resp.S("Register success"))
}

func AuthRefreshHandle(c *gin.Context) {
	token, err := utils.GetHeaderToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(400, resp.E(err.Error()))
		return
	}
	reToken, err := utils.RefreshToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.E(err.Error()))
		return
	}
	c.Header("Token", reToken)
	c.JSON(200, resp.S("success", reToken))
}

func init() {
	RegisterRoute(func(g *gin.Engine) {
		authGroup := g.Group("/auth")
		{
			authGroup.POST("/login", AuthLoginHandle)
			authGroup.POST("/register", AuthRegisterHandle)
			authGroup.GET("/refresh", AuthRefreshHandle, mw.JWTAuthMiddleware)
		}
	})
}
