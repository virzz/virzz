package web

import "github.com/gin-gonic/gin"

type Router func(*gin.Engine)

var routers = []Router{}

func RegisterRoute(routes ...Router) {
	routers = append(routers, routes...)
}

// func init() {
// 	http.RegisterRoute(func(g *gin.Engine) {
// 		v1Group := g.Group("/v1")
// 		{
// 			v1Group.GET("/test", TestHandle)
// 		}
// 	})
// }
