// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVRouterRouteTable updates VRouterRouteTable
func (cli *ZSClient) UpdateVRouterRouteTable(uuid string, params param.UpdateVRouterRouteTableParam) (*view.UpdateVRouterRouteTableEventView, error) {
	resp := view.UpdateVRouterRouteTableEventView{}
	if err := cli.Put("v1/vrouter-route-tables/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
