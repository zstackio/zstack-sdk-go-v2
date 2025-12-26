// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateFirewallRule updates FirewallRule
func (cli *ZSClient) UpdateFirewallRule(uuid string, params param.UpdateFirewallRuleParam) (*view.UpdateFirewallRuleEventView, error) {
	resp := view.UpdateFirewallRuleEventView{}
	if err := cli.Put("v1/vpcfirewalls/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
