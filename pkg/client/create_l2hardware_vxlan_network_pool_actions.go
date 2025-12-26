// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateL2HardwareVxlanNetworkPool creates L2HardwareVxlanNetworkPool
func (cli *ZSClient) CreateL2HardwareVxlanNetworkPool(params param.CreateL2HardwareVxlanNetworkPoolParam) (*view.CreateL2HardwareVxlanNetworkPoolEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkPoolEventView{}
	if err := cli.Post("v1/l2-networks/hardware-vxlan-pool", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
