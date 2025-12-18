// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAccountPriceTableRef queries AccountPriceTableRef list
func (cli *ZSClient) QueryAccountPriceTableRef(params param.QueryParam) ([]view.AccountPriceTableRefInventoryView, error) {
	var resp []view.AccountPriceTableRefInventoryView
	return resp, cli.List("v1/accounts/price-tables/refs", &params, &resp)
}
