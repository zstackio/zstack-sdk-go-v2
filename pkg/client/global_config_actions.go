// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateGlobalConfig updates GlobalConfig
func (cli *ZSClient) UpdateGlobalConfig(uuid string, params param.UpdateGlobalConfigParam) (*view.GlobalConfigInventoryView, error) {
	var resp view.UpdateGlobalConfigEventView
	if err := cli.Put("v1/global-configurations/{category}/{name}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ResetGlobalConfig operates on GlobalConfig
func (cli *ZSClient) ResetGlobalConfig(uuid string, params param.ResetGlobalConfigParam) (*view.GlobalConfigInventoryView, error) {
	resp := view.GlobalConfigInventoryView{}
	if err := cli.Put("v1/global-configurations/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryGlobalConfig queries GlobalConfig list
func (cli *ZSClient) QueryGlobalConfig(params *param.QueryParam) ([]view.GlobalConfigInventoryView, error) {
	var resp []view.GlobalConfigInventoryView
	return resp, cli.List("v1/global-configurations", params, &resp)
}

func (cli *ZSClient) GetGlobalConfig(uuid string) (*view.GlobalConfigInventoryView, error) {
	var resp view.GlobalConfigInventoryView
	if err := cli.Get("v1/global-configurations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
