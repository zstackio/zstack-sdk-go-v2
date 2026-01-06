// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySftpBackupStorage queries SftpBackupStorage list
func (cli *ZSClient) QuerySftpBackupStorage(params *param.QueryParam) ([]view.SftpBackupStorageInventoryView, error) {
	var resp []view.SftpBackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/sftp", params, &resp)
}
// UpdateSftpBackupStorage updates SftpBackupStorage
func (cli *ZSClient) UpdateSftpBackupStorage(uuid string, params param.UpdateSftpBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.UpdateBackupStorageEventView
	if err := cli.Put("v1/backup-storage/sftp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ReconnectSftpBackupStorage operates on SftpBackupStorage
func (cli *ZSClient) ReconnectSftpBackupStorage(uuid string, params param.ReconnectSftpBackupStorageParam) (*view.SftpBackupStorageInventoryView, error) {
	var resp view.ReconnectSftpBackupStorageEventView
	if err := cli.Put("v1/backup-storage/sftp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddSftpBackupStorage adds SftpBackupStorage
func (cli *ZSClient) AddSftpBackupStorage(params param.AddSftpBackupStorageParam) (*view.SftpBackupStorageInventoryView, error) {
	resp := view.SftpBackupStorageInventoryView{}
	if err := cli.Post("v1/backup-storage/sftp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
