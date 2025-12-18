// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryFiberChannelStorage queries FiberChannelStorage list
func (cli *ZSClient) QueryFiberChannelStorage(params param.QueryParam) ([]view.FiberChannelStorageInventoryView, error) {
	var resp []view.FiberChannelStorageInventoryView
	return resp, cli.List("v1/storage-devices/fiber-channel/controllers", &params, &resp)
}
