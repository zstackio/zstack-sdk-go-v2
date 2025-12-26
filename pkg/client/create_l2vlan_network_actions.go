// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateL2VlanNetwork creates L2VlanNetwork
func (cli *ZSClient) CreateL2VlanNetwork(params param.CreateL2VlanNetworkParam) (*view.CreateL2VlanNetworkEventView, error) {
	resp := view.CreateL2VlanNetworkEventView{}
	if err := cli.Post("v1/l2-networks/vlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
