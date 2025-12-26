// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachFirewallRuleSetToL3 operates on FirewallRuleSetToL3
func (cli *ZSClient) AttachFirewallRuleSetToL3(params param.AttachFirewallRuleSetToL3Param) (*view.AttachFirewallRuleSetToL3EventView, error) {
	resp := view.AttachFirewallRuleSetToL3EventView{}
	if err := cli.Post("v1/vpcfirewalls/ruleSets/{ruleSetUuid}/l3networks/{l3Uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
