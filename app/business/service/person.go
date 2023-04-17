package service

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-hik-sdk/app/business/entity"
	"slowcom-hik-sdk/basic"
	"slowcom-hik-sdk/gerror"
)

type personRequest struct {
	basic.BaseRequest
}

var PersonRequest = new(personRequest)

// Add 添加
func (s *personRequest) Add(add *entity.PersonAdd) (personAddRes *entity.PersonAddRes, err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/persons/create`), add)
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	hikResponse, err := s.CheckResponse(res)
	bytes, _ := json.Marshal(hikResponse.Data)
	err = json.Unmarshal(bytes, &personAddRes)
	return
}

// Update 更新
func (s *personRequest) Update(update *entity.PersonUpdate) (personUpdateRes *entity.PersonUpdateRes, err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/persons/update`), update)
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	hikResponse, err := s.CheckResponse(res)
	bytes, _ := json.Marshal(hikResponse.Data)
	err = json.Unmarshal(bytes, &personUpdateRes)
	return
}

// Delete 删除人员
func (s *personRequest) Delete(employeeNo string) (rr error) {
	res, err := httpclient.Post(s.BuildUrl(fmt.Sprintf(`/api/v1/open/basic/persons/delete?employeeNo=%s`, employeeNo)), nil)
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// Get 获取人员单个信息
func (s *personRequest) Get(employeeNo string) (person *entity.Person, err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf(`/api/v1/open/basic/persons/get?employeeNo=%s`, employeeNo)))
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	hikResponse, err := s.CheckResponse(res)
	bytes, _ := json.Marshal(hikResponse.Data)
	err = json.Unmarshal(bytes, &person)
	return
}

// FaceUpdate 更新人员的人脸
// employeeNo 人员编号
// faceImageBase64 人脸的base64字符串
// verifyImage 是否校验人脸质量，默认校验人脸质量
func (s *personRequest) FaceUpdate(employeeNo string, faceImageBase64 string, verifyImage bool) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/faces/update`), map[string]interface{}{
		`employeeNo`:      employeeNo,
		`faceImageBase64`: faceImageBase64,
		`verifyImage`:     verifyImage,
	})
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// FaceDelete 删除人脸
func (s *personRequest) FaceDelete(employeeNo string) (err error) {
	res, err := httpclient.Post(s.BuildUrl(fmt.Sprintf(`/api/v1/open/basic/faces/delete?employeeNo=%s`, employeeNo)), nil)
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}
