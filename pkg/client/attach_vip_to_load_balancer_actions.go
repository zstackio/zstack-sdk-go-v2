// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachVipToLoadBalancer operates on VipToLoadBalancer
func (cli *ZSClient) AttachVipToLoadBalancer(params param.AttachVipToLoadBalancerParam) (*view.AttachVipToLoadBalancerEventView, error) {
	resp := view.AttachVipToLoadBalancerEventView{}
	if err := cli.Post("v1/load-balancers/{loadBalancerUuid}/vip/{vipUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
