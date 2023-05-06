package service

import (
	"encoding/json"
	"fmt"
	"slowcom-hik-sdk/app/access_control/entity"
	"slowcom-hik-sdk/gerror"
	"slowcom-hik-sdk/http"
)

type DeviceControlRequest struct {
	HikClient *http.HikHttpClient
}

// DoorControl 控制设备开门
// deviceSerial 设备序列号
// employeeNo 工号
// cmd 操作命令：open（开门），close（关门），alwaysOpen（常开），alwaysClose（常关），visitorCallLadder（访客呼梯），householdCallLadder（住户呼梯）
func (s *DeviceControlRequest) DoorControl(deviceSerial string, employeeNo string, cmd string) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/accessControl/remoteControl/actions/open`, map[string]interface{}{
		`deviceSerial`: deviceSerial,
		`employeeNo`:   employeeNo,
		`cmd`:          cmd,
	})
	return
}

// GetQrcode 根据卡号信息生成访客二维码
func (s *DeviceControlRequest) GetQrcode(cardNo string, effectTime string, expireTime string, openTimes int) (deviceControlGetQrcodeRes *entity.DeviceControlGetQrcodeRes, error error) {
	res, err := s.HikClient.Post(`/api/v1/community/access/visitors/actions/getQrcode`,
		fmt.Sprintf(`cardNo=%s&effectTime=%s&expireTime=%s&openTimes=%d`, cardNo, effectTime, expireTime, openTimes))
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &deviceControlGetQrcodeRes)
	return
}
