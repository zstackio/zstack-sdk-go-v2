// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVRouterRouteTable creates VRouterRouteTable
func (cli *ZSClient) CreateVRouterRouteTable(params param.CreateVRouterRouteTableParam) (*view.CreateVRouterRouteTableEventView, error) {
	resp := view.CreateVRouterRouteTableEventView{}
	if err := cli.Post("v1/vrouter-route-tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
