// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPolicyRouteTableVRouterRef queries PolicyRouteTableVRouterRef list
func (cli *ZSClient) QueryPolicyRouteTableVRouterRef(params *param.QueryParam) ([]view.PolicyRouteTableVRouterRefInventoryView, error) {
	var resp []view.PolicyRouteTableVRouterRefInventoryView
	return resp, cli.List("v1/policy-routes/tables/vrouters/refs", params, &resp)
}
