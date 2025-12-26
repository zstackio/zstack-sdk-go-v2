// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPolicyRouteTableRouteEntry queries PolicyRouteTableRouteEntry list
func (cli *ZSClient) QueryPolicyRouteTableRouteEntry(params *param.QueryParam) ([]view.PolicyRouteTableRouteEntryInventoryView, error) {
	var resp []view.PolicyRouteTableRouteEntryInventoryView
	return resp, cli.List("v1/policy-routes/routes", params, &resp)
}
