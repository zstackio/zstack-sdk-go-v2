// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateLoadBalancerServerGroup creates LoadBalancerServerGroup
func (cli *ZSClient) CreateLoadBalancerServerGroup(params param.CreateLoadBalancerServerGroupParam) (*view.CreateLoadBalancerServerGroupEventView, error) {
	resp := view.CreateLoadBalancerServerGroupEventView{}
	if err := cli.Post("v1/load-balancers/{loadBalancerUuid}/servergroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
