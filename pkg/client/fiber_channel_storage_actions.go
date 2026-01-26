// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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

func (cli *ZSClient) GetFiberChannelStorage(uuid string) (*view.FiberChannelStorageInventoryView, error) {
	var resp view.FiberChannelStorageInventoryView
	if err := cli.Get("v1/storage-devices/fiber-channel/controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFiberChannelStorage Pagination
func (cli *ZSClient) PageFiberChannelStorage(params *param.QueryParam) ([]view.FiberChannelStorageInventoryView, int, error) {
	var fiberChannelStorages []view.FiberChannelStorageInventoryView
	total, err := cli.Page("v1/storage-devices/fiber-channel/controllers", params, &fiberChannelStorages)
	return fiberChannelStorages, total, err
}
