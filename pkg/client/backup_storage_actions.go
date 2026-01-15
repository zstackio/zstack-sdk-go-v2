// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBackupStorage queries BackupStorage list
func (cli *ZSClient) QueryBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List("v1/backup-storage", params, &resp)
}

func (cli *ZSClient) GetBackupStorage(uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get("v1/backup-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBackupStorage Pagination
func (cli *ZSClient) PageBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, int, error) {
	var backupStorages []view.BackupStorageInventoryView
	total, err := cli.Page("v1/backup-storage", params, &backupStorages)
	return backupStorages, total, err
}
// UpdateBackupStorage updates BackupStorage
func (cli *ZSClient) UpdateBackupStorage(uuid string, params param.UpdateBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Put("v1/backup-storage", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBackupStorage deletes BackupStorage
func (cli *ZSClient) DeleteBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage", uuid, string(deleteMode))
}
// ReconnectBackupStorage operates on BackupStorage
func (cli *ZSClient) ReconnectBackupStorage(uuid string, params param.ReconnectBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Put("v1/backup-storage", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
