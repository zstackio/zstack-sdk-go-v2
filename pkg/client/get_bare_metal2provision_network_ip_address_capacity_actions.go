// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetBareMetal2ProvisionNetworkIpAddressCapacity gets BareMetal2ProvisionNetworkIpAddressCapacity by uuid
func (cli *ZSClient) GetBareMetal2ProvisionNetworkIpAddressCapacity(uuid string) (*view.GetBareMetal2ProvisionNetworkIpAddressCapacityView, error) {
	var resp view.GetBareMetal2ProvisionNetworkIpAddressCapacityView
	if err := cli.Get("v1/baremetal2/provision-networks/ip-capacity", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
