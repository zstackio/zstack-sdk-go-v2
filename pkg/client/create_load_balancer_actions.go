// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateLoadBalancer creates LoadBalancer
func (cli *ZSClient) CreateLoadBalancer(params param.CreateLoadBalancerParam) (*view.CreateLoadBalancerEventView, error) {
	resp := view.CreateLoadBalancerEventView{}
	if err := cli.Post("v1/load-balancers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
