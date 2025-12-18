// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddVmNicToLoadBalancer 操作AddVmNicToLoadBalancer
func (cli *ZSClient) AddVmNicToLoadBalancer(params param.AddVmNicToLoadBalancerParam) (*view.AddVmNicToLoadBalancerEventView, error) {
	resp := view.AddVmNicToLoadBalancerEventView{}
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/vm-instances/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

