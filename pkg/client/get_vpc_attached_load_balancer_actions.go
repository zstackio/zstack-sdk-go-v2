// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetVpcAttachedLoadBalancer gets VpcAttachedLoadBalancer by uuid
func (cli *ZSClient) GetVpcAttachedLoadBalancer(uuid string) (*view.GetVpcAttachedLoadBalancerView, error) {
	var resp view.GetVpcAttachedLoadBalancerView
	if err := cli.Get("v1/vpc/virtual-routers/{uuid}/attached-lb", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
