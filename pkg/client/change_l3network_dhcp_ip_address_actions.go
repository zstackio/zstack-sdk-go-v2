// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeL3NetworkDhcpIpAddress changes L3NetworkDhcpIpAddress
func (cli *ZSClient) ChangeL3NetworkDhcpIpAddress(uuid string, params param.ChangeL3NetworkDhcpIpAddressParam) (*view.ChangeL3NetworkDhcpIpAddressEventView, error) {
	resp := view.ChangeL3NetworkDhcpIpAddressEventView{}
	if err := cli.Put("v1/l3-networks/{l3NetworkUuid}/dhcp-ip", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
