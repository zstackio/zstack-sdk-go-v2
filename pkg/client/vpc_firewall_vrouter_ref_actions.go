// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcFirewallVRouterRef 查询VpcFirewallVRouterRef列表
func (cli *ZSClient) QueryVpcFirewallVRouterRef(params param.QueryParam) ([]view.QueryVpcFirewallVRouterRefView, error) {
	var resp []view.QueryVpcFirewallVRouterRefView
	return resp, cli.List("v1/vpcfirewalls/vrouters/refs", &params, &resp)
}

