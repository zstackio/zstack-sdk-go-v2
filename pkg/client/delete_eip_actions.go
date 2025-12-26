// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteEip deletes Eip
func (cli *ZSClient) DeleteEip(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/eips/{uuid}", uuid, string(deleteMode))
}
