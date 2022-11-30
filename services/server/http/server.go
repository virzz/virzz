package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Host  string
	Port  int
	Debug bool
}

type Resp struct {
	Code int
	Msg  string
	Data interface{} `json:"data,omitempty"`
}

// NewWebServer - New Web Server
func NewWebServer(conf *Config) *http.Server {
	if conf.Debug {
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

	s := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}
