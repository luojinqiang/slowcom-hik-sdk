package service

import (
	"encoding/json"
	"fmt"
	"github.com/luojinqiang/slowcom-hik-sdk/app/common/entity"
	"github.com/luojinqiang/slowcom-hik-sdk/http"
)

type EventRequest struct {
	HikClient *http.HikHttpClient
}

// CreateConsumer 创建消费者 https://pic.hik-cloud.com/opencustom/apidoc/online/open/9e27913442f548a984e86b7d7978c53f.html?timestamp=1683343883586#%E5%88%9B%E5%BB%BA%E6%B6%88%E8%B4%B9%E8%80%85SPU1hk
// groupNo 消息分组编号，1，2，3共3个分组
func (s *EventRequest) CreateConsumer(groupNo string) (consumerId string, err error) {
	res, err := s.HikClient.Post(fmt.Sprintf(`/api/v1/mq/consumer/group%s`, groupNo), nil)
	if err != nil {
		return
	}
	var data map[string]string
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &data)
	consumerId = data["consumerId"]
	return
}

// MessageConsumer 消费消息
// autoCommit 是否自动提交偏移量
// consumerId 消费者ID
func (s *EventRequest) MessageConsumer(autoCommit bool, consumerId string) (list []*entity.EventMsg, err error) {
	res, err := s.HikClient.Post(`/api/v1/mq/consumer/messages`,
		fmt.Sprintf(`autoCommit=%t&consumerId=%s`, autoCommit, consumerId))
	if err != nil {
		return nil, err
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	return
}

// SubmitOffsets 提交消息偏移量
// consumerId 消费者ID
func (s *EventRequest) SubmitOffsets(consumerId string) (err error) {
	_, err = s.HikClient.Post(`/api/v1/mq/consumer/offsets`,
		fmt.Sprintf(`consumerId=%s`, consumerId))
	return
}
