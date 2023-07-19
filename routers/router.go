package routers

import (
	"github.com/gilbertlim/member-service-go/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", api.GetMembers)
	r.GET("/:memberId", api.GetMembers)

	return r
}
