// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateBareMetal2ProvisionNetwork creates BareMetal2ProvisionNetwork
func (cli *ZSClient) CreateBareMetal2ProvisionNetwork(params param.CreateBareMetal2ProvisionNetworkParam) (*view.CreateBareMetal2ProvisionNetworkEventView, error) {
	resp := view.CreateBareMetal2ProvisionNetworkEventView{}
	if err := cli.Post("v1/baremetal2/provision-networks", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
