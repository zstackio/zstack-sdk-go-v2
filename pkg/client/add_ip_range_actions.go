// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddIpRange adds IpRange
func (cli *ZSClient) AddIpRange(params param.AddIpRangeParam) (*view.AddIpRangeEventView, error) {
	resp := view.AddIpRangeEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/ip-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
