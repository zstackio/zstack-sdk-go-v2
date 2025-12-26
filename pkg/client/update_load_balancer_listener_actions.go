// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateLoadBalancerListener updates LoadBalancerListener
func (cli *ZSClient) UpdateLoadBalancerListener(uuid string, params param.UpdateLoadBalancerListenerParam) (*view.UpdateLoadBalancerListenerEventView, error) {
	resp := view.UpdateLoadBalancerListenerEventView{}
	if err := cli.Put("v1/load-balancers/listeners/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
