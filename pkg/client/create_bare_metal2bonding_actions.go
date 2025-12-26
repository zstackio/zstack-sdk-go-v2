// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateBareMetal2Bonding creates BareMetal2Bonding
func (cli *ZSClient) CreateBareMetal2Bonding(params param.CreateBareMetal2BondingParam) (*view.CreateBareMetal2BondingEventView, error) {
	resp := view.CreateBareMetal2BondingEventView{}
	if err := cli.Post("v1/baremetal2/chassis/bond", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
