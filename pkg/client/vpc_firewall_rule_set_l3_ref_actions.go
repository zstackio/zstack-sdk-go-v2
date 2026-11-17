// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// QueryFirewallRuleSetL3Ref queries FirewallRuleSetL3Ref list
func (cli *ZSClient) QueryFirewallRuleSetL3Ref(params *param.QueryParam) ([]view.VpcFirewallRuleSetL3RefInventoryView, error) {
	var resp []view.VpcFirewallRuleSetL3RefInventoryView
	return resp, cli.List("v1/vpcfirewalls/l3networks/rulesets/refs", params, &resp)
}

func (cli *ZSClient) GetFirewallRuleSetL3Ref(uuid string) (*view.VpcFirewallRuleSetL3RefInventoryView, error) {
	var resp view.VpcFirewallRuleSetL3RefInventoryView
	if err := cli.Get("v1/vpcfirewalls/l3networks/rulesets/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRuleSetL3Ref Pagination
func (cli *ZSClient) PageFirewallRuleSetL3Ref(params *param.QueryParam) ([]view.VpcFirewallRuleSetL3RefInventoryView, int, error) {
	var firewallRuleSetL3Refs []view.VpcFirewallRuleSetL3RefInventoryView
	total, err := cli.Page("v1/vpcfirewalls/l3networks/rulesets/refs", params, &firewallRuleSetL3Refs)
	return firewallRuleSetL3Refs, total, err
}
