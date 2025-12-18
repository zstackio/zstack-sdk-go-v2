// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateL2VlanNetwork creates L2VlanNetwork
func (cli *ZSClient) CreateL2VlanNetwork(params param.CreateL2VlanNetworkParam) (*view.CreateL2VlanNetworkEventView, error) {
	resp := view.CreateL2VlanNetworkEventView{}
	if err := cli.Post("v1/l2-networks/vlan", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
