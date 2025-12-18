// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVirtualBorderRouterFromLocal 查询VirtualBorderRouterFromLocal列表
func (cli *ZSClient) QueryVirtualBorderRouterFromLocal(params param.QueryParam) ([]view.QueryVirtualBorderRouterFromLocalView, error) {
	var resp []view.QueryVirtualBorderRouterFromLocalView
	return resp, cli.List("v1/hybrid/aliyun/border-router", &params, &resp)
}

