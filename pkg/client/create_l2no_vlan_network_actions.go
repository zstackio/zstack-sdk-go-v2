// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateL2NoVlanNetwork creates L2NoVlanNetwork
func (cli *ZSClient) CreateL2NoVlanNetwork(params param.CreateL2NoVlanNetworkParam) (*view.CreateL2NetworkEventView, error) {
	resp := view.CreateL2NetworkEventView{}
	if err := cli.Post("v1/l2-networks/no-vlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
