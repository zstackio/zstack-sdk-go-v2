// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVRouterRouteTable queries VRouterRouteTable list
func (cli *ZSClient) QueryVRouterRouteTable(params *param.QueryParam) ([]view.VRouterRouteTableInventoryView, error) {
	var resp []view.VRouterRouteTableInventoryView
	return resp, cli.List("v1/vrouter-route-tables", params, &resp)
}
