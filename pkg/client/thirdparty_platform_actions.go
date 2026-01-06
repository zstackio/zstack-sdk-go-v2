// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryThirdpartyPlatform queries ThirdpartyPlatform list
func (cli *ZSClient) QueryThirdpartyPlatform(params *param.QueryParam) ([]view.ThirdpartyPlatformInventoryView, error) {
	var resp []view.ThirdpartyPlatformInventoryView
	return resp, cli.List("v1/zwatch/third-party/platforms", params, &resp)
}
// UpdateThirdpartyPlatform updates ThirdpartyPlatform
func (cli *ZSClient) UpdateThirdpartyPlatform(uuid string, params param.UpdateThirdpartyPlatformParam) (*view.ThirdpartyPlatformInventoryView, error) {
	var resp view.UpdateThirdpartyPlatformEventView
	if err := cli.Put("v1/zwatch/third-party/platforms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddThirdpartyPlatform adds ThirdpartyPlatform
func (cli *ZSClient) AddThirdpartyPlatform(params param.AddThirdpartyPlatformParam) (*view.ThirdpartyPlatformInventoryView, error) {
	var resp view.AddThirdpartyPlatformEventView
	if err := cli.Post("v1/zwatch/third-party/platforms", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteThirdpartyPlatform deletes ThirdpartyPlatform
func (cli *ZSClient) DeleteThirdpartyPlatform(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/third-party/platforms/{uuid}", uuid, string(deleteMode))
}
