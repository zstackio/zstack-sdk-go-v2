// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2ChassisPciDevice queries BareMetal2ChassisPciDevice list
func (cli *ZSClient) QueryBareMetal2ChassisPciDevice(params *param.QueryParam) ([]view.BareMetal2ChassisPciDeviceInventoryView, error) {
	var resp []view.BareMetal2ChassisPciDeviceInventoryView
	return resp, cli.List("v1/baremetal2/chassis/pci-device/pci-devices", params, &resp)
}
// UpdateBareMetal2ChassisPciDevice updates BareMetal2ChassisPciDevice
func (cli *ZSClient) UpdateBareMetal2ChassisPciDevice(uuid string, params param.UpdateBareMetal2ChassisPciDeviceParam) (*view.BareMetal2ChassisPciDeviceInventoryView, error) {
	var resp view.UpdateBareMetal2ChassisPciDeviceEventView
	if err := cli.Put("v1/baremetal2/chassis/pci-devices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
