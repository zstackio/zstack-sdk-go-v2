// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryDGpuProfile queries DGpuProfile list
func (cli *ZSClient) QueryDGpuProfile(ctx context.Context, params *param.QueryParam) ([]view.DGpuProfileInventoryView, error) {
	var resp []view.DGpuProfileInventoryView
	return resp, cli.List(ctx, "v1/gpu-device/dgpu-profiles", params, &resp)
}

func (cli *ZSClient) GetDGpuProfile(ctx context.Context, uuid string) (*view.DGpuProfileInventoryView, error) {
	var resp view.DGpuProfileInventoryView
	if err := cli.Get(ctx, "v1/gpu-device/dgpu-profiles", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDGpuProfile Pagination
func (cli *ZSClient) PageDGpuProfile(ctx context.Context, params *param.QueryParam) ([]view.DGpuProfileInventoryView, int, error) {
	var dGpuProfiles []view.DGpuProfileInventoryView
	total, err := cli.Page(ctx, "v1/gpu-device/dgpu-profiles", params, &dGpuProfiles)
	return dGpuProfiles, total, err
}
// SetDGpuProfile operates on DGpuProfile
func (cli *ZSClient) SetDGpuProfile(ctx context.Context, gpuSpecUuid string, params param.SetDGpuProfileParam) (*view.DGpuProfileInventoryView, error) {
	resp := view.DGpuProfileInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/gpu-device/gpu-device-specs", gpuSpecUuid, "", map[string]interface{}{
		"setDGpuProfile": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
