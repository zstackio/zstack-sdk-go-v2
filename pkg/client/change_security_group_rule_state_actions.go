// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSecurityGroupRuleState changes SecurityGroupRuleState
func (cli *ZSClient) ChangeSecurityGroupRuleState(uuid string, params param.ChangeSecurityGroupRuleStateParam) (*view.ChangeSecurityGroupRuleStateEventView, error) {
	resp := view.ChangeSecurityGroupRuleStateEventView{}
	if err := cli.Put("v1/security-groups/{securityGroupUuid}/rules/state/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
