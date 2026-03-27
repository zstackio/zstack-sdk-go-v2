// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNativeHost queries NativeHost list
func (cli *ZSClient) QueryNativeHost(ctx context.Context, params *param.QueryParam) ([]view.NativeHostInventoryView, error) {
	var resp []view.NativeHostInventoryView
	return resp, cli.List(ctx, "v1/container/native/host", params, &resp)
}

func (cli *ZSClient) GetNativeHost(ctx context.Context, uuid string) (*view.NativeHostInventoryView, error) {
	var resp view.NativeHostInventoryView
	if err := cli.Get(ctx, "v1/container/native/host", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNativeHost Pagination
func (cli *ZSClient) PageNativeHost(ctx context.Context, params *param.QueryParam) ([]view.NativeHostInventoryView, int, error) {
	var nativeHosts []view.NativeHostInventoryView
	total, err := cli.Page(ctx, "v1/container/native/host", params, &nativeHosts)
	return nativeHosts, total, err
}
