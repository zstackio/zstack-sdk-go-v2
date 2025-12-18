// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachFirewallRuleSetFromL3 操作FirewallRuleSetFromL3
func (cli *ZSClient) DetachFirewallRuleSetFromL3(params param.DetachFirewallRuleSetFromL3Param) (*view.DetachFirewallRuleSetFromL3EventView, error) {
	resp := view.DetachFirewallRuleSetFromL3EventView{}
	if err := cli.Post("v1/vpcfirewalls/l3networks/{l3Uuid}/ruleSets/{ruleSetUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

