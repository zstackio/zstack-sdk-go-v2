// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPolicyRouteTableVRouterRef 查询PolicyRouteTableVRouterRef列表
func (cli *ZSClient) QueryPolicyRouteTableVRouterRef(params param.QueryParam) ([]view.QueryPolicyRouteTableVRouterRefView, error) {
	var resp []view.QueryPolicyRouteTableVRouterRefView
	return resp, cli.List("v1/policy-routes/tables/vrouters/refs", &params, &resp)
}

