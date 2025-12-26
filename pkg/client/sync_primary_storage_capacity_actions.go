// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncPrimaryStorageCapacity operates on SyncPrimaryStorageCapacity
func (cli *ZSClient) SyncPrimaryStorageCapacity(uuid string, params param.SyncPrimaryStorageCapacityParam) (*view.SyncPrimaryStorageCapacityEventView, error) {
	resp := view.SyncPrimaryStorageCapacityEventView{}
	if err := cli.Put("v1/primary-storage/{primaryStorageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
