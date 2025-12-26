// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddIpv6RangeByNetworkCidr adds Ipv6RangeByNetworkCidr
func (cli *ZSClient) AddIpv6RangeByNetworkCidr(params param.AddIpv6RangeByNetworkCidrParam) (*view.AddIpRangeByNetworkCidrEventView, error) {
	resp := view.AddIpRangeByNetworkCidrEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/ipv6-ranges/by-cidr", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
