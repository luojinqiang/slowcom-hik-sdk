package service

import (
	"github.com/ddliu/go-httpclient"
	"slowcom-hik-sdk/app/access_control/entity"
	"slowcom-hik-sdk/basic"
	"slowcom-hik-sdk/gerror"
)

type cardRequest struct {
	basic.BaseRequest
}

var CardRequest = new(cardRequest)

// BatchAdd 批量添加
func (s *cardRequest) BatchAdd(list []*entity.CardAdd) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/cards/batchCreate`), map[string]interface{}{
		`cards`: list,
	})
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// BatchDelete 批量删除卡号
func (s *cardRequest) BatchDelete(list []string) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/cards/batchDelete`), map[string]interface{}{
		`cardNos`: list,
	})
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// PersonCardsDelete 根据人员编号删除对应的卡片
func (s *cardRequest) PersonCardsDelete(employeeNo string, list []string) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/cards/batchDelete`), map[string]interface{}{
		`employeeNo`: employeeNo,
		`cardNos`:    list,
	})
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}
