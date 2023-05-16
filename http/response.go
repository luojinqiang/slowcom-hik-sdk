package http

import (
	"encoding/json"
	"github.com/ddliu/go-httpclient"
	"github.com/luojinqiang/slowcom-hik-sdk/gerror"
)

// HikResponse 返回
type HikResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// PageEntity 分页条件
type PageEntity struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
}

// checkResponse 校验请求
func checkResponse(res *httpclient.Response) (hikResponse *HikResponse, err error) {
	if res.StatusCode == 401 { //token 过期，重新请求
		return nil, gerror.ErrIs授权过期
	}
	if res.StatusCode != 200 {
		return nil, gerror.New(res.StatusCode, "请求状态异常")
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, gerror.ErrIs数据解析异常
	}
	err = json.Unmarshal(bytes, &hikResponse)
	if err != nil {
		return nil, gerror.ErrIs数据解析异常
	}
	if hikResponse.Code == 200 || hikResponse.Code == 0 {
		return
	} else {
		return hikResponse, gerror.New(hikResponse.Code, hikResponse.Message)
	}
}
