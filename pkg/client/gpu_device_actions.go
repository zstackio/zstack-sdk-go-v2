// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryGpuDevice 查询GpuDevice列表
func (cli *ZSClient) QueryGpuDevice(params param.QueryParam) ([]view.QueryGpuDeviceView, error) {
	var resp []view.QueryGpuDeviceView
	return resp, cli.List("v1/gpu-device/gpu-devices", &params, &resp)
}

