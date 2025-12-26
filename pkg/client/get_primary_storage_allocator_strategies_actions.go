// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPrimaryStorageAllocatorStrategies gets PrimaryStorageAllocatorStrategies by uuid
func (cli *ZSClient) GetPrimaryStorageAllocatorStrategies(uuid string) (*view.GetPrimaryStorageAllocatorStrategiesView, error) {
	var resp view.GetPrimaryStorageAllocatorStrategiesView
	if err := cli.Get("v1/primary-storage/allocators/strategies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
