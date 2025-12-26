// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteDataVolume deletes DataVolume
func (cli *ZSClient) DeleteDataVolume(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes/{uuid}", uuid, string(deleteMode))
}
