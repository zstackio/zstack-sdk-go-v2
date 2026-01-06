// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBackupStorage queries BackupStorage list
func (cli *ZSClient) QueryBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List("v1/backup-storage", params, &resp)
}
// UpdateBackupStorage updates BackupStorage
func (cli *ZSClient) UpdateBackupStorage(uuid string, params param.UpdateBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.UpdateBackupStorageEventView
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteBackupStorage deletes BackupStorage
func (cli *ZSClient) DeleteBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage/{uuid}", uuid, string(deleteMode))
}
// ReconnectBackupStorage operates on BackupStorage
func (cli *ZSClient) ReconnectBackupStorage(uuid string, params param.ReconnectBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.ReconnectBackupStorageEventView
	if err := cli.Put("v1/backup-storage/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
