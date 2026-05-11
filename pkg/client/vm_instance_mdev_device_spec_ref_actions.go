// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstanceMdevDeviceSpecRef queries VmInstanceMdevDeviceSpecRef list
func (cli *ZSClient) QueryVmInstanceMdevDeviceSpecRef(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstanceMdevDeviceSpecRefInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", params, &resp)
}

func (cli *ZSClient) GetVmInstanceMdevDeviceSpecRef(ctx context.Context, uuid string) (*view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp view.VmInstanceMdevDeviceSpecRefInventoryView
	if err := cli.Get(ctx, "v1/vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstanceMdevDeviceSpecRef Pagination
func (cli *ZSClient) PageVmInstanceMdevDeviceSpecRef(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, int, error) {
	var vmInstanceMdevDeviceSpecRefs []view.VmInstanceMdevDeviceSpecRefInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", params, &vmInstanceMdevDeviceSpecRefs)
	return vmInstanceMdevDeviceSpecRefs, total, err
}
