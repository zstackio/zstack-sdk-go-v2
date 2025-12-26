// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveMdevDeviceSpecFromVmInstance removes MdevDeviceSpecFromVmInstance
func (cli *ZSClient) RemoveMdevDeviceSpecFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/mdev-device-specs/{mdevSpecUuid}/vm-instances/{vmInstanceUuid}", uuid, string(deleteMode))
}
