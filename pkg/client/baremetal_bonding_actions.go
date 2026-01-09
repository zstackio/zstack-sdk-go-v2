// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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

func (cli *ZSClient) GetBaremetalBonding(uuid string) (*view.BaremetalBondingInventoryView, error) {
	var resp view.BaremetalBondingInventoryView
	if err := cli.Get("v1/baremetal/network/bondings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
