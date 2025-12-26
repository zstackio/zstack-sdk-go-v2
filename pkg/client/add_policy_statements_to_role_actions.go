// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddPolicyStatementsToRole adds PolicyStatementsToRole
func (cli *ZSClient) AddPolicyStatementsToRole(params param.AddPolicyStatementsToRoleParam) (*view.AddPolicyStatementsToRoleEventView, error) {
	resp := view.AddPolicyStatementsToRoleEventView{}
	if err := cli.Post("v1/identities/roles/{uuid}/policy-statements", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
