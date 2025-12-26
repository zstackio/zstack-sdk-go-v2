// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPriceTable queries PriceTable list
func (cli *ZSClient) QueryPriceTable(params *param.QueryParam) ([]view.PriceTableInventoryView, error) {
	var resp []view.PriceTableInventoryView
	return resp, cli.List("v1/billings/price-tables", params, &resp)
}
