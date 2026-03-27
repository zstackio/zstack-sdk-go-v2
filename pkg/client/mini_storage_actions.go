// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMiniStorage queries MiniStorage list
func (cli *ZSClient) QueryMiniStorage(ctx context.Context, params *param.QueryParam) ([]view.MiniStorageInventoryView, error) {
	var resp []view.MiniStorageInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/mini", params, &resp)
}

func (cli *ZSClient) GetMiniStorage(ctx context.Context, uuid string) (*view.MiniStorageInventoryView, error) {
	var resp view.MiniStorageInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/mini", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMiniStorage Pagination
func (cli *ZSClient) PageMiniStorage(ctx context.Context, params *param.QueryParam) ([]view.MiniStorageInventoryView, int, error) {
	var miniStorages []view.MiniStorageInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/mini", params, &miniStorages)
	return miniStorages, total, err
}
// AddMiniStorage adds MiniStorage
func (cli *ZSClient) AddMiniStorage(ctx context.Context, params param.AddMiniStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(ctx, "v1/primary-storage/mini", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
