// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachRoleToAccount 操作RoleToAccount
func (cli *ZSClient) AttachRoleToAccount(params param.AttachRoleToAccountParam) (*view.AttachRoleToAccountEventView, error) {
	resp := view.AttachRoleToAccountEventView{}
	if err := cli.Post("v1/identities/accounts/{accountUuid}/roles/{roleUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

