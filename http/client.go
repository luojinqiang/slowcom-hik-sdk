package http

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	url2 "net/url"
	"slowcom-hik-sdk/gerror"
	"sync"
)

// HikHttpClient 海康httpClient
type HikHttpClient struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
	rwLock       sync.RWMutex
	accessToken  string
}

const (
	USERAGENT       = "slowcom_agent"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 10
)

// hikAccessTokenRes 海康
type hikAccessTokenRes struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
	Scope       string `json:"scope"`
}

// NewHikHttpClient 创建海康httpClient
func NewHikHttpClient(clientId, clientSecret string) *HikHttpClient {
	return &HikHttpClient{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
}

// 生成一个http请求客户端
func buildHttpClient() *httpclient.HttpClient {
	return httpclient.NewHttpClient().Defaults(httpclient.Map{
		"opt_useragent":      USERAGENT,
		"opt_timeout":        TIMEOUT,
		"Accept-Encoding":    "gzip,deflate,sdch",
		"opt_connecttimeout": CONNECT_TIMEOUT,
		"OPT_DEBUG":          true,
	})
}

// getAccessToken 获取access_token
func (s *HikHttpClient) getAccessToken() string {
	defer func() {
		s.rwLock.Unlock()
	}()
	s.rwLock.Lock()
	if s.accessToken != `` {
		return s.accessToken
	}
	response, err := buildHttpClient().WithHeader("Content-Type", "application/x-www-form-urlencoded").Post(s.BaseUrl+`/oauth/token`,
		fmt.Sprintf(`client_id=%s&client_secret=%s&grant_type=%s&scope=%s`, s.ClientId, s.ClientSecret, `client_credentials`, `app`))
	if err != nil {
		panic(fmt.Sprintf(`hik请求access_token失败：%s`, err.Error()))
	}
	if response.StatusCode != 200 {
		panic(fmt.Sprintf(`hik请求access_token失败：StatusCode = %d`, response.StatusCode))
	}
	bytes, err := response.ReadAll()
	if err != nil {
		panic(fmt.Sprintf(`hik请求access_token失败：%s`, err.Error()))
	}
	var tokenRes hikAccessTokenRes
	err = json.Unmarshal(bytes, &tokenRes)
	s.accessToken = tokenRes.AccessToken
	return s.accessToken
}

// PostJson json请求
func (s *HikHttpClient) PostJson(url string, data interface{}) (response *HikResponse, err error) {
	token := s.getAccessToken()
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+token).PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	if err != nil {
		if err == gerror.ErrIs授权过期 {
			s.accessToken = ``
			token = s.getAccessToken()
			res, err = buildHttpClient().WithHeader("Authorization", "Bearer "+token).PostJson(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	response, err = checkResponse(res)
	return
}

// Post post普通请求
func (s *HikHttpClient) Post(url, data interface{}) (response *HikResponse, err error) {
	token := s.getAccessToken()
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+token).Post(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
	if err != nil {
		if err == gerror.ErrIs授权过期 {
			s.accessToken = ``
			token = s.getAccessToken()
			res, err = buildHttpClient().WithHeader("Authorization", "Bearer "+token).Post(fmt.Sprintf("%s%s", s.BaseUrl, url), data)
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	response, err = checkResponse(res)
	return
}

// Get get请求
func (s *HikHttpClient) Get(url string) (response *HikResponse, err error) {
	token := s.getAccessToken()
	res, err := buildHttpClient().WithHeader("Authorization", "Bearer "+token).Get(fmt.Sprintf("%s%s", s.BaseUrl, url), url2.Values{})
	if err != nil {
		if err == gerror.ErrIs授权过期 {
			s.accessToken = ``
			token = s.getAccessToken()
			res, err = buildHttpClient().WithHeader("Authorization", "Bearer "+token).Get(fmt.Sprintf("%s%s", s.BaseUrl, url), url2.Values{})
			if err != nil {
				return
			}
		} else {
			return
		}
	}
	response, err = checkResponse(res)
	return
}
