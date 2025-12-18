// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPolicyRouteTable 查询PolicyRouteTable列表
func (cli *ZSClient) QueryPolicyRouteTable(params param.QueryParam) ([]view.QueryPolicyRouteTableView, error) {
	var resp []view.QueryPolicyRouteTableView
	return resp, cli.List("v1/policy-routes/tables", &params, &resp)
}

