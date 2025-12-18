// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmInstanceDeviceAddressArchive 查询VmInstanceDeviceAddressArchive列表
func (cli *ZSClient) QueryVmInstanceDeviceAddressArchive(params param.QueryParam) ([]view.QueryVmInstanceDeviceAddressArchiveView, error) {
	var resp []view.QueryVmInstanceDeviceAddressArchiveView
	return resp, cli.List("v1/vmInstance/device/address/archive", &params, &resp)
}

