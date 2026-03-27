// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstanceDeviceAddressArchive queries VmInstanceDeviceAddressArchive list
func (cli *ZSClient) QueryVmInstanceDeviceAddressArchive(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceDeviceAddressArchiveInventoryView, error) {
	var resp []view.VmInstanceDeviceAddressArchiveInventoryView
	return resp, cli.List(ctx, "v1/vmInstance/device/address/archive", params, &resp)
}

func (cli *ZSClient) GetVmInstanceDeviceAddressArchive(ctx context.Context, uuid string) (*view.VmInstanceDeviceAddressArchiveInventoryView, error) {
	var resp view.VmInstanceDeviceAddressArchiveInventoryView
	if err := cli.Get(ctx, "v1/vmInstance/device/address/archive", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstanceDeviceAddressArchive Pagination
func (cli *ZSClient) PageVmInstanceDeviceAddressArchive(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceDeviceAddressArchiveInventoryView, int, error) {
	var vmInstanceDeviceAddressArchives []view.VmInstanceDeviceAddressArchiveInventoryView
	total, err := cli.Page(ctx, "v1/vmInstance/device/address/archive", params, &vmInstanceDeviceAddressArchives)
	return vmInstanceDeviceAddressArchives, total, err
}
