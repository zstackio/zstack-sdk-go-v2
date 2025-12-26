// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteDirectory deletes Directory
func (cli *ZSClient) DeleteDirectory(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/delete/directory", uuid, string(deleteMode))
}
