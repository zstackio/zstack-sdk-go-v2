// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLoadBalancerListener queries LoadBalancerListener list
func (cli *ZSClient) QueryLoadBalancerListener(params param.QueryParam) ([]view.LoadBalancerListenerInventoryView, error) {
	var resp []view.LoadBalancerListenerInventoryView
	return resp, cli.List("v1/load-balancers/listeners", &params, &resp)
}
