// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshLoadBalancer operates on RefreshLoadBalancer
func (cli *ZSClient) RefreshLoadBalancer(uuid string, params param.RefreshLoadBalancerParam) (*view.RefreshLoadBalancerEventView, error) {
	resp := view.RefreshLoadBalancerEventView{}
	if err := cli.Put("v1/load-balancers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
