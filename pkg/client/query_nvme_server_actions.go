// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryNvmeServer queries NvmeServer list
func (cli *ZSClient) QueryNvmeServer(params *param.QueryParam) ([]view.NvmeServerInventoryView, error) {
	var resp []view.NvmeServerInventoryView
	return resp, cli.List("v1/storage-devices/nvme/servers", params, &resp)
}
