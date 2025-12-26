// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmInstanceDeviceAddressGroup queries VmInstanceDeviceAddressGroup list
func (cli *ZSClient) QueryVmInstanceDeviceAddressGroup(params *param.QueryParam) ([]view.VmInstanceDeviceAddressGroupInventoryView, error) {
	var resp []view.VmInstanceDeviceAddressGroupInventoryView
	return resp, cli.List("v1/vmInstance/device/address/group", params, &resp)
}
