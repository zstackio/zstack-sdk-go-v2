// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPciDevice 查询PciDevice列表
func (cli *ZSClient) QueryPciDevice(params param.QueryParam) ([]view.QueryPciDeviceView, error) {
	var resp []view.QueryPciDeviceView
	return resp, cli.List("v1/pci-device/pci-devices", &params, &resp)
}

