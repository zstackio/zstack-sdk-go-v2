// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateFirewallRuleSet creates FirewallRuleSet
func (cli *ZSClient) CreateFirewallRuleSet(params param.CreateFirewallRuleSetParam) (*view.CreateFirewallRuleSetEventView, error) {
	resp := view.CreateFirewallRuleSetEventView{}
	if err := cli.Post("v1/vpcfirewalls/ruleSets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
