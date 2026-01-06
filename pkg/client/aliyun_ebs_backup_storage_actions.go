// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAliyunEbsBackupStorage queries AliyunEbsBackupStorage list
func (cli *ZSClient) QueryAliyunEbsBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/aliyun/ebs", params, &resp)
}
// AddAliyunEbsBackupStorage adds AliyunEbsBackupStorage
func (cli *ZSClient) AddAliyunEbsBackupStorage(params param.AddAliyunEbsBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.AddBackupStorageEventView
	if err := cli.Post("v1/backup-storage/aliyun/ebs", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateAliyunEbsBackupStorage updates AliyunEbsBackupStorage
func (cli *ZSClient) UpdateAliyunEbsBackupStorage(uuid string, params param.UpdateAliyunEbsBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.UpdateBackupStorageEventView
	if err := cli.Put("v1/backup-storage/aliyun/ebs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
