// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteNvmeServer deletes NvmeServer
func (cli *ZSClient) DeleteNvmeServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/storage-devices/nvme/servers/{uuid}", uuid, string(deleteMode))
}
