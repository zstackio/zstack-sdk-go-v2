// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteAccount deletes Account
func (cli *ZSClient) DeleteAccount(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/{uuid}", uuid, string(deleteMode))
}
