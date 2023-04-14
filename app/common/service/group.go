package service

import (
	"encoding/json"
	"fmt"
	"github.com/ddliu/go-httpclient"
	"slowcom-hik-sdk/app/common/entity"
	"slowcom-hik-sdk/basic"
	"slowcom-hik-sdk/gerror"
)

/**
API文档	https://pic.hik-cloud.com/opencustom/apidoc/online/open/4f65c0b1db3a4ea48e9c7d90e8fc18d1.html?timestamp=1681441984213
*/
type groupRequest struct {
	basic.BaseRequest
}

var GroupRequest = new(groupRequest)

// Add 该接口用于通过编号来新增组。最多支持3000个组，最多支持5层嵌套。
func (s *groupRequest) Add(groupAdd *entity.GroupAdd) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/groups/create`), groupAdd)
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// Delete 该接口用于通过组编号来删除组。组下面挂有下级节点或者设备的不可以删除，需清空后进行删除。
func (s *groupRequest) Delete(groupNo string) (err error) {
	res, err := httpclient.Post(s.BuildUrl(fmt.Sprintf(`/api/v1/open/basic/groups/delete?groupNo=%s`, groupNo)), nil)
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// Update 该接口用于通过组编号修改组的名称
func (s *groupRequest) Update(groupUpdate *entity.GroupUpdate) (err error) {
	res, err := httpclient.PostJson(s.BuildUrl(`/api/v1/open/basic/groups/update`), groupUpdate)
	if err != nil {
		return gerror.ErrIs系统异常
	}
	_, err = s.CheckResponse(res)
	return
}

// Get 获取组详细信息
func (s *groupRequest) Get(groupNo string) (group *entity.Group, err error) {
	res, err := httpclient.Get(s.BuildUrl(fmt.Sprintf(`/api/v1/open/basic/groups/delete?groupNo=%s`, groupNo)))
	if err != nil {
		return nil, gerror.ErrIs系统异常
	}
	hikResponse, err := s.CheckResponse(res)
	bytes, _ := json.Marshal(hikResponse.Data)
	err = json.Unmarshal(bytes, &group)
	return
}

// List 获取节点列表
// groupNo 不传该参数则获取全部，传该参数则获取该组下的下级节点列表
func (s *groupRequest) List(groupNo string) (list []*entity.Group, err error) {
	var res *httpclient.Response
	if groupNo == `` {
		res, err = httpclient.Get(s.BuildUrl(fmt.Sprintf(`/api/v1/open/basic/groups/actions/listAll`)))
		if err != nil {
			return nil, gerror.ErrIs系统异常
		}
	} else {
		res, err = httpclient.Get(s.BuildUrl(fmt.Sprintf(`/api/v1/open/basic/groups/actions/childrenList?parentNo=%s`, groupNo)))
		if err != nil {
			return nil, gerror.ErrIs系统异常
		}
	}
	hikResponse, err := s.CheckResponse(res)
	bytes, _ := json.Marshal(hikResponse.Data)
	err = json.Unmarshal(bytes, &list)
	return
}
