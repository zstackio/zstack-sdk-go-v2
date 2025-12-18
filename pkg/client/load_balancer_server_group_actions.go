// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLoadBalancerServerGroup 查询LoadBalancerServerGroup列表
func (cli *ZSClient) QueryLoadBalancerServerGroup(params param.QueryParam) ([]view.QueryLoadBalancerServerGroupView, error) {
	var resp []view.QueryLoadBalancerServerGroupView
	return resp, cli.List("v1/load-balancers/servergroups", &params, &resp)
}

