// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateVolumeBackup creates VolumeBackup
func (cli *ZSClient) CreateVolumeBackup(params param.CreateVolumeBackupParam) (*view.VolumeBackupInventoryView, error) {
	var resp view.CreateVolumeBackupEventView
	if err := cli.Post("v1/volumes/{volumeUuid}/volume-backups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// SyncVolumeBackup operates on VolumeBackup
func (cli *ZSClient) SyncVolumeBackup(uuid string, params param.SyncVolumeBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Put("v1/volume-backups/imageStore/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVolumeBackup queries VolumeBackup list
func (cli *ZSClient) QueryVolumeBackup(params *param.QueryParam) ([]view.VolumeBackupInventoryView, error) {
	var resp []view.VolumeBackupInventoryView
	return resp, cli.List("v1/volume-backups", params, &resp)
}
// DeleteVolumeBackup deletes VolumeBackup
func (cli *ZSClient) DeleteVolumeBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-backups/{uuid}", uuid, string(deleteMode))
}
