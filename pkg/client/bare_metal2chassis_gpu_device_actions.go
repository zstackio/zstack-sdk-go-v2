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

func (cli *ZSClient) GetBareMetal2ChassisGpuDevice(uuid string) (*view.BareMetal2ChassisGpuDeviceInventoryView, error) {
	var resp view.BareMetal2ChassisGpuDeviceInventoryView
	if err := cli.Get("v1/baremetal2/chassis/gpu-device/gpu-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
