// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeletePciDeviceOffering deletes PciDeviceOffering
func (cli *ZSClient) DeletePciDeviceOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device/pci-device-offerings/{uuid}", uuid, string(deleteMode))
}
