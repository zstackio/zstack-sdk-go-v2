// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMiniStorageHostRef queries MiniStorageHostRef list
func (cli *ZSClient) QueryMiniStorageHostRef(ctx context.Context, params *param.QueryParam) ([]view.MiniStorageHostRefInventoryView, error) {
	var resp []view.MiniStorageHostRefInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/mini/host-refs", params, &resp)
}

func (cli *ZSClient) GetMiniStorageHostRef(ctx context.Context, uuid string) (*view.MiniStorageHostRefInventoryView, error) {
	var resp view.MiniStorageHostRefInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/mini/host-refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMiniStorageHostRef Pagination
func (cli *ZSClient) PageMiniStorageHostRef(ctx context.Context, params *param.QueryParam) ([]view.MiniStorageHostRefInventoryView, int, error) {
	var miniStorageHostRefs []view.MiniStorageHostRefInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/mini/host-refs", params, &miniStorageHostRefs)
	return miniStorageHostRefs, total, err
}
