// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySftpBackupStorage queries SftpBackupStorage list
func (cli *ZSClient) QuerySftpBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.SftpBackupStorageInventoryView, error) {
	var resp []view.SftpBackupStorageInventoryView
	return resp, cli.List(ctx, "v1/backup-storage/sftp", params, &resp)
}

func (cli *ZSClient) GetSftpBackupStorage(ctx context.Context, uuid string) (*view.SftpBackupStorageInventoryView, error) {
	var resp view.SftpBackupStorageInventoryView
	if err := cli.Get(ctx, "v1/backup-storage/sftp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSftpBackupStorage Pagination
func (cli *ZSClient) PageSftpBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.SftpBackupStorageInventoryView, int, error) {
	var sftpBackupStorages []view.SftpBackupStorageInventoryView
	total, err := cli.Page(ctx, "v1/backup-storage/sftp", params, &sftpBackupStorages)
	return sftpBackupStorages, total, err
}
// UpdateSftpBackupStorage updates SftpBackupStorage
func (cli *ZSClient) UpdateSftpBackupStorage(ctx context.Context, uuid string, params param.UpdateSftpBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/sftp", uuid, "", map[string]interface{}{
		"updateSftpBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ReconnectSftpBackupStorage operates on SftpBackupStorage
func (cli *ZSClient) ReconnectSftpBackupStorage(ctx context.Context, uuid string, params param.ReconnectSftpBackupStorageParam) (*view.SftpBackupStorageInventoryView, error) {
	resp := view.SftpBackupStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/sftp", uuid, "", map[string]interface{}{
		"reconnectSftpBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddSftpBackupStorage adds SftpBackupStorage
func (cli *ZSClient) AddSftpBackupStorage(ctx context.Context, params param.AddSftpBackupStorageParam) (*view.SftpBackupStorageInventoryView, error) {
	resp := view.SftpBackupStorageInventoryView{}
	if err := cli.Post(ctx, "v1/backup-storage/sftp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
