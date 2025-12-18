// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddVRouterRouteEntry 操作AddVRouterRouteEntry
func (cli *ZSClient) AddVRouterRouteEntry(params param.AddVRouterRouteEntryParam) (*view.AddVRouterRouteEntryEventView, error) {
	resp := view.AddVRouterRouteEntryEventView{}
	if err := cli.Post("v1/vrouter-route-tables/{routeTableUuid}/route-entries", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

