// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNvmeServer queries NvmeServer list
func (cli *ZSClient) QueryNvmeServer(params param.QueryParam) ([]view.NvmeServerInventoryView, error) {
	var resp []view.NvmeServerInventoryView
	return resp, cli.List("v1/storage-devices/nvme/servers", &params, &resp)
}
