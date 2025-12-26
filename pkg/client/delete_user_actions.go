// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteUser deletes User
func (cli *ZSClient) DeleteUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users/{uuid}", uuid, string(deleteMode))
}
