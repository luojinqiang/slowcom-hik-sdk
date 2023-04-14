package entity

// GroupAdd 添加组entity
type GroupAdd struct {
	GroupName string `json:"groupName"` // 分组名称
	GroupNo   string `json:"groupNo"`   // 分组编号
	ParentNo  string `json:"parentNo"`  // 父节点组编号
}

// GroupUpdate 更新组
type GroupUpdate struct {
	GroupName string `json:"groupName"` // 分组名称
	GroupNo   string `json:"groupNo"`   // 分组编号
}

// Group 组详细信息
type Group struct {
	GroupId   string `json:"groupId"`   // 组id
	GroupName string `json:"groupName"` // 分组名称
	GroupNo   string `json:"groupNo"`   // 分组编号
	ParentId  string `json:"parentId"`  // 父节点id
}
