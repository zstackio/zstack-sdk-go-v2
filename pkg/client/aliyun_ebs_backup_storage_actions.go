// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAliyunEbsBackupStorage queries AliyunEbsBackupStorage list
func (cli *ZSClient) QueryAliyunEbsBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List(ctx, "v1/backup-storage/aliyun/ebs", params, &resp)
}

func (cli *ZSClient) GetAliyunEbsBackupStorage(ctx context.Context, uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get(ctx, "v1/backup-storage/aliyun/ebs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAliyunEbsBackupStorage Pagination
func (cli *ZSClient) PageAliyunEbsBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.BackupStorageInventoryView, int, error) {
	var aliyunEbsBackupStorages []view.BackupStorageInventoryView
	total, err := cli.Page(ctx, "v1/backup-storage/aliyun/ebs", params, &aliyunEbsBackupStorages)
	return aliyunEbsBackupStorages, total, err
}
// AddAliyunEbsBackupStorage adds AliyunEbsBackupStorage
func (cli *ZSClient) AddAliyunEbsBackupStorage(ctx context.Context, params param.AddAliyunEbsBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Post(ctx, "v1/backup-storage/aliyun/ebs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAliyunEbsBackupStorage updates AliyunEbsBackupStorage
func (cli *ZSClient) UpdateAliyunEbsBackupStorage(ctx context.Context, uuid string, params param.UpdateAliyunEbsBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/backup-storage/aliyun/ebs", uuid, "", map[string]interface{}{
		"updateAliyunEbsBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
