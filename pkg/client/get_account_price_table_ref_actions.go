// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetAccountPriceTableRef gets AccountPriceTableRef by uuid
func (cli *ZSClient) GetAccountPriceTableRef(uuid string) (*view.GetAccountPriceTableRefView, error) {
	var resp view.GetAccountPriceTableRefView
	if err := cli.Get("v1/billings/price-tables/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
