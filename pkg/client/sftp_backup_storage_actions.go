// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySftpBackupStorage queries SftpBackupStorage list
func (cli *ZSClient) QuerySftpBackupStorage(params *param.QueryParam) ([]view.SftpBackupStorageInventoryView, error) {
	var resp []view.SftpBackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/sftp", params, &resp)
}

func (cli *ZSClient) GetSftpBackupStorage(uuid string) (*view.SftpBackupStorageInventoryView, error) {
	var resp view.SftpBackupStorageInventoryView
	if err := cli.Get("v1/backup-storage/sftp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSftpBackupStorage Pagination
func (cli *ZSClient) PageSftpBackupStorage(params *param.QueryParam) ([]view.SftpBackupStorageInventoryView, int, error) {
	var sftpBackupStorages []view.SftpBackupStorageInventoryView
	total, err := cli.Page("v1/backup-storage/sftp", params, &sftpBackupStorages)
	return sftpBackupStorages, total, err
}
// UpdateSftpBackupStorage updates SftpBackupStorage
func (cli *ZSClient) UpdateSftpBackupStorage(uuid string, params param.UpdateSftpBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Put("v1/backup-storage/sftp", uuid, map[string]interface{}{
		"updateSftpBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ReconnectSftpBackupStorage operates on SftpBackupStorage
func (cli *ZSClient) ReconnectSftpBackupStorage(uuid string, params param.ReconnectSftpBackupStorageParam) (*view.SftpBackupStorageInventoryView, error) {
	resp := view.SftpBackupStorageInventoryView{}
	if err := cli.Put("v1/backup-storage/sftp", uuid, map[string]interface{}{
		"reconnectSftpBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddSftpBackupStorage adds SftpBackupStorage
func (cli *ZSClient) AddSftpBackupStorage(params param.AddSftpBackupStorageParam) (*view.SftpBackupStorageInventoryView, error) {
	resp := view.SftpBackupStorageInventoryView{}
	if err := cli.Post("v1/backup-storage/sftp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
