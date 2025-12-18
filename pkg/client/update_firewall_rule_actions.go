// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateFirewallRule updates FirewallRule
func (cli *ZSClient) UpdateFirewallRule(uuid string, params param.UpdateFirewallRuleParam) (*view.UpdateFirewallRuleEventView, error) {
	resp := view.UpdateFirewallRuleEventView{}
	if err := cli.Put("v1/vpcfirewalls/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
