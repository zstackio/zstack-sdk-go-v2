// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2ChassisGpuDevice queries BareMetal2ChassisGpuDevice list
func (cli *ZSClient) QueryBareMetal2ChassisGpuDevice(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ChassisGpuDeviceInventoryView, error) {
	var resp []view.BareMetal2ChassisGpuDeviceInventoryView
	return resp, cli.List(ctx, "v1/baremetal2/chassis/gpu-device/gpu-devices", params, &resp)
}

func (cli *ZSClient) GetBareMetal2ChassisGpuDevice(ctx context.Context, uuid string) (*view.BareMetal2ChassisGpuDeviceInventoryView, error) {
	var resp view.BareMetal2ChassisGpuDeviceInventoryView
	if err := cli.Get(ctx, "v1/baremetal2/chassis/gpu-device/gpu-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2ChassisGpuDevice Pagination
func (cli *ZSClient) PageBareMetal2ChassisGpuDevice(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ChassisGpuDeviceInventoryView, int, error) {
	var bareMetal2ChassisGpuDevices []view.BareMetal2ChassisGpuDeviceInventoryView
	total, err := cli.Page(ctx, "v1/baremetal2/chassis/gpu-device/gpu-devices", params, &bareMetal2ChassisGpuDevices)
	return bareMetal2ChassisGpuDevices, total, err
}
