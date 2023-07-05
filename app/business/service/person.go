package service

import (
	"encoding/json"
	"fmt"
	"github.com/luojinqiang/slowcom-hik-sdk/app/business/entity"
	"github.com/luojinqiang/slowcom-hik-sdk/http"
)

type PersonRequest struct {
	HikClient *http.HikHttpClient
}

// Add 添加
func (s *PersonRequest) Add(add *entity.PersonAdd) (personAddRes *entity.PersonAddRes, err error) {
	res, err := s.HikClient.PostJson(`/api/v1/open/basic/persons/create`, add)
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &personAddRes)
	return
}

// Update 更新
func (s *PersonRequest) Update(update *entity.PersonUpdate) (personUpdateRes *entity.PersonUpdateRes, err error) {
	mp := make(map[string]interface{})
	mp["personName"] = update.PersonName
	mp["employeeNo"] = update.EmployeeNo
	if update.PersonPhone != `` {
		mp["personPhone"] = update.PersonPhone
	}
	if update.FaceImageBase64 != `` {
		mp["faceImageBase64"] = update.FaceImageBase64
	}
	mp["verifyImage"] = update.VerifyImage
	res, err := s.HikClient.PostJson(`/api/v1/open/basic/persons/update`, mp)
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &personUpdateRes)
	return
}

// Delete 删除人员
func (s *PersonRequest) Delete(employeeNo string) (err error) {
	_, err = s.HikClient.PostJson(fmt.Sprintf(`/api/v1/open/basic/persons/delete?employeeNo=%s`, employeeNo), nil)
	return
}

// Get 获取人员单个信息
func (s *PersonRequest) Get(employeeNo string) (person *entity.Person, err error) {
	res, err := s.HikClient.Get(fmt.Sprintf(`/api/v1/open/basic/persons/get?employeeNo=%s`, employeeNo))
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &person)
	return
}

// FaceUpdate 更新人员的人脸
// employeeNo 人员编号
// faceImageBase64 人脸的base64字符串
// verifyImage 是否校验人脸质量，默认校验人脸质量
func (s *PersonRequest) FaceUpdate(employeeNo string, faceImageBase64 string, verifyImage bool) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/basic/faces/update`, map[string]interface{}{
		`employeeNo`:      employeeNo,
		`faceImageBase64`: faceImageBase64,
		`verifyImage`:     verifyImage,
	})
	return
}

// FaceDelete 删除人脸
func (s *PersonRequest) FaceDelete(employeeNo string) (err error) {
	_, err = s.HikClient.Post(fmt.Sprintf(`/api/v1/open/basic/faces/delete?employeeNo=%s`, employeeNo), nil)
	return
}
