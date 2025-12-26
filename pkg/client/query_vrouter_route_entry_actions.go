// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVRouterRouteEntry queries VRouterRouteEntry list
func (cli *ZSClient) QueryVRouterRouteEntry(params *param.QueryParam) ([]view.VRouterRouteEntryInventoryView, error) {
	var resp []view.VRouterRouteEntryInventoryView
	return resp, cli.List("v1/vrouter-route-tables/route-entries", params, &resp)
}
