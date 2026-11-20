// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryPrimaryStorage queries PrimaryStorage list
func (cli *ZSClient) QueryPrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List("v1/primary-storage", params, &resp)
}

func (cli *ZSClient) GetPrimaryStorage(uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.Get("v1/primary-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePrimaryStorage Pagination
func (cli *ZSClient) PagePrimaryStorage(params *param.QueryParam) ([]view.PrimaryStorageInventoryView, int, error) {
	var primaryStorages []view.PrimaryStorageInventoryView
	total, err := cli.Page("v1/primary-storage", params, &primaryStorages)
	return primaryStorages, total, err
}
// ReconnectPrimaryStorage operates on PrimaryStorage
func (cli *ZSClient) ReconnectPrimaryStorage(uuid string, params param.ReconnectPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/primary-storage", uuid, "", map[string]interface{}{
		"reconnectPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePrimaryStorage updates PrimaryStorage
func (cli *ZSClient) UpdatePrimaryStorage(uuid string, params param.UpdatePrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/primary-storage", uuid, "", map[string]interface{}{
		"updatePrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePrimaryStorage deletes PrimaryStorage
func (cli *ZSClient) DeletePrimaryStorage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/primary-storage", uuid, string(deleteMode))
}

// AttachPrimaryStorageToCluster operates on PrimaryStorageToCluster
func (cli *ZSClient) AttachPrimaryStorageToCluster(clusterUuid string, primaryStorageUuid string, params param.AttachPrimaryStorageToClusterParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/clusters/%s/primary-storage/%s", clusterUuid, primaryStorageUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageCandidatesForVmMigration gets PrimaryStorageCandidatesForVmMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVmMigration(uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVolume operates on PrimaryStorageMigrateVolume
func (cli *ZSClient) PrimaryStorageMigrateVolume(volumeUuid string, params param.PrimaryStorageMigrateVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/primary-storage/volumes", volumeUuid, "", map[string]interface{}{
		"primaryStorageMigrateVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVolumeAsync Async
func (cli *ZSClient) PrimaryStorageMigrateVolumeAsync(volumeUuid string, params param.PrimaryStorageMigrateVolumeParam) (string, error) {

	resource := fmt.Sprintf("v1/primary-storage/volumes/%s/actions", volumeUuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetPrimaryStorageTypes gets PrimaryStorageTypes by uuid
func (cli *ZSClient) GetPrimaryStorageTypes() (*view.GetPrimaryStorageTypesView, error) {
	var resp view.GetPrimaryStorageTypesView
	if err := cli.GetWithRespKey("v1/primary-storage/types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageCandidatesForVolumeMigration gets PrimaryStorageCandidatesForVolumeMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVolumeMigration(uuid string) (*view.PrimaryStorageInventoryView, error) {
	var resp view.PrimaryStorageInventoryView
	if err := cli.GetWithRespKey("v1/primary-storage/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageLicenseInfo gets PrimaryStorageLicenseInfo by uuid
func (cli *ZSClient) GetPrimaryStorageLicenseInfo(uuid string) (*view.GetPrimaryStorageLicenseInfoView, error) {
	var resp view.GetPrimaryStorageLicenseInfoView
	if err := cli.GetWithRespKey("v1/primary-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddNfsPrimaryStorage adds NfsPrimaryStorage
func (cli *ZSClient) AddNfsPrimaryStorage() (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post("v1/primary-storage/nfs", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangePrimaryStorageState changes PrimaryStorageState
func (cli *ZSClient) ChangePrimaryStorageState(uuid string, params param.ChangePrimaryStorageStateParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/primary-storage", uuid, "", map[string]interface{}{
		"changePrimaryStorageState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSimulatorPrimaryStorage adds SimulatorPrimaryStorage
func (cli *ZSClient) AddSimulatorPrimaryStorage(params param.AddSimulatorPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post("v1/primary-storage/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpStorageTrashOnPrimaryStorage operates on UpStorageTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpStorageTrashOnPrimaryStorage(uuid string, params param.CleanUpStorageTrashOnPrimaryStorageParam) (*view.CleanUpStorageTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpStorageTrashOnPrimaryStorageEventView{}
	if err := cli.PutWithRespKey("v1/primary-storage", uuid, "", map[string]interface{}{
		"cleanUpStorageTrashOnPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPrimaryStorageFromCluster operates on PrimaryStorageFromCluster
func (cli *ZSClient) DetachPrimaryStorageFromCluster(clusterUuid string, primaryStorageUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/clusters", clusterUuid, fmt.Sprintf("primary-storage/%s", primaryStorageUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddSharedMountPointPrimaryStorage adds SharedMountPointPrimaryStorage
func (cli *ZSClient) AddSharedMountPointPrimaryStorage() (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post("v1/primary-storage/smp", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTrashOnPrimaryStorage gets TrashOnPrimaryStorage by uuid
func (cli *ZSClient) GetTrashOnPrimaryStorage() (*view.InstallPathRecycleInventoryView, error) {
	var resp view.InstallPathRecycleInventoryView
	if err := cli.GetWithRespKey("v1/primary-storage/trash", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidatePrimaryStoragesForCreatingVm gets CandidatePrimaryStoragesForCreatingVm by uuid
func (cli *ZSClient) GetCandidatePrimaryStoragesForCreatingVm() (*view.GetCandidatePrimaryStoragesForCreatingVmView, error) {
	var resp view.GetCandidatePrimaryStoragesForCreatingVmView
	if err := cli.GetWithRespKey("v1/vm-instances/candidate-storages", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanUpTrashOnPrimaryStorage operates on UpTrashOnPrimaryStorage
func (cli *ZSClient) CleanUpTrashOnPrimaryStorage(uuid string, params param.CleanUpTrashOnPrimaryStorageParam) (*view.CleanUpTrashOnPrimaryStorageEventView, error) {
	resp := view.CleanUpTrashOnPrimaryStorageEventView{}
	if err := cli.PutWithRespKey("v1/primary-storage", uuid, "", map[string]interface{}{
		"cleanUpTrashOnPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddLocalPrimaryStorage adds LocalPrimaryStorage
func (cli *ZSClient) AddLocalPrimaryStorage() (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Post("v1/primary-storage/local-storage", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageAllocatorStrategies gets PrimaryStorageAllocatorStrategies by uuid
func (cli *ZSClient) GetPrimaryStorageAllocatorStrategies() (*view.GetPrimaryStorageAllocatorStrategiesView, error) {
	var resp view.GetPrimaryStorageAllocatorStrategiesView
	if err := cli.GetWithRespKey("v1/primary-storage/allocators/strategies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVm operates on PrimaryStorageMigrateVm
func (cli *ZSClient) PrimaryStorageMigrateVm(vmInstanceUuid string, params param.PrimaryStorageMigrateVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"primaryStorageMigrateVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PrimaryStorageMigrateVmAsync Async
func (cli *ZSClient) PrimaryStorageMigrateVmAsync(vmInstanceUuid string, params param.PrimaryStorageMigrateVmParam) (string, error) {

	resource := fmt.Sprintf("v1/vm-instances/%s/actions", vmInstanceUuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// CleanUpImageCacheOnPrimaryStorage operates on UpImageCacheOnPrimaryStorage
func (cli *ZSClient) CleanUpImageCacheOnPrimaryStorage(uuid string, params param.CleanUpImageCacheOnPrimaryStorageParam) (*view.CleanUpImageCacheOnPrimaryStorageEventView, error) {
	resp := view.CleanUpImageCacheOnPrimaryStorageEventView{}
	if err := cli.PutWithRespKey("v1/primary-storage", uuid, "", map[string]interface{}{
		"cleanUpImageCacheOnPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrimaryStorageUsageReport gets PrimaryStorageUsageReport by uuid
func (cli *ZSClient) GetPrimaryStorageUsageReport(uuid string) (*view.GetPrimaryStorageUsageReportView, error) {
	var resp view.GetPrimaryStorageUsageReportView
	if err := cli.GetWithRespKey("v1/primary-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
