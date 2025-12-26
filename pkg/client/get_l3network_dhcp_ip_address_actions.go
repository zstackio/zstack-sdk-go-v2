// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetL3NetworkDhcpIpAddress gets L3NetworkDhcpIpAddress by uuid
func (cli *ZSClient) GetL3NetworkDhcpIpAddress(uuid string) (*view.GetL3NetworkDhcpIpAddressView, error) {
	var resp view.GetL3NetworkDhcpIpAddressView
	if err := cli.Get("v1/l3-networks/{l3NetworkUuid}/dhcp-ip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
