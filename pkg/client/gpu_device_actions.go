// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGpuDevice queries GpuDevice list
func (cli *ZSClient) QueryGpuDevice(params *param.QueryParam) ([]view.GpuDeviceInventoryView, error) {
	var resp []view.GpuDeviceInventoryView
	return resp, cli.List("v1/gpu-device/gpu-devices", params, &resp)
}
