// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateBaremetalBonding creates BaremetalBonding
func (cli *ZSClient) CreateBaremetalBonding(params param.CreateBaremetalBondingParam) (*view.BaremetalBondingInventoryView, error) {
	var resp view.CreateBaremetalBondingEventView
	if err := cli.Post("v1/baremetal/network/bondings", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryBaremetalBonding queries BaremetalBonding list
func (cli *ZSClient) QueryBaremetalBonding(params *param.QueryParam) ([]view.BaremetalBondingInventoryView, error) {
	var resp []view.BaremetalBondingInventoryView
	return resp, cli.List("v1/baremetal/network/bondings", params, &resp)
}
