// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmInstancePciDeviceSpecRef 查询VmInstancePciDeviceSpecRef列表
func (cli *ZSClient) QueryVmInstancePciDeviceSpecRef(params param.QueryParam) ([]view.QueryVmInstancePciDeviceSpecRefView, error) {
	var resp []view.QueryVmInstancePciDeviceSpecRefView
	return resp, cli.List("v1/vm-instances/{vmInstanceUuid}/pci-device-specs", &params, &resp)
}

