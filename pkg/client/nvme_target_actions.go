// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNvmeTarget 查询NvmeTarget列表
func (cli *ZSClient) QueryNvmeTarget(params param.QueryParam) ([]view.QueryNvmeTargetView, error) {
	var resp []view.QueryNvmeTargetView
	return resp, cli.List("v1/storage-devices/nvme/controllers", &params, &resp)
}

