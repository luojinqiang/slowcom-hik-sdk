package entity

import "time"

// PermissionGroupAddRes 添加权限组
type PermissionGroupAddRes struct {
	GroupId   string `json:"groupId"`   //组ID
	GroupName string `json:"groupName"` // 组名称
}

// PermissionGroup 权限组
type PermissionGroup struct {
	GroupId    string    `json:"groupId"`
	GroupName  string    `json:"groupName"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
}

// PermissionGroupPageRes 分页查询
type PermissionGroupPageRes struct {
	PageNo   int                `json:"pageNo"`
	PageSize int                `json:"pageSize"`
	Total    int                `json:"total"`
	Rows     []*PermissionGroup `json:"rows"`
}

// PermissionGroupPersonPageRes 权限组人员列表
type PermissionGroupPersonPageRes struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	Rows     []struct {
		EmployeeNo string `json:"employeeNo"`
	} `json:"rows"`
}

// PermissionGroupDevicePageRes 权限组设备列表
type PermissionGroupDevicePageRes struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	Rows     []struct {
		DeviceSerial  string `json:"deviceSerial"`
		HasCapability bool   `json:"hasCapability"`
	} `json:"rows"`
}
