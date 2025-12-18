// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCpuMemoryCapacity 获取CpuMemoryCapacity详情
func (cli *ZSClient) GetCpuMemoryCapacity(uuid string) (*view.GetCpuMemoryCapacityView, error) {
	var resp view.GetCpuMemoryCapacityView
	if err := cli.Get("v1/hosts/capacities/cpu-memory", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

