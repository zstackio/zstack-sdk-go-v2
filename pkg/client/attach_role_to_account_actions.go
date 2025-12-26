// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachRoleToAccount operates on RoleToAccount
func (cli *ZSClient) AttachRoleToAccount(params param.AttachRoleToAccountParam) (*view.AttachRoleToAccountEventView, error) {
	resp := view.AttachRoleToAccountEventView{}
	if err := cli.Post("v1/identities/accounts/{accountUuid}/roles/{roleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
