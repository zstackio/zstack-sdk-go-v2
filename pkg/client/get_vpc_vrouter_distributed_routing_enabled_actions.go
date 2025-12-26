// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcVRouterDistributedRoutingEnabled gets VpcVRouterDistributedRoutingEnabled by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingEnabled(uuid string) (*view.GetVpcVRouterDistributedRoutingEnabledView, error) {
	var resp view.GetVpcVRouterDistributedRoutingEnabledView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/distributed-routing", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
