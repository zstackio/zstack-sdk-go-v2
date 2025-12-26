// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ValidateSecurityGroupRule operates on ValidateSecurityGroupRule
func (cli *ZSClient) ValidateSecurityGroupRule(params param.ValidateSecurityGroupRuleParam) (*view.ValidateSecurityGroupRuleView, error) {
	var resp view.ValidateSecurityGroupRuleView
	if err := cli.Get("v1/security-groups/{securityGroupUuid}/rules/validation", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
