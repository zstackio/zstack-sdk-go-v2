// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBareMetal2ChassisGpuDevice queries BareMetal2ChassisGpuDevice list
func (cli *ZSClient) QueryBareMetal2ChassisGpuDevice(params param.QueryParam) ([]view.BareMetal2ChassisGpuDeviceInventoryView, error) {
	var resp []view.BareMetal2ChassisGpuDeviceInventoryView
	return resp, cli.List("v1/baremetal2/chassis/gpu-device/gpu-devices", &params, &resp)
}
