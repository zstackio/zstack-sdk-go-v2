// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddImageStoreBackupStorage adds ImageStoreBackupStorage
func (cli *ZSClient) AddImageStoreBackupStorage(params param.AddImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.Post("v1/backup-storage/image-store", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateImageStoreBackupStorage updates ImageStoreBackupStorage
func (cli *ZSClient) UpdateImageStoreBackupStorage(uuid string, params param.UpdateImageStoreBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/backup-storage/image-store", uuid, "", map[string]interface{}{
		"updateImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryImageStoreBackupStorage queries ImageStoreBackupStorage list
func (cli *ZSClient) QueryImageStoreBackupStorage(params *param.QueryParam) ([]view.ImageStoreBackupStorageInventoryView, error) {
	var resp []view.ImageStoreBackupStorageInventoryView
	return resp, cli.List("v1/backup-storage/image-store", params, &resp)
}

func (cli *ZSClient) GetImageStoreBackupStorage(uuid string) (*view.ImageStoreBackupStorageInventoryView, error) {
	var resp view.ImageStoreBackupStorageInventoryView
	if err := cli.Get("v1/backup-storage/image-store", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImageStoreBackupStorage Pagination
func (cli *ZSClient) PageImageStoreBackupStorage(params *param.QueryParam) ([]view.ImageStoreBackupStorageInventoryView, int, error) {
	var imageStoreBackupStorages []view.ImageStoreBackupStorageInventoryView
	total, err := cli.Page("v1/backup-storage/image-store", params, &imageStoreBackupStorages)
	return imageStoreBackupStorages, total, err
}
// ReconnectImageStoreBackupStorage operates on ImageStoreBackupStorage
func (cli *ZSClient) ReconnectImageStoreBackupStorage(uuid string, params param.ReconnectImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.PutWithRespKey("v1/backup-storage/image-store", uuid, "", map[string]interface{}{
		"reconnectImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetImageStoreBackupStorageQuota operates on ImageStoreBackupStorageQuota
func (cli *ZSClient) SetImageStoreBackupStorageQuota(params param.SetImageStoreBackupStorageQuotaParam) (*view.SetImageStoreBackupStorageQuotaEventView, error) {
	resp := view.SetImageStoreBackupStorageQuotaEventView{}
	if err := cli.PutWithRespKey("v1/backup-storage/image-store/actions", "", "", map[string]interface{}{
		"setImageStoreBackupStorageQuota": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImagesFromImageStoreBackupStorage gets ImagesFromImageStoreBackupStorage by uuid
func (cli *ZSClient) GetImagesFromImageStoreBackupStorage(uuid string, params param.GetImagesFromImageStoreBackupStorageParam) (*view.GetImagesFromImageStoreBackupStorageView, error) {
	resp := view.GetImagesFromImageStoreBackupStorageView{}
	if err := cli.PutWithRespKey("v1/backup-storage", uuid, "", map[string]interface{}{
		"getImagesFromImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReclaimSpaceFromImageStore operates on ReclaimSpaceFromImageStore
func (cli *ZSClient) ReclaimSpaceFromImageStore(uuid string, params param.ReclaimSpaceFromImageStoreParam) (*view.ReclaimSpaceFromImageStoreEventView, error) {
	resp := view.ReclaimSpaceFromImageStoreEventView{}
	if err := cli.PutWithRespKey("v1/backup-storage/image-store", uuid, "", map[string]interface{}{
		"reclaimSpaceFromImageStore": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReclaimSpaceFromImageStoreAsync Async
func (cli *ZSClient) ReclaimSpaceFromImageStoreAsync(uuid string, params param.ReclaimSpaceFromImageStoreParam) (string, error) {

	resource := fmt.Sprintf("v1/backup-storage/image-store/%s/actions", uuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// AddDisasterImageStoreBackupStorage adds DisasterImageStoreBackupStorage
func (cli *ZSClient) AddDisasterImageStoreBackupStorage(params param.AddDisasterImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.Post("v1/backup-storage/image-store/disaster", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
