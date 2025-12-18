// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryMulticastRouter 查询MulticastRouter列表
func (cli *ZSClient) QueryMulticastRouter(params param.QueryParam) ([]view.QueryMulticastRouterView, error) {
	var resp []view.QueryMulticastRouterView
	return resp, cli.List("v1/multicast/virtual-routers", &params, &resp)
}

