// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBackupStorage queries BackupStorage list
func (cli *ZSClient) QueryBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List("v1/backup-storage", params, &resp)
}

func (cli *ZSClient) GetBackupStorage(uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get("v1/backup-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBackupStorage Pagination
func (cli *ZSClient) PageBackupStorage(params *param.QueryParam) ([]view.BackupStorageInventoryView, int, error) {
	var backupStorages []view.BackupStorageInventoryView
	total, err := cli.Page("v1/backup-storage", params, &backupStorages)
	return backupStorages, total, err
}
// UpdateBackupStorage updates BackupStorage
func (cli *ZSClient) UpdateBackupStorage(uuid string, params param.UpdateBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/backup-storage", uuid, "", map[string]interface{}{
		"updateBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBackupStorage deletes BackupStorage
func (cli *ZSClient) DeleteBackupStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/backup-storage", uuid, string(deleteMode))
}
// ReconnectBackupStorage operates on BackupStorage
func (cli *ZSClient) ReconnectBackupStorage(uuid string, params param.ReconnectBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/backup-storage", uuid, "", map[string]interface{}{
		"reconnectBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeBackupStorageState changes BackupStorageState
func (cli *ZSClient) ChangeBackupStorageState(uuid string, params param.ChangeBackupStorageStateParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/backup-storage", uuid, "", map[string]interface{}{
		"changeBackupStorageState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateBackupStorageForCreatingImage gets CandidateBackupStorageForCreatingImage by uuid
func (cli *ZSClient) GetCandidateBackupStorageForCreatingImage() (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.GetWithRespKey("v1/images/candidate-backup-storage", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTrashOnBackupStorage gets TrashOnBackupStorage by uuid
func (cli *ZSClient) GetTrashOnBackupStorage() (*view.InstallPathRecycleInventoryView, error) {
	var resp view.InstallPathRecycleInventoryView
	if err := cli.GetWithRespKey("v1/backup-storage/trash", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportImageFromBackupStorage operates on ImageFromBackupStorage
func (cli *ZSClient) ExportImageFromBackupStorage(backupStorageUuid string, params param.ExportImageFromBackupStorageParam) (*view.ExportImageFromBackupStorageEventView, error) {
	resp := view.ExportImageFromBackupStorageEventView{}
	if err := cli.PutWithRespKey("v1/backup-storage", backupStorageUuid, "", map[string]interface{}{
		"exportImageFromBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportImageFromBackupStorageAsync Async
func (cli *ZSClient) ExportImageFromBackupStorageAsync(backupStorageUuid string, params param.ExportImageFromBackupStorageParam) (string, error) {

	resource := fmt.Sprintf("v1/backup-storage/%s/actions", backupStorageUuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetBackupStorageCapacity gets BackupStorageCapacity by uuid
func (cli *ZSClient) GetBackupStorageCapacity() (*view.GetBackupStorageCapacityView, error) {
	var resp view.GetBackupStorageCapacityView
	if err := cli.GetWithRespKey("v1/backup-storage/capacities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BackupStorageMigrateImage operates on StorageMigrateImage
func (cli *ZSClient) BackupStorageMigrateImage(imageUuid string, params param.BackupStorageMigrateImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/backup-storage/images", imageUuid, "", map[string]interface{}{
		"backupStorageMigrateImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BackupStorageMigrateImageAsync Async
func (cli *ZSClient) BackupStorageMigrateImageAsync(imageUuid string, params param.BackupStorageMigrateImageParam) (string, error) {

	resource := fmt.Sprintf("v1/backup-storage/images/%s/actions", imageUuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// AttachBackupStorageToZone operates on BackupStorageToZone
func (cli *ZSClient) AttachBackupStorageToZone(zoneUuid string, backupStorageUuid string, params param.AttachBackupStorageToZoneParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/zones/%s/backup-storage/%s", zoneUuid, backupStorageUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpTrashOnBackupStorage operates on UpTrashOnBackupStorage
func (cli *ZSClient) CleanUpTrashOnBackupStorage(uuid string, params param.CleanUpTrashOnBackupStorageParam) (*view.CleanUpTrashOnBackupStorageEventView, error) {
	resp := view.CleanUpTrashOnBackupStorageEventView{}
	if err := cli.PutWithRespKey("v1/backup-storage", uuid, "", map[string]interface{}{
		"cleanUpTrashOnBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBackupStorageTypes gets BackupStorageTypes by uuid
func (cli *ZSClient) GetBackupStorageTypes() (*view.GetBackupStorageTypesView, error) {
	var resp view.GetBackupStorageTypesView
	if err := cli.GetWithRespKey("v1/backup-storage/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteExportedImageFromBackupStorage deletes ExportedImageFromBackupStorage
func (cli *ZSClient) DeleteExportedImageFromBackupStorage(backupStorageUuid string, imageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/backup-storage", backupStorageUuid, fmt.Sprintf("exported-images/%s", imageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// DetachBackupStorageFromZone operates on BackupStorageFromZone
func (cli *ZSClient) DetachBackupStorageFromZone(zoneUuid string, backupStorageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/zones", zoneUuid, fmt.Sprintf("backup-storage/%s", backupStorageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddSimulatorBackupStorage adds SimulatorBackupStorage
func (cli *ZSClient) AddSimulatorBackupStorage(params param.AddSimulatorBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.Post("v1/backup-storage/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MergeDataOnBackupStorage operates on MergeDataOnBackupStorage
func (cli *ZSClient) MergeDataOnBackupStorage(backupStorageUuid string, params param.MergeDataOnBackupStorageParam) (*view.MergeDataOnBackupStorageEventView, error) {
	resp := view.MergeDataOnBackupStorageEventView{}
	if err := cli.PutWithRespKey("v1/cdp-task/mergedata", backupStorageUuid, "", map[string]interface{}{
		"mergeDataOnBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCdpBackupStorageRequirement gets CdpBackupStorageRequirement by uuid
func (cli *ZSClient) GetCdpBackupStorageRequirement(uuid string) (*view.GetCdpBackupStorageRequirementView, error) {
	var resp view.GetCdpBackupStorageRequirementView
	if err := cli.GetWithRespKey("v1/cdp-backup-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBackupStorageCandidatesForImageMigration gets BackupStorageCandidatesForImageMigration by uuid
func (cli *ZSClient) GetBackupStorageCandidatesForImageMigration(uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.GetWithRespKey("v1/backup-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
