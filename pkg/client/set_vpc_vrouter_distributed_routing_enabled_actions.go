// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVpcVRouterDistributedRoutingEnabled operates on SetVpcVRouterDistributedRoutingEnabled
func (cli *ZSClient) SetVpcVRouterDistributedRoutingEnabled(params param.SetVpcVRouterDistributedRoutingEnabledParam) (*view.SetVpcVRouterDistributedRoutingEnabledEventView, error) {
	resp := view.SetVpcVRouterDistributedRoutingEnabledEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/distributed-routing", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
