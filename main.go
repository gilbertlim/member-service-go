package main

import (
	"fmt"
	"github.com/gilbertlim/member-service-go/pkg/setting"
	"github.com/gilbertlim/member-service-go/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	fmt.Println(endPoint)

	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
