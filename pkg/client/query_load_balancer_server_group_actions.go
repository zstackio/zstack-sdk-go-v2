// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLoadBalancerServerGroup queries LoadBalancerServerGroup list
func (cli *ZSClient) QueryLoadBalancerServerGroup(params param.QueryParam) ([]view.LoadBalancerServerGroupInventoryView, error) {
	var resp []view.LoadBalancerServerGroupInventoryView
	return resp, cli.List("v1/load-balancers/servergroups", &params, &resp)
}
