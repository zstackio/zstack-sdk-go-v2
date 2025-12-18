// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateL2HardwareVxlanNetwork 创建L2HardwareVxlanNetwork
func (cli *ZSClient) CreateL2HardwareVxlanNetwork(params param.CreateL2HardwareVxlanNetworkParam) (*view.CreateL2HardwareVxlanNetworkEventView, error) {
	resp := view.CreateL2HardwareVxlanNetworkEventView{}
	if err := cli.Post("v1/l2-networks/hardware-vxlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

