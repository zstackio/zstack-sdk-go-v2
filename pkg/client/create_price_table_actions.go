// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreatePriceTable creates PriceTable
func (cli *ZSClient) CreatePriceTable(params param.CreatePriceTableParam) (*view.CreatePriceTableEventView, error) {
	resp := view.CreatePriceTableEventView{}
	if err := cli.Post("v1/billings/price-tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
