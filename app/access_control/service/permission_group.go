package service

import (
	"encoding/json"
	"fmt"
	"github.com/luojinqiang/slowcom-hik-sdk/app/access_control/entity"
	"github.com/luojinqiang/slowcom-hik-sdk/http"
)

type PermissionGroupRequest struct {
	HikClient *http.HikHttpClient
}

// Add 权限组新增 权限组数量上限与门禁设备接入授权数量一致，默认为500个；权限组名称不支持除了.、-_[]【】()#@~<>以外的特殊字符 16位；
func (s *PermissionGroupRequest) Add(groupName string) (permissionGroupAddRes *entity.PermissionGroupAddRes, err error) {
	res, err := s.HikClient.PostJson(`/api/v1/open/accessControl/permissionGroups/create`, map[string]interface{}{
		`groupName`: groupName,
	})
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &permissionGroupAddRes)
	return
}

// Update 权限组更新
func (s *PermissionGroupRequest) Update(groupId, groupName string) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/accessControl/permissionGroups/update`, map[string]interface{}{
		`groupId`:   groupId,
		`groupName`: groupName,
	})
	return
}

// Delete 删除权限组 如果权限组包含设备或人员，同时解绑所有人员和设备；删除权限组，会同时移除权限组下人员和设备下发记录以及已下发到设备的权限（静默移除权限，不会通过消息通过告知用户）；
func (s *PermissionGroupRequest) Delete(groupId string) (err error) {
	_, err = s.HikClient.Post(fmt.Sprintf(`/api/v1/open/accessControl/permissionGroups/delete?groupId=%s`, groupId), nil)
	return
}

// Page 分页获取权限组列表
func (s *PermissionGroupRequest) Page(pageNo, pageSize int) (list []*entity.PermissionGroupPageRes, err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/accessControl/permissionGroups/actions/page`, map[string]interface{}{
		`pageNo`:   pageNo,
		`pageSize`: pageSize,
	})
	return
}

// BindPerson 权限组绑定人员
// groupId 权限组ID
// employeeNos 关联的人员编号列表
// autoIssue 是否自动下发,true或者不填默认调用本接口后自动下发权限组,false为手动下发(需要调用根据权限组下发接口下发)
func (s *PermissionGroupRequest) BindPerson(groupId string, employeeNos []string, autoIssue bool) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/accessControl/permissionGroups/actions/addPersons`, map[string]interface{}{
		`groupId`:     groupId,
		`employeeNos`: employeeNos,
		`autoIssue`:   autoIssue,
	})
	return
}

// ReleasePerson 将人员从权限组移除 移除人员的权限下发记录，同时清理已下发到设备上的用户权限；如果是删除人员，不需要调用该接口，会自动移除人员相关权限；
func (s *PermissionGroupRequest) ReleasePerson(groupId string, employeeNos []string) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/accessControl/permissionGroups/actions/removePersons`, map[string]interface{}{
		`groupId`:     groupId,
		`employeeNos`: employeeNos,
	})
	return
}

// GetEmployeeNosPage 权限组获取人员编号列表 分页,返回人员对象列表
func (s *PermissionGroupRequest) GetEmployeeNosPage(pageNo int, pageSize int, groupId string) (permissionGroupPersonPageRes *entity.PermissionGroupPersonPageRes, err error) {
	res, err := s.HikClient.PostJson(`/api/v1/open/accessControl/permissionGroups/actions/getEmployeeNos`, map[string]interface{}{
		`pageNo`:   pageNo,
		`pageSize`: pageSize,
		`groupId`:  groupId,
	})
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &permissionGroupPersonPageRes)
	return
}

// BindDevice 权限组绑定设备 将设备添加到权限组中；权限组中的人员权限将自动下发到该设备
// checkCapability 绑定时校验是否有能力集。该配置建议开启，用于过滤没有能力集的设备并告知开发者。
// BindDevice 关联的设备序列号列表
// BindDevice 是否自动下发,true或者不填默认调用本接口后自动下发权限组,false为手动下发(需要调用根据权限组下发接口下发)
func (s *PermissionGroupRequest) BindDevice(groupId string, checkCapability bool, deviceSerials []string, autoIssue bool) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/accessControl/permissionGroups/actions/addDevices`, map[string]interface{}{
		`groupId`:         groupId,
		`checkCapability`: checkCapability,
		`deviceSerials`:   deviceSerials,
		`autoIssue`:       autoIssue,
	})
	return
}

// ReleaseDevice 解绑设备
func (s *PermissionGroupRequest) ReleaseDevice(groupId string, deviceSerials []string) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/accessControl/permissionGroups/actions/removeDevices`, map[string]interface{}{
		`groupId`:       groupId,
		`deviceSerials`: deviceSerials,
	})
	return
}

// GetDeviceSerialsPage 权限组获取设备序列号列表 分页,返回设备序列号和是否有能力集对象列表；
func (s *PermissionGroupRequest) GetDeviceSerialsPage(pageNo int, pageSize int, groupId string) (PermissionGroupDevicePageRes *entity.PermissionGroupDevicePageRes, err error) {
	res, err := s.HikClient.PostJson(`/api/v1/open/accessControl/permissionGroups/actions/getEmployeeNos`, map[string]interface{}{
		`pageNo`:   pageNo,
		`pageSize`: pageSize,
		`groupId`:  groupId,
	})
	if err != nil {
		return
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &PermissionGroupDevicePageRes)
	return
}

// PermissionIssue 根据权限组下发权限 该接口禁止并发调用！！！
// 通过该方式会下发权限组中未下发过的人员权限和下发失败的权限记录；
func (s *PermissionGroupRequest) PermissionIssue(groupId string) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/accessControl/allots/actions/issuedByGroup`, map[string]interface{}{
		`groupId`: groupId,
	})
	return
}
