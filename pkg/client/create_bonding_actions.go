// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateBonding creates Bonding
func (cli *ZSClient) CreateBonding(params param.CreateBondingParam) (*view.CreateBondingEventView, error) {
	resp := view.CreateBondingEventView{}
	if err := cli.Post("v1/hosts/bondings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
