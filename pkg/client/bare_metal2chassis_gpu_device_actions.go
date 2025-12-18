// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBareMetal2ChassisGpuDevice 查询BareMetal2ChassisGpuDevice列表
func (cli *ZSClient) QueryBareMetal2ChassisGpuDevice(params param.QueryParam) ([]view.QueryBareMetal2ChassisGpuDeviceView, error) {
	var resp []view.QueryBareMetal2ChassisGpuDeviceView
	return resp, cli.List("v1/baremetal2/chassis/gpu-device/gpu-devices", &params, &resp)
}

