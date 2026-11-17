// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// CreateVmBackup creates VmBackup
func (cli *ZSClient) CreateVmBackup(rootVolumeUuid string, params param.CreateVmBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/volumes/%s/vm-backups", rootVolumeUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmBackupAsync Async
func (cli *ZSClient) CreateVmBackupAsync(rootVolumeUuid string, params param.CreateVmBackupParam) (string, error) {

	resource := fmt.Sprintf("v1/volumes/%s/vm-backups", rootVolumeUuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// RecoverVmBackupFromImageStoreBackupStorage operates on VmBackupFromImageStoreBackupStorage
func (cli *ZSClient) RecoverVmBackupFromImageStoreBackupStorage(groupUuid string, params param.RecoverVmBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-backups", groupUuid, "", map[string]interface{}{
		"recoverVmBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmBackup deletes VmBackup
func (cli *ZSClient) DeleteVmBackup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-backups", uuid, string(deleteMode))
}

// SyncVmBackup operates on VmBackup
func (cli *ZSClient) SyncVmBackup(imageStoreUuid string, params param.SyncVmBackupParam) (*view.SyncVmBackupEventView, error) {
	resp := view.SyncVmBackupEventView{}
	if err := cli.PutWithRespKey("v1/vm-backups/imageStore", imageStoreUuid, "", map[string]interface{}{
		"syncVmBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVmBackupFromImageStoreBackupStorage operates on VmBackupFromImageStoreBackupStorage
func (cli *ZSClient) SyncVmBackupFromImageStoreBackupStorage(groupUuid string, params param.SyncVmBackupFromImageStoreBackupStorageParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-backups", groupUuid, "", map[string]interface{}{
		"syncVmBackupFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVmFromVmBackup operates on VmFromVmBackup
func (cli *ZSClient) RevertVmFromVmBackup(groupUuid string, params param.RevertVmFromVmBackupParam) (*view.RevertVmFromVmBackupEventView, error) {
	resp := view.RevertVmFromVmBackupEventView{}
	if err := cli.PutWithRespKey("v1/vm-backups", groupUuid, "", map[string]interface{}{
		"revertVmFromVmBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
