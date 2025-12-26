// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// LogOut operates on LogOut
func (cli *ZSClient) LogOut(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/sessions/{sessionUuid}", uuid, string(deleteMode))
}
