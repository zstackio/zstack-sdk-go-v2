// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPriceTable queries PriceTable list
func (cli *ZSClient) QueryPriceTable(params param.QueryParam) ([]view.PriceTableInventoryView, error) {
	var resp []view.PriceTableInventoryView
	return resp, cli.List("v1/billings/price-tables", &params, &resp)
}
