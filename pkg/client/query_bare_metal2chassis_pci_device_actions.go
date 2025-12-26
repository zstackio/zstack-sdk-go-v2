// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBareMetal2ChassisPciDevice queries BareMetal2ChassisPciDevice list
func (cli *ZSClient) QueryBareMetal2ChassisPciDevice(params *param.QueryParam) ([]view.BareMetal2ChassisPciDeviceInventoryView, error) {
	var resp []view.BareMetal2ChassisPciDeviceInventoryView
	return resp, cli.List("v1/baremetal2/chassis/pci-device/pci-devices", params, &resp)
}
