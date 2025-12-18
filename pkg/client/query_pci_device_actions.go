// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPciDevice queries PciDevice list
func (cli *ZSClient) QueryPciDevice(params param.QueryParam) ([]view.PciDeviceInventoryView, error) {
	var resp []view.PciDeviceInventoryView
	return resp, cli.List("v1/pci-device/pci-devices", &params, &resp)
}
