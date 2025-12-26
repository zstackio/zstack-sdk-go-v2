// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPolicyRouteTable queries PolicyRouteTable list
func (cli *ZSClient) QueryPolicyRouteTable(params *param.QueryParam) ([]view.PolicyRouteTableInventoryView, error) {
	var resp []view.PolicyRouteTableInventoryView
	return resp, cli.List("v1/policy-routes/tables", params, &resp)
}
