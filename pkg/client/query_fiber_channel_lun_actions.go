// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFiberChannelLun queries FiberChannelLun list
func (cli *ZSClient) QueryFiberChannelLun(params *param.QueryParam) ([]view.FiberChannelLunInventoryView, error) {
	var resp []view.FiberChannelLunInventoryView
	return resp, cli.List("v1/storage-devices/fiber-channel/luns", params, &resp)
}
