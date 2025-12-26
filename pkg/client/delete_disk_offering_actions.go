// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteDiskOffering deletes DiskOffering
func (cli *ZSClient) DeleteDiskOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/disk-offerings/{uuid}", uuid, string(deleteMode))
}
