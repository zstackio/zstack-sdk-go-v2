// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateFirewallRule creates FirewallRule
func (cli *ZSClient) CreateFirewallRule(params param.CreateFirewallRuleParam) (*view.CreateFirewallRuleEventView, error) {
	resp := view.CreateFirewallRuleEventView{}
	if err := cli.Post("v1/vpcfirewalls/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
