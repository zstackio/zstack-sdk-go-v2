// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVpcVRouterDistributedRoutingEnabled 操作SetVpcVRouterDistributedRoutingEnabled
func (cli *ZSClient) SetVpcVRouterDistributedRoutingEnabled(params param.SetVpcVRouterDistributedRoutingEnabledParam) (*view.SetVpcVRouterDistributedRoutingEnabledEventView, error) {
	resp := view.SetVpcVRouterDistributedRoutingEnabledEventView{}
	if err := cli.Post("v1/vpc/virtual-routers/{uuid}/distributed-routing", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

