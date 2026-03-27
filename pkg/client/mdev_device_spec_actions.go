// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateMdevDeviceSpec updates MdevDeviceSpec
func (cli *ZSClient) UpdateMdevDeviceSpec(ctx context.Context, uuid string, params param.UpdateMdevDeviceSpecParam) (*view.MdevDeviceSpecInventoryView, error) {
	resp := view.MdevDeviceSpecInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/mdev-device-specs", uuid, "", map[string]interface{}{
		"updateMdevDeviceSpec": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryMdevDeviceSpec queries MdevDeviceSpec list
func (cli *ZSClient) QueryMdevDeviceSpec(ctx context.Context, params *param.QueryParam) ([]view.MdevDeviceSpecInventoryView, error) {
	var resp []view.MdevDeviceSpecInventoryView
	return resp, cli.List(ctx, "v1/mdev-device-specs", params, &resp)
}

func (cli *ZSClient) GetMdevDeviceSpec(ctx context.Context, uuid string) (*view.MdevDeviceSpecInventoryView, error) {
	var resp view.MdevDeviceSpecInventoryView
	if err := cli.Get(ctx, "v1/mdev-device-specs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMdevDeviceSpec Pagination
func (cli *ZSClient) PageMdevDeviceSpec(ctx context.Context, params *param.QueryParam) ([]view.MdevDeviceSpecInventoryView, int, error) {
	var mdevDeviceSpecs []view.MdevDeviceSpecInventoryView
	total, err := cli.Page(ctx, "v1/mdev-device-specs", params, &mdevDeviceSpecs)
	return mdevDeviceSpecs, total, err
}
