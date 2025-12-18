// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryFirewallRuleSetL3Ref 查询FirewallRuleSetL3Ref列表
func (cli *ZSClient) QueryFirewallRuleSetL3Ref(params param.QueryParam) ([]view.QueryFirewallRuleSetL3RefView, error) {
	var resp []view.QueryFirewallRuleSetL3RefView
	return resp, cli.List("v1/vpcfirewalls/l3networks/rulesets/refs", &params, &resp)
}

