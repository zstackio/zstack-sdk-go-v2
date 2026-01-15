// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2ChassisGpuDevice queries BareMetal2ChassisGpuDevice list
func (cli *ZSClient) QueryBareMetal2ChassisGpuDevice(params *param.QueryParam) ([]view.BareMetal2ChassisGpuDeviceInventoryView, error) {
	var resp []view.BareMetal2ChassisGpuDeviceInventoryView
	return resp, cli.List("v1/baremetal2/chassis/gpu-device/gpu-devices", params, &resp)
}

// PageBareMetal2ChassisGpuDevice Pagination
func (cli *ZSClient) PageBareMetal2ChassisGpuDevice(params *param.QueryParam) ([]view.BareMetal2ChassisGpuDeviceInventoryView, int, error) {
	var bareMetal2ChassisGpuDevices []view.BareMetal2ChassisGpuDeviceInventoryView
	total, err := cli.Page("v1/baremetal2/chassis/gpu-device/gpu-devices", params, &bareMetal2ChassisGpuDevices)
	return bareMetal2ChassisGpuDevices, total, err
}
