// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmInstanceDeviceAddressGroup 查询VmInstanceDeviceAddressGroup列表
func (cli *ZSClient) QueryVmInstanceDeviceAddressGroup(params param.QueryParam) ([]view.QueryVmInstanceDeviceAddressGroupView, error) {
	var resp []view.QueryVmInstanceDeviceAddressGroupView
	return resp, cli.List("v1/vmInstance/device/address/group", &params, &resp)
}

