package api

import (
	"github.com/gilbertlim/member-service-go/pkg/app"
	"github.com/gilbertlim/member-service-go/pkg/e"
	"github.com/gilbertlim/member-service-go/service/member_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMembers(c *gin.Context) {
	appG := app.Gin{C: c}

	memberService := member_service.Member{}

	members, err := memberService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = members

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
