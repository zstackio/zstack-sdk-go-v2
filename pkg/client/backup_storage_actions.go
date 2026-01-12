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
// UpdateBackupStorage updates BackupStorage
func (cli *ZSClient) UpdateBackupStorage(uuid string, params param.UpdateBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.UpdateBackupStorageEventView
	err := cli.PutWithSpec("v1/backup-storage", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteBackupStorage deletes BackupStorage
func (cli *ZSClient) DeleteBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/backup-storage", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// ReconnectBackupStorage operates on BackupStorage
func (cli *ZSClient) ReconnectBackupStorage(uuid string, params param.ReconnectBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.ReconnectBackupStorageEventView
	err := cli.PutWithSpec("v1/backup-storage", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
