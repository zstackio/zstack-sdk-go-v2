// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVRouterRouteTable creates VRouterRouteTable
func (cli *ZSClient) CreateVRouterRouteTable(params param.CreateVRouterRouteTableParam) (*view.CreateVRouterRouteTableEventView, error) {
	resp := view.CreateVRouterRouteTableEventView{}
	if err := cli.Post("v1/vrouter-route-tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
