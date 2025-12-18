// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcMulticastRoute gets VpcMulticastRoute by uuid
func (cli *ZSClient) GetVpcMulticastRoute(uuid string) (*view.GetVpcMulticastRouteView, error) {
	var resp view.GetVpcMulticastRouteView
	if err := cli.Get("v1/multicast/virtual-routers/{uuid}/routes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
