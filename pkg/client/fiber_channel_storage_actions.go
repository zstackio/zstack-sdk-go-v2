// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RefreshFiberChannelStorage operates on FiberChannelStorage
func (cli *ZSClient) RefreshFiberChannelStorage(params param.RefreshFiberChannelStorageParam) (*view.FiberChannelStorageInventoryView, error) {
	resp := view.FiberChannelStorageInventoryView{}
	if err := cli.Post("v1/storage-devices/fiber-channel/controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryFiberChannelStorage queries FiberChannelStorage list
func (cli *ZSClient) QueryFiberChannelStorage(params *param.QueryParam) ([]view.FiberChannelStorageInventoryView, error) {
	var resp []view.FiberChannelStorageInventoryView
	return resp, cli.List("v1/storage-devices/fiber-channel/controllers", params, &resp)
}
