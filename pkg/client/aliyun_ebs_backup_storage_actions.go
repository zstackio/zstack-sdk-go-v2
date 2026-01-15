// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAliyunEbsBackupStorage queries AliyunEbsBackupStorage list
func (cli *ZSClient) QueryAliyunEbsBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/aliyun/ebs", params, &resp)
}

func (cli *ZSClient) GetAliyunEbsBackupStorage(uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get("v1/backup-storage/aliyun/ebs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunEbsBackupStorage Pagination
func (cli *ZSClient) PageAliyunEbsBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, int, error) {
	var aliyunEbsBackupStorages []view.BackupStorageInventoryView
	total, err := cli.Page("v1/backup-storage/aliyun/ebs", params, &aliyunEbsBackupStorages)
	return aliyunEbsBackupStorages, total, err
}
// AddAliyunEbsBackupStorage adds AliyunEbsBackupStorage
func (cli *ZSClient) AddAliyunEbsBackupStorage(params param.AddAliyunEbsBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Post("v1/backup-storage/aliyun/ebs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAliyunEbsBackupStorage updates AliyunEbsBackupStorage
func (cli *ZSClient) UpdateAliyunEbsBackupStorage(uuid string, params param.UpdateAliyunEbsBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Put("v1/backup-storage/aliyun/ebs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
