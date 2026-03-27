// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstancePciDeviceSpecRef queries VmInstancePciDeviceSpecRef list
func (cli *ZSClient) QueryVmInstancePciDeviceSpecRef(ctx context.Context, params *param.QueryParam) ([]view.VmInstancePciDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstancePciDeviceSpecRefInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/{vmInstanceUuid}/pci-device-specs", params, &resp)
}

func (cli *ZSClient) GetVmInstancePciDeviceSpecRef(ctx context.Context, uuid string) (*view.VmInstancePciDeviceSpecRefInventoryView, error) {
	var resp view.VmInstancePciDeviceSpecRefInventoryView
	if err := cli.Get(ctx, "v1/vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstancePciDeviceSpecRef Pagination
func (cli *ZSClient) PageVmInstancePciDeviceSpecRef(ctx context.Context, params *param.QueryParam) ([]view.VmInstancePciDeviceSpecRefInventoryView, int, error) {
	var vmInstancePciDeviceSpecRefs []view.VmInstancePciDeviceSpecRefInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/{vmInstanceUuid}/pci-device-specs", params, &vmInstancePciDeviceSpecRefs)
	return vmInstancePciDeviceSpecRefs, total, err
}
