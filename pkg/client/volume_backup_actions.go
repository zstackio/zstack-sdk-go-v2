// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateVolumeBackup creates VolumeBackup
func (cli *ZSClient) CreateVolumeBackup(volumeUuid string, params param.CreateVolumeBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/volumes/%s/volume-backups", volumeUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolumeBackupAsync Async
func (cli *ZSClient) CreateVolumeBackupAsync(volumeUuid string, params param.CreateVolumeBackupParam) (string, error) {

	resource := fmt.Sprintf("v1/volumes/%s/volume-backups", volumeUuid)
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
	if err := cli.PutWithRespKey("v1/volume-backups/imageStore", imageStoreUuid, "", map[string]interface{}{
		"syncVolumeBackup": params.Params,
	}, &resp); err != nil {
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
// DeleteVolumeBackup sends the action with its default options.
func (cli *ZSClient) DeleteVolumeBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteVolumeBackupWithParams(uuid, deleteMode, param.DeleteVolumeBackupParamDetail{})
}

// DeleteVolumeBackupWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteVolumeBackupWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteVolumeBackupParamDetail) error {
	query := url.Values{
		"backupStorageUuids": params.BackupStorageUuids,
	}
	if params.HandleDependency != nil {
		query.Set("handleDependency", strconv.FormatBool(*params.HandleDependency))
	}
	return cli.deleteResourceWithQuery("v1/volume-backups", uuid, deleteModeOrDefault(deleteMode, params.DeleteMode), query)
}

// RecoverBackupFromImageStoreBackupStorage operates on BackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverBackupFromImageStoreBackupStorage(uuid string, params param.RecoverBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithRespKey("v1/volume-backups", uuid, "", map[string]interface{}{
		"recoverBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncBackupFromImageStoreBackupStorage operates on BackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncBackupFromImageStoreBackupStorage(uuid string, params param.SyncBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithRespKey("v1/volume-backups", uuid, "", map[string]interface{}{
		"syncBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVolumeFromVolumeBackup operates on VolumeFromVolumeBackup
func (cli *ZSClient) RevertVolumeFromVolumeBackup(uuid string, params param.RevertVolumeFromVolumeBackupParam) (*view.RevertVolumeFromVolumeBackupEventView, error) {
	resp := view.RevertVolumeFromVolumeBackupEventView{}
	if err := cli.PutWithRespKey("v1/volume-backups", uuid, "", map[string]interface{}{
		"revertVolumeFromVolumeBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
