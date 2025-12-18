// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSecurityGroupRulePriority 更新SecurityGroupRulePriority
func (cli *ZSClient) UpdateSecurityGroupRulePriority(uuid string, params param.UpdateSecurityGroupRulePriorityParam) (*view.UpdateSecurityGroupRulePriorityEventView, error) {
	resp := view.UpdateSecurityGroupRulePriorityEventView{}
	if err := cli.Put("v1/security-groups/{securityGroupUuid}/rules/priority/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

