// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateUserProxyConfig updates UserProxyConfig
func (cli *ZSClient) UpdateUserProxyConfig(ctx context.Context, uuid string, params param.UpdateUserProxyConfigParam) (*view.UserProxyConfigInventoryView, error) {
	resp := view.UserProxyConfigInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/user-proxy-configs", uuid, "", map[string]interface{}{
		"updateUserProxyConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteUserProxyConfig deletes UserProxyConfig
func (cli *ZSClient) DeleteUserProxyConfig(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/user-proxy-configs", uuid, string(deleteMode))
}
// QueryUserProxyConfig queries UserProxyConfig list
func (cli *ZSClient) QueryUserProxyConfig(ctx context.Context, params *param.QueryParam) ([]view.UserProxyConfigInventoryView, error) {
	var resp []view.UserProxyConfigInventoryView
	return resp, cli.List(ctx, "v1/user-proxy-configs", params, &resp)
}

func (cli *ZSClient) GetUserProxyConfig(ctx context.Context, uuid string) (*view.UserProxyConfigInventoryView, error) {
	var resp view.UserProxyConfigInventoryView
	if err := cli.Get(ctx, "v1/user-proxy-configs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUserProxyConfig Pagination
func (cli *ZSClient) PageUserProxyConfig(ctx context.Context, params *param.QueryParam) ([]view.UserProxyConfigInventoryView, int, error) {
	var userProxyConfigs []view.UserProxyConfigInventoryView
	total, err := cli.Page(ctx, "v1/user-proxy-configs", params, &userProxyConfigs)
	return userProxyConfigs, total, err
}
// CreateUserProxyConfig creates UserProxyConfig
func (cli *ZSClient) CreateUserProxyConfig(ctx context.Context, params param.CreateUserProxyConfigParam) (*view.UserProxyConfigInventoryView, error) {
	resp := view.UserProxyConfigInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/user-proxy-configs", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
