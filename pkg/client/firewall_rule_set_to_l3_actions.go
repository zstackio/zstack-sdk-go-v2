// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachFirewallRuleSetToL3 操作FirewallRuleSetToL3
func (cli *ZSClient) AttachFirewallRuleSetToL3(params param.AttachFirewallRuleSetToL3Param) (*view.AttachFirewallRuleSetToL3EventView, error) {
	resp := view.AttachFirewallRuleSetToL3EventView{}
	if err := cli.Post("v1/vpcfirewalls/ruleSets/{ruleSetUuid}/l3networks/{l3Uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

