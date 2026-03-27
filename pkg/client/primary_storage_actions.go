// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPrimaryStorage queries PrimaryStorage list
func (cli *ZSClient) QueryPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List(ctx, "v1/primary-storage", params, &resp)
}

func (cli *ZSClient) GetPrimaryStorage(ctx context.Context, uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.Get(ctx, "v1/primary-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePrimaryStorage Pagination
func (cli *ZSClient) PagePrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.PrimaryStorageInventoryView, int, error) {
	var primaryStorages []view.PrimaryStorageInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage", params, &primaryStorages)
	return primaryStorages, total, err
}
// ReconnectPrimaryStorage operates on PrimaryStorage
func (cli *ZSClient) ReconnectPrimaryStorage(ctx context.Context, uuid string, params param.ReconnectPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage", uuid, "", map[string]interface{}{
		"reconnectPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePrimaryStorage updates PrimaryStorage
func (cli *ZSClient) UpdatePrimaryStorage(ctx context.Context, uuid string, params param.UpdatePrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/primary-storage", uuid, "", map[string]interface{}{
		"updatePrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePrimaryStorage deletes PrimaryStorage
func (cli *ZSClient) DeletePrimaryStorage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/primary-storage", uuid, string(deleteMode))
}
