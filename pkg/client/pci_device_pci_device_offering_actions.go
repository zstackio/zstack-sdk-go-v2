// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPciDevicePciDeviceOffering 查询PciDevicePciDeviceOffering列表
func (cli *ZSClient) QueryPciDevicePciDeviceOffering(params param.QueryParam) ([]view.QueryPciDevicePciDeviceOfferingView, error) {
	var resp []view.QueryPciDevicePciDeviceOfferingView
	return resp, cli.List("v1/pci-devices/pci-devices/pci-device-offerings", &params, &resp)
}

