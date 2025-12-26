// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateBaremetalBonding creates BaremetalBonding
func (cli *ZSClient) CreateBaremetalBonding(params param.CreateBaremetalBondingParam) (*view.CreateBaremetalBondingEventView, error) {
	resp := view.CreateBaremetalBondingEventView{}
	if err := cli.Post("v1/baremetal/network/bondings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
