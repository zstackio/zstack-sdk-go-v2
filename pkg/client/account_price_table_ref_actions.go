// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAccountPriceTableRef queries AccountPriceTableRef list
func (cli *ZSClient) QueryAccountPriceTableRef(params *param.QueryParam) ([]view.AccountPriceTableRefInventoryView, error) {
	var resp []view.AccountPriceTableRefInventoryView
	return resp, cli.List("v1/accounts/price-tables/refs", params, &resp)
}

func (cli *ZSClient) GetAccountPriceTableRef(uuid string) (*view.AccountPriceTableRefInventoryView, error) {
	var resp view.AccountPriceTableRefInventoryView
	if err := cli.Get("v1/accounts/price-tables/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccountPriceTableRef Pagination
func (cli *ZSClient) PageAccountPriceTableRef(params *param.QueryParam) ([]view.AccountPriceTableRefInventoryView, int, error) {
	var accountPriceTableRefs []view.AccountPriceTableRefInventoryView
	total, err := cli.Page("v1/accounts/price-tables/refs", params, &accountPriceTableRefs)
	return accountPriceTableRefs, total, err
}
