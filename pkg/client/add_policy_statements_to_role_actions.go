// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddPolicyStatementsToRole 操作AddPolicyStatementsToRole
func (cli *ZSClient) AddPolicyStatementsToRole(params param.AddPolicyStatementsToRoleParam) (*view.AddPolicyStatementsToRoleEventView, error) {
	resp := view.AddPolicyStatementsToRoleEventView{}
	if err := cli.Post("v1/identities/roles/{uuid}/policy-statements", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

