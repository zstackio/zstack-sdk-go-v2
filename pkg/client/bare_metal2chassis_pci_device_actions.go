// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2ChassisPciDevice queries BareMetal2ChassisPciDevice list
func (cli *ZSClient) QueryBareMetal2ChassisPciDevice(params *param.QueryParam) ([]view.BareMetal2ChassisPciDeviceInventoryView, error) {
	var resp []view.BareMetal2ChassisPciDeviceInventoryView
	return resp, cli.List("v1/baremetal2/chassis/pci-device/pci-devices", params, &resp)
}

func (cli *ZSClient) GetBareMetal2ChassisPciDevice(uuid string) (*view.BareMetal2ChassisPciDeviceInventoryView, error) {
	var resp view.BareMetal2ChassisPciDeviceInventoryView
	if err := cli.Get("v1/baremetal2/chassis/pci-device/pci-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateBareMetal2ChassisPciDevice updates BareMetal2ChassisPciDevice
func (cli *ZSClient) UpdateBareMetal2ChassisPciDevice(uuid string, params param.UpdateBareMetal2ChassisPciDeviceParam) (*view.BareMetal2ChassisPciDeviceInventoryView, error) {
	var resp view.UpdateBareMetal2ChassisPciDeviceEventView
	err := cli.PutWithSpec("v1/baremetal2/chassis/pci-devices", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
