// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateFirewallRuleTemplate creates FirewallRuleTemplate
func (cli *ZSClient) CreateFirewallRuleTemplate(params param.CreateFirewallRuleTemplateParam) (*view.CreateFirewallRuleTemplateEventView, error) {
	resp := view.CreateFirewallRuleTemplateEventView{}
	if err := cli.Post("v1/vpcfirewalls/rules/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
