// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetServiceTypeOnHostNetworkBonding operates on SetServiceTypeOnHostNetworkBonding
func (cli *ZSClient) SetServiceTypeOnHostNetworkBonding(params param.SetServiceTypeOnHostNetworkBondingParam) (*view.SetServiceTypeOnHostNetworkBondingEventView, error) {
	resp := view.SetServiceTypeOnHostNetworkBondingEventView{}
	if err := cli.Post("v1/hosts/bondings/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
