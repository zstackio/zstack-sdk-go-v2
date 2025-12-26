// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryFiberChannelStorage queries FiberChannelStorage list
func (cli *ZSClient) QueryFiberChannelStorage(params *param.QueryParam) ([]view.FiberChannelStorageInventoryView, error) {
	var resp []view.FiberChannelStorageInventoryView
	return resp, cli.List("v1/storage-devices/fiber-channel/controllers", params, &resp)
}
