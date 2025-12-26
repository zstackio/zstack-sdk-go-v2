// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPciDeviceOffering queries PciDeviceOffering list
func (cli *ZSClient) QueryPciDeviceOffering(params *param.QueryParam) ([]view.PciDeviceOfferingInventoryView, error) {
	var resp []view.PciDeviceOfferingInventoryView
	return resp, cli.List("v1/pci-device/pci-device-offerings", params, &resp)
}
