package entity

// PersonAdd 添加人员
type PersonAdd struct {
	EmployeeNo      string `json:"employeeNo"`      //人员编号，同一个租户下不能重复，只支持英文、数字
	PersonName      string `json:"personName"`      //人员姓名
	PersonPhone     string `json:"personPhone"`     //人员手机号
	FaceImageBase64 string `json:"faceImageBase64"` //base64编码的人脸图片，图片大小需小于200kB，数据需去除base64前缀
	VerifyImage     bool   `json:"verifyImage"`     //是否进行人脸质量校验，默认：true；如果调用方能保证人脸的质量符合标准那么可以配置为false,注意：如果调用方不能保证人脸质量，而且不对人脸质量校验，那么图片质量差的人脸可能会下发到设备，可能导致人脸下发失败
}

// PersonAddRes 添加人员返回
type PersonAddRes struct {
	EmployeeNo      string   `json:"employeeNo"`
	PersonName      string   `json:"personName"`
	PersonPhone     string   `json:"personPhone"`
	FaceUrl         string   `json:"faceUrl"`
	PersonType      string   `json:"personType"`
	FloorNo         int      `json:"floorNo"`
	RoomNo          int      `json:"roomNo"`
	ValidBeginTime  string   `json:"validBeginTime"`
	ValidEndTime    string   `json:"validEndTime"`
	PlanTemplateNos []int    `json:"planTemplateNos"`
	DynamicCode     string   `json:"dynamicCode"`
	MaxOpenDoorTime int      `json:"maxOpenDoorTime"`
	CallNumbers     []string `json:"callNumbers"`
	FloorNumbers    []int    `json:"floorNumbers"`
	LocalUIRight    int      `json:"localUIRight"`
}

// PersonUpdate 更新人员
type PersonUpdate struct {
	EmployeeNo      string `json:"employeeNo"`
	PersonName      string `json:"personName"`
	PersonPhone     string `json:"personPhone"`
	FaceImageBase64 string `json:"faceImageBase64"`
	VerifyImage     bool   `json:"verifyImage"`
}

// PersonUpdateRes 更新人员返回
type PersonUpdateRes struct {
	PersonStatisticsId string `json:"personStatisticsId"` //人员下发结果记录id
	FaceStatisticsId   string `json:"faceStatisticsId"`   //人脸下发结果记录id
}

// Person 获取人员信息
type Person struct {
	EmployeeNo      string `json:"employeeNo"`      //人员姓名
	PersonName      string `json:"personName"`      //人员编号，同一个租户下不能重复，只支持英文、数字
	PersonPhone     string `json:"personPhone"`     //人员手机号
	FaceUrl         string `json:"faceUrl"`         //人员类型， normal：普通人，visitor:访客，blackList:黑名单，默认：normal
	PersonType      string `json:"personType"`      //人脸图片url，有效期为24小时
	FloorNo         int    `json:"floorNo"`         //楼层号，取值1-999
	RoomNo          int    `json:"roomNo"`          //房间号，取值1-99
	ValidBeginTime  string `json:"validBeginTime"`  //人员权限开始时间，日期格式只支持如下两种格式：(1)权限精确到天，示例：2020-01-01 (2)权限精确到秒并加上时区，示例：2019-01-29T00:00:59+08:00
	ValidEndTime    string `json:"validEndTime"`    //人员权限结束时间，日期格式只支持如下两种格式：(1)权限精确到天，示例：2020-01-01 (2)权限精确到秒并加上时区，例如：2019-01-29T00:00:59+08:00 ; 结束时间需早于2037-12-31
	PlanTemplateNos string `json:"planTemplateNos"` //人员计划模板列表，不填时默认为[1] 代表计划模板 1(全天候模板 )； 调用计划模板修改接口，修改计划模板1可能会影响默认计划模板人员的权限。
}
