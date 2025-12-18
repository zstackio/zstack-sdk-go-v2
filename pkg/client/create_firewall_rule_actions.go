// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateFirewallRule creates FirewallRule
func (cli *ZSClient) CreateFirewallRule(params param.CreateFirewallRuleParam) (*view.CreateFirewallRuleEventView, error) {
	resp := view.CreateFirewallRuleEventView{}
	if err := cli.Post("v1/vpcfirewalls/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
