// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSecurityGroupRule 操作AddSecurityGroupRule
func (cli *ZSClient) AddSecurityGroupRule(params param.AddSecurityGroupRuleParam) (*view.AddSecurityGroupRuleEventView, error) {
	resp := view.AddSecurityGroupRuleEventView{}
	if err := cli.Post("v1/security-groups/{securityGroupUuid}/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

