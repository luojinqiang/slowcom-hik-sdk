package entity

import (
	"github.com/luojinqiang/slowcom-hik-sdk/http"
)

// DeviceCreate 注册设备 注册设备到对应分组内。注册设备时首先会将设备添加到平台，
type DeviceCreate struct {
	DeviceSerial string `json:"deviceSerial"` //设备序列号
	GroupNo      string `json:"groupNo"`      // 组编号
	ValidateCode string `json:"validateCode"` //验证码
}

// DeviceUpdate 更新设备
type DeviceUpdate struct {
	DeviceName   string `json:"deviceName"`   //设备名称
	DeviceSerial string `json:"deviceSerial"` //设备号
}

// Device 设备
type Device struct {
	DeviceId      string `json:"deviceId"`      //设备id
	DeviceModel   string `json:"deviceModel"`   //设备型号
	DeviceName    string `json:"deviceName"`    //设备名称
	DeviceSerial  string `json:"deviceSerial"`  //设备序列号
	DeviceStatus  int    `json:"deviceStatus"`  //设备状态 0：离线 1：在线
	GroupId       string `json:"groupId"`       //组id
	IsEncrypt     int    `json:"isEncrypt"`     //设备加密状态：0-关闭，1-开启
	DeviceVersion string `json:"deviceVersion"` //设备固件版本号
}

// DevicePageRes 分页
type DevicePageRes struct {
	http.PageEntity
	Rows []*Device `json:"rows"`
}
