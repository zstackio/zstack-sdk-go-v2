// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmInstanceDeviceAddressArchive queries VmInstanceDeviceAddressArchive list
func (cli *ZSClient) QueryVmInstanceDeviceAddressArchive(params *param.QueryParam) ([]view.VmInstanceDeviceAddressArchiveInventoryView, error) {
	var resp []view.VmInstanceDeviceAddressArchiveInventoryView
	return resp, cli.List("v1/vmInstance/device/address/archive", params, &resp)
}
