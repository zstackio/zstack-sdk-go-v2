// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachVRouterRouteTableToVRouter operates on VRouterRouteTableToVRouter
func (cli *ZSClient) AttachVRouterRouteTableToVRouter(params param.AttachVRouterRouteTableToVRouterParam) (*view.AttachVRouterRouteTableToVRouterEventView, error) {
	resp := view.AttachVRouterRouteTableToVRouterEventView{}
	if err := cli.Post("v1/vrouter-route-tables/{routeTableUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
