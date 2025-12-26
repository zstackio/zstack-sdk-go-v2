// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdatePriceTable updates PriceTable
func (cli *ZSClient) UpdatePriceTable(uuid string, params param.UpdatePriceTableParam) (*view.UpdatePriceTableEventView, error) {
	resp := view.UpdatePriceTableEventView{}
	if err := cli.Put("v1/billings/price-tables/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
