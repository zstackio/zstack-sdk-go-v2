// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryLoadBalancerListener queries LoadBalancerListener list
func (cli *ZSClient) QueryLoadBalancerListener(params *param.QueryParam) ([]view.LoadBalancerListenerInventoryView, error) {
	var resp []view.LoadBalancerListenerInventoryView
	return resp, cli.List("v1/load-balancers/listeners", params, &resp)
}
