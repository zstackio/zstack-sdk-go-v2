// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryCephBackupStorage queries CephBackupStorage list
func (cli *ZSClient) QueryCephBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/ceph", params, &resp)
}
// AddCephBackupStorage adds CephBackupStorage
func (cli *ZSClient) AddCephBackupStorage(params param.AddCephBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	var resp view.AddBackupStorageEventView
	if err := cli.Post("v1/backup-storage/ceph", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
