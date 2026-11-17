// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryLocalStorageResourceRef queries LocalStorageResourceRef list
func (cli *ZSClient) QueryLocalStorageResourceRef(params *param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, error) {
	var resp []view.LocalStorageResourceRefInventoryView
	return resp, cli.List("v1/primary-storage/local-storage/resource-refs", params, &resp)
}

func (cli *ZSClient) GetLocalStorageResourceRef(uuid string) (*view.LocalStorageResourceRefInventoryView, error) {
	var resp view.LocalStorageResourceRefInventoryView
	if err := cli.Get("v1/primary-storage/local-storage/resource-refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLocalStorageResourceRef Pagination
func (cli *ZSClient) PageLocalStorageResourceRef(params *param.QueryParam) ([]view.LocalStorageResourceRefInventoryView, int, error) {
	var localStorageResourceRefs []view.LocalStorageResourceRefInventoryView
	total, err := cli.Page("v1/primary-storage/local-storage/resource-refs", params, &localStorageResourceRefs)
	return localStorageResourceRefs, total, err
}

// GetLocalStorageHostDiskCapacity gets LocalStorageHostDiskCapacity by uuid
func (cli *ZSClient) GetLocalStorageHostDiskCapacity(uuid string) (*view.HostDiskCapacityView, error) {
	var resp view.HostDiskCapacityView
	if err := cli.GetWithRespKey("v1/primary-storage/local-storage", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LocalStorageMigrateVolume operates on LocalStorageMigrateVolume
func (cli *ZSClient) LocalStorageMigrateVolume(volumeUuid string, params param.LocalStorageMigrateVolumeParam) (*view.LocalStorageResourceRefInventoryView, error) {
	resp := view.LocalStorageResourceRefInventoryView{}
	if err := cli.PutWithRespKey("v1/primary-storage/local-storage/volumes", volumeUuid, "", map[string]interface{}{
		"localStorageMigrateVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// LocalStorageGetVolumeMigratableHosts operates on LocalStorageGetVolumeMigratableHosts
func (cli *ZSClient) LocalStorageGetVolumeMigratableHosts(uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.GetWithRespKey("v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
