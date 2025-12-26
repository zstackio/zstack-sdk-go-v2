// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCpuMemoryCapacity gets CpuMemoryCapacity by uuid
func (cli *ZSClient) GetCpuMemoryCapacity(uuid string) (*view.GetCpuMemoryCapacityView, error) {
	var resp view.GetCpuMemoryCapacityView
	if err := cli.Get("v1/hosts/capacities/cpu-memory", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
