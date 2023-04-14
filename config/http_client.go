package config

import "github.com/ddliu/go-httpclient"

const (
	USERAGENT       = "slowcom_agent"
	TIMEOUT         = 30
	CONNECT_TIMEOUT = 5
)

func init() {
	httpclient.Defaults(httpclient.Map{
		"opt_useragent":   USERAGENT,
		"opt_timeout":     TIMEOUT,
		"Accept-Encoding": "gzip,deflate,sdch",
	})
}
