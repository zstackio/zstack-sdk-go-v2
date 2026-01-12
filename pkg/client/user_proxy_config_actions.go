// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateUserProxyConfig updates UserProxyConfig
func (cli *ZSClient) UpdateUserProxyConfig(uuid string, params param.UpdateUserProxyConfigParam) (*view.UserProxyConfigInventoryView, error) {
	var resp view.UpdateUserProxyConfigEventView
	err := cli.PutWithSpec("v1/user-proxy-configs", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteUserProxyConfig deletes UserProxyConfig
func (cli *ZSClient) DeleteUserProxyConfig(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/user-proxy-configs", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryUserProxyConfig queries UserProxyConfig list
func (cli *ZSClient) QueryUserProxyConfig(params *param.QueryParam) ([]view.UserProxyConfigInventoryView, error) {
	var resp []view.UserProxyConfigInventoryView
	return resp, cli.List("v1/user-proxy-configs", params, &resp)
}

func (cli *ZSClient) GetUserProxyConfig(uuid string) (*view.UserProxyConfigInventoryView, error) {
	var resp view.UserProxyConfigInventoryView
	if err := cli.Get("v1/user-proxy-configs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateUserProxyConfig creates UserProxyConfig
func (cli *ZSClient) CreateUserProxyConfig(params param.CreateUserProxyConfigParam) (*view.UserProxyConfigInventoryView, error) {
	var resp view.CreateUserProxyConfigEventView
	if err := cli.Post("v1/user-proxy-configs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
