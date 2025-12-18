// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddIpv6Range adds Ipv6Range
func (cli *ZSClient) AddIpv6Range(params param.AddIpv6RangeParam) (*view.AddIpRangeEventView, error) {
	resp := view.AddIpRangeEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/ipv6-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
