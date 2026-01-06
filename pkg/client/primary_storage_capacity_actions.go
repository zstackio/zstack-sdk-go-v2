// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// SyncPrimaryStorageCapacity operates on PrimaryStorageCapacity
func (cli *ZSClient) SyncPrimaryStorageCapacity(uuid string, params param.SyncPrimaryStorageCapacityParam) (*view.PrimaryStorageInventoryView, error) {
	var resp view.SyncPrimaryStorageCapacityEventView
	if err := cli.Put("v1/primary-storage/{primaryStorageUuid}/actions", uuid, params, &resp); err != nil {
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
