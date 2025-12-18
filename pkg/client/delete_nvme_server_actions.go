// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteNvmeServer deletes NvmeServer
func (cli *ZSClient) DeleteNvmeServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/storage-devices/nvme/servers/{uuid}", uuid, string(deleteMode))
}
