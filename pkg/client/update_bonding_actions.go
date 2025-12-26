// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateBonding updates Bonding
func (cli *ZSClient) UpdateBonding(uuid string, params param.UpdateBondingParam) (*view.UpdateBondingEventView, error) {
	resp := view.UpdateBondingEventView{}
	if err := cli.Put("v1/hosts/bondings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
