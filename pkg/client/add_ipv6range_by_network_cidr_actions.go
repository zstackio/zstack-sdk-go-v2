// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIpv6RangeByNetworkCidr 操作AddIpv6RangeByNetworkCidr
func (cli *ZSClient) AddIpv6RangeByNetworkCidr(params param.AddIpv6RangeByNetworkCidrParam) (*view.AddIpRangeByNetworkCidrEventView, error) {
	resp := view.AddIpRangeByNetworkCidrEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/ipv6-ranges/by-cidr", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

