package api

import (
	"fmt"
	"github.com/gilbertlim/member-service-go/pkg/app"
	"github.com/gilbertlim/member-service-go/pkg/e"
	"github.com/gilbertlim/member-service-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MemberDto struct {
	MemberId string `json:"memberId" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Address  string `json:"address" binding:"required"`
}

func CreateMember(c *gin.Context) {
	appG := app.Gin{C: c}

	var memberDto MemberDto

	if err := c.ShouldBind(&memberDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("%v", err),
		})
	} else {
		memberService := service.Member{
			MemberId: memberDto.MemberId,
			Name:     memberDto.Name,
			Email:    memberDto.Email,
			Phone:    memberDto.Phone,
			Address:  memberDto.Address,
		}

		rowsAffected, err := memberService.CreateMember()
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

	memberService := service.Member{}

	members, err := memberService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = members

	appG.Response(http.StatusOK, e.SUCCESS, data)
}
