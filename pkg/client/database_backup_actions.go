// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteDatabaseBackup deletes DatabaseBackup
func (cli *ZSClient) DeleteDatabaseBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/database-backups/{uuid}", uuid, string(deleteMode))
}
// CreateDatabaseBackup creates DatabaseBackup
func (cli *ZSClient) CreateDatabaseBackup(params param.CreateDatabaseBackupParam) (*view.DatabaseBackupInventoryView, error) {
	var resp view.CreateDatabaseBackupEventView
	if err := cli.Post("v1/database-backups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryDatabaseBackup queries DatabaseBackup list
func (cli *ZSClient) QueryDatabaseBackup(params *param.QueryParam) ([]view.DatabaseBackupInventoryView, error) {
	var resp []view.DatabaseBackupInventoryView
	return resp, cli.List("v1/database-backups", params, &resp)
}
// SyncDatabaseBackup operates on DatabaseBackup
func (cli *ZSClient) SyncDatabaseBackup(uuid string, params param.SyncDatabaseBackupParam) (*view.DatabaseBackupInventoryView, error) {
	resp := view.DatabaseBackupInventoryView{}
	if err := cli.Put("v1/database-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
