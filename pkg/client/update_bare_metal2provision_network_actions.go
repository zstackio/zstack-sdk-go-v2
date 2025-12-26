// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBareMetal2ProvisionNetwork updates BareMetal2ProvisionNetwork
func (cli *ZSClient) UpdateBareMetal2ProvisionNetwork(uuid string, params param.UpdateBareMetal2ProvisionNetworkParam) (*view.UpdateBareMetal2ProvisionNetworkEventView, error) {
	resp := view.UpdateBareMetal2ProvisionNetworkEventView{}
	if err := cli.Put("v1/baremetal2/provision-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
