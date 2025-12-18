// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachRoleFromAccount operates on RoleFromAccount
func (cli *ZSClient) DetachRoleFromAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/accounts/{accountUuid}/roles/{roleUuid}", uuid, string(deleteMode))
}
