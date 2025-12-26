// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachScsiLunFromVmInstance operates on ScsiLunFromVmInstance
func (cli *ZSClient) DetachScsiLunFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances/{vmInstanceUuid}/scsi-lun/{uuid}", uuid, string(deleteMode))
}
