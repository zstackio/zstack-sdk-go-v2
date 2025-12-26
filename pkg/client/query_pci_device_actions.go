// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPciDevice queries PciDevice list
func (cli *ZSClient) QueryPciDevice(params *param.QueryParam) ([]view.PciDeviceInventoryView, error) {
	var resp []view.PciDeviceInventoryView
	return resp, cli.List("v1/pci-device/pci-devices", params, &resp)
}
