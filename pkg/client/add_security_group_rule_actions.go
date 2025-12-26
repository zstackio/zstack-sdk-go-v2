// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSecurityGroupRule adds SecurityGroupRule
func (cli *ZSClient) AddSecurityGroupRule(params param.AddSecurityGroupRuleParam) (*view.AddSecurityGroupRuleEventView, error) {
	resp := view.AddSecurityGroupRuleEventView{}
	if err := cli.Post("v1/security-groups/{securityGroupUuid}/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
