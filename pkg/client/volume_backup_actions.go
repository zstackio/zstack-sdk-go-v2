// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateVolumeBackup creates VolumeBackup
func (cli *ZSClient) CreateVolumeBackup(params param.CreateVolumeBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/volume-backups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolumeBackupAsync Async
func (cli *ZSClient) CreateVolumeBackupAsync(params param.CreateVolumeBackupParam) (string, error) {

	resource := "v1/volumes/{volumeUuid}/volume-backups"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// SyncVolumeBackup operates on VolumeBackup
func (cli *ZSClient) SyncVolumeBackup(imageStoreUuid string, params param.SyncVolumeBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Put("v1/volume-backups/imageStore", imageStoreUuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVolumeBackup queries VolumeBackup list
func (cli *ZSClient) QueryVolumeBackup(params *param.QueryParam) ([]view.VolumeBackupInventoryView, error) {
	var resp []view.VolumeBackupInventoryView
	return resp, cli.List("v1/volume-backups", params, &resp)
}

func (cli *ZSClient) GetVolumeBackup(uuid string) (*view.VolumeBackupInventoryView, error) {
	var resp view.VolumeBackupInventoryView
	if err := cli.Get("v1/volume-backups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVolumeBackup Pagination
func (cli *ZSClient) PageVolumeBackup(params *param.QueryParam) ([]view.VolumeBackupInventoryView, int, error) {
	var volumeBackups []view.VolumeBackupInventoryView
	total, err := cli.Page("v1/volume-backups", params, &volumeBackups)
	return volumeBackups, total, err
}
// DeleteVolumeBackup deletes VolumeBackup
func (cli *ZSClient) DeleteVolumeBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-backups", uuid, string(deleteMode))
}
