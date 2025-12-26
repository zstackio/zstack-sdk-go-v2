// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePciDevice deletes PciDevice
func (cli *ZSClient) DeletePciDevice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device/pci-devices/{uuid}", uuid, string(deleteMode))
}
