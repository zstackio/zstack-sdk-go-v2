// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2BondingNicRef queries BareMetal2BondingNicRef list
func (cli *ZSClient) QueryBareMetal2BondingNicRef(params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, error) {
	var resp []view.BareMetal2ChassisInventoryView
	return resp, cli.List("v1/baremetal2/bonding/nic/refs", params, &resp)
}

// PageBareMetal2BondingNicRef Pagination
func (cli *ZSClient) PageBareMetal2BondingNicRef(params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, int, error) {
	var bareMetal2BondingNicRefs []view.BareMetal2ChassisInventoryView
	total, err := cli.Page("v1/baremetal2/bonding/nic/refs", params, &bareMetal2BondingNicRefs)
	return bareMetal2BondingNicRefs, total, err
}
