// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSecurityGroupRulePriority updates SecurityGroupRulePriority
func (cli *ZSClient) UpdateSecurityGroupRulePriority(uuid string, params param.UpdateSecurityGroupRulePriorityParam) (*view.UpdateSecurityGroupRulePriorityEventView, error) {
	resp := view.UpdateSecurityGroupRulePriorityEventView{}
	if err := cli.Put("v1/security-groups/{securityGroupUuid}/rules/priority/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
