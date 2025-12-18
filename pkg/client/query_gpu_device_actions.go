// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGpuDevice queries GpuDevice list
func (cli *ZSClient) QueryGpuDevice(params param.QueryParam) ([]view.GpuDeviceInventoryView, error) {
	var resp []view.GpuDeviceInventoryView
	return resp, cli.List("v1/gpu-device/gpu-devices", &params, &resp)
}
