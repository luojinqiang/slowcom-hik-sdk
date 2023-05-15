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
		ClientId:     "1232456",
		ClientSecret: "321654",
	}
	//deviceControlRequest := service.DeviceControlRequest{HikClient: hikClient}
	//
	//err := deviceControlRequest.DoorControl("G70731081", "1002", "open")
	//
	//fmt.Println(`err1`, err)

	eventRequest := service2.EventRequest{HikClient: hikClient}

	//consumerId, err := eventRequest.CreateConsumer(`1`)
	//fmt.Println(`consumerId`, consumerId)
	//fmt.Println(`err`, err)
	list, err := eventRequest.MessageConsumer(true, "f408b39cf3b54397b99bdcb0a19ed46f")
	fmt.Println(`err`, err)
	fmt.Println(`list`, list)

	//deviceControlRequest := service.DeviceControlRequest{HikClient: hikClient}
	//url, err := deviceControlRequest.GetQrcode("1001", "2023-05-15 00:00:00", "2023-05-16 00:00:00", 1)
	//fmt.Println(`err`, err)
	//fmt.Println(`url`, url)
}
