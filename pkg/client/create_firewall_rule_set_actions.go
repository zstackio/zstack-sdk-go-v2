// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateFirewallRuleSet creates FirewallRuleSet
func (cli *ZSClient) CreateFirewallRuleSet(params param.CreateFirewallRuleSetParam) (*view.CreateFirewallRuleSetEventView, error) {
	resp := view.CreateFirewallRuleSetEventView{}
	if err := cli.Post("v1/vpcfirewalls/ruleSets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
