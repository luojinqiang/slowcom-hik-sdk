package service

import (
	"slowcom-hik-sdk/app/access_control/entity"
	"slowcom-hik-sdk/gerror"
	"slowcom-hik-sdk/http"
)

type CardRequest struct {
	hikClient *http.HikHttpClient
}

// BatchAdd 批量添加
func (s *CardRequest) BatchAdd(list []*entity.CardAdd) (err error) {
	_, err = s.hikClient.PostJson(`/api/v1/open/basic/cards/batchCreate`, map[string]interface{}{
		`cards`: list,
	})
	return
}

// BatchDelete 批量删除卡号
func (s *CardRequest) BatchDelete(list []string) (err error) {
	_, err = s.hikClient.PostJson(`/api/v1/open/basic/cards/batchDelete`, map[string]interface{}{
		`cardNos`: list,
	})
	if err != nil {
		return gerror.ErrIs系统异常
	}
	return
}

// PersonCardsDelete 根据人员编号删除对应的卡片
func (s *CardRequest) PersonCardsDelete(employeeNo string, list []string) (err error) {
	_, err = s.hikClient.PostJson(`/api/v1/open/basic/cards/batchDelete`, map[string]interface{}{
		`employeeNo`: employeeNo,
		`cardNos`:    list,
	})
	return
}
