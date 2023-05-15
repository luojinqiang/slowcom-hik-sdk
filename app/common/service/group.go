package service

import (
	"encoding/json"
	"fmt"
	"github.com/luojinqiang/slowcom-hik-sdk/app/common/entity"
	"github.com/luojinqiang/slowcom-hik-sdk/http"
)

/**
API文档	https://pic.hik-cloud.com/opencustom/apidoc/online/open/4f65c0b1db3a4ea48e9c7d90e8fc18d1.html?timestamp=1681441984213
*/
type GroupRequest struct {
	HikClient *http.HikHttpClient
}

// Add 该接口用于通过编号来新增组。最多支持3000个组，最多支持5层嵌套。
func (s *GroupRequest) Add(groupAdd *entity.GroupAdd) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/basic/groups/create`, groupAdd)
	return
}

// Delete 该接口用于通过组编号来删除组。组下面挂有下级节点或者设备的不可以删除，需清空后进行删除。
func (s *GroupRequest) Delete(groupNo string) (err error) {
	_, err = s.HikClient.Post(fmt.Sprintf(`/api/v1/open/basic/groups/delete?groupNo=%s`, groupNo), nil)
	return
}

// Update 该接口用于通过组编号修改组的名称
func (s *GroupRequest) Update(groupUpdate *entity.GroupUpdate) (err error) {
	_, err = s.HikClient.PostJson(`/api/v1/open/basic/groups/update`, groupUpdate)
	return
}

// Get 获取组详细信息
func (s *GroupRequest) Get(groupNo string) (group *entity.Group, err error) {
	hikResponse, err := s.HikClient.Get(fmt.Sprintf(`/api/v1/open/basic/groups/delete?groupNo=%s`, groupNo))
	if err != nil {
		return nil, err
	}
	bytes, _ := json.Marshal(hikResponse.Data)
	err = json.Unmarshal(bytes, &group)
	return
}

// List 获取组列表
// groupNo 不传该参数则获取全部，传该参数则获取该组下的下级节点列表
func (s *GroupRequest) List(groupNo string) (list []*entity.Group, err error) {
	var res *http.HikResponse
	if groupNo == `` {
		res, err = s.HikClient.Get(fmt.Sprintf(`/api/v1/open/basic/groups/actions/listAll`))
		if err != nil {
			return nil, err
		}
	} else {
		res, err = s.HikClient.Get(fmt.Sprintf(`/api/v1/open/basic/groups/actions/childrenList?parentNo=%s`, groupNo))
		if err != nil {
			return nil, err
		}
	}
	bytes, _ := json.Marshal(res.Data)
	err = json.Unmarshal(bytes, &list)
	return
}
