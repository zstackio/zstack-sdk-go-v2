// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddIpRangeByNetworkCidr adds IpRangeByNetworkCidr
func (cli *ZSClient) AddIpRangeByNetworkCidr(params param.AddIpRangeByNetworkCidrParam) (*view.AddIpRangeByNetworkCidrEventView, error) {
	resp := view.AddIpRangeByNetworkCidrEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/ip-ranges/by-cidr", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
