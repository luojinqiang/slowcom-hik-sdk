package main

import (
	"fmt"
	"slowcom-hik-sdk/app/access_control/service"
	"slowcom-hik-sdk/config"
	"slowcom-hik-sdk/http"
)

func main() {
	hikClient := &http.HikHttpClient{
		BaseUrl:      config.BaseUrl,
		ClientId:     "123456465",
		ClientSecret: "134313213433",
	}
	deviceControlRequest := service.DeviceControlRequest{HikClient: hikClient}

	err := deviceControlRequest.DoorControl("G70731081", "1002", "open")

	fmt.Println(`err1`, err)

}