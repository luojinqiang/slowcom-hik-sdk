package basic

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-hik-sdk/config"
	"slowcom-hik-sdk/g"
	"slowcom-hik-sdk/gerror"
)

type BaseRequest struct {
}

// BuildUrl 构建URL
func (s *BaseRequest) BuildUrl(url string) string {
	return fmt.Sprintf("%s%s", config.BaseUrl, url)
}

// CheckResponse 校验请求
func (s *BaseRequest) CheckResponse(res *httpclient.Response) (hikResponse *g.HikResponse, err error) {
	if res.StatusCode != 200 {
		return nil, gerror.ErrIs请求状态异常
	}
	bytes, err := res.ReadAll()
	if err != nil {
		return nil, gerror.ErrIs数据解析异常
	}
	err = json.Unmarshal(bytes, &hikResponse)
	if err != nil {
		return nil, gerror.ErrIs数据解析异常
	}
	if hikResponse.Code != 200 {
		return
	} else {
		return nil, gerror.New(hikResponse.Message)
	}
}
