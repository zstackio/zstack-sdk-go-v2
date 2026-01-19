// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// SyncPrimaryStorageCapacity operates on PrimaryStorageCapacity
func (cli *ZSClient) SyncPrimaryStorageCapacity(primaryStorageUuid string, params param.SyncPrimaryStorageCapacityParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.Put("v1/primary-storage", primaryStorageUuid, map[string]interface{}{
		"syncPrimaryStorageCapacity": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// GetPrimaryStorageCapacity gets PrimaryStorageCapacity by uuid
func (cli *ZSClient) GetPrimaryStorageCapacity(uuid string) (*view.PrimaryStorageCapacityInventoryView, error) {
	var resp view.PrimaryStorageCapacityInventoryView
	if err := cli.Get("v1/primary-storage/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
