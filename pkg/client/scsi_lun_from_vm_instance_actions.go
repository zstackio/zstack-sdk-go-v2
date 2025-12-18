// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachScsiLunFromVmInstance 操作ScsiLunFromVmInstance
func (cli *ZSClient) DetachScsiLunFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/scsi-lun/{uuid}", uuid, string(deleteMode))
}

