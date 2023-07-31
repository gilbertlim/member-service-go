package main

import (
	"fmt"
	"github.com/gilbertlim/member-service-go/config"
	"github.com/gilbertlim/member-service-go/models"
	"github.com/gilbertlim/member-service-go/pkg/setting"
	"github.com/gilbertlim/member-service-go/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	gin.SetMode(config.RuntimeConf.Server.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := config.RuntimeConf.Server.ReadTimeout
	writeTimeout := config.RuntimeConf.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", config.RuntimeConf.Server.HttpPort)

	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
