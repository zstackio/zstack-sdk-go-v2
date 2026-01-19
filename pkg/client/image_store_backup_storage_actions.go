// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddImageStoreBackupStorage adds ImageStoreBackupStorage
func (cli *ZSClient) AddImageStoreBackupStorage(params param.AddImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.Post("v1/backup-storage/image-store", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateImageStoreBackupStorage updates ImageStoreBackupStorage
func (cli *ZSClient) UpdateImageStoreBackupStorage(uuid string, params param.UpdateImageStoreBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Put("v1/backup-storage/image-store", uuid, map[string]interface{}{
		"updateImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryImageStoreBackupStorage queries ImageStoreBackupStorage list
func (cli *ZSClient) QueryImageStoreBackupStorage(params *param.QueryParam) ([]view.ImageStoreBackupStorageInventoryView, error) {
	var resp []view.ImageStoreBackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/image-store", params, &resp)
}

func (cli *ZSClient) GetImageStoreBackupStorage(uuid string) (*view.ImageStoreBackupStorageInventoryView, error) {
	var resp view.ImageStoreBackupStorageInventoryView
	if err := cli.Get("v1/backup-storage/image-store", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImageStoreBackupStorage Pagination
func (cli *ZSClient) PageImageStoreBackupStorage(params *param.QueryParam) ([]view.ImageStoreBackupStorageInventoryView, int, error) {
	var imageStoreBackupStorages []view.ImageStoreBackupStorageInventoryView
	total, err := cli.Page("v1/backup-storage/image-store", params, &imageStoreBackupStorages)
	return imageStoreBackupStorages, total, err
}
// ReconnectImageStoreBackupStorage operates on ImageStoreBackupStorage
func (cli *ZSClient) ReconnectImageStoreBackupStorage(uuid string, params param.ReconnectImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.Put("v1/backup-storage/image-store", uuid, map[string]interface{}{
		"reconnectImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
