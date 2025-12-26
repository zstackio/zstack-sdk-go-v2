// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemovePciDeviceSpecFromVmInstance removes PciDeviceSpecFromVmInstance
func (cli *ZSClient) RemovePciDeviceSpecFromVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device-specs/{pciSpecUuid}/vm-instances/{vmInstanceUuid}", uuid, string(deleteMode))
}
