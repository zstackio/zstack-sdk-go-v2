// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddImageStoreBackupStorage adds ImageStoreBackupStorage
func (cli *ZSClient) AddImageStoreBackupStorage(ctx context.Context, params param.AddImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.Post(ctx, "v1/backup-storage/image-store", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateImageStoreBackupStorage updates ImageStoreBackupStorage
func (cli *ZSClient) UpdateImageStoreBackupStorage(ctx context.Context, uuid string, params param.UpdateImageStoreBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/image-store", uuid, "", map[string]interface{}{
		"updateImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryImageStoreBackupStorage queries ImageStoreBackupStorage list
func (cli *ZSClient) QueryImageStoreBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.ImageStoreBackupStorageInventoryView, error) {
	var resp []view.ImageStoreBackupStorageInventoryView
	return resp, cli.List(ctx, "v1/backup-storage/image-store", params, &resp)
}

func (cli *ZSClient) GetImageStoreBackupStorage(ctx context.Context, uuid string) (*view.ImageStoreBackupStorageInventoryView, error) {
	var resp view.ImageStoreBackupStorageInventoryView
	if err := cli.Get(ctx, "v1/backup-storage/image-store", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImageStoreBackupStorage Pagination
func (cli *ZSClient) PageImageStoreBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.ImageStoreBackupStorageInventoryView, int, error) {
	var imageStoreBackupStorages []view.ImageStoreBackupStorageInventoryView
	total, err := cli.Page(ctx, "v1/backup-storage/image-store", params, &imageStoreBackupStorages)
	return imageStoreBackupStorages, total, err
}
// ReconnectImageStoreBackupStorage operates on ImageStoreBackupStorage
func (cli *ZSClient) ReconnectImageStoreBackupStorage(ctx context.Context, uuid string, params param.ReconnectImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/image-store", uuid, "", map[string]interface{}{
		"reconnectImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
