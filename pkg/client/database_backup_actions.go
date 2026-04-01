// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteDatabaseBackup deletes DatabaseBackup
func (cli *ZSClient) DeleteDatabaseBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/database-backups", uuid, string(deleteMode))
}
// CreateDatabaseBackup creates DatabaseBackup
func (cli *ZSClient) CreateDatabaseBackup(params param.CreateDatabaseBackupParam) (*view.DatabaseBackupInventoryView, error) {
	resp := view.DatabaseBackupInventoryView{}
	if err := cli.Post("v1/database-backups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDatabaseBackupAsync Async
func (cli *ZSClient) CreateDatabaseBackupAsync(params param.CreateDatabaseBackupParam) (string, error) {

	resource := "v1/database-backups"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// QueryDatabaseBackup queries DatabaseBackup list
func (cli *ZSClient) QueryDatabaseBackup(params *param.QueryParam) ([]view.DatabaseBackupInventoryView, error) {
	var resp []view.DatabaseBackupInventoryView
	return resp, cli.List("v1/database-backups", params, &resp)
}

func (cli *ZSClient) GetDatabaseBackup(uuid string) (*view.DatabaseBackupInventoryView, error) {
	var resp view.DatabaseBackupInventoryView
	if err := cli.Get("v1/database-backups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDatabaseBackup Pagination
func (cli *ZSClient) PageDatabaseBackup(params *param.QueryParam) ([]view.DatabaseBackupInventoryView, int, error) {
	var databaseBackups []view.DatabaseBackupInventoryView
	total, err := cli.Page("v1/database-backups", params, &databaseBackups)
	return databaseBackups, total, err
}
// SyncDatabaseBackup operates on DatabaseBackup
func (cli *ZSClient) SyncDatabaseBackup(imageStoreUuid string, params param.SyncDatabaseBackupParam) (*view.DatabaseBackupInventoryView, error) {
	resp := view.DatabaseBackupInventoryView{}
	if err := cli.PutWithRespKey("v1/database-backups/imageStore", imageStoreUuid, "", map[string]interface{}{
		"syncDatabaseBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
