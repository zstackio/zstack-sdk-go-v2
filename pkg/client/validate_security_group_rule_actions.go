// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidateSecurityGroupRule 操作ValidateSecurityGroupRule
func (cli *ZSClient) ValidateSecurityGroupRule(params param.ValidateSecurityGroupRuleParam) (*view.ValidateSecurityGroupRuleView, error) {
	var resp view.ValidateSecurityGroupRuleView
	if err := cli.Get("v1/security-groups/{securityGroupUuid}/rules/validation", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

