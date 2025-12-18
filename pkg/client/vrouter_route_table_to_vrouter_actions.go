// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachVRouterRouteTableToVRouter 操作VRouterRouteTableToVRouter
func (cli *ZSClient) AttachVRouterRouteTableToVRouter(params param.AttachVRouterRouteTableToVRouterParam) (*view.AttachVRouterRouteTableToVRouterEventView, error) {
	resp := view.AttachVRouterRouteTableToVRouterEventView{}
	if err := cli.Post("v1/vrouter-route-tables/{routeTableUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

