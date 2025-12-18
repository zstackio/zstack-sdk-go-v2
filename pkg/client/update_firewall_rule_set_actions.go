// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateFirewallRuleSet updates FirewallRuleSet
func (cli *ZSClient) UpdateFirewallRuleSet(uuid string, params param.UpdateFirewallRuleSetParam) (*view.UpdateFirewallRuleSetEventView, error) {
	resp := view.UpdateFirewallRuleSetEventView{}
	if err := cli.Put("v1/vpcfirewalls/ruleSets/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
