// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryCephBackupStorage queries CephBackupStorage list
func (cli *ZSClient) QueryCephBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/ceph", params, &resp)
}

func (cli *ZSClient) GetCephBackupStorage(uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get("v1/backup-storage/ceph", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddCephBackupStorage adds CephBackupStorage
func (cli *ZSClient) AddCephBackupStorage(params param.AddCephBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.AddBackupStorageEventView
	if err := cli.Post("v1/backup-storage/ceph", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
