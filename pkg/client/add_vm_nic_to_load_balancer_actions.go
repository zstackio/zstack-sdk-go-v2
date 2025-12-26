// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddVmNicToLoadBalancer adds VmNicToLoadBalancer
func (cli *ZSClient) AddVmNicToLoadBalancer(params param.AddVmNicToLoadBalancerParam) (*view.AddVmNicToLoadBalancerEventView, error) {
	resp := view.AddVmNicToLoadBalancerEventView{}
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
