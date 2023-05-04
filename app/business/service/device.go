package service

import (
	"encoding/json"
	"fmt"
	"slowcom-hik-sdk/app/business/entity"
	"slowcom-hik-sdk/gerror"
	"slowcom-hik-sdk/http"
)

type DeviceRequest struct {
	hikClient *http.HikHttpClient
}

// Create 注册设备到对应分组内。注册设备时首先会将设备添加到平台，然后异步同步设备通道。如果设备添加成功而同步设备通道失败，则可以先获取设备列表信息，再手动调用通道同步接口同步设备下的通道
func (s *DeviceRequest) Create(create *entity.DeviceCreate) (err error) {
	_, err = s.hikClient.PostJson(`/api/v1/open/basic/devices/create`, create)
	return
}

// Delete 从某一分组内删除设备
func (s *DeviceRequest) Delete(deviceSerial string) (err error) {
	_, err = s.hikClient.Post(fmt.Sprintf(`/api/v1/open/basic/devices/delete?deviceSerial=%s`, deviceSerial), nil)
	return
}

// Update 该接口用于修改设备名称
func (s *DeviceRequest) Update(deviceUpdate *entity.DeviceUpdate) (err error) {
	_, err = s.hikClient.PostJson(`/api/v1/open/basic/groups/update`, deviceUpdate)
	return
}

// Get 该接口用于根据设备序列号获取单个设备详细信息
func (s *DeviceRequest) Get(deviceSerial string) (device *entity.Device, err error) {
	res, err := s.hikClient.Get(fmt.Sprintf(`/api/v1/open/basic/devices/get?deviceSerial=%s`, deviceSerial))
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &device)
	return
}

// Page 该接口用于查询某组下设备列表信息 分页
func (s *DeviceRequest) Page(groupNo string, pageNo, pageSize int) (page *entity.DevicePageRes, err error) {
	res, err := s.hikClient.Get(fmt.Sprintf(`https://api2.hik-cloud.com/api/v1/open/basic/devices/list?groupNo=%s&pageNo=%d&pageSize=%d`, groupNo, pageNo, pageSize))
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &page)
	return
}

// OffLineConfirm 设备下线确认
func (s *DeviceRequest) OffLineConfirm(deviceSerial string) (err error) {
	_, err = s.hikClient.Get(fmt.Sprintf(`/v1/carrier/wing/endpoint/confirm/right/offlineconfirm?deviceSerial=%s`, deviceSerial))
	return
}

// OnLineConfirm 设备上线确认
func (s *DeviceRequest) OnLineConfirm(deviceSerial string) (err error) {
	_, err = s.hikClient.Get(fmt.Sprintf(`/v1/carrier/wing/endpoint/confirm/right/onlineconfirm?deviceSerial=%s`, deviceSerial))
	return
}
