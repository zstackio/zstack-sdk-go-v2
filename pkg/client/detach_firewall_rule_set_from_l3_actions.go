// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DetachFirewallRuleSetFromL3 operates on FirewallRuleSetFromL3
func (cli *ZSClient) DetachFirewallRuleSetFromL3(params param.DetachFirewallRuleSetFromL3Param) (*view.DetachFirewallRuleSetFromL3EventView, error) {
	resp := view.DetachFirewallRuleSetFromL3EventView{}
	if err := cli.Post("v1/vpcfirewalls/l3networks/{l3Uuid}/ruleSets/{ruleSetUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
