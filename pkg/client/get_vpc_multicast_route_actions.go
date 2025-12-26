// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcMulticastRoute gets VpcMulticastRoute by uuid
func (cli *ZSClient) GetVpcMulticastRoute(uuid string) (*view.GetVpcMulticastRouteView, error) {
	var resp view.GetVpcMulticastRouteView
	if err := cli.Get("v1/multicast/virtual-routers/{uuid}/routes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
