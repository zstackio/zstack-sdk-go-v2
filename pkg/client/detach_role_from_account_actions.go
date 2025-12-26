// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachRoleFromAccount operates on RoleFromAccount
func (cli *ZSClient) DetachRoleFromAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/accounts/{accountUuid}/roles/{roleUuid}", uuid, string(deleteMode))
}
