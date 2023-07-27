package api

import (
	"fmt"
	"github.com/gilbertlim/member-service-go/pkg/app"
	"github.com/gilbertlim/member-service-go/pkg/dto"
	"github.com/gilbertlim/member-service-go/pkg/e"
	"github.com/gilbertlim/member-service-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateMember(c *gin.Context) {
	appG := app.Gin{C: c}

	var memberDto dto.MemberDto

	if err := c.ShouldBind(&memberDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
	} else {
		rowsAffected, err := service.CreateMember(memberDto)
		data := make(map[string]interface{})

		if err != nil {
			data["rowsAffected"] = 0
			data["errorMessage"] = err.Error()
			appG.Response(http.StatusInternalServerError, e.ERROR, data)
			return
		}

		data["rowsAffected"] = rowsAffected
		appG.Response(http.StatusOK, e.SUCCESS, data)
	}
}

func GetMembers(c *gin.Context) {
	appG := app.Gin{C: c}

	members, err := service.GetMembers()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = members

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func GetMember(c *gin.Context) {
	appG := app.Gin{C: c}

	memberId := c.Param("memberId")

	member, err := service.GetMember(memberId)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
	}

	data := make(map[string]interface{})
	data["lists"] = member

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func DeleteMember(c *gin.Context) {
	appG := app.Gin{C: c}

	data := make(map[string]interface{})

	memberId := c.Param("memberId")
	rowsAffected, err := service.DeleteMember(memberId)
	if err != nil {
		data["rowsAffected"] = 0
		data["errorMessage"] = err.Error()
		appG.Response(http.StatusInternalServerError, e.ERROR, data)
		return
	}

	data["rowsAffected"] = rowsAffected
	appG.Response(http.StatusOK, e.SUCCESS, data)
}
