package main

import (
	"slowcom-hik-sdk/app/common/entity"
	"slowcom-hik-sdk/app/common/service"
	"slowcom-hik-sdk/config"
	"slowcom-hik-sdk/http"
)

func main() {
	hikClient := &http.HikHttpClient{
		BaseUrl:      config.BaseUrl,
		ClientId:     "e69535b3f3b04b1ea51437d92dcf7b80",
		ClientSecret: "e9aaa9408d6343a1945aa420074375dc",
	}
	groupRequest := &service.GroupRequest{
		HikClient: hikClient,
	}
	_ = groupRequest.Add(&entity.GroupAdd{
		GroupName: "测试分组",
		GroupNo:   "分组编号",
	})

}
