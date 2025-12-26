// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetHostAllocatorStrategies gets HostAllocatorStrategies by uuid
func (cli *ZSClient) GetHostAllocatorStrategies(uuid string) (*view.GetHostAllocatorStrategiesView, error) {
	var resp view.GetHostAllocatorStrategiesView
	if err := cli.Get("v1/hosts/allocators/strategies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
