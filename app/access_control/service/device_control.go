package service

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-hik-sdk/app/access_control/entity"
	"slowcom-hik-sdk/basic"
	"slowcom-hik-sdk/gerror"
)

type deviceControlRequest struct {
	basic.BaseRequest
}

var DeviceControlRequest = new(deviceControlRequest)

// DoorControl 控制设备开门
// deviceSerial 设备序列号
// employeeNo 工号
// cmd 操作命令：open（开门），close（关门），alwaysOpen（常开），alwaysClose（常关），visitorCallLadder（访客呼梯），householdCallLadder（住户呼梯）
func (s *deviceControlRequest) DoorControl(deviceSerial string, employeeNo string, cmd string) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/accessControl/permissionGroups/update`), map[string]interface{}{
		`deviceSerial`: deviceSerial,
		`employeeNo`:   employeeNo,
		`cmd`:          cmd,
	})
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// GetQrcode 根据卡号信息生成访客二维码
func (s *deviceControlRequest) GetQrcode(cardNo string, effectTime string, expireTime string, openTimes int) (deviceControlGetQrcodeRes *entity.DeviceControlGetQrcodeRes, error error) {
	res, err := httpclient.
		WithHeader("Content-Type", "application/x-www-form-urlencoded").Post(s.BuildUrl(`/api/v1/community/access/visitors/actions/getQrcode`),
		fmt.Sprintf(`cardNo=%s&effectTime=%s&expireTime=%s&openTimes=%d`, cardNo, effectTime, expireTime, openTimes))
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	hikResponse, err := s.CheckResponse(res)
	bytes, _ := json.Marshal(hikResponse.Data)
	err = json.Unmarshal(bytes, &deviceControlGetQrcodeRes)
	return
}
