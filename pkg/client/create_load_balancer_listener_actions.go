// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateLoadBalancerListener creates LoadBalancerListener
func (cli *ZSClient) CreateLoadBalancerListener(params param.CreateLoadBalancerListenerParam) (*view.CreateLoadBalancerListenerEventView, error) {
	resp := view.CreateLoadBalancerListenerEventView{}
	if err := cli.Post("v1/load-balancers/{loadBalancerUuid}/listeners", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
