// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNvmeTarget queries NvmeTarget list
func (cli *ZSClient) QueryNvmeTarget(params param.QueryParam) ([]view.NvmeTargetInventoryView, error) {
	var resp []view.NvmeTargetInventoryView
	return resp, cli.List("v1/storage-devices/nvme/controllers", &params, &resp)
}
