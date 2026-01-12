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
	var resp view.AddImageStoreBackupStorageEventView
	if err := cli.Post("v1/backup-storage/image-store", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateImageStoreBackupStorage updates ImageStoreBackupStorage
func (cli *ZSClient) UpdateImageStoreBackupStorage(uuid string, params param.UpdateImageStoreBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.UpdateBackupStorageEventView
	err := cli.PutWithSpec("v1/backup-storage/image-store", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
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
// ReconnectImageStoreBackupStorage operates on ImageStoreBackupStorage
func (cli *ZSClient) ReconnectImageStoreBackupStorage(uuid string, params param.ReconnectImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	var resp view.ReconnectImageStoreBackupStorageEventView
	err := cli.PutWithSpec("v1/backup-storage/image-store", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
