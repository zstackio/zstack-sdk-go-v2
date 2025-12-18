// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVirtualRouterVRouterRouteTableRef 查询VirtualRouterVRouterRouteTableRef列表
func (cli *ZSClient) QueryVirtualRouterVRouterRouteTableRef(params param.QueryParam) ([]view.QueryVirtualRouterVRouterRouteTableRefView, error) {
	var resp []view.QueryVirtualRouterVRouterRouteTableRefView
	return resp, cli.List("v1/vrouter-route-tables/virtual-router-refs", &params, &resp)
}

