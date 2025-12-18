// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetPrimaryStorageAllocatorStrategies 获取PrimaryStorageAllocatorStrategies详情
func (cli *ZSClient) GetPrimaryStorageAllocatorStrategies(uuid string) (*view.GetPrimaryStorageAllocatorStrategiesView, error) {
	var resp view.GetPrimaryStorageAllocatorStrategiesView
	if err := cli.Get("v1/primary-storage/allocators/strategies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

