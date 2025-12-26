// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteRole deletes Role
func (cli *ZSClient) DeleteRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/roles/{uuid}", uuid, string(deleteMode))
}
