// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2Bonding queries BareMetal2Bonding list
func (cli *ZSClient) QueryBareMetal2Bonding(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2BondingInventoryView, error) {
	var resp []view.BareMetal2BondingInventoryView
	return resp, cli.List(ctx, "v1/baremetal2/bonding", params, &resp)
}

func (cli *ZSClient) GetBareMetal2Bonding(ctx context.Context, uuid string) (*view.BareMetal2BondingInventoryView, error) {
	var resp view.BareMetal2BondingInventoryView
	if err := cli.Get(ctx, "v1/baremetal2/bonding", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2Bonding Pagination
func (cli *ZSClient) PageBareMetal2Bonding(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2BondingInventoryView, int, error) {
	var bareMetal2Bondings []view.BareMetal2BondingInventoryView
	total, err := cli.Page(ctx, "v1/baremetal2/bonding", params, &bareMetal2Bondings)
	return bareMetal2Bondings, total, err
}
// CreateBareMetal2Bonding creates BareMetal2Bonding
func (cli *ZSClient) CreateBareMetal2Bonding(ctx context.Context, params param.CreateBareMetal2BondingParam) (*view.BareMetal2BondingInventoryView, error) {
	resp := view.BareMetal2BondingInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/baremetal2/chassis/bond", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
