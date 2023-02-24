package dns

import (
	"github.com/gin-gonic/gin"
	"github.com/virzz/virzz/services/web"
	"github.com/virzz/virzz/services/web/mw"
	"github.com/virzz/virzz/services/web/resp"
)

func GetDNSLogHandle(c *gin.Context) {
	// TODO: GetDNSLogHandle
	c.JSON(200, resp.S("todo"))
}

func DelDNSLogHandle(c *gin.Context) {
	// TODO: DelDNSLogHandle
	c.JSON(200, resp.S("todo"))
}

func todoHandle(c *gin.Context) {}

func init() {
	web.RegisterRoute(func(g *gin.Engine) {
		group := g.Group("/dns", mw.JWTAuthMiddleware)
		{
			// dns log
			group.GET("/log", GetDNSLogHandle)
			group.GET("/log/del/:ids", DelDNSLogHandle)
			// dns custom record
			group.GET("/record", todoHandle)          // get all
			group.GET("/record/new", todoHandle)      // new
			group.GET("/record/:id", todoHandle)      // update
			group.GET("/record/del/:ids", todoHandle) // delete
		}
	})
}
