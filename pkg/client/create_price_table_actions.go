// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreatePriceTable creates PriceTable
func (cli *ZSClient) CreatePriceTable(params param.CreatePriceTableParam) (*view.CreatePriceTableEventView, error) {
	resp := view.CreatePriceTableEventView{}
	if err := cli.Post("v1/billings/price-tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
