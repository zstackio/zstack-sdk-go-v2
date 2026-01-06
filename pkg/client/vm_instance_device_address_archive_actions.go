// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstanceDeviceAddressArchive queries VmInstanceDeviceAddressArchive list
func (cli *ZSClient) QueryVmInstanceDeviceAddressArchive(params *param.QueryParam) ([]view.VmInstanceDeviceAddressArchiveInventoryView, error) {
	var resp []view.VmInstanceDeviceAddressArchiveInventoryView
	return resp, cli.List("v1/vmInstance/device/address/archive", params, &resp)
}
