// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteDiskOffering deletes DiskOffering
func (cli *ZSClient) DeleteDiskOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/disk-offerings/{uuid}", uuid, string(deleteMode))
}
