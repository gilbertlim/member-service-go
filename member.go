package main

import (
	"github.com/gilbertlim/member-service-go/pkg/app"
	"github.com/gilbertlim/member-service-go/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMembers(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"message": "good test",
	})
}
