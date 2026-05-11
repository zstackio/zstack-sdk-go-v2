// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryDGpuDevice queries DGpuDevice list
func (cli *ZSClient) QueryDGpuDevice(ctx context.Context, params *param.QueryParam) ([]view.DGpuDeviceInventoryView, error) {
	var resp []view.DGpuDeviceInventoryView
	return resp, cli.List(ctx, "v1/gpu-device/dgpu-devices", params, &resp)
}

func (cli *ZSClient) GetDGpuDevice(ctx context.Context, uuid string) (*view.DGpuDeviceInventoryView, error) {
	var resp view.DGpuDeviceInventoryView
	if err := cli.Get(ctx, "v1/gpu-device/dgpu-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDGpuDevice Pagination
func (cli *ZSClient) PageDGpuDevice(ctx context.Context, params *param.QueryParam) ([]view.DGpuDeviceInventoryView, int, error) {
	var dGpuDevices []view.DGpuDeviceInventoryView
	total, err := cli.Page(ctx, "v1/gpu-device/dgpu-devices", params, &dGpuDevices)
	return dGpuDevices, total, err
}
