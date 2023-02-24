package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/virzz/virzz/services/web/mw"
	"github.com/virzz/virzz/services/web/resp"
)

// NewServer - New Web Server
func NewServer(debug bool) *http.Server {

	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, resp.S("pong"))
	})

	// Middleware
	engine.Use(mw.Middlewares...)

	// Router
	for _, route := range routers {
		route(engine)
	}

	return &http.Server{
		Addr:           fmt.Sprintf("%s:%d", viper.GetString("web.host"), viper.GetInt("web.port")),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
