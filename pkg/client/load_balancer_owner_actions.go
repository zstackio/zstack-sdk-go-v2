// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLoadBalancerOwner 获取LoadBalancerOwner详情
func (cli *ZSClient) GetLoadBalancerOwner(uuid string) (*view.GetLoadBalancerOwnerView, error) {
	var resp view.GetLoadBalancerOwnerView
	if err := cli.Get("v1/load-balancers/{loadBalancerUuid}/owner", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

