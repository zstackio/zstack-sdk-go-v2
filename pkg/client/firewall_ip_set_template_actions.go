// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryFirewallIpSetTemplate 查询FirewallIpSetTemplate列表
func (cli *ZSClient) QueryFirewallIpSetTemplate(params param.QueryParam) ([]view.QueryFirewallIpSetTemplateView, error) {
	var resp []view.QueryFirewallIpSetTemplateView
	return resp, cli.List("v1/vpcfirewalls/ipset/templates", &params, &resp)
}

