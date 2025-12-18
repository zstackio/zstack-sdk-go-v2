// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcVRouterDistributedRoutingConnections gets VpcVRouterDistributedRoutingConnections by uuid
func (cli *ZSClient) GetVpcVRouterDistributedRoutingConnections(uuid string) (*view.GetVpcVRouterDistributedRoutingConnectionsView, error) {
	var resp view.GetVpcVRouterDistributedRoutingConnectionsView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/tracked-connections", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
