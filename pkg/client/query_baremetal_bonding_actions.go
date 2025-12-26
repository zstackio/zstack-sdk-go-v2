// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBaremetalBonding queries BaremetalBonding list
func (cli *ZSClient) QueryBaremetalBonding(params *param.QueryParam) ([]view.BaremetalBondingInventoryView, error) {
	var resp []view.BaremetalBondingInventoryView
	return resp, cli.List("v1/baremetal/network/bondings", params, &resp)
}
