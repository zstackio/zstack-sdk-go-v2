// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateLoadBalancerServerGroup creates LoadBalancerServerGroup
func (cli *ZSClient) CreateLoadBalancerServerGroup(params param.CreateLoadBalancerServerGroupParam) (*view.CreateLoadBalancerServerGroupEventView, error) {
	resp := view.CreateLoadBalancerServerGroupEventView{}
	if err := cli.Post("v1/load-balancers/{loadBalancerUuid}/servergroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
