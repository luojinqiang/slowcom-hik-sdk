package entity

// CardAdd 新增卡片
type CardAdd struct {
	CardNo     string `json:"cardNo"`     //卡号(只支持纯数字)
	CardType   string `json:"cardType"`   //员工工号
	EmployeeNo string `json:"employeeNo"` //卡片类型，normalCard：普通卡，hijackCard: 胁迫卡，patrolCard：巡更卡，superCard： 超级卡
}
