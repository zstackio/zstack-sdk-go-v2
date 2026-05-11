// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteDatabaseBackup deletes DatabaseBackup
func (cli *ZSClient) DeleteDatabaseBackup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/database-backups", uuid, string(deleteMode))
}
// CreateDatabaseBackup creates DatabaseBackup
func (cli *ZSClient) CreateDatabaseBackup(ctx context.Context, params param.CreateDatabaseBackupParam) (*view.DatabaseBackupInventoryView, error) {
	resp := view.DatabaseBackupInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/database-backups", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDatabaseBackupAsync Async
func (cli *ZSClient) CreateDatabaseBackupAsync(ctx context.Context, params param.CreateDatabaseBackupParam) (string, error) {

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
func (cli *ZSClient) QueryDatabaseBackup(ctx context.Context, params *param.QueryParam) ([]view.DatabaseBackupInventoryView, error) {
	var resp []view.DatabaseBackupInventoryView
	return resp, cli.List(ctx, "v1/database-backups", params, &resp)
}

func (cli *ZSClient) GetDatabaseBackup(ctx context.Context, uuid string) (*view.DatabaseBackupInventoryView, error) {
	var resp view.DatabaseBackupInventoryView
	if err := cli.Get(ctx, "v1/database-backups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDatabaseBackup Pagination
func (cli *ZSClient) PageDatabaseBackup(ctx context.Context, params *param.QueryParam) ([]view.DatabaseBackupInventoryView, int, error) {
	var databaseBackups []view.DatabaseBackupInventoryView
	total, err := cli.Page(ctx, "v1/database-backups", params, &databaseBackups)
	return databaseBackups, total, err
}
// SyncDatabaseBackup operates on DatabaseBackup
func (cli *ZSClient) SyncDatabaseBackup(ctx context.Context, imageStoreUuid string, params param.SyncDatabaseBackupParam) (*view.DatabaseBackupInventoryView, error) {
	resp := view.DatabaseBackupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/database-backups/imageStore", imageStoreUuid, "", map[string]interface{}{
		"syncDatabaseBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
