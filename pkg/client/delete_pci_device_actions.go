// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeletePciDevice deletes PciDevice
func (cli *ZSClient) DeletePciDevice(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/pci-device/pci-devices/{uuid}", uuid, string(deleteMode))
}
