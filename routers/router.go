package routers

import (
	"github.com/gilbertlim/member-service-go/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	m := r.Group("members")
	{
		m.POST("", api.CreateMember)
		m.GET("", api.GetMembers)
		m.GET("/:memberId", api.GetMember)
		m.DELETE("/:memberId", api.DeleteMember)
	}

	return r
}
