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
	var resp view.SyncPrimaryStorageCapacityEventView
	err := cli.PutWithSpec("v1/primary-storage", fmt.Sprintf(\"%s/actions\", primaryStorageUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// GetPrimaryStorageCapacity gets PrimaryStorageCapacity by uuid
func (cli *ZSClient) GetPrimaryStorageCapacity(uuid string) (*view.PrimaryStorageCapacityInventoryView, error) {
	var resp view.PrimaryStorageCapacityInventoryView
	if err := cli.Get("v1/primary-storage/capacities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
