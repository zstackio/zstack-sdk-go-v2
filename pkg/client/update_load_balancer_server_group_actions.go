// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateLoadBalancerServerGroup updates LoadBalancerServerGroup
func (cli *ZSClient) UpdateLoadBalancerServerGroup(uuid string, params param.UpdateLoadBalancerServerGroupParam) (*view.UpdateLoadBalancerServerGroupEventView, error) {
	resp := view.UpdateLoadBalancerServerGroupEventView{}
	if err := cli.Put("v1/load-balancers/servergroups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
