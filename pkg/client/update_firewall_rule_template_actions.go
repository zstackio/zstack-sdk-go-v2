// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateFirewallRuleTemplate updates FirewallRuleTemplate
func (cli *ZSClient) UpdateFirewallRuleTemplate(uuid string, params param.UpdateFirewallRuleTemplateParam) (*view.UpdateFirewallRuleTemplateEventView, error) {
	resp := view.UpdateFirewallRuleTemplateEventView{}
	if err := cli.Put("v1/vpcfirewalls/rules/template/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
