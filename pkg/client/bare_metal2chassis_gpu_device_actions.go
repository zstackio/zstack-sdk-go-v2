// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2ChassisGpuDevice queries BareMetal2ChassisGpuDevice list
func (cli *ZSClient) QueryBareMetal2ChassisGpuDevice(params *param.QueryParam) ([]view.BareMetal2ChassisGpuDeviceInventoryView, error) {
	var resp []view.BareMetal2ChassisGpuDeviceInventoryView
	return resp, cli.List("v1/baremetal2/chassis/gpu-device/gpu-devices", params, &resp)
}
