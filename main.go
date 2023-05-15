package main

import (
	"fmt"
	service2 "github.com/luojinqiang/slowcom-hik-sdk/app/common/service"
	"github.com/luojinqiang/slowcom-hik-sdk/config"
	"github.com/luojinqiang/slowcom-hik-sdk/http"
)

func main() {
	hikClient := &http.HikHttpClient{
		BaseUrl:      config.BaseUrl,
		ClientId:     "123456",
		ClientSecret: "321654",
	}
	//deviceControlRequest := service.DeviceControlRequest{HikClient: hikClient}
	//
	//err := deviceControlRequest.DoorControl("G70731081", "1002", "open")
	//
	//fmt.Println(`err1`, err)

	eventRequest := service2.EventRequest{HikClient: hikClient}

	consumerId, err := eventRequest.CreateConsumer(`1`)
	fmt.Println(`consumerId`, consumerId)
	fmt.Println(`err`, err)

}
