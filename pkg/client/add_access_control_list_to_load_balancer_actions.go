// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddAccessControlListToLoadBalancer adds AccessControlListToLoadBalancer
func (cli *ZSClient) AddAccessControlListToLoadBalancer(params param.AddAccessControlListToLoadBalancerParam) (*view.AddAccessControlListToLoadBalancerEventView, error) {
	resp := view.AddAccessControlListToLoadBalancerEventView{}
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/access-control-lists", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
