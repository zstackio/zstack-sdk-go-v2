// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddReservedIpRange 操作AddReservedIpRange
func (cli *ZSClient) AddReservedIpRange(params param.AddReservedIpRangeParam) (*view.AddReservedIpRangeEventView, error) {
	resp := view.AddReservedIpRangeEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/reserved-ip-ranges", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

