// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddVRouterRouteEntry adds VRouterRouteEntry
func (cli *ZSClient) AddVRouterRouteEntry(params param.AddVRouterRouteEntryParam) (*view.AddVRouterRouteEntryEventView, error) {
	resp := view.AddVRouterRouteEntryEventView{}
	if err := cli.Post("v1/vrouter-route-tables/{routeTableUuid}/route-entries", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
