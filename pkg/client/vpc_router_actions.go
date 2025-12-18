// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcRouter 查询VpcRouter列表
func (cli *ZSClient) QueryVpcRouter(params param.QueryParam) ([]view.QueryVpcRouterView, error) {
	var resp []view.QueryVpcRouterView
	return resp, cli.List("v1/vpc/virtual-routers", &params, &resp)
}

