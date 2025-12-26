// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeLoadBalancerListener changes LoadBalancerListener
func (cli *ZSClient) ChangeLoadBalancerListener(uuid string, params param.ChangeLoadBalancerListenerParam) (*view.ChangeLoadBalancerListenerEventView, error) {
	resp := view.ChangeLoadBalancerListenerEventView{}
	if err := cli.Put("v1/load-balancers/listeners/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
