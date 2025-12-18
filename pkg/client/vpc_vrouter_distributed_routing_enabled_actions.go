// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetVpcVRouterDistributedRoutingEnabled 获取VpcVRouterDistributedRoutingEnabled详情
func (cli *ZSClient) GetVpcVRouterDistributedRoutingEnabled(uuid string) (*view.GetVpcVRouterDistributedRoutingEnabledView, error) {
	var resp view.GetVpcVRouterDistributedRoutingEnabledView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/distributed-routing", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

