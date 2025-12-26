// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVirtualRouterVRouterRouteTableRef queries VirtualRouterVRouterRouteTableRef list
func (cli *ZSClient) QueryVirtualRouterVRouterRouteTableRef(params *param.QueryParam) ([]view.VirtualRouterVRouterRouteTableRefInventoryView, error) {
	var resp []view.VirtualRouterVRouterRouteTableRefInventoryView
	return resp, cli.List("v1/vrouter-route-tables/virtual-router-refs", params, &resp)
}
