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

// PageFiberChannelLun Pagination
func (cli *ZSClient) PageFiberChannelLun(params *param.QueryParam) ([]view.FiberChannelLunInventoryView, int, error) {
	var fiberChannelLuns []view.FiberChannelLunInventoryView
	total, err := cli.Page("v1/storage-devices/fiber-channel/luns", params, &fiberChannelLuns)
	return fiberChannelLuns, total, err
}
