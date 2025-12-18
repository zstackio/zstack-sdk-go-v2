// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVRouterRouteTable gets VRouterRouteTable by uuid
func (cli *ZSClient) GetVRouterRouteTable(uuid string) (*view.GetVRouterRouteTableView, error) {
	var resp view.GetVRouterRouteTableView
	if err := cli.Get("v1/vrouter-route-tables/vrouter/{virtualRouterVmUuid}", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
