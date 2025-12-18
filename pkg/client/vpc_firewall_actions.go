// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVpcFirewall 查询VpcFirewall列表
func (cli *ZSClient) QueryVpcFirewall(params param.QueryParam) ([]view.QueryVpcFirewallView, error) {
	var resp []view.QueryVpcFirewallView
	return resp, cli.List("v1/vpcfirewalls", &params, &resp)
}

