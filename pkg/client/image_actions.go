// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddImage adds Image
func (cli *ZSClient) AddImage(params param.AddImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post("v1/images", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddImageAsync Async
func (cli *ZSClient) AddImageAsync(params param.AddImageParam) (string, error) {

	resource := "v1/images"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// QueryImage queries Image list
func (cli *ZSClient) QueryImage(params *param.QueryParam) ([]view.ImageInventoryView, error) {
	var resp []view.ImageInventoryView
	return resp, cli.List("v1/images", params, &resp)
}

func (cli *ZSClient) GetImage(uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.Get("v1/images", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImage Pagination
func (cli *ZSClient) PageImage(params *param.QueryParam) ([]view.ImageInventoryView, int, error) {
	var images []view.ImageInventoryView
	total, err := cli.Page("v1/images", params, &images)
	return images, total, err
}
// SyncImage operates on Image
func (cli *ZSClient) SyncImage(imageStoreUuid string, params param.SyncImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/backup-storage/image-store", imageStoreUuid, "", map[string]interface{}{
		"syncImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RecoverImage operates on Image
func (cli *ZSClient) RecoverImage(imageUuid string, params param.RecoverImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/images", imageUuid, "", map[string]interface{}{
		"recoverImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CloneImage operates on Image
func (cli *ZSClient) CloneImage(imageUuid string, params param.CloneImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/image/clone/%s", imageUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteImage sends the action with its default options.
func (cli *ZSClient) DeleteImage(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteImageWithParams(uuid, deleteMode, param.DeleteImageParamDetail{})
}

// DeleteImageWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteImageWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteImageParamDetail) error {
	query := url.Values{
		"backupStorageUuids": params.BackupStorageUuids,
	}
	return cli.deleteResourceWithQuery("v1/images", uuid, deleteModeOrDefault(deleteMode, params.DeleteMode), query)
}
// UpdateImage updates Image
func (cli *ZSClient) UpdateImage(uuid string, params param.UpdateImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/images", uuid, "", map[string]interface{}{
		"updateImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ExpungeImage operates on Image
func (cli *ZSClient) ExpungeImage(uuid string) error {
	params := map[string]interface{}{
		"expungeImage": map[string]interface{}{},
	}
	return cli.Put("v1/images", uuid, params, nil)
}

// GetImageQga gets ImageQga by uuid
func (cli *ZSClient) GetImageQga(uuid string) (*view.GetImageQgaView, error) {
	var resp view.GetImageQgaView
	if err := cli.GetWithRespKey("v1/images", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromVolumeBackup creates RootVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeBackup(backupUuid string, params param.CreateRootVolumeTemplateFromVolumeBackupParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/images/root-volume-templates/from/volume-template/%s", backupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromVolumeSnapshot creates RootVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateRootVolumeTemplateFromVolumeSnapshot(snapshotUuid string, params param.CreateRootVolumeTemplateFromVolumeSnapshotParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/images/root-volume-templates/from/volume-snapshots/%s", snapshotUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoveryImageFromImageStoreBackupStorage operates on yImageFromImageStoreBackupStorage
func (cli *ZSClient) RecoveryImageFromImageStoreBackupStorage(uuid string, params param.RecoveryImageFromImageStoreBackupStorageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/backup-storage", uuid, "", map[string]interface{}{
		"recoveryImageFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageQga operates on ImageQga
func (cli *ZSClient) SetImageQga(uuid string, params param.SetImageQgaParam) (*view.SetImageQgaEventView, error) {
	resp := view.SetImageQgaEventView{}
	if err := cli.PutWithRespKey("v1/images", uuid, "", map[string]interface{}{
		"setImageQga": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncImageFromImageStoreBackupStorage operates on ImageFromImageStoreBackupStorage
func (cli *ZSClient) SyncImageFromImageStoreBackupStorage(uuid string, params param.SyncImageFromImageStoreBackupStorageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/images", uuid, "", map[string]interface{}{
		"syncImageFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageBootMode operates on ImageBootMode
func (cli *ZSClient) SetImageBootMode(uuid string, params param.SetImageBootModeParam) (*view.SetImageBootModeEventView, error) {
	resp := view.SetImageBootModeEventView{}
	if err := cli.PutWithRespKey("v1/images", uuid, "", map[string]interface{}{
		"setImageBootMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolumeBackup creates DataVolumeTemplateFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeBackup(backupUuid string, params param.CreateDataVolumeTemplateFromVolumeBackupParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/images/data-volume-templates/from/volume-template/%s", backupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncImageSize operates on ImageSize
func (cli *ZSClient) SyncImageSize(uuid string, params param.SyncImageSizeParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/images", uuid, "", map[string]interface{}{
		"syncImageSize": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUploadImageJobDetails gets UploadImageJobDetails by uuid
func (cli *ZSClient) GetUploadImageJobDetails(uuid string) (*view.GetUploadImageJobDetailsView, error) {
	var resp view.GetUploadImageJobDetailsView
	if err := cli.GetWithRespKey("v1/images/upload-job/details", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageSecurityLevel operates on ImageSecurityLevel
func (cli *ZSClient) SetImageSecurityLevel(uuid string, params param.SetImageSecurityLevelParam) (*view.SetImageSecurityLevelEventView, error) {
	resp := view.SetImageSecurityLevelEventView{}
	if err := cli.PutWithRespKey("v1/images", uuid, "", map[string]interface{}{
		"setImageSecurityLevel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImageCandidatesForVmToChange gets ImageCandidatesForVmToChange by uuid
func (cli *ZSClient) GetImageCandidatesForVmToChange(uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeImageState changes ImageState
func (cli *ZSClient) ChangeImageState(uuid string, params param.ChangeImageStateParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/images", uuid, "", map[string]interface{}{
		"changeImageState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolume creates DataVolumeTemplateFromVolume
func (cli *ZSClient) CreateDataVolumeTemplateFromVolume(volumeUuid string, params param.CreateDataVolumeTemplateFromVolumeParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/images/data-volume-templates/from/volumes/%s", volumeUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolumeAsync Async
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeAsync(volumeUuid string, params param.CreateDataVolumeTemplateFromVolumeParam) (string, error) {

	resource := fmt.Sprintf("v1/images/data-volume-templates/from/volumes/%s", volumeUuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// CalculateImageHash operates on ImageHash
func (cli *ZSClient) CalculateImageHash(uuid string, params param.CalculateImageHashParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/images", uuid, "", map[string]interface{}{
		"calculateImageHash": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeTemplateFromVolumeSnapshot creates DataVolumeTemplateFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeTemplateFromVolumeSnapshot(snapshotUuid string, params param.CreateDataVolumeTemplateFromVolumeSnapshotParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/images/data-volume-templates/from/volume-snapshots/%s", snapshotUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromRootVolume creates RootVolumeTemplateFromRootVolume
func (cli *ZSClient) CreateRootVolumeTemplateFromRootVolume(rootVolumeUuid string, params param.CreateRootVolumeTemplateFromRootVolumeParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/images/root-volume-templates/from/volumes/%s", rootVolumeUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRootVolumeTemplateFromRootVolumeAsync Async
func (cli *ZSClient) CreateRootVolumeTemplateFromRootVolumeAsync(rootVolumeUuid string, params param.CreateRootVolumeTemplateFromRootVolumeParam) (string, error) {

	resource := fmt.Sprintf("v1/images/root-volume-templates/from/volumes/%s", rootVolumeUuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
