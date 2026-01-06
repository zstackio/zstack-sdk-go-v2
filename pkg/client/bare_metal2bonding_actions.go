// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2Bonding queries BareMetal2Bonding list
func (cli *ZSClient) QueryBareMetal2Bonding(params *param.QueryParam) ([]view.BareMetal2BondingInventoryView, error) {
	var resp []view.BareMetal2BondingInventoryView
	return resp, cli.List("v1/baremetal2/bonding", params, &resp)
}
// CreateBareMetal2Bonding creates BareMetal2Bonding
func (cli *ZSClient) CreateBareMetal2Bonding(params param.CreateBareMetal2BondingParam) (*view.BareMetal2BondingInventoryView, error) {
	var resp view.CreateBareMetal2BondingEventView
	if err := cli.Post("v1/baremetal2/chassis/bond", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
