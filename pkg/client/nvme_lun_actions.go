// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNvmeLun 查询NvmeLun列表
func (cli *ZSClient) QueryNvmeLun(params param.QueryParam) ([]view.QueryNvmeLunView, error) {
	var resp []view.QueryNvmeLunView
	return resp, cli.List("v1/storage-devices/nvme/luns", &params, &resp)
}

