package entity

import "github.com/luojinqiang/slowcom-hik-sdk/http"

// EventMsg 消息类型
type EventMsg struct {
	MsgId     string `json:"msgId"`     //消息ID ，唯一ID
	MsgType   string `json:"msgType"`   //消息类型，参见海康文档
	Content   string `json:"content"`   //消息内容，参见文档
	Timestamp int64  `json:"timestamp"` // 消息发送到消息通道时的时间戳
}

// EventMsgResponse 事件消息返回response
type EventMsgResponse struct {
	HikResponse *http.HikResponse
	List        []*EventMsg
}
