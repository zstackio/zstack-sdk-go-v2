// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryLoadBalancer queries LoadBalancer list
func (cli *ZSClient) QueryLoadBalancer(params *param.QueryParam) ([]view.LoadBalancerInventoryView, error) {
	var resp []view.LoadBalancerInventoryView
	return resp, cli.List("v1/load-balancers", params, &resp)
}
