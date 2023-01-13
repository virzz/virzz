package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Resp struct {
	Code int
	Msg  string
	Data interface{} `json:"data,omitempty"`
}

// NewWebServer - New Web Server
func NewWebServer(debug bool) *http.Server {

	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
	}
	engine := gin.Default()

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, Resp{Code: 0, Msg: "pong"})
	})

	// Middleware
	engine.Use(middlewares...)

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
