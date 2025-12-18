// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPciDeviceSpec 查询PciDeviceSpec列表
func (cli *ZSClient) QueryPciDeviceSpec(params param.QueryParam) ([]view.QueryPciDeviceSpecView, error) {
	var resp []view.QueryPciDeviceSpecView
	return resp, cli.List("v1/pci-device-specs", &params, &resp)
}

