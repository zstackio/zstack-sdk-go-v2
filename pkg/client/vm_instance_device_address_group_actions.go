// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstanceDeviceAddressGroup queries VmInstanceDeviceAddressGroup list
func (cli *ZSClient) QueryVmInstanceDeviceAddressGroup(params *param.QueryParam) ([]view.VmInstanceDeviceAddressGroupInventoryView, error) {
	var resp []view.VmInstanceDeviceAddressGroupInventoryView
	return resp, cli.List("v1/vmInstance/device/address/group", params, &resp)
}

// PageVmInstanceDeviceAddressGroup Pagination
func (cli *ZSClient) PageVmInstanceDeviceAddressGroup(params *param.QueryParam) ([]view.VmInstanceDeviceAddressGroupInventoryView, int, error) {
	var vmInstanceDeviceAddressGroups []view.VmInstanceDeviceAddressGroupInventoryView
	total, err := cli.Page("v1/vmInstance/device/address/group", params, &vmInstanceDeviceAddressGroups)
	return vmInstanceDeviceAddressGroups, total, err
}
