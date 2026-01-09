// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryFiberChannelLun queries FiberChannelLun list
func (cli *ZSClient) QueryFiberChannelLun(params *param.QueryParam) ([]view.FiberChannelLunInventoryView, error) {
	var resp []view.FiberChannelLunInventoryView
	return resp, cli.List("v1/storage-devices/fiber-channel/luns", params, &resp)
}

func (cli *ZSClient) GetFiberChannelLun(uuid string) (*view.FiberChannelLunInventoryView, error) {
	var resp view.FiberChannelLunInventoryView
	if err := cli.Get("v1/storage-devices/fiber-channel/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
