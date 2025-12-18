// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeLoadBalancerListener changes LoadBalancerListener
func (cli *ZSClient) ChangeLoadBalancerListener(uuid string, params param.ChangeLoadBalancerListenerParam) (*view.ChangeLoadBalancerListenerEventView, error) {
	resp := view.ChangeLoadBalancerListenerEventView{}
	if err := cli.Put("v1/load-balancers/listeners/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
