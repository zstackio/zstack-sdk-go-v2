// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAccountPriceTableRef queries AccountPriceTableRef list
func (cli *ZSClient) QueryAccountPriceTableRef(params *param.QueryParam) ([]view.AccountPriceTableRefInventoryView, error) {
	var resp []view.AccountPriceTableRefInventoryView
	return resp, cli.List("v1/accounts/price-tables/refs", params, &resp)
}
