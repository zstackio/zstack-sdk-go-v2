// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteNasFileSystem deletes NasFileSystem
func (cli *ZSClient) DeleteNasFileSystem(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage/nas/{uuid}", uuid, string(deleteMode))
}
