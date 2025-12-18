// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryFiberChannelLun 查询FiberChannelLun列表
func (cli *ZSClient) QueryFiberChannelLun(params param.QueryParam) ([]view.QueryFiberChannelLunView, error) {
	var resp []view.QueryFiberChannelLunView
	return resp, cli.List("v1/storage-devices/fiber-channel/luns", &params, &resp)
}

