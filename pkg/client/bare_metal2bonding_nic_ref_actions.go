// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryBareMetal2BondingNicRef queries BareMetal2BondingNicRef list
func (cli *ZSClient) QueryBareMetal2BondingNicRef(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, error) {
	var resp []view.BareMetal2ChassisInventoryView
	return resp, cli.List(ctx, "v1/baremetal2/bonding/nic/refs", params, &resp)
}

func (cli *ZSClient) GetBareMetal2BondingNicRef(ctx context.Context, uuid string) (*view.BareMetal2ChassisInventoryView, error) {
	var resp view.BareMetal2ChassisInventoryView
	if err := cli.Get(ctx, "v1/baremetal2/bonding/nic/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBareMetal2BondingNicRef Pagination
func (cli *ZSClient) PageBareMetal2BondingNicRef(ctx context.Context, params *param.QueryParam) ([]view.BareMetal2ChassisInventoryView, int, error) {
	var bareMetal2BondingNicRefs []view.BareMetal2ChassisInventoryView
	total, err := cli.Page(ctx, "v1/baremetal2/bonding/nic/refs", params, &bareMetal2BondingNicRefs)
	return bareMetal2BondingNicRefs, total, err
}
