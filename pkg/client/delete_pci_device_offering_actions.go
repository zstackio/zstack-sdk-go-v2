// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePciDeviceOffering deletes PciDeviceOffering
func (cli *ZSClient) DeletePciDeviceOffering(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device/pci-device-offerings/{uuid}", uuid, string(deleteMode))
}
