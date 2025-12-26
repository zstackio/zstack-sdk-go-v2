// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddServerGroupToLoadBalancerListener adds ServerGroupToLoadBalancerListener
func (cli *ZSClient) AddServerGroupToLoadBalancerListener(params param.AddServerGroupToLoadBalancerListenerParam) (*view.AddServerGroupToLoadBalancerListenerEventView, error) {
	resp := view.AddServerGroupToLoadBalancerListenerEventView{}
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/servergroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
