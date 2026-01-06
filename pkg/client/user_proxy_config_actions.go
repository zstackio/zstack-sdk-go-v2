// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateUserProxyConfig updates UserProxyConfig
func (cli *ZSClient) UpdateUserProxyConfig(uuid string, params param.UpdateUserProxyConfigParam) (*view.UserProxyConfigInventoryView, error) {
	var resp view.UpdateUserProxyConfigEventView
	if err := cli.Put("v1/user-proxy-configs/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteUserProxyConfig deletes UserProxyConfig
func (cli *ZSClient) DeleteUserProxyConfig(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/user-proxy-configs/{uuid}", uuid, string(deleteMode))
}
// QueryUserProxyConfig queries UserProxyConfig list
func (cli *ZSClient) QueryUserProxyConfig(params *param.QueryParam) ([]view.UserProxyConfigInventoryView, error) {
	var resp []view.UserProxyConfigInventoryView
	return resp, cli.List("v1/user-proxy-configs", params, &resp)
}
// CreateUserProxyConfig creates UserProxyConfig
func (cli *ZSClient) CreateUserProxyConfig(params param.CreateUserProxyConfigParam) (*view.UserProxyConfigInventoryView, error) {
	var resp view.CreateUserProxyConfigEventView
	if err := cli.Post("v1/user-proxy-configs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
