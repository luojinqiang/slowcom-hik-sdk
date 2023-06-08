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

// PersonType 人员类型
type PersonType string

const (
	PersonTypIsNormal    = "normal"
	PersonTypIsVisitor   = "visitor"
	PersonTypIsBlackList = "blackList"
)

// BindPersonParams 权限组绑定人员参数
// GroupId 权限组ID
// EmployeeNo 人员编号
// AutoIssue 是否自动下发权限
// Mobile 人员手机号
// PersonType 人员类型
// ValidBeginTime 人员权限开始时间，日期格式只支持如下两种格式：(1)权限精确到天，默认为当前时间，示例：2020-01-01 (2)权限精确到秒并加上时区，示例：2019-01-29T00:00:59+08:00
// ValidEndTime 人员权限结束时间，日期格式只支持如下两种格式：(1)权限精确到天，默认为2037-12-19T16:00:00+08:00，示例：2020-01-01 (2)权限精确到秒并加上时区，例如：2019-01-29T00:00:59+08:00 ; 结束时间需早于2037-12-31
// MaxOpenDoorTime 人员最大开门次数，0-代表无次数限制，默认无次数限制；门禁设备支持该字段，可视对讲设备不支持；只有当personType=visitor时该参数才会生效
type BindPersonParams struct {
	GroupId         string     `json:"groupId"`
	EmployeeNo      string     `json:"employeeNo"`
	AutoIssue       bool       `json:"autoIssue"`
	Mobile          string     `json:"mobile"`
	PersonType      PersonType `json:"personType"`
	ValidBeginTime  string     `json:"validBeginTime"`
	ValidEndTime    string     `json:"validEndTime"`
	MaxOpenDoorTime int        `json:"maxOpenDoorTime"`
}
