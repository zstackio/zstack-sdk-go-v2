// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeSecurityGroupRule changes SecurityGroupRule
func (cli *ZSClient) ChangeSecurityGroupRule(uuid string, params param.ChangeSecurityGroupRuleParam) (*view.ChangeSecurityGroupRuleEventView, error) {
	resp := view.ChangeSecurityGroupRuleEventView{}
	if err := cli.Put("v1/security-groups/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
