// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddReservedIpRange adds ReservedIpRange
func (cli *ZSClient) AddReservedIpRange(params param.AddReservedIpRangeParam) (*view.AddReservedIpRangeEventView, error) {
	resp := view.AddReservedIpRangeEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/reserved-ip-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
