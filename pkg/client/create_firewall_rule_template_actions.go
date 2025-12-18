// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateFirewallRuleTemplate creates FirewallRuleTemplate
func (cli *ZSClient) CreateFirewallRuleTemplate(params param.CreateFirewallRuleTemplateParam) (*view.CreateFirewallRuleTemplateEventView, error) {
	resp := view.CreateFirewallRuleTemplateEventView{}
	if err := cli.Post("v1/vpcfirewalls/rules/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
