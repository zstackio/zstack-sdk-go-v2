// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateL2HardwareVxlanNetwork creates L2HardwareVxlanNetwork
func (cli *ZSClient) CreateL2HardwareVxlanNetwork(params param.CreateL2HardwareVxlanNetworkParam) (*view.CreateL2HardwareVxlanNetworkEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkEventView{}
	if err := cli.Post("v1/l2-networks/hardware-vxlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
