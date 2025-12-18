// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryFiberChannelStorage 查询FiberChannelStorage列表
func (cli *ZSClient) QueryFiberChannelStorage(params param.QueryParam) ([]view.QueryFiberChannelStorageView, error) {
	var resp []view.QueryFiberChannelStorageView
	return resp, cli.List("v1/storage-devices/fiber-channel/controllers", &params, &resp)
}

