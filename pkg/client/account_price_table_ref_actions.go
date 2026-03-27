// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAccountPriceTableRef queries AccountPriceTableRef list
func (cli *ZSClient) QueryAccountPriceTableRef(ctx context.Context, params *param.QueryParam) ([]view.AccountPriceTableRefInventoryView, error) {
	var resp []view.AccountPriceTableRefInventoryView
	return resp, cli.List(ctx, "v1/accounts/price-tables/refs", params, &resp)
}

func (cli *ZSClient) GetAccountPriceTableRef(ctx context.Context, uuid string) (*view.AccountPriceTableRefInventoryView, error) {
	var resp view.AccountPriceTableRefInventoryView
	if err := cli.Get(ctx, "v1/accounts/price-tables/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccountPriceTableRef Pagination
func (cli *ZSClient) PageAccountPriceTableRef(ctx context.Context, params *param.QueryParam) ([]view.AccountPriceTableRefInventoryView, int, error) {
	var accountPriceTableRefs []view.AccountPriceTableRefInventoryView
	total, err := cli.Page(ctx, "v1/accounts/price-tables/refs", params, &accountPriceTableRefs)
	return accountPriceTableRefs, total, err
}
