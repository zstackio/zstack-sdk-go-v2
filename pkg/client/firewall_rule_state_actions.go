// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeFirewallRuleState 操作FirewallRuleState
func (cli *ZSClient) ChangeFirewallRuleState(uuid string, params param.ChangeFirewallRuleStateParam) (*view.ChangeFirewallRuleStateEventView, error) {
	resp := view.ChangeFirewallRuleStateEventView{}
	if err := cli.Put("v1/vpcfirewalls/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

