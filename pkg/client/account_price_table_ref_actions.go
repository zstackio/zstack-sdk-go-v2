// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAccountPriceTableRef queries AccountPriceTableRef list
func (cli *ZSClient) QueryAccountPriceTableRef(params *param.QueryParam) ([]view.AccountPriceTableRefInventoryView, error) {
	var resp []view.AccountPriceTableRefInventoryView
	return resp, cli.List("v1/accounts/price-tables/refs", params, &resp)
}
// GetAccountPriceTableRef gets AccountPriceTableRef by uuid
func (cli *ZSClient) GetAccountPriceTableRef(uuid string) (*view.AccountPriceTableRefInventoryView, error) {
	var resp view.AccountPriceTableRefInventoryView
	if err := cli.Get("v1/billings/price-tables/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
