// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBareMetal2ChassisPciDevice 查询BareMetal2ChassisPciDevice列表
func (cli *ZSClient) QueryBareMetal2ChassisPciDevice(params param.QueryParam) ([]view.QueryBareMetal2ChassisPciDeviceView, error) {
	var resp []view.QueryBareMetal2ChassisPciDeviceView
	return resp, cli.List("v1/baremetal2/chassis/pci-device/pci-devices", &params, &resp)
}

