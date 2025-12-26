// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetIpAddressCapacity gets IpAddressCapacity by uuid
func (cli *ZSClient) GetIpAddressCapacity(uuid string) (*view.GetIpAddressCapacityView, error) {
	var resp view.GetIpAddressCapacityView
	if err := cli.Get("v1/ip-capacity", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
