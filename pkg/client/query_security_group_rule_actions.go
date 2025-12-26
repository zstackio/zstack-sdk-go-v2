// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QuerySecurityGroupRule queries SecurityGroupRule list
func (cli *ZSClient) QuerySecurityGroupRule(params *param.QueryParam) ([]view.SecurityGroupRuleInventoryView, error) {
	var resp []view.SecurityGroupRuleInventoryView
	return resp, cli.List("v1/security-groups/rules", params, &resp)
}
