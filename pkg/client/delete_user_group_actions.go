// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteUserGroup deletes UserGroup
func (cli *ZSClient) DeleteUserGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/groups/{uuid}", uuid, string(deleteMode))
}
