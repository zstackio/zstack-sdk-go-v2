// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstanceDeviceAddressArchive queries VmInstanceDeviceAddressArchive list
func (cli *ZSClient) QueryVmInstanceDeviceAddressArchive(params *param.QueryParam) ([]view.VmInstanceDeviceAddressArchiveInventoryView, error) {
	var resp []view.VmInstanceDeviceAddressArchiveInventoryView
	return resp, cli.List("v1/vmInstance/device/address/archive", params, &resp)
}

// PageVmInstanceDeviceAddressArchive Pagination
func (cli *ZSClient) PageVmInstanceDeviceAddressArchive(params *param.QueryParam) ([]view.VmInstanceDeviceAddressArchiveInventoryView, int, error) {
	var vmInstanceDeviceAddressArchives []view.VmInstanceDeviceAddressArchiveInventoryView
	total, err := cli.Page("v1/vmInstance/device/address/archive", params, &vmInstanceDeviceAddressArchives)
	return vmInstanceDeviceAddressArchives, total, err
}
