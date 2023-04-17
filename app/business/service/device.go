package service

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-hik-sdk/app/business/entity"
	"slowcom-hik-sdk/basic"
	"slowcom-hik-sdk/gerror"
)

type deviceRequest struct {
	basic.BaseRequest
}

var DeviceRequest = new(deviceRequest)

// Create 注册设备到对应分组内。注册设备时首先会将设备添加到平台，然后异步同步设备通道。如果设备添加成功而同步设备通道失败，则可以先获取设备列表信息，再手动调用通道同步接口同步设备下的通道
func (s *deviceRequest) Create(create *entity.DeviceCreate) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/devices/create`), create)
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// Delete 从某一分组内删除设备
func (s *deviceRequest) Delete(deviceSerial string) (err error) {
	res, err := httpclient.Post(s.BuildUrl(fmt.Sprintf(`/api/v1/open/basic/devices/delete?deviceSerial=%s`, deviceSerial)), nil)
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// Update 该接口用于修改设备名称
func (s *deviceRequest) Update(deviceUpdate *entity.DeviceUpdate) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/groups/update`), deviceUpdate)
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// Get 该接口用于根据设备序列号获取单个设备详细信息
func (s *deviceRequest) Get(deviceSerial string) (device *entity.Device, err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf(`/api/v1/open/basic/devices/get?deviceSerial=%s`, deviceSerial)))
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	hikResponse, err := s.CheckResponse(res)
	bytes, _ := json.Marshal(hikResponse.Data)
	err = json.Unmarshal(bytes, &device)
	return
}

// Page 该接口用于查询某组下设备列表信息 分页
func (s *deviceRequest) Page(groupNo string, pageNo, pageSize int) (page *entity.DevicePage, err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf(`https://api2.hik-cloud.com/api/v1/open/basic/devices/list?groupNo=%s&pageNo=%d&pageSize=%d`, groupNo, pageNo, pageSize)))
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	hikResponse, err := s.CheckResponse(res)
	bytes, _ := json.Marshal(hikResponse.Data)
	err = json.Unmarshal(bytes, &page)
	return
}

// OffLineConfirm 设备下线确认
func (s *deviceRequest) OffLineConfirm(deviceSerial string) (err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf(`/v1/carrier/wing/endpoint/confirm/right/offlineconfirm?deviceSerial=%s`, deviceSerial)))
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)

	return
}

// OnLineConfirm 设备上线确认
func (s *deviceRequest) OnLineConfirm(deviceSerial string) (err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf(`/v1/carrier/wing/endpoint/confirm/right/onlineconfirm?deviceSerial=%s`, deviceSerial)))
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}
