// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBareMetal2BondingNicRef queries BareMetal2BondingNicRef list
func (cli *ZSClient) QueryBareMetal2BondingNicRef(params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, error) {
	var resp []view.BareMetal2ChassisInventoryView
	return resp, cli.List("v1/baremetal2/bonding/nic/refs", params, &resp)
}
