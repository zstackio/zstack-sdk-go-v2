// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddCephPrimaryStorage adds CephPrimaryStorage
func (cli *ZSClient) AddCephPrimaryStorage(ctx context.Context, params param.AddCephPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/primary-storage/ceph", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCephPrimaryStorage queries CephPrimaryStorage list
func (cli *ZSClient) QueryCephPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/ceph", params, &resp)
}

func (cli *ZSClient) GetCephPrimaryStorage(ctx context.Context, uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/ceph", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCephPrimaryStorage Pagination
func (cli *ZSClient) PageCephPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.PrimaryStorageInventoryView, int, error) {
	var cephPrimaryStorages []view.PrimaryStorageInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/ceph", params, &cephPrimaryStorages)
	return cephPrimaryStorages, total, err
}
