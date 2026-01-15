// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGpuDevice queries GpuDevice list
func (cli *ZSClient) QueryGpuDevice(params *param.QueryParam) ([]view.GpuDeviceInventoryView, error) {
	var resp []view.GpuDeviceInventoryView
	return resp, cli.List("v1/gpu-device/gpu-devices", params, &resp)
}

// PageGpuDevice Pagination
func (cli *ZSClient) PageGpuDevice(params *param.QueryParam) ([]view.GpuDeviceInventoryView, int, error) {
	var gpuDevices []view.GpuDeviceInventoryView
	total, err := cli.Page("v1/gpu-device/gpu-devices", params, &gpuDevices)
	return gpuDevices, total, err
}
