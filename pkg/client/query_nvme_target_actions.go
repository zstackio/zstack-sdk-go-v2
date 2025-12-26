// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryNvmeTarget queries NvmeTarget list
func (cli *ZSClient) QueryNvmeTarget(params *param.QueryParam) ([]view.NvmeTargetInventoryView, error) {
	var resp []view.NvmeTargetInventoryView
	return resp, cli.List("v1/storage-devices/nvme/controllers", params, &resp)
}
