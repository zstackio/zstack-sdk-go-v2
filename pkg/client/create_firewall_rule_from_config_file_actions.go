// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateFirewallRuleFromConfigFile creates FirewallRuleFromConfigFile
func (cli *ZSClient) CreateFirewallRuleFromConfigFile(params param.CreateFirewallRuleFromConfigFileParam) (*view.CreateFirewallRuleFromConfigFileEventView, error) {
	resp := view.CreateFirewallRuleFromConfigFileEventView{}
	if err := cli.Post("v1/vpcfirewalls/rules/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
