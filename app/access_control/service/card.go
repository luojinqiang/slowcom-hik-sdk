package service

import (
	"github.com/luojinqiang/slowcom-hik-sdk/app/access_control/entity"
	"github.com/luojinqiang/slowcom-hik-sdk/http"
)

type CardRequest struct {
	HikClient *http.HikHttpClient
}

// BatchAdd 批量添加
func (s *CardRequest) BatchAdd(list []*entity.CardAdd) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/basic/cards/batchCreate`, map[string]interface{}{
		`cards`: list,
	})
	return
}

// BatchDelete 批量删除卡号
func (s *CardRequest) BatchDelete(list []string) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/basic/cards/batchDelete`, map[string]interface{}{
		`cardNos`: list,
	})
	return
}

// PersonCardsDelete 根据人员编号删除对应的卡片
func (s *CardRequest) PersonCardsDelete(employeeNo string, list []string) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/basic/cards/batchDelete`, map[string]interface{}{
		`employeeNo`: employeeNo,
		`cardNos`:    list,
	})
	return
}
