// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVRouterRouteTable 查询VRouterRouteTable列表
func (cli *ZSClient) QueryVRouterRouteTable(params param.QueryParam) ([]view.QueryVRouterRouteTableView, error) {
	var resp []view.QueryVRouterRouteTableView
	return resp, cli.List("v1/vrouter-route-tables", &params, &resp)
}

