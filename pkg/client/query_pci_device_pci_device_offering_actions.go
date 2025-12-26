// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPciDevicePciDeviceOffering queries PciDevicePciDeviceOffering list
func (cli *ZSClient) QueryPciDevicePciDeviceOffering(params *param.QueryParam) ([]view.PciDevicePciDeviceOfferingRefInventoryView, error) {
	var resp []view.PciDevicePciDeviceOfferingRefInventoryView
	return resp, cli.List("v1/pci-devices/pci-devices/pci-device-offerings", params, &resp)
}
