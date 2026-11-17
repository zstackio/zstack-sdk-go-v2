// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

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

// GetDatabaseBackupFromImageStore gets DatabaseBackupFromImageStore by uuid
func (cli *ZSClient) GetDatabaseBackupFromImageStore() (*view.GetDatabaseBackupFromImageStoreView, error) {
	var resp view.GetDatabaseBackupFromImageStoreView
	if err := cli.GetWithRespKey("v1/database-backups/image-store", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncDatabaseBackupFromImageStoreBackupStorage operates on DatabaseBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncDatabaseBackupFromImageStoreBackupStorage(uuid string, params param.SyncDatabaseBackupFromImageStoreBackupStorageParam) (*view.DatabaseBackupInventoryView, error) {
	resp := view.DatabaseBackupInventoryView{}
	if err := cli.PutWithRespKey("v1/database-backups", uuid, "", map[string]interface{}{
		"syncDatabaseBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteExportedDatabaseBackupFromBackupStorage deletes ExportedDatabaseBackupFromBackupStorage
func (cli *ZSClient) DeleteExportedDatabaseBackupFromBackupStorage(databaseBackupUuid string, backupStorageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/exported-database-backup", databaseBackupUuid, fmt.Sprintf("backup-storage/%s", backupStorageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// ExportDatabaseBackupFromBackupStorage operates on DatabaseBackupFromBackupStorage
func (cli *ZSClient) ExportDatabaseBackupFromBackupStorage(databaseBackupUuid string, backupStorageUuid string, params param.ExportDatabaseBackupFromBackupStorageParam) (*view.ExportDatabaseBackupFromBackupStorageEventView, error) {
	resp := view.ExportDatabaseBackupFromBackupStorageEventView{}
	err := cli.PutWithSpec("v1/database-backups", databaseBackupUuid, fmt.Sprintf("backup-storage/%s/actions", backupStorageUuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverDatabaseFromBackup operates on DatabaseFromBackup
func (cli *ZSClient) RecoverDatabaseFromBackup(params param.RecoverDatabaseFromBackupParam) (*view.RecoverDatabaseFromBackupEventView, error) {
	resp := view.RecoverDatabaseFromBackupEventView{}
	if err := cli.PutWithRespKey("v1/database-backups/actions", "", "", map[string]interface{}{
		"recoverDatabaseFromBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
