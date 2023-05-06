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
		ClientId:     "e69535b3f3b04b1ea51437d92dcf7b80",
		ClientSecret: "e9aaa9408d6343a1945aa420074375dc",
	}
	deviceControlRequest := service.DeviceControlRequest{HikClient: hikClient}

	err := deviceControlRequest.DoorControl("G70731081", "1002", "open")

	fmt.Println(`err1`, err)

}
