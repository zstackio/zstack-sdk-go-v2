// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetBareMetal2GatewayAllocatorStrategies 获取BareMetal2GatewayAllocatorStrategies详情
func (cli *ZSClient) GetBareMetal2GatewayAllocatorStrategies(uuid string) (*view.GetBareMetal2GatewayAllocatorStrategiesView, error) {
	var resp view.GetBareMetal2GatewayAllocatorStrategiesView
	if err := cli.Get("v1/baremetal2/gateways/allocators/strategies", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

