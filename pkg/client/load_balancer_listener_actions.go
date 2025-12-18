// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateLoadBalancerListener 更新LoadBalancerListener
func (cli *ZSClient) UpdateLoadBalancerListener(uuid string, params param.UpdateLoadBalancerListenerParam) (*view.UpdateLoadBalancerListenerEventView, error) {
	resp := view.UpdateLoadBalancerListenerEventView{}
	if err := cli.Put("v1/load-balancers/listeners/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

