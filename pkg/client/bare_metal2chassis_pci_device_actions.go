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

// PageBareMetal2ChassisPciDevice Pagination
func (cli *ZSClient) PageBareMetal2ChassisPciDevice(params *param.QueryParam) ([]view.BareMetal2ChassisPciDeviceInventoryView, int, error) {
	var bareMetal2ChassisPciDevices []view.BareMetal2ChassisPciDeviceInventoryView
	total, err := cli.Page("v1/baremetal2/chassis/pci-device/pci-devices", params, &bareMetal2ChassisPciDevices)
	return bareMetal2ChassisPciDevices, total, err
}
// UpdateBareMetal2ChassisPciDevice updates BareMetal2ChassisPciDevice
func (cli *ZSClient) UpdateBareMetal2ChassisPciDevice(uuid string, params param.UpdateBareMetal2ChassisPciDeviceParam) (*view.BareMetal2ChassisPciDeviceInventoryView, error) {
	resp := view.BareMetal2ChassisPciDeviceInventoryView{}
	if err := cli.PutWithRespKey("v1/baremetal2/chassis/pci-devices", uuid, "", map[string]interface{}{
		"updateBareMetal2ChassisPciDevice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
